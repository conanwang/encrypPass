# encrypPass


使用aes对称加密算法，提供将明文加密为明文和将密文解密为明文的功能 主要功能函数为两个： AesEncrypt aes加密函数 AesDecrypt aes解密函数

示范代码如下：

加密明文： encryptCode := AesEncrypt(orig, key)

解密密文： decryptCode := AesDecrypt(encryptCode, key)
