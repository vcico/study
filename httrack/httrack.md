	
## 扒皮软件httrack开发

	wget http://download.httrack.com/cserv.php3?File=httrack.tar.gz
	mv cserv.php3\?File\=httrack.tar.gz  httrack.tar.gz
	tar -zxvf httrack.tar.gz 
	cd httrack-3.49.2/
	./configure && make -j8 && make install DESTDIR=/

	test.c https://www.httrack.com/html/plug.html

	gcc -O -g3 -fPIC -shared -o test.so test.c
	httrack --wrapper ./test.so http://www.st-dongxing.com/
	httrack --wrapper ./test.so,user-define-args http://www.st-dongxing.com/
	du -h --max-depth 1 #查看文件夹大小


## 功能开发

- 首页 列表页 内容页TDK
- 统计代码
- 页面内容替换(正则)
- 补全站内网址
- 友链管理(页面内容替换)

