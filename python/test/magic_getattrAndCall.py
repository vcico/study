#coding=utf-8



"""
__getattr__  和 __call__ 
"""
class Chain(object):

    def __init__(self,path=''):
        self._path = path
        
    def __getattr__(self,path):
        print type(path)
        return Chain('%s/%s' % (self._path,path))
        
    def __str__(self):
        return self._path
        
    def __call__(self, s):
        print type(s)
        return Chain('%s:%s' % (self._path, s))
        
if __name__ == '__main__':
    
    #>>> Chain().status.user.timeline.list
    #'/status/user/timeline/list'
    
    """
        通过两步完成的 
        第一步 获取users时  返回的是一个 Chain对象
        第二部 调用users时 用到了__call__  又返回了一个Chain对象
    """
    
    Chain().users('michael').repos
    
    #/users/:user/repos