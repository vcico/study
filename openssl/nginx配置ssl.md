



    #server {
    #    listen       443;
    #    server_name  localhost;

    #    ssl                  on;
    #    ssl_certificate      cert.pem;  # 证书
    #    ssl_certificate_key  cert.key;	 # 非对称密钥对

    #    ssl_session_timeout  5m;

		 要注意 访问终端是否支持该版本 协议
    #    ssl_protocols  SSLv2 SSLv3 TLSv1;
		 支持/不支持的加密算法  ！xxx  不支持 xxx算法 
			!anull 算法不能是空
			！md5 不能是MD5
    #    ssl_ciphers  HIGH:!aNULL:!MD5;
    #    ssl_prefer_server_ciphers   on;

    #    location / {
    #        root   html;
    #        index  index.html index.htm;
    #    }
    #}

	查看所支持的算法
	openssl ciphers -v 排版  【 all / HIGH:!aNULL:!MD5 / LOW 】 

	kx : 秘钥交换算法
	Au : 认证算法
	Enc：加密算法
	Mac: 完整校验算法
	这四个算法可以随意组合

	SEED-SHA                SSLv3 Kx=RSA      Au=RSA  Enc=SEED(128) Mac=SHA1
	CAMELLIA128-SHA         SSLv3 Kx=RSA      Au=RSA  Enc=Camellia(128) Mac=SHA1
	PSK-AES128-CBC-SHA      SSLv3 Kx=PSK      Au=PSK  Enc=AES(128)  Mac=SHA1
	ECDHE-RSA-DES-CBC3-SHA  SSLv3 Kx=ECDH     Au=RSA  Enc=3DES(168) Mac=SHA1
	ECDHE-ECDSA-DES-CBC3-SHA SSLv3 Kx=ECDH     Au=ECDSA Enc=3DES(168) Mac=SHA1
	EDH-RSA-DES-CBC3-SHA    SSLv3 Kx=DH       Au=RSA  Enc=3DES(168) Mac=SHA1
	EDH-DSS-DES-CBC3-SHA    SSLv3 Kx=DH       Au=DSS  Enc=3DES(168) Mac=SHA1
	ECDH-RSA-DES-CBC3-SHA   SSLv3 Kx=ECDH/RSA Au=ECDH Enc=3DES(168) Mac=SHA1
	ECDH-ECDSA-DES-CBC3-SHA SSLv3 Kx=ECDH/ECDSA Au=ECDH Enc=3DES(168) Mac=SHA1

	测试 协议 版本
	openssl s_client -connect www.baidu.com:443  -lts1/ssl2 只使用TLSv1/??协议版本访问
	
	测试 加密算法 是否支持 
	# AES128-SHA              SSLv3 Kx=RSA      Au=RSA  Enc=AES(128)  Mac=SHA1
	#ECDHE-ECDSA-DES-CBC3-SHA SSLv3 Kx= ECDH     Au=ECDSA Enc=3DES(168)
	openssl s_client -connect www.baidu.com:443 -cipher AES128-SHA / ECDH

	-cipher 参数如果没有写全： ECDH 它会自动添加其他的 也就是在支持 ecdh里面的选择一个
	如果写全了 就没的选 


## 利用nmap测试服务器支持的协议和算法 (有可能出错 效率高)


> nmap -PN -p 443 -sT --script --ssl-enum-ciphers www.baidu.com

## 一张证书 多个域名 （使用者备用名称）

	查看淘宝的证书
	# openssl x509 -inform pem -in taobao.der -noout -text

	配置文件
	openssl.cnf

	openssl req -new -key rsa.key -out rsa.csr 
	直接在配置文件里面写上 上述命令生成签名请求文件时 不在提示输入
	openssl req -new -config md.cnf -key md.key -out md.csr

	如下内容写入的md.cnf

	[req]
	prompt = no  # 不提示
	distinguished_name=dn 分配的名字 【dn】
	req_extensions = ext 加上扩展 【ext】
	[dn]  # 提示的输入
	CN = ?
	emaiAddress = ?
	O = ?
	L = ?
	ST = ?
	C = ?
	[ext]   扩展
	使用者备用名称
	subjectAltName = DNS:www.xx.com,DNS:www.ddd.com,DNS:WWW.XX.COM,IP:xx.xx.xxx.xxx

	某些CA签证中心不支持IP签证 这是规则问题 不是技术问题
	


	自签名 默认也没有使用扩展 （证书版本 V1 x V3 ok）
	文本 md.ext
	subjectAltName = DNS:www.xx.com,DNS:www.ddd.com,DNS:WWW.XX.COM,IP:xx.xx.xxx.xxx

	openssl x509 -req -days 365 -in md.csr -signkey md.key -out text.der -extfile md.ext