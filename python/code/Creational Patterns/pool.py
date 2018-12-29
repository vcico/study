#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
 池  : 线程池/连接池
 http://www.bkjia.com/Pythonjc/1230462.html
https://gist.github.com/BeginMan/0afc01a5a01470372a0e3399322d233d
http://www.iteye.com/topic/938193
http://www.iteye.com/topic/938193
"""
try:
    import queue
except ImportError:  # python 2.x compatibility
    import Queue as queue

class ObjectPool(object):

    def __init__(self,queue,auto_get=False):
        self._queue = queue
        self.item = self._queue.get() if auto_get else None

    def __enter__(self): # with 打开
        if self.item is None:
            self.item = self._queue.get()
        print('enter: %s ' % self.item)
        return self.item

    def __exit__(self, exc_type, exc_val, exc_tb): # with 关闭
        if self.item is not None:
            self._queue.put(self.item)
            print("exit: %s" % self.item)
            self.item = None

    def __del__(self):      # 删除变量 \ 析构器
        if self.item is not None:
            self._queue.put(self.item)
            print("del: %s" % self.item)
            self.item = None

def main():
    sample_queue = queue.Queue()
    sample_queue.put('yam0')
    sample_queue.put('yam1')
    sample_queue.put('yam2')
    sample_queue.put('yam3')
    p = ObjectPool(sample_queue)
    for i in  range(10):
        with p as item:
            print item

if __name__ == '__main__':
    main()