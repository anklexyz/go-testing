package file_process

import (
	"bufio"
	"os"
	"strings"
)

func ReadFirstLine() string {
	//Open打开指定文件进行读取。如果成功，就可以使用返回文件的方法进行读取。
	//如果有错误，它的类型将是*PathError。
	open, err := os.Open("log")
	defer open.Close()
	if err != nil {
		return ""
	}
	//NewScanner扫描器读取输入，这里输入源是open
	scanner := bufio.NewScanner(open)
	//Scan() 默认对每行进行读取，可利用for语句进行循环遍历文件
	//Text() 将scanner读取的内容变为string
	for scanner.Scan() {
		return scanner.Text() //第一行读到就返回
	}
	return ""
}

func ProcessFirstLine() string {
	line := ReadFirstLine()
	destLine := strings.ReplaceAll(line, "11", "00")
	return destLine
}
