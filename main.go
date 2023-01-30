package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//io read from ./sum.txt

func main() {
	//io read from ./sum.txt
	res, err := readTxt(*FilePath)
	if err != nil {
		fmt.Println("readTxt err=", err)
		return
	}
	//sum
	var sum float64
	for _, v := range res {
		sum += v
	}
	fmt.Printf("sum=%v\n", sum)
}

func readTxt(filePath string) ([]float64, error) {
	var res []float64
	//read from ./sum.txt
	bts, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("filePath=[%v]"+"os.ReadFile err=[%v]\n", filePath, err)
		return nil, err
	}
	rd := bytes.NewReader(bts)
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		lineStr := string(line)
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		lineStr = strings.TrimPrefix(lineStr, "+")
		//string to float64
		var f float64
		f, err = strconv.ParseFloat(lineStr, 64)
		if err != nil {
			fmt.Printf("line=[%s], strconv.ParseFloat err=%v\n", line, err)
			return nil, err
		}
		fmt.Println("line-f=", f)
		res = append(res, f)
	}
	fmt.Println("结束read")
	return res, nil
}
