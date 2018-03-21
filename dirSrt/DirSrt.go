package dirSrt

import (
	"fmt"
	"time"
	"path/filepath"
	"os"
	"log"
	"path"
	"sync"
	"GoogleTranslate/srt"
)

// 翻译路径下英文字母的主方法
func TranslateDirEnToCh(dir string) {
	fmt.Println("Start Translate: ", time.Now())
	fmt.Println("Dir: ", dir)
	pathArr := walkDir(dir)

	//fmt.Println(len(pathArr),pathArr)
	var wg sync.WaitGroup
	wg.Add(len(pathArr))

	for _, u := range pathArr {
		go func(u string) {
			srt.TranslateSrt(u)
			wg.Done()
		}(u)
	}
	wg.Wait()
	fmt.Println("Translate Done: ", time.Now())
	fmt.Println("Translate File Num: ", len(pathArr))
}

//遍历文件夹
func walkDir(DirName string) (pathArr []string) {
	pathArr = []string{}
	filepath.Walk(DirName, func(path string, info os.FileInfo, err error) error {
		if _, err := os.Stat(DirName); err != nil {
			log.Fatal("文件夹不存在")
		}
		if getExt(path) == ".srt" {
			pathArr = append(pathArr, path)
		}
		return nil
	})
	return

}

//获取文件拓展名
func getExt(fileName string) string {
	return path.Ext(fileName)
}
