package GoogleTranslate

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"errors"
	"io"
	"os"
	"fmt"
	"github.com/robertkrimen/otto"
	"reflect"
	"log"
	"GoogleTranslate/translate"
	"GoogleTranslate/token"
)

func main() {
	s := "Cool."
	//s = "all right priorities come into focus."
	//_, err := getTranslateTokenTest(s)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(str)

	fmt.Println(s)
	//s = strings.Replace(s, " ", "%20", -1)

	key := token.GetGoogleToken(s)
	fmt.Println(key)

	res, err := translate.TranslateEnToChApi(s, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}

func getGoogleTokenTest(str string) otto.Value {
	vm := otto.New()

	vm.Set("a", str)
	vm.Run(`
    function VL(a) {
        var b = a.trim();
        var res = TL(b);
        return res;
    }

    function TL(a) {
        var k = "";
        var b = 406644;
        var b1 = 3293161072;

        var jd = ".";
        var $b = "+-a^+6";
        var Zb = "+-3^+b+-f";

        for (var e = [], f = 0, g = 0; g < a.length; g++) {
            var m = a.charCodeAt(g);
            128 > m ? e[f++] = m : (2048 > m ? e[f++] = m >> 6 | 192 : (55296 == (m & 64512) && g + 1 < a.length && 56320 == (a.charCodeAt(g + 1) & 64512) ? (m = 65536 + ((m & 1023) << 10) + (a.charCodeAt(++g) & 1023),
                e[f++] = m >> 18 | 240,
                e[f++] = m >> 12 & 63 | 128) : e[f++] = m >> 12 | 224,
                e[f++] = m >> 6 & 63 | 128),
                e[f++] = m & 63 | 128)
        }
        a = b;
        for (f = 0; f < e.length; f++) a += e[f],
            a = RL(a, $b);
        a = RL(a, Zb);
        a ^= b1 || 0;
        0 > a && (a = (a & 2147483647) + 2147483648);
        a %= 1E6;
        return a.toString() + jd + (a ^ b)
    };

    function RL(a, b) {
        var t = "a";
        var Yb = "+";
        for (var c = 0; c < b.length - 2; c += 3) {
            var d = b.charAt(c + 2),
                d = d >= t ? d.charCodeAt(0) - 87 : Number(d),
                d = b.charAt(c + 1) == Yb ? a >>> d : a << d;
            a = b.charAt(c) == Yb ? a + d & 4294967295 : a ^ d
        }
        return a
    }
`)

	value, _ := vm.Run("VL(a)")
	{
		// value is an int64 with a value of 16
		value, _ := value.ToFloat()
		fmt.Println(value)
	}
	fmt.Println(reflect.TypeOf(value))
	return value
}

// 抓取谷歌翻译的key
func getTranslateTokenTest(str string) (string, error) {
	token := ""
	BaseUrl := "http://localhost:8080/google/token.html?str="
	url := BaseUrl + str
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	// 关闭链接
	defer resp.Body.Close()
	// 如果状态码不等于200
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status) // 例如: 404 Not Found
	}
	// 爬取(打印)body
	io.Copy(os.Stdout, resp.Body)
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return "", err
	}
	token = doc.Find("#result").Text()
	return token, nil
}
