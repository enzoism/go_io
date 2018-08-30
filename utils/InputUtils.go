package utils

import "fmt"

func InputString()  {
	var username string
	fmt.Print("-------------------请用户输入当前用户名：")
	fmt.Scanln(&username)
	fmt.Println("username is:", username)
}