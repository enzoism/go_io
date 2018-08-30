package utils

import (
	"bytes"
	"encoding/gob"
	"io"
	"fmt"
	"os"
	"io/ioutil"
)
//参考网址：https://mp.weixin.qq.com/s/uS24t5xnuPPPrCqVqLiTEg
//参考网址：https://mp.weixin.qq.com/s/wZRAbYNkUsI4qjkT0XXw5g
//准备编码的数据
type User struct {
	X, Y, Z int
	Name    string
}

//接收解码结果的结构
type UserReceive struct {
	X, Y *int32
	Name string
}

func encode(data interface{}) *bytes.Buffer {
	//Buffer类型实现了io.Writer接口
	var buf bytes.Buffer
	//得到编码器
	enc := gob.NewEncoder(&buf)
	//调用编码器的Encode方法来编码数据data
	enc.Encode(data)
	//编码后的结果放在buf中
	return &buf
}

func decode(data interface{}) *UserReceive {
	d := data.(io.Reader)
	//获取一个解码器，参数需要实现io.Reader接口
	dec := gob.NewDecoder(d)
	var receive UserReceive
	//调用解码器的Decode方法将数据解码，用Q类型的q的指针来接收
	dec.Decode(&receive)
	return &receive
}

// 序列化对象
func CodeObject() {
	fmt.Println("---------------序列化对象")
	//初始化一个数据
	data := User{3, 4, 5, "CloudGeek"}
	//编码后得到buf字节切片
	buf := encode(data)
	//用于接收解码数据
	var receiveUser *UserReceive
	//解码操作
	receiveUser = decode(buf)
	//"CloudGeek": {3,4}
	fmt.Printf("%q: {%d,%d}\n", receiveUser.Name, *receiveUser.X, *receiveUser.Y)
}

//******************************************序列化对象到文件*********************************************//
//试验用的数据类型
type Address struct {
	City    string
	Country string
}
//将数据序列号后写到文件中
func encodeToAddress(filePath string) {
	pa := &Address{"Chengdu", "China"}
	//打开文件，不存在的时候新建
	file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	//encode后写到这个文件中
	enc := gob.NewEncoder(file)
	enc.Encode(pa)
}

//从文件中读取数据并反序列化
func decodeFromAddress(filePath string) *Address {
	file, _ := os.Open(filePath)
	defer file.Close()
	var pa Address
	//decode操作
	dec := gob.NewDecoder(file)
	dec.Decode(&pa)
	return &pa
}

func CodeObjectWithAddress(filePath string) {
	fmt.Println("---------------序列化对象到文件(当前文件必须先存在)")
	encodeToAddress(filePath)
	pa := decodeFromAddress(filePath)
	fmt.Println(*pa) //{Chengdu China}
}
//******************************************从文件中读取数据*********************************************//

func fileRead(filePath string) {
	fmt.Println("---------------从文件中读取数据")
	//打开一个文件，文件里面存的是数字123
	f, _ := os.Open(filePath)
	defer f.Close()
	//用一个长度为5的byte切片来读取数据
	b := make([]byte, 5)
	//n也就是读取到的数据长度
	n, _ := f.Read(b)
	//输出内容是：3 [49 50 51 0 0] 123
	fmt.Println(n, b, string(ConvertByte2String(b,GB18030)))
}

func fileWrite(filePath string) {
	//创建文件
	fmt.Println("---------------从文件中存储数据")
	f, _ := os.Create(filePath)
	//待写入的数据
	b := []byte("CloudGeek")
	//执行写操作，n为写入的字符数
	n, _ := f.Write(b)
	//输出结果是：9
	fmt.Println("数据写入完毕，字节数：",n)
}
func FileOperationAddress(filePath string) {
	fileWrite(filePath)
	fileRead(filePath)
}

//******************************************远程的IOUtils*********************************************//
func IOReadFile(filePath string) {
	b, err := ioutil.ReadFile(filePath)
	if err == nil {
		fmt.Println(string(b))
	}
}
func IOWriteFile(filePath string) {
	//文件不存在会新建，权限通过perm指定，文件存在会被清空后再写入数据
	err := ioutil.WriteFile(filePath, []byte("CloudGeek"), 0666)
	if err != nil {
		fmt.Println(err)
	}
}

//******************************************遍历文件目录*********************************************//

func FileMenuRead(filePath string) {
	//目录遍历
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, file := range files {
			fmt.Println("文件名：", file.Name(), "文件大小：", file.Size())
		}
	}
}

//******************************************遍历文件目录*********************************************//
func ReadInByteArray(filePath string) {
	//file是os.File类型，实现了io.Reader接口
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	} else {
		//ReadAll方法接收io.Reader类型的参数，返回一个[]byte类型的结果
		b, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		} else {
			//[]byte类型的数据转换为string后输出
			fmt.Println(string(b))
		}
	}
}