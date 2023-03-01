package main

import "fmt"

func main() {
	//数组的声明[容量]type{元素}
	a1 := [3]int{7, 8, 9} //声明一个数组 [容量]type{元素}
	fmt.Println(len(a1), a1)
	//切片[]type{元素}
	//切片的声明[]type,与数组的区别，切片不需要声明容量，可以理解为切了一大片 给他自己用去分配使用
	array := [3]int{1, 2, 3} //数组的声明
	slice := []int{4, 4, 5}  //切片的声明
	fmt.Println(array, slice)

}
