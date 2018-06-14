
文件
	f = open('/Users/michael/test.txt', 'r')
	f.read()
	f.close()

	with open('/path/to/file', 'r') as f:
    	print f.read()

	for line in f.readlines():
    	print(line.strip())

目录

	import os
	os.name # 操作系统名字
	os.uname()
	os.environ 环境变量
	os.getenv('PATH') 获取环境变量
	

	# 查看当前目录的绝对路径:
	>>> os.path.abspath('.')
	'/Users/michael'
	# 在某个目录下创建一个新目录，
	# 首先把新目录的完整路径表示出来:
	>>> os.path.join('/Users/michael', 'testdir')
	'/Users/michael/testdir'
	# 然后创建一个目录:
	>>> os.mkdir('/Users/michael/testdir')
	# 删掉一个目录:
	>>> os.rmdir('/Users/michael/testdir')

	# 对文件重命名:
	>>> os.rename('test.txt', 'test.py')
	# 删掉文件:
	>>> os.remove('test.py')

	os.path.join()
	part-1/part-2
	os.path.split()
	('/Users/michael/testdir', 'file.txt')
	os.path.splitext('/path/to/file.txt') 
	('/path/to/file', '.txt')


##  shutil可以看做是os模块的补充 提供了copyfile()