package encrypPass

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)


//aes���ܺ�����orig����Ϊ�����ַ�����keyΪ����ʱʹ�õ���Կ
func AesEncrypt(orig string, key string) string {
	// ת���ֽ�����
	origData := []byte(orig)
	k := []byte(key)
	// ������Կ
	block, _ := aes.NewCipher(k)
	// ��ȡ��Կ��ĳ���
	blockSize := block.BlockSize()
	// ��ȫ��
	origData = PKCS7Padding(origData, blockSize)
	// ����ģʽ������cbc����
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// ��������
	cryted := make([]byte, len(origData))
	// ����
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

//����
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}


//ȥ��
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//aes���ܺ�����cryted����Ϊ�����ַ�����keyΪ����ʱʹ�õ���Կ
func AesDecrypt(cryted string, key string) string {
	// ת���ֽ�����
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// ������Կ
	block, _ := aes.NewCipher(k)
	// ��ȡ��Կ��ĳ���
	blockSize := block.BlockSize()
	// ����ģʽ
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// ��������
	orig := make([]byte, len(crytedByte))
	// ����
	blockMode.CryptBlocks(orig, crytedByte)
	// ȥ��ȫ��
	orig = PKCS7UnPadding(orig)
	return string(orig)
}