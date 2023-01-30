package main

import (
	"flag"
	"fmt"
)

// 相对路径 求和文件相对go程序的路径 默认是./sum.txt
var FilePath = flag.String("path", "./sum.txt", "the path of sum file; for example: ./sum.txt")

func init() {
	fmt.Println("init()...")
	flag.Parse()
}
