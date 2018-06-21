#coding=utf-8


class ListMetaclass(type):
    
    def __new__(cls,name,bases,attrs):
        print cls.__name__
        print name
        print bases
        print attrs
        return type.__new__(cls,name,bases,attrs)
        # return type(name,bases,attrs)
        
class MyList(list):
    __metaclass__ = ListMetaclass
    dasf = 'fdasf'
    
    def remove(self):
        pass
    
    
if __name__ == '__main__':
    
    L = MyList()
    print L
