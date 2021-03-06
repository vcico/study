#coding=utf-8




class Field(object):
    
    def __init__(self,name,colunm_type):
        self.name = name
        self.colunm_type = colunm_type
        
    def __str__(self):
        return '<%s:%s>' % (self.__class__.__name__, self.name)
        
class StringField(Field):
    
    def __init__(self,name):
        super(StringField,self).__init__(name,'varchar(100)')


        
class IntegerField(Field):
    def __init__(self, name):
        super(IntegerField, self).__init__(name, 'bigint')

        
"""
在ModelMetaclass中，一共做了几件事情：
    排除掉对Model类的修改；
    在当前类（比如User）中查找定义的类的所有属性，如果找到一个Field属性，就把它保存到一个__mappings__的dict中，同时从类属性中删除该Field属性，否则，容易造成运行时错误；
    把表名保存到__table__中，这里简化为表名默认为类名。
"""
class ModelMetaclass(type):
    
    def __new__(cls,name,bases,attrs):
        if name == 'Model':
            return type.__new__(cls,name,bases,attrs)
        mappings = dict()
        for k, v in attrs.iteritems():
            if isinstance(v,Field):
                #print 'Found mapping: %s ==> %s' % (k,v)
                mappings[k] = v
        for k in mappings.iterkeys():
            attrs.pop(k)
        attrs['__table__'] = name # 假设表明和类名一致
        attrs['__mappings__'] = mappings   # 保存属性和列的映射关系
        return type.__new__(cls,name,bases,attrs)
        
        
class Model(dict):

    __metaclass__ = ModelMetaclass
    
    def __init__(self,**kw):
        super(Model,self).__init__(**kw)
        
    def __getattr__(self,key):
        try :
            return self[key]
        except KeyError:
            raise AttributeError(r"'Model' object has no attribute '%s'" % key)
            
    def __setattr__(self,key,value):
        self[key] = value
        
    def save(self):
        fields = []
        params = []
        args = []
        for k, v in self.__mappings__.iteritems():
            fields.append(v.name)
            params.append('?')
            args.append(getattr(self, k, None))
        sql = 'insert into %s (%s) values (%s)' % (self.__table__, ','.join(fields), ','.join(params))
        print('SQL: %s' % sql)
        print('ARGS: %s' % str(args))
    
        
        
class User(Model):
    
    id = IntegerField('id')
    name = StringField('username')
    email = StringField('email')
    password = StringField('password')
    
    
if __name__ == '__main__':
    
    u = User(id=123,name='Mich',email='adsf@fda.com',password='yuikmnbhj')
    u.save()

