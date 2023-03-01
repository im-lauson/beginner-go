package main

import "fmt"

func main() {
	//数组的声明[容量]type{元素}
	a1 := [3]int{7, 8, 9} //声明一个数组 [容量]type{元素}
	fmt.Println(len(a1), cap(a1))
	//切片[]type{元素}
	//切片的声明[]type,与数组的区别，切片不需要声明容量，可以理解为切了一大片 给他自己用去分配使用
	array := [3]int{1, 2, 3} //数组的声明
	slice := []int{4, 4, 5}  //切片的声明:append(追加元素)、len获取元素数量、cap获取切片容量
	fmt.Println(array, slice)
	//推荐写法(追加)是：=make([]type,初始化长度0,预估一个容量)
	//make([]type,len,cap)
	s1 := make([]int, 0, 4)
	s1 = append(s1, 1, 2, 3, 4, 5)
	s1 = append(s1, []int{1, 2, 3, 4, 5, 6, 7, 8}...)
	fmt.Println(s1)
	//获取子切片  array\slice[star:end]
	fmt.Println(s1[0:9])
	//切片:最直观的对比：ArrayList ，即基于数组的List的实现，切片的底层也是数组

}
