package main

import (
	"fmt"
	"enzoism/go_io/utils"
)

func main() {
	fmt.Println("---------------------xiaoming")
	// 对象编码
	//utils.CodeObject()
	//utils.CodeObjectWithAddress("C:/AppData/a.txt")

	// 对象文件写入
	//utils.FileOperationAddress("C:/AppData/a2.txt")

	// 原生的IO写入
	//utils.IOWriteFile("C:/AppData/a2.txt")
	//utils.IOReadFile("C:/AppData/a2.txt")

	// 获取用户输入
	//utils.InputString()

	// Buffer缓存区IO
	//utils.BufIOToFile("C:/AppData/a2.txt")

	// 文件目录遍历
	//utils.FileMenuRead("C://TT_Install+")

	// 读取数据到[]byte中
	utils.ReadInByteArray("C:/AppData/a2.txt");

}
