package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
)

// EncryptStr 对字符串进行MD5加密
func EncryptStr(str string) string {
	if str == "" {
		return ""
	}
	md5Bytes := md5.Sum([]byte(str))
	return hex.EncodeToString(md5Bytes[:])
}

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBAK62ZF4C6DxaSt0c74k1X4U36sipzk59PyOhUPIcFnNfG1qY+pu9
gkMPE1O75s0Q+dlekmdwLTsAIui6p8gmyKMCAwEAAQJACltblcscUz+TAoqNlJeq
Yu5Op7iRN0vra0RL1R5fIlWe0NSj8kTDBezVYxAVwebOQBO3XFFlnI3T9QUOSnll
IQIhAObiIQ7IfxdTAbHEhWg/4+2ZkG6VEoOuVrU3FbwY81tRAiEAwbf17RVKD9OH
h2PB33utRLqAwEDBTKQ5HFME1671P7MCIQCbslR+pqBl9zkGSzN3yNYI7Wzj1a2F
lXStgbcrgFvj8QIgC+Z7KxdVt2ctOjn8nPgCCujSI/1WYpjsETtgXseWtVUCIBbc
+MiSwmPz28oPljZ1c2OBayU34vDPCQc1d1FK0iwp
-----END RSA PRIVATE KEY-----
`)

// 公钥
var publicKey = []byte(`
-----BEGIN RSA PUBLIC KEY-----
MEgCQQCutmReAug8WkrdHO+JNV+FN+rIqc5OfT8joVDyHBZzXxtamPqbvYJDDxNT
u+bNEPnZXpJncC07ACLouqfIJsijAgMBAAE=
-----END RSA PUBLIC KEY-----
`)

// 加密
func RsaEncrypt(str string) (string, error) {
	block, _ := pem.Decode(publicKey) //将密钥解析成公钥实例
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKCS1PublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return "", err
	}
	encStrBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubInterface, []byte(str)) //RSA算法加密
	return hex.EncodeToString(encStrBytes), err
}

func RsaDecrypt(encStr string) (string, error) {
	block, _ := pem.Decode(privateKey) //将密钥解析成私钥实例
	if block == nil {
		return "", errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return "", err
	}

	encStrBytes, err := hex.DecodeString(encStr)
	decStr, err := rsa.DecryptPKCS1v15(rand.Reader, priv, encStrBytes) //RSA算法解密
	return string(decStr), err
}
