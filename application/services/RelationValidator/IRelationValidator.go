package RelationValidator

type IRelationValidator interface {
	GenerateToken(userID int64, relationType int, idx int) (string, error)
	GetTokenInfo(token string) (userID int64, relationType int, err error)
}
