package srt

import (
	"path"
	"strings"
	"os"
	"bufio"
	"io"
	"fmt"
	"strconv"
	"GoogleTranslate/translate"
	"log"
	"time"
	"errors"
)

// 重写srt文件
func RewriteSrtFile(name string, resArr map[int]string) error {
	fileName := path.Base(name)
	createFileName := path.Join(path.Dir(name), strings.Replace(fileName, ".srt", ".ch.srt", -1))

	preLine := 0
	lineNum := 0
	file, err := os.Open(name)
	createFile, e := os.Create(createFileName)
	if err != nil {
		return err
	}
	if e != nil {
		return e
	}
	defer file.Close()
	defer createFile.Close()

	br := bufio.NewReader(file)
	for {
		line, isPrefix, err1 := br.ReadLine()
		lineNum += 1

		if err1 != nil {
			if err1 == io.EOF {
				break // OK
			}
			err = err1
			return err
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return err
		}

		// 转换字符数组为字符串
		str := string(line)

		// 取模2的行数
		if (lineNum % 4) == 1 {
			preLine, _ = strconv.Atoi(str)
		}

		// 取模3的英文字母
		io.Copy(createFile, strings.NewReader(str+"\r\n"))
		if (lineNum % 4) == 3 {
			io.Copy(createFile, strings.NewReader(resArr[preLine]+"\r\n"))
		}
	}

	return nil
}

// 读取srt文件中的文字
func GetSrtLines(name string) (map[int]string, error) {
	var lineArr = make(map[int]string)

	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lineNum := 0
	tmpStr := ""
	tmpArr := []int{}
	preLine := 0

	br := bufio.NewReader(file)
	for {
		line, isPrefix, err1 := br.ReadLine()
		lineNum += 1

		if err1 != nil {
			if err1 == io.EOF {
				break // OK
			}
			err = err1
			return nil, err
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return nil, err
		}

		// 转换字符数组为字符串
		str := string(line)

		// 取模2的行数
		if (lineNum % 4) == 1 {
			preLine, _ = strconv.Atoi(str)
		}

		// 取模3的英文字母
		if (lineNum % 4) == 3 {
			// 判断字符串最后是否有.
			ext := str[len(str)-1:]
			if ext == "." {
				// 加入map
				totalStr := tmpStr + str
				if len(tmpArr) != 0 {
					for _, v := range tmpArr {
						lineArr[v] = totalStr
					}
				}

				lineArr[preLine] = totalStr
				// 清除tmp数据
				tmpArr = []int{}
				tmpStr = ""
			} else {
				tmpArr = append(tmpArr, preLine)
				tmpStr += str
			}
		}
	}
	return lineArr, nil
}

func TranslateSrtFunc(srtName string) {
	// 读取srt文件中的文字
	fmt.Println("Start Translate: ", time.Now())

	err := TranslateSrt(srtName)
	if err != nil {
		fmt.Println("Translate Error: ")
		log.Fatal(err)
	} else {
		fmt.Println("Translate Done: ", time.Now())
	}
}

func TranslateSrt(srtName string) error {
	if path.Ext(srtName) != ".srt" {
		log.Fatal(errors.New("字母文件后缀必须为.srt"))
	}
	fmt.Println("File Name: ", srtName)
	lineArr, err := GetSrtLines(srtName)
	if err != nil {
		log.Fatal(err)
	}
	num := (len(lineArr) / 10) + 1
	var ch = make(chan map[int]string, num)

	for i := 1; i <= num; i++ {
		start := (i-1)*10 + 1
		end := i * 10
		if end > len(lineArr) {
			end = len(lineArr)
		}
		go translate.GoGetTranslate(lineArr, start, end, ch)
	}

	var resArr = make(map[int]string)
	for i := 0; i < num; i++ {
		select {
		case res := <-ch:
			for k, v := range res {
				resArr[k] = v
			}
		}
	}

	err = RewriteSrtFile(srtName, resArr)
	return err
}
