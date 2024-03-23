package fio

import "os"

// FileIO 标准系统文件 IO
type FileIO struct {
	fd *os.File // 系统文件描述符
}

// NewFileIOManager 初始化标准文件 IO
func NewFileIOManager(fileName string) (*FileIO, error) {
	fd, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		DataFilePerm,
	)
	if err != nil {
		return nil, err
	}
	return &FileIO{fd: fd}, nil
}

func (f FileIO) Read(bytes []byte, offset int64) (int, error) {
	return f.fd.ReadAt(bytes, offset)
}

func (f FileIO) Write(bytes []byte) (int, error) {
	return f.fd.Write(bytes)
}

func (f FileIO) Sync() error {
	return f.fd.Sync()
}

func (f FileIO) Close() error {
	return f.fd.Close()
}

func (f FileIO) Size() (int64, error) {
	stat, err := f.fd.Stat()
	if err != nil {
		return 0, err
	}

	return stat.Size(), nil
}
