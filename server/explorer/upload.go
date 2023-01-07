package explorer

import (
	"github.com/YogeLiu/CloudDisk/dao"
	model "github.com/YogeLiu/CloudDisk/model"
	"github.com/YogeLiu/CloudDisk/pkg/filesystem"
	"github.com/YogeLiu/CloudDisk/pkg/util"

	"context"
)

type CreateUploadSessionService struct {
	Path string `json:"path" binding:"required"`
	Name string `json:"name" binding:"required"`
	Size uint64 `json:"size,string" binding:"required"`
}

func (svc *CreateUploadSessionService) Create(ctx context.Context) (credential *model.Credential, err error) {
	user := ctx.Value("user").(*dao.User)
	fs := filesystem.NewFileSystem(user)
	credential, err = fs.Handler.Token(ctx)
	if err != nil {
		util.Log().Error("acquire credential error: %s", err)
	}
	return
}
