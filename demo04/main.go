package main

import "fmt"

func main() {
	forR()
}
func forR() {
	arr := []int{1, 2, 3, 4}
	for i, i2 := range arr { //i为下标，i2为数据 for rang不要忘记
		fmt.Println(i, i2)
	}
	for _, i := range arr { //i为下标，i2为数据,可以省略，根本不需要括号
		fmt.Println(i)
	}

}
