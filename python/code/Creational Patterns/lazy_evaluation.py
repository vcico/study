#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
惰性初始化
方法当成属性  且只实例化一次
举个实际场景 ...在做ORM最后生成SQL的时候，可以利用lazy computation
"""
# future 未来 版本兼容

from __future__ import print_function ## 强制 print 使用() . print 'error'   print('right')
import functools

"""
参考： http://www.wklken.me/posts/2013/08/18/python-extra-functools.html
functools.partial 【 partial偏爱 】  重定义函数签名并设置默认参数 int2(num) 等于 int(num,2) 
functool.update_wrapper  【wrapper包装器】从原始对象拷贝或加入现有partial对象（避免装饰器函数的__name__、module、__doc__和 __dict__ 变成 闭包函数的）
functool.wraps
functools.reduce
functools.cmp_to_key
functools.total_ordering
"""

class lazy_property(object):

    def __init__(self,function):
        self.function = function
        functools.update_wrapper(self,function)

    def __get__(self, obj, type_):
        if obj is None:
            return self
        value = self.function(obj)
        obj.__dict__[self.function.__name__] = value
        return value

def lazy_property2(fn):
    attr = '__lazy__' + fn.__name__

    @property
    def _lazy_property(self):
        if not hasattr(self,attr):
            setattr(self,attr,fn(self))
        return getattr(self,attr)
    return _lazy_property

class test:

    @lazy_property
    def index(self):
        print("index init")
        return  ':index page'

    @lazy_property2
    def admin(self):
        print("admin init")
        return ":admin page"


if __name__ == '__main__':
    t = test()
    print(t.index)
    print(t.index)
    print(t.admin)
    print(t.admin)
