package repository

import (
	"bytes"
	"log"
	"os"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/minio/minio-go"
)

var _ contracts.IFileSystem = (*MinIO)(nil)

type MinIO struct {
	client *minio.Client
}

func NewMinIO() *MinIO {
	endpoint := os.Getenv("MINIO_DATA_LAKE_URL")
	accessKeyID := os.Getenv("MINIO_ROOT_USER")
	secretAccessKey := os.Getenv("MINIO_ROOT_PASSWORD")
	useSSL := true

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	return &MinIO{client: minioClient}

}

func (m *MinIO) Upload(bucketName string, fileName string, buf bytes.Buffer) {
	info, err := m.client.PutObject(
		bucketName,
		fileName,
		&buf,
		int64(buf.Len()),
		minio.PutObjectOptions{ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
	)

	log.Printf("Successfully uploaded %s of size %d\n", bucketName, info)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", fileName, buf.Len())
}
