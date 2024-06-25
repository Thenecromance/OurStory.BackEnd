package JWT

import (
	"errors"
	"fmt"
	"github.com/Thenecromance/OurStories/constants"
	"github.com/Thenecromance/OurStories/middleware/Authorization"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/utility/cache/lru"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	inst = New()
)

const (
	duration = time.Hour * time.Duration(1)
)

// claim is a struct that will be used to store the user information
type claim struct {
	Obj                  interface{} `json:"info"` // the user information
	jwt.RegisteredClaims             // embedded unmodified `jwt.RegisteredClaims`
}

type Service struct {
	cache *lru.Cache
}

func (s *Service) StoreToken(token_ string, claim_ any, duration_ time.Duration) error {
	s.cache.Add(token_, claim_, time.Now().Add(duration_)) // store the token to the cache
	return nil
}

func defaultRegisteredClaim() jwt.RegisteredClaims {
	return jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
}

func (s *Service) SignTokenToUser(claim_ any) (string, error) {
	c := claim{
		Obj: claim_,
	}
	c.RegisteredClaims = defaultRegisteredClaim()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenStr, err := token.SignedString([]byte(Authorization.EncryptKey))

	if err = s.StoreToken(tokenStr, claim_, duration); err != nil {
		log.Error("Error while storing token ", err)
		return "", err
	}

	return tokenStr, nil
}

func (s *Service) GetClaimFromToken(token_ string) (any, error) {
	claims := &claim{}
	_, err := jwt.ParseWithClaims(token_, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Authorization.EncryptKey), nil
	})
	if err != nil {
		log.Error("Error while parsing token ", err)
		return nil, err
	}
	return claims, nil
}

func (s *Service) GetUserClaimFromToken(token_ string) (any, error) {
	claim_, err := s.GetClaimFromToken(token_) // get full claim from token
	if err != nil {
		log.Error("Error while getting claim from token ", err)
		return nil, err
	}
	return claim_.(*claim).Obj, nil
}

func (s *Service) MarkTokenExpired(token_ string) error {
	if !s.HasToken(token_) {
		return fmt.Errorf("token %s not found", token_)
	}
	s.cache.Remove(token_)
	return nil
}

func (s *Service) HasToken(token_ string) bool {
	_, ok := s.cache.Get(token_) // just check if the token is exist or not
	return ok
}

func (s *Service) TokenValid(token_ string) bool {
	// just precheck if the token is exist or not , if not exist means the token is invalid , just don't need to check anymore
	if !s.HasToken(token_) {
		return false
	}
	_, err := s.GetClaimFromToken(token_) // just check if the token is valid or not
	return err == nil
}

func (s *Service) TokenExpired(token_ string) bool {
	claims, err := s.GetClaimFromToken(token_)
	if err != nil {
		return false
	}
	return claims.(*claim).ExpiresAt.Time.Before(time.Now())
}

func (s *Service) MiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {

		token, err := c.Cookie("Authorization")

		if err != nil {
			log.Error("Error while getting token from cookie ", err)
			resp := response.New()
			resp.SetCode(response.NetworkAuthenticationRequired).AddData("Invalid token provided")
			resp.Send(c)
			c.Abort()
			return
		}
		// check if the token is empty
		if token == "" {
			resp := response.New()
			resp.SetCode(response.NetworkAuthenticationRequired).AddData("No token provided")
			c.Abort()
			return
		}

		// 7 is the length of "Bearer "
		// sub string the token to get the real token

		//token = token[7:]

		// check if the token is expired
		if s.TokenExpired(token) {
			resp := response.New()
			resp.SetCode(response.NetworkAuthenticationRequired).AddData("your token has been expired")
			resp.Send(c)
			c.Abort()
			return
		}

		// check if the token is valid
		if !s.TokenValid(token) {
			resp := response.New()
			log.Warn("Invalid token provided")
			resp.Error("Invalid token provided")
			resp.Send(c)
			c.Abort()
			return
		}

		// if the token is valid and not expired, then continue the request
		userClaim, err := s.GetUserClaimFromToken(token)
		if err != nil {
			resp := response.New()
			log.Error("Error while getting user claim from token ", err)
			resp.SetCode(response.NetworkAuthenticationRequired).AddData("Something goes wrong while getting user claim from token")
			resp.Send(c)
			c.Abort()
			return
		}
		//based on type to do the type assertion
		if userClaim == nil {
			log.Error("Error while getting user claim from token ", err)
			resp := response.New()
			resp.SetCode(response.NetworkAuthenticationRequired).AddData("Invalid token provided")
			resp.Send(c)
			c.Abort()
			return
		}

		c.Set(constants.AuthObject, userClaim) // set the user claim to the context

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
	return Instance().MiddleWare()
}

func TokenValid(ctx *gin.Context) (bool, error) {
	//{
	//	cookie, _ := ctx.Cookie("Authorization")
	//	/*if err != nil {
	//		log.Error("Error while getting token from cookie ", err)
	//		resp.SetCode(response.BadRequest).AddData("Something wrong with the cookie")
	//		return
	//	}*/
	//	if cookie != "" {
	//		if Instance().TokenValid(cookie) {
	//			return
	//		}
	//	}
	//
	//}

	cookie, err := ctx.Cookie("Authorization")
	if err != nil || cookie == "" {
		log.Error("Error while getting token from cookie ", err)
		return false, errors.New("please login first")
	}

	return Instance().TokenValid(cookie), nil
}

func ValidAndGetResult(ctx *gin.Context) (any, error) {
	ok, err := TokenValid(ctx)
	if err != nil || !ok {
		return nil, errors.New("please login first")
	}
	claims, err := Instance().GetUserClaimFromToken(ctx.GetString("Authorization"))
	if err != nil {
		return nil, errors.New("please login first")
	}
	return claims, nil
}
