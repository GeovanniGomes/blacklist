package contracts

import "bytes"

type IFileSystem interface {
	Upload(bucketName, fileName string, buf bytes.Buffer)
}
