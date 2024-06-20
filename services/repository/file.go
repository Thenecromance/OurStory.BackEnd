package repository

import "gopkg.in/gorp.v2"

type FileRepository interface {
}

type fileRepository struct {
	db *gorp.DbMap
}

func NewFileRepository(db *gorp.DbMap) FileRepository {
	return &fileRepository{db}
}
