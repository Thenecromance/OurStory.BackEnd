package filestore

import (
	response "github.com/Thenecromance/OurStories/backend/Response"
	"github.com/Thenecromance/OurStories/base/log"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/gin-gonic/gin"
)

type controller struct {
	group *Interface.GroupNode
}

func (ctrl *controller) Name() string {
	return "file"
}

func (ctrl *controller) BuildRoutes() {
	ctrl.group.Router.GET("/", ctrl.downloadFile)
	ctrl.group.Router.POST("/", ctrl.uploadFile)
	ctrl.group.Router.DELETE("/", ctrl.deleteFile)
}

func (ctrl *controller) RequestGroup(cb Interface.NodeCallback) {
	ctrl.group = cb(ctrl.Name(), "api")
}

func NewController(i ...Interface.Controller) Interface.Controller {
	ctrl := &controller{}
	return ctrl
}

func (ctrl *controller) uploadFile(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

	file, _ := ctx.FormFile("file")
	log.Debug("file:", file.Filename)
	err := ctx.SaveUploadedFile(file, "upload/"+file.Filename)
	if err != nil {
		log.Error(err)
		return
	}
	resp.SetCode(response.SUCCESS).AddData(file.Filename + " upload success")
}

func (ctrl *controller) downloadFile(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

	fileName := ctx.Param("filename")
	ctx.File("upload/" + fileName)
}

func (ctrl *controller) deleteFile(ctx *gin.Context) {
	resp := response.New(ctx)
	defer resp.Send()

}
