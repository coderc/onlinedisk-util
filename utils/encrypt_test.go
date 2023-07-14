package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	str := "123456"
	encStr := EncryptStr(str)
	encStr2 := EncryptStr(str)

	assert.Equal(t, encStr, encStr2)

	// 将加密后的字符串解密
	var err error
	encStr, err = RsaEncrypt(str)
	assert.NoError(t, err)
	decStr, err := RsaDecrypt(encStr)
	assert.NoError(t, err)
	assert.Equal(t, str, decStr)
}
