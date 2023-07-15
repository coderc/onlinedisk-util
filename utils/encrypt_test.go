package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02-15:04:05.000"))
	str := "123456"
	encStr := EncryptStrMD5(str)
	encStr2 := EncryptStrMD5(str)

	assert.Equal(t, encStr, encStr2)

	// 将加密后的字符串解密
	var err error
	encStr, err = RsaEncrypt(str)
	assert.NoError(t, err)
	decStr, err := RsaDecrypt(encStr)
	assert.NoError(t, err)
	assert.Equal(t, str, decStr)
}
