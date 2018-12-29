#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
建造者模式
参考：  https://xionchen.github.io/2017/07/01/python-patterns_03/

通过建造者 统一创建 ， 屏蔽复杂对象的创建细节

创建者 通过符合 《创建步骤标准(抽象类)的建造细节类》 创建出实例  返还给实例需求者

建筑公司(创建者)  选择 别墅/写字楼(建造细节类) 建筑团队 根据 步骤标准(抽象类) 创建 房子(实例) 然后返回给客户

"""

import  ABC

class company(object):
    """
    建筑公司
    我知道怎样去建造一个房子
    """

    def __init__(self):
        self.builder = None

    def construct_building(self):
        self.builder.new_building()
        self.builder.build_floor()
        self.builder.build_size()

    def get_building(self):
        return self.builder.building

class builder(object):
    """
    创建步骤标准
    """
    __metaclass__ = ABCMeta

    def __init__(self):
        self.building = None

    def new_building(self):
        # 开始创建一个新的房子
        self.building = Building()

    @ABC.abstractmethod
    def build_floor(self):
        raise NotImplementedError

    @ABC.abstractmethod
    def build_size(self):
        raise NotImplementedError

class BuilderOffice(builder):
    """
    写字楼
    """
    def build_floor(self):
        self.building.floor = ''
    def build_size(self):
        self.building.size = ''

class BuilderVilla(builder):
    """
    别墅
    """
    def build_floor(self):
        self.building.floor = ''
    def build_size(self):
        self.building.size = ''


class Building(object):
    """
    房子
    """
    def __init__(self):
        self.floor = None
        self.size = None

    def __repr__(self):
        return 'Floor: {0.floor} | Size: {0.size}'.format(self)


