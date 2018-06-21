开发库/工具集

随机数
hash函数 
加解密  - enc 、 aes-128-cbc
证书管理
tls\ssl - s-client、 s-server

证书的格式

DER/CER    原始格式 二进制的 打开看 乱码
PEM        base64格式 可以打开看  可读 输入 比如需要贴到配置文件里 

证书格式转换

openssl x509 -inform PEM -in xx.pem -outform DER -out xx.der/cer
openssl x509 -inform DER -in xx.der -outform PEM -out xx.pen

openssl制作自签名的证书

1. 非对称密钥对的生成  ： 公钥 私钥

	  RSA 经常使用 浏览器兼容更好 
		openssl genrsa -out rsa.key 2048 
			genrsa ： generation rsa 生成 RSA key
			rsa.key 有私钥/公钥/其他相关的数据 
			2048 长度
		openssl rsa -text -in rsa.key 结构化查看
		openssl rsa -in rsa.key -pubout -out rsa-public.key # 导出公钥
	  DSA 
		openssl dsaparam -genkey 2048 | openssl dsa -out dsa.key 生成 DSA key
		
	  ECDSA  如果用VPN 客户端服务器都是自己的 可以选择其他两项
		openssl ecparam -genkey -name secp256r1 | openssl ec -out ec.key 生成 ECDSA key

2. 创建证书签名请求文件：CSR: certificate signing requests 
	
	包含： 公钥 证书使用者名字
	RSA：
		openssl req -new -key rsa.key -out rsa.csr # 需要输入相关信息 对用 使用者信息
			common name 是别人能看到来确认你的 认真填写

	openssl req -text -in rsa.csr -noout 【不输出本身的base64文本】 查看Key
	
3. 【自】签名证书
	
	包括： 版本号 序列号 签名算法 指纹算法(谁签的)
	
	把csr文件 交给 CA中心 或者自己签 来生成
	cer der  pem crt等
	
	自签名
	openssl x509 -req -days 365 -in res.csr -signkey rsa.key -out rsa_xxx.cer
	

openssl 查看证书

	openssl x509 -in baidu.pem | -text 结构化  -noout 不输出base64文本内容 


