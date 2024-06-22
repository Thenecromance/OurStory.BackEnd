package repository

import "gopkg.in/gorp.v2"

type RelationshipRepository interface {
}

type relationshipRepositoryImpl struct {
	db *gorp.DbMap
}
