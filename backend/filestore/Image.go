package filestore

import "github.com/Thenecromance/OurStories/base/utils"

var (
	ImageFolder = "resource/images"
)

func init() {
	if !utils.DirExists("resource") {
		utils.CreateIfNotExist("resource")
	}

	utils.CreateIfNotExist(ImageFolder)
}

type imageStore struct {
}
