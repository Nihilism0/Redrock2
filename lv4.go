package main

import "fmt"

type Interface interface { //题目给定的
	Len() int           //求长度
	Less(i, j int) bool //元素的index
	Swap(i, j int)      //交换元素
}
type STR []int

//求切片长度

func (l STR) Len() int {
	return len(l)
}

//如果index为j的数小于i 则返回true

func (l STR) Less(i, j int) bool {
	return l[i] > l[j]
}

//交换!

func (l STR) Swap(i, j int) {
	temp := l[i]
	l[i] = l[j]
	l[j] = temp
}

//将满足于这个接口的类型赋予给这个接口 对接口类型的变量进行操作 返回一个接口类型的值(一个切片)

func DoIt(wow Interface) Interface {
	lens := wow.Len()
	for i := 0; i < lens-1; i++ {
		for j := i; j < lens; j++ {
			ret := wow.Less(i, j)
			if ret {
				wow.Swap(i, j)
			}
		}
	}
	return wow
}

func main() {
	var wow STR
	wow = []int{9, 6, -8, 5, 2, 1, 11, 0, 54, 2, 1, 34, 1, 9}
	result := DoIt(wow)
	fmt.Println(result)
}
