package main

import (
	"fmt"
)

/**
 * 稀疏数组  （一种压缩方式）

实际需求： 五子棋程序，存盘退出和续上盘残局的功能

默认为0 黑子为1 白子为2
使用二位数组记录棋盘
但默认值都是0 记录了太多没有意义的数据
解决办法就是稀疏数组 只记录有意义的数据：
	记录数组几行几列 有多少个不同的值
	把不同值的元素行列记录在小规模的数组中从而缩小程序的规模


// 以值为基准
type Point struct{
	x,y int 坐标
}

type Values struct{
	value int // 值
	point []Point
}

type sparseArray struct{
	x,y int 几行几列
	defaultValue int 默认值
	valueList []Values
}

以坐标为基准
type Point struct{
	x,y int  坐标
	value int 值
}

type sparseArray struct {
	x int 几行
	y int  几列
	defaultValue int 默认值
	point []Point
}

*/

type Point struct {
	PointX, PointY int
	Value          int
}

type SparseArr struct {
	X, Y     int
	DefaultV int
	Point    []Point
}

func (sArray *SparseArr) save(array *[][]int) {
	// 不能range 指针
	for r, row := range *array {
		for p, point := range row {
			if point != sArray.DefaultV {
				sArray.Point = append(sArray.Point, Point{
					PointX: r, PointY: p, Value: point,
				})
			}
		}
	}
}

// 只能常量来定义数组长度
func (sArray *SparseArr) restore() *[][]int {
	var (
		array        [][]int
		next         = 0
		count        = len(sArray.Point)
		row          []int
		currentPoint Point
	)
	for i := 0; i < sArray.X; i++ {
		row = []int{}
		for j := 0; j < sArray.Y; j++ {
			if next < count {
				currentPoint = sArray.Point[next]
				if i == currentPoint.PointX && j == currentPoint.PointY {
					row = append(row, currentPoint.Value)
					next++
				} else {
					row = append(row, sArray.DefaultV)
				}
			} else {
				row = append(row, sArray.DefaultV)
			}
		}
		array = append(array, row)
	}
	return &array
}

/**
* 使用稀疏数组 存盘 棋盘或地图等
  复盘

@todo 保存内容并恢复数据结构
*/
func SparseArray() {

	var (
		array  [][]int
		sArray SparseArr
	)

	array = [][]int{
		{0, 0, 0, 22, 0, 0, 15},
		{0, 11, 0, 0, 0, 17, 0},
		{0, 0, 0, -6, 0, 0, 0},
		{0, 0, 0, 0, 0, 39, 0},
		{91, 0, 0, 0, 0, 0, 0},
		{0, 0, 28, 0, 0, 0, 0},
	}
	sArray = SparseArr{
		DefaultV: 0,
		X:        len(array),
		Y:        len(array[0]),
	}
	sArray.save(&array)
	// [{0 3 22} {0 6 15} {1 1 11} {1 5 17} {2 3 -6} {3 5 39} {4 0 91} {5 2 28}]
	fmt.Println(sArray.Point)

	arrayP := sArray.restore()
	// [[0 0 0 22 0 0 15] [0 11 0 0 0 17 0] [0 0 0 -6 0 0 0] [0 0 0 0 0 39 0] [91 0 0 0 0 0 0] [0 0 28 0 0 0 0]]
	fmt.Println(*arrayP)

}
