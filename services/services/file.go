package services

import "github.com/Thenecromance/OurStories/services/repository"

type FileService interface {
}

type fileServiceImpl struct {
	repo repository.FileRepository
}

func NewFileService(repo repository.FileRepository) FileService {
	return &fileServiceImpl{repo}
}