openssl 获取服务器证书

	openssl s_client -connect --help 

	
	>$ openssl s_client -connect www.baidu.com:443  【-showcerts 打印整个证书链】 

	CONNECTED(00000003)
	depth=3 C = US, O = "VeriSign, Inc.", OU = Class 3 Public Primary Certification Authority
	verify return:1
	depth=2 C = US, O = "VeriSign, Inc.", OU = VeriSign Trust Network, OU = "(c) 2006 VeriSign, Inc. - For authorized use only", CN = VeriSign Class 3 Public Primary Certification Authority - G5
	verify return:1
	depth=1 C = US, O = Symantec Corporation, OU = Symantec Trust Network, CN = Symantec Class 3 Secure Server CA - G4
	verify return:1
	depth=0 C = CN, ST = beijing, L = beijing, O = "BeiJing Baidu Netcom Science Technology Co., Ltd", OU = service operation department., CN = baidu.com
	verify return:1
	---
	Certificate chain  证书链 某公司颁发给某子证书再颁发给百度的证书 总共3个证书  s-subject-拥有这个证书的主题  i-issue-颁发这张证书的机构
	 0 s:/C=CN/ST=beijing/L=beijing/O=BeiJing Baidu Netcom Science Technology Co., Ltd/OU=service operation department./CN=baidu.com
	   i:/C=US/O=Symantec Corporation/OU=Symantec Trust Network/CN=Symantec Class 3 Secure Server CA - G4
	 1 s:/C=US/O=Symantec Corporation/OU=Symantec Trust Network/CN=Symantec Class 3 Secure Server CA - G4
	   i:/C=US/O=VeriSign, Inc./OU=VeriSign Trust Network/OU=(c) 2006 VeriSign, Inc. - For authorized use only/CN=VeriSign Class 3 Public Primary Certification Authority - G5
	 2 s:/C=US/O=VeriSign, Inc./OU=VeriSign Trust Network/OU=(c) 2006 VeriSign, Inc. - For authorized use only/CN=VeriSign Class 3 Public Primary Certification Authority - G5
	   i:/C=US/O=VeriSign, Inc./OU=Class 3 Public Primary Certification Authority
	---
	Server certificate   # 默认只打印了一张证书
	-----BEGIN CERTIFICATE-----
	MIIIdDCCB1ygAwIBAgIQRgvtzGxo+wBn8JgNuE2/gjANBgkqhkiG9w0BAQsFADB+
	MQswCQYDVQQGEwJVUzEdMBsGA1UEChMUU3ltYW50ZWMgQ29ycG9yYXRpb24xHzAd
	BgNVBAsTFlN5bWFudGVjIFRydXN0IE5ldHdvcmsxLzAtBgNVBAMTJlN5bWFudGVj
	IENsYXNzIDMgU2VjdXJlIFNlcnZlciBDQSAtIEc0MB4XDTE3MDYyOTAwMDAwMFoX
	DTE4MDgxNzIzNTk1OVowgagxCzAJBgNVBAYTAkNOMRAwDgYDVQQIDAdiZWlqaW5n
	MRAwDgYDVQQHDAdiZWlqaW5nMTkwNwYDVQQKDDBCZWlKaW5nIEJhaWR1IE5ldGNv
	bSBTY2llbmNlIFRlY2hub2xvZ3kgQ28uLCBMdGQxJjAkBgNVBAsMHXNlcnZpY2Ug
	b3BlcmF0aW9uIGRlcGFydG1lbnQuMRIwEAYDVQQDDAliYWlkdS5jb20wggEiMA0G
	CSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDImQ8LQt6/ovSyE1jc5M7knA7yDHOb
	ZUA8OgG8ciml3aPKDbuVQ0JLcuFkRBb74nVxe9EAwgN7NMrSL0dSTl2pYonMfklj
	7ymfr+LKUigoO8So1XKUsntq6KU7/qXRI4Tpd9zoWoTXUSMgM3BvjN9tyows9GSm
	v2foMURsnO00az4YITnM2bVuklui6j4TfkRuA/uypcLPRrZM6XnxRvHrznsMCZ5N
	wqpOQeDQuQCGaFqkKPFlbi8Rb+LPFAICUmmqFplit6ac7gvQeLBN3BeJx1Pu4Jmg
	oSTPKU0OtRtQtuXaY6+oqpGGP9gqvE5eWYy2YPA2AeojmiNsSPrzf8bDAgMBAAGj
	ggTBMIIEvTCCAm8GA1UdEQSCAmYwggJiggsqLmJhaWR1LmNvbYIOKi5iYWlmdWJh
	by5jb22CEiouYmFpZHVjb250ZW50LmNvbYIRKi5iYWlkdXN0YXRpYy5jb22CDiou
	YmRzdGF0aWMuY29tggsqLmJkaW1nLmNvbYIMKi5iYWlmYWUuY29tggwqLmhhbzEy
	My5jb22CCyoubnVvbWkuY29tgg0qLmNodWFua2UuY29tgg0qLnRydXN0Z28uY29t
	gg4qLmJhaWR1YmNlLmNvbYIMKi5taXBjZG4uY29tgg4qLmJhaWR1cGNzLmNvbYIN
	Ki5iY2Vob3N0LmNvbYIMKi5haXBhZ2UuY29tggsqLmFpcGFnZS5jboILKi5kbG5l
	bC5jb22CDyouYmNlLmJhaWR1LmNvbYIQKi5leXVuLmJhaWR1LmNvbYIPKi5tYXAu
	YmFpZHUuY29tgg8qLm1iZC5iYWlkdS5jb22CECoubmV3cy5iYWlkdS5jb22CESou
	ZmFueWkuYmFpZHUuY29tgg4qLmltLmJhaWR1LmNvbYIQKi5zYWZlLmJhaWR1LmNv
	bYIRKi5zc2wyLmR1YXBwcy5jb22CCWJhaWR1LmNvbYIMYmFpZnViYW8uY29tggpi
	YWlmYWUuY29tggx3d3cuYmFpZHUuY26CEHd3dy5iYWlkdS5jb20uY26CEmNsaWNr
	LmhtLmJhaWR1LmNvbYIQbG9nLmhtLmJhaWR1LmNvbYIQY20ucG9zLmJhaWR1LmNv
	bYIQd24ucG9zLmJhaWR1LmNvbYIUdXBkYXRlLnBhbi5iYWlkdS5jb22CD21jdC55
	Lm51b21pLmNvbTAJBgNVHRMEAjAAMA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAU
	BggrBgEFBQcDAQYIKwYBBQUHAwIwYQYDVR0gBFowWDBWBgZngQwBAgIwTDAjBggr
	BgEFBQcCARYXaHR0cHM6Ly9kLnN5bWNiLmNvbS9jcHMwJQYIKwYBBQUHAgIwGQwX
	aHR0cHM6Ly9kLnN5bWNiLmNvbS9ycGEwHwYDVR0jBBgwFoAUX2DPYZBV34RDFIpg
	KrL1evRDGO8wKwYDVR0fBCQwIjAgoB6gHIYaaHR0cDovL3NzLnN5bWNiLmNvbS9z
	cy5jcmwwVwYIKwYBBQUHAQEESzBJMB8GCCsGAQUFBzABhhNodHRwOi8vc3Muc3lt
	Y2QuY29tMCYGCCsGAQUFBzAChhpodHRwOi8vc3Muc3ltY2IuY29tL3NzLmNydDCC
	AQIGCisGAQQB1nkCBAIEgfMEgfAA7gB1AN3rHSt6DU+mIIuBrYFocH4ujp0B1VyI
	jT0RxM227L7MAAABXPLl8Q4AAAQDAEYwRAIgYmI+Xnom5nr+k+p8RW/BdqobpmTN
	ubZW8eM1DTueRfcCIAXdDURAie+w+eZMDfgxHlo/r9wZ8gpQVoTib5w3b/6+AHUA
	pLkJkLQYWBSHuxOizGdwCjw1mAT5G9+443fNDsgN3BAAAAFc8uXxRAAABAMARjBE
	AiA+PkSvZOvIBmxaBLXLceHH1NqW+Mcrz+xtXNmJOl2ElQIgfVeG4xeuWrb7bpoU
	smR+LStIG1TPCwimn3JRE3we/fYwDQYJKoZIhvcNAQELBQADggEBADjrCz8a7cax
	h7vpyuUFZ/fiKBHE7VLqfppgf3XYNBoqh21qM6gTGzdiSeZj+vx+KOUn38f080Rg
	N2aEkag3n03cufIXR8Yn8haXcusz5PONSlMQnN5rZBwpZ8obItiO8KGOh5lgHQ+s
	SloX/j8nDDCQgrNkcG2A78nUT+VxGGENxnPmqajP/O2h/kg02qjcnPoj6Elmm/At
	5dWWANX374yS7c0fgLZZ1mfZoIqooaRxsSJl5RzyRNU3Bzv5CZCJCGYFqC3RS28Q
	vTCjde7TMsAQiWkZ97IKlUMXdbHManm7K85aWcG4Wg8isr9d2GPUZYgcUSc8KfWY
	aP5MzoeU6ug=
	-----END CERTIFICATE-----
	subject=/C=CN/ST=beijing/L=beijing/O=BeiJing Baidu Netcom Science Technology Co., Ltd/OU=service operation department./CN=baidu.com
	issuer=/C=US/O=Symantec Corporation/OU=Symantec Trust Network/CN=Symantec Class 3 Secure Server CA - G4
	---
	No client certificate CA names sent
	Server Temp Key: ECDH, prime256v1, 256 bits
	---
	SSL handshake has read 5396 bytes and written 373 bytes
	---
	New, TLSv1/SSLv3, Cipher is ECDHE-RSA-AES128-GCM-SHA256
	Server public key is 2048 bit
	Secure Renegotiation IS supported
	Compression: NONE
	Expansion: NONE
	SSL-Session:
	    Protocol  : TLSv1.2			# 协议
	    Cipher    : ECDHE-RSA-AES128-GCM-SHA256  # 算法
	    Session-ID: D6F07F23090FAB9D1CB25C2ECABB4890C98CF00E0FE21D825155C3BD8DA2568A
	    Session-ID-ctx: 
	    Master-Key: A746103CBE9FE941F4199658D9463D19CEDE3085D06F81ABD7E1BFF940838051EDA3A10F371B42FA9BD1BD373FA3505D
	    Key-Arg   : None
	    Krb5 Principal: None
	    PSK identity: None
	    PSK identity hint: None
	    TLS session ticket:
	    0000 - 5b 53 02 28 e2 f4 bc 6b-99 09 03 66 59 21 50 29   [S.(...k...fY!P)
	    0010 - 54 4a 05 e7 25 d8 40 3d-6e 1b 32 3e 1e 11 49 91   TJ..%.@=n.2>..I.
	    0020 - b9 4f d3 fe 06 fe cc 33-6c d2 25 6c 91 d7 f9 e0   .O.....3l.%l....
	    0030 - ae a0 b0 cd db 4e 3a e5-15 48 9e 25 3c 94 b9 f2   .....N:..H.%<...
	    0040 - ca 7d d3 41 15 5c b3 2c-88 d3 52 21 5b 62 8f 98   .}.A.\.,..R![b..
	    0050 - 37 e1 0f 69 47 0f a9 43-26 59 d9 94 72 ff 74 2c   7..iG..C&Y..r.t,
	    0060 - f5 14 53 43 95 5f 91 a1-38 97 3e be 70 3e ea b9   ..SC._..8.>.p>..
	    0070 - 1b 2c 30 e3 8f a7 9f 76-e4 75 8e 16 68 a1 9e 15   .,0....v.u..h...
	    0080 - 2d 24 04 40 99 da fb 47-e4 e9 5a 75 4f 07 7a 39   -$.@...G..ZuO.z9
	    0090 - 93 d6 19 81 aa 6d af 7b-52 f8 5d ec 0c e8 1a d8   .....m.{R.].....
	
	    Start Time: 1527580250
	    Timeout   : 300 (sec)
	    Verify return code: 0 (ok)
	---
	HEAD / HTTP/1.1     # 输入HTTP请求
	HOST:www.baidu.com 						
	HTTP/1.1 400 Bad Request
	
	closed
