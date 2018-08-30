package utils

import (
	"bufio"
	"fmt"
	"os"
)

func BufIOToFile(filePath string)  {
	fmt.Println("--------------缓冲区数据写入")
	//创建文件(忽略错误处理过程)
	file, _ := os.Create(filePath)
	defer file.Close()

	//写入2个字符串到缓存区
	w := bufio.NewWriter(file)
	n1, _ := w.WriteString("CloudGeek!")
	n2, _ := w.WriteString("Hello，CloudGeek!")

	//将缓存中内容写到文件
	w.Flush()
	fmt.Println("--------------第一次写入长度", n1)
	fmt.Println("--------------第二次写入长度", n2)
}