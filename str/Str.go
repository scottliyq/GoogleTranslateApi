package str

import (
	"fmt"
	"time"
	"GoogleTranslate/translate"
	"log"
)

// 翻译输入的英文到中文
// 输出英文,中文
func TranslateStrEnToChFunc(en string) {
	fmt.Println("Start Translate: ", time.Now())
	fmt.Println("English: ", en)

	ch, err := translate.GetTranslateEnToChContent(en)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Chinese:", ch)
	fmt.Println("Translate Done: ", time.Now())
}

// 翻译中文到英文
func TranslateStrChToEnFunc(ch string) {
	fmt.Println("Start Translate: ", time.Now())
	fmt.Println("Chinese: ", ch)

	en, err := translate.GetTranslateChToEnContent(ch)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("English:", en)
	fmt.Println("Translate Done: ", time.Now())
}
