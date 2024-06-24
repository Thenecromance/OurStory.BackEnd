package RelationValidator

type IRelationValidator interface {
	GenerateToken(userID int, relationType int, idx int) (string, error)
	GetTokenInfo(token string) (userID int, relationType int, err error)
}
