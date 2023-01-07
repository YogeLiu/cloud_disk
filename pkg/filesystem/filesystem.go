package filesystem

import (
	"context"
	"sync"

	"github.com/YogeLiu/CloudDisk/model"
	"github.com/YogeLiu/CloudDisk/pkg/filesystem/driver"

	"github.com/YogeLiu/CloudDisk/dao"
	"github.com/YogeLiu/CloudDisk/pkg/util"
)

var FSPool = sync.Pool{
	New: func() any { return &FileSystem{FileDao: dao.NewFileDao()} },
}

type FileSystem struct {
	PolicyType *string
	User       *dao.User
	FileDao    *dao.FileDao
	Handler    driver.Handler
}

func NewFileSystem(user *dao.User) *FileSystem {
	return FSPool.Get().(*FileSystem)
}

func (fs *FileSystem) RecyleFileSystem() {
	fs.reset()
	FSPool.Put(fs)
}

func (fs *FileSystem) reset() {
	fs.Handler = nil
	fs.User = nil
	fs.PolicyType = nil
}

func (fs *FileSystem) Create(ctx context.Context, file *model.FileCreateDTO) (err error) {
	return
}

func (fs *FileSystem) Delete(ctx context.Context, fids []int64) (err error) {
	err = fs.FileDao.Delete(ctx, fids)
	if err != nil {
		util.Log().Error("delete file error: %s", err)
	}
	return
}

func (fs *FileSystem) Rename(ctx context.Context, fid int64, name string) (err error) {
	err = fs.FileDao.Update(ctx, fid, map[string]interface{}{"name": name})
	if err != nil {
		util.Log().Error("rename file error: %s", err)
	}
	return
}

func (fs *FileSystem) Download(ctx context.Context) {}

func (fs *FileSystem) Move(ctx context.Context, from []int64, to int64) (err error) {
	err = fs.FileDao.Updates(ctx, from, map[string]interface{}{"parent_id": to})
	if err != nil {
		util.Log().Error("move file error: %s", err)
	}
	return
}

func (fs *FileSystem) Compress(ctx context.Context) {}

func (fs *FileSystem) Decompress(ctx context.Context) {}

func (fs *FileSystem) Search(ctx context.Context) {}
