package main

func main() {
}

type User struct {
	Name string
	Age  int
}

//指针接收器

func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}
