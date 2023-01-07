package driver

import (
	"context"

	"github.com/YogeLiu/CloudDisk/model"
	"github.com/YogeLiu/CloudDisk/server/fsctx"
)

// Handler 存储策略适配器
type Handler interface {
	// 上传文件, dst为文件存储路径，size 为文件大小。上下文关闭
	// 时，应取消上传并清理临时文件
	Put(ctx context.Context, file fsctx.FileHeader) error

	// 删除一个或多个给定路径的文件，返回删除失败的文件路径列表及错误
	Delete(ctx context.Context, files []string) ([]string, error)

	Token(ctx context.Context) (*model.Credential, error)
	// 获取缩略图，可直接在ContentResponse中返回文件数据流，也可指
	// 定为重定向
	// Thumb(ctx context.Context, path string) (*response.ContentResponse, error)

	// 获取外链/下载地址，
	// url - 站点本身地址,
	// isDownload - 是否直接下载
	// Source(ctx context.Context, path string, url url.URL, ttl int64, isDownload bool, speed int) (string, error)

	// Token 获取有效期为ttl的上传凭证和签名
	// Token(ctx context.Context, ttl int64, uploadSession *serializer.UploadSession, file fsctx.FileHeader) (*serializer.UploadCredential, error)

	// CancelToken 取消已经创建的有状态上传凭证
	// CancelToken(ctx context.Context, uploadSession *serializer.UploadSession) error
}
