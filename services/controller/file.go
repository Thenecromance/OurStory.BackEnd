package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/route"
	"github.com/Thenecromance/OurStories/services/services"
	"github.com/gin-gonic/gin"
)

type fileRoutes struct {
	fileInfos      Interface.IRoute
	fileOperations Interface.IRoute
	AllFileList    Interface.IRoute
}

type FileController struct {
	fileRoutes
	service services.FileService
}

func (fc *FileController) Initialize() {
	fc.fileOperations = route.NewREST("/api/file")
	{
		fc.fileOperations.SetHandler(
			fc.getFile,
			fc.addFile,
			fc.editFile,
			fc.deleteFile,
		)
	}
	fc.fileInfos = route.NewREST("/api/file/info")
	{
		fc.fileInfos.SetHandler(
			fc.getFileInfo,
			fc.addFileInfo,
			fc.editFileInfo,
			fc.deleteFileInfo,
		)
	}

}

//------------------------------------------------------------
// handlers
//------------------------------------------------------------

func (fc *FileController) getFile(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}
func (fc *FileController) addFile(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}
func (fc *FileController) editFile(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}
func (fc *FileController) deleteFile(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}

//------------------------------------------------------------
// infos
//------------------------------------------------------------

func (fc *FileController) getFileInfo(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}
func (fc *FileController) addFileInfo(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}
func (fc *FileController) editFileInfo(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

}
func (fc *FileController) deleteFileInfo(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

}

// ------------------------------------------------------------
// list
// ------------------------------------------------------------
func (fc *FileController) getAllFileList(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}

func NewFileController(service services.FileService) *FileController {
	fc := &FileController{
		service: service,
	}
	fc.Initialize()
	return fc
}
