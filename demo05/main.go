package main

import "fmt"

func main() {
	forR()
	fruit()
}
func forR() {
	arr := []int{1, 2, 3, 4}
	arr2 := []int{1, 2, 3, 4}
	if arr == nil {
		fmt.Println("false")
	} else if arr2 == nil {
		fmt.Println("false")
	} else if len(arr2) == len(arr) {
		fmt.Println("个数相等")
	} else {
		fmt.Println("true")
	}
	//for i, i2 := range arr { //i为下标，i2为数据 for rang不要忘记
	//	fmt.Println(i, i2)
	//}
	//for _, i := range arr { //i为下标，i2为数据,可以省略，根本不需要括号
	//	fmt.Println(i)
	//}

}
func fruit() {
	fruit := "蓝莓"
	switch fruit {
	case "苹果":
		fmt.Println("答对了")
	case "草莓", "蓝莓":
		fmt.Println("这个是梅梅")
	default:
		fmt.Println("牛逼")
	}
}
