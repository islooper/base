package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// PKCS7Padding 填充数据
func PKCS7Padding(ciphertext []byte) []byte {

	bs := aes.BlockSize
	padding := bs - len(ciphertext)%bs
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, paddingText...)
}

// PKCS7UnPadding 放出数据
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt 加密
// origData 原始数据
// key 密码
// iv 偏移量
func AesEncrypt(origData, key, iv string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	newOrigData := []byte(origData)
	newOrigData = PKCS7Padding(newOrigData)
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(newOrigData))
	blockMode.CryptBlocks(crypted, newOrigData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

// AesDecrypt 解密
// crypted 加密值
// key 密码
// iv 偏移量
func AesDecrypt(crypted, key, iv string) (string, error) {

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	newCrypted, _ := base64.StdEncoding.DecodeString(crypted)
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(newCrypted))
	blockMode.CryptBlocks(origData, newCrypted)
	origData = PKCS7UnPadding(origData)
	return string(origData), nil
}
