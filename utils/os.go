package utils

import (
	"fmt"
	"os"
	"time"
)

// GetFileBytes 获取文件字节
func GetFileBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()
	bytes := make([]byte, fileSize)

	_, err = file.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func CreateFileName(router string) string {
	// {router}-{timeStamp}
	return fmt.Sprintf("%s-%s", router, time.Now().Format("2006-01-02-15:04:05.000"))
}
