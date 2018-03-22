package main

import (
	"GoogleTranslate/srt"
	"flag"
	"fmt"
	str2 "GoogleTranslate/str"
	"GoogleTranslate/dirSrt"
)

// string 类型的flag
//参数 1.相当于-s 名称 2. 默认值为空 3. --help 展示的内容 说明文档
// 返回的类型是字符串参数
var sep = flag.String("srt", "", "翻译一个后缀为srt的文件,需输入文件名(文件名中空格前加\\或用双引号引出)")
var dir = flag.String("d", "", "输入一个文件路径,翻译其下面的所有srt格式文件")
var strEn = flag.String("e", "", "翻译一段英文,需输入字符串(字符串中空格前加\\或用双引号引出)")
var strCh = flag.String("c", "", "翻译一段中文,需输入字符串(字符串中空格前加\\或用双引号引出)")

func main() {
	flag.Parse()
	switch {
	case *sep != "":
		srt.TranslateSrtFunc(*sep)
	case *strEn != "":
		if len(flag.Args()) != 0 {
			fmt.Println("翻译一段英文,需输入字符串(字符串中空格前加\\或用双引号引出)")
		} else {
			str2.TranslateStrEnToChFunc(*strEn)
		}
	case *strCh != "":
		if len(flag.Args()) != 0 {
			fmt.Println("翻译一段中文,需输入字符串(字符串中空格前加\\或用双引号引出)")
		} else {
			str2.TranslateStrChToEnFunc(*strCh)
		}
	case *dir != "":
		if len(flag.Args()) != 0 {
			fmt.Println("输入一个文件路径,翻译其下面的所有srt格式文件")
		} else {
			dirSrt.TranslateDirEnToCh(*dir)
		}
	default:
		fmt.Println("请输入正确的指令. -srt/-d/-s")
	}
}
