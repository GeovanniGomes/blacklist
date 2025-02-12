package service

import "github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"

type FileSystemService struct {
	client contracts.IFileSystem
}

func NewFileSystemService(client contracts.IFileSystem) *FileSystemService {
	return &FileSystemService{client: client}
}
