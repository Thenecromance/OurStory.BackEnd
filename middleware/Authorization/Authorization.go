package Authorization

type IAuth interface {
	AuthorizeUser(claim interface{}) (string, error)

	AuthorizeToken(token string) (interface{}, error)

	TokenExpired(token string) (bool, error)

	RefreshToken(oldToken string) (string, error)

	MarkTokenExpired(token string) error
}
