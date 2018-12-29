# coding=utf-8
"""
参考 https://xionchen.github.io/2017/03/01/python-patterns_02/

__dict__可以重新绑定
在__init__中绑定了之后 所有对象就是类对象 拥有 we are one 的性质
所以所有实例中的状态都是同一个(最新状态)

跟单例模式不同的区别： 单例模式从对象实例上都是相同的， borg模式 是所有的 __dict__ （所有的实例属性）是相同的

所谓的单例模式，就是确保这个类只有一个对象。虽然单例听起来感觉好像很专业的样子，这可能并非是个好主意。有时候我们可能确实想要很多对象，但是这些对象有共享的状态。我们在乎的是共享的状态，毕竟谁会在乎无限的可能

## 单例模式
class Singleton(object):
    def __new__(cls, *p, **k):
        if not '_the_instance' in cls.__dict__:
            cls._the_instance = object.__new__(cls)
        return cls._the_instance

## 多线程

无论是单例模式还是Brog模式，如果要支持多线程，在有效性和效率方面都需要考量一个平衡

## 使用场景

比如：数据库链接

子类处理方式不一样(多态)，参数不一样 但 其他状态都一样

单例：不能多态????????

"""

class Borg(object):

    __shared_state = {}

    def __init__(self):
        self.__dict__ = self.__shared_state
        self.state = 'Init'

    def __str__(self):
        # print self.__dict__
        return self.state

class YourBorg(Borg):

    def __init__(self):
        super(YourBorg, self).__init__()
        self.NewState = 'init'

    def __str__(self):
        print self.__dict__
        return '%s  %s' % (self.state,self.NewState)

if __name__ == '__main__':

    a = YourBorg()
    b = YourBorg()

    b.state = 'B-STATE'
    a.NewState = 'A-NEW-STATE'

    # a = Borg()
    # b = Borg()
    print a
    print b
    # a.__dict__ = {}
    # b.state = 'ASTATE'
    # print a
    # print b