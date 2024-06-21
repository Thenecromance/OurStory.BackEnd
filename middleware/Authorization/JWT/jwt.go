package JWT

import (
	"github.com/Thenecromance/OurStories/middleware/Authorization"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/utility/cache/lru"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	constKey = "M0nkey_Cl1cker"
)

var (
	inst = New()
)

// Claim is a struct that will be used to store the user information
type Claim struct {
	Obj                  interface{} `json:"info"` // the user information
	jwt.RegisteredClaims             // embedded unmodified `jwt.RegisteredClaims`
}

type Service struct {
	cache *lru.Cache
}

func (s *Service) MarkTokenExpired(token string) error {
	s.cache.Remove(token)
	return nil
}

func (s *Service) AuthorizeUser(claim interface{}) (string, error) {
	token_, err := s.authToken(claim, 1, "secret")
	if err != nil {
		return "", err
	}
	s.cache.Add(token_, claim, time.Now().Add(time.Hour*time.Duration(1)))
	return token_, err
}

func (s *Service) AuthorizeToken(token string) (interface{}, error) {
	if v, ok := s.cache.Get(token); ok {
		return v, nil
	}
	return s.parseToken(token, constKey)
}

func (s *Service) TokenExpired(token string) (bool, error) {
	claims, err := s.parseToken(token, constKey)
	if err != nil {
		return false, err
	}
	return claims.ExpiresAt.Time.Before(time.Now()), nil
}

func (s *Service) RefreshToken(oldToken string) (string, error) {

	claims, err := s.parseToken(oldToken, constKey)
	if err != nil {
		return "", err
	}
	return s.authToken(claims.Obj, 1, constKey)
}

func (s *Service) parseToken(token, key string) (*Claim, error) {
	claims := &Claim{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	log.Info("claims ", claims)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (s *Service) authToken(obj interface{}, expireTime int64, key string) (string, error) {
	claim := Claim{
		Obj: obj,
	}
	claim.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expireTime))),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(key))
}

func (s *Service) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Authorization")

		if token == "" {
			log.Warn("token is empty")
			resp := response.New()
			resp.SetCode(response.NotAcceptable).AddData("token is invalid")
			c.Abort()
			resp.Send(c)
			return
		}

		token = token[7:]

		expired, err := s.TokenExpired(token)
		log.Info("token  ", token, expired)
		if expired == true || err != nil {
			log.Warn("token is expired")
			resp := response.New()
			resp.SetCode(response.NotAcceptable).AddData("token is invalid")
			c.Abort()
			resp.Send(c)
			return
		}

		obj, err := s.AuthorizeToken(token)
		if err != nil {
			log.Warn("token is invalid", err)
			resp := response.New()
			resp.SetCode(response.NotAcceptable).AddData("token is invalid")
			c.Abort()
			resp.Send(c)
			return
		}
		log.Info("obj ", obj)

		c.Set(Authorization.AuthObject, obj.(*Claim).Obj)
		c.Next()
	}
}

func New() Authorization.IAuth {
	return &Service{
		cache: lru.New(0),
	}
}

// Instance will return the singleton instance of the Authorization.IAuth
func Instance() Authorization.IAuth {
	if inst == nil {
		inst = New()
	}
	return inst
}

func Middleware() gin.HandlerFunc {
	return inst.Middleware()
}
