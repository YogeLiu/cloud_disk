package fsctx

import "io"

type UploadFileInfo struct{}

type FileHeader interface {
	io.Closer
	io.Reader
	Info() *UploadFileInfo
}

type FileStream struct {
	File io.ReadCloser
}

func (f *FileStream) Info() *UploadFileInfo {
	return &UploadFileInfo{}
}

func (f *FileStream) Close() error {
	if f.File != nil {
		return f.File.Close()
	}
	return nil
}

func (f *FileStream) Read(p []byte) (n int, err error) {
	if f.File != nil {
		return f.File.Read(p)
	}
	return
}
