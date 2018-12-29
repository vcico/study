
# 函数和类是运行时使用type()函数动态创建的

	print type(Class)
	<type 'type'>

> type 函数既可以返回一个对象的类型 也可以创建新类型

> type 创建类

创建class对象 三个参数

1. class名称
2. 集成的父类集合
3. class的方法名称与函数绑定 

	def fn(self,name='x'):
		print 'hello %s' % name

	# 创建 Hello 类
	Hello = type('Hello',(object,), dict(hello=fn))
	
	
*用type创建和直接写class一样。 python解释器扫描class语法 并用type创建*

# metaclass (元类) 控制类的创建行为 创建类/修改类

如同类创建实例一样  先定义 metaclass 然后创建类 也可以吧类看做metaclass创建出来的实例








