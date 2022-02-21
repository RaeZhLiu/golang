package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	//w 代表大小写字母+数字+下划线
	reEmail = `\w+@\w+\.\w+`
	//s? 有或者没有s
	//+ 代表匹配1次或多次
	// \s\S各种字符
	// +?代表贪婪模式
	reLink   = `href="(https?://[\s\S]+?)"`
	rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
	reIDCard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|1[012])((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	reImage  = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

//HandleError 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

//GetPageContent 根据传入的url，Get获取url的返回内容
func GetPageContent(url string) (pageStr string) {
	resp, err := http.Get(url)
	HandleError(err, "http.get Error!")
	defer func() { _ = resp.Close }()
	//2. 读取页面内容
	pageByte, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll Error!")
	//字节转字符串
	pageStr = string(pageByte)
	return
}

//GetEmail 从指定的URL中爬取邮箱信息
func GetEmail(url string) {
	pageStr := GetPageContent(url)
	re := regexp.MustCompile(reEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

//GetLink 爬取链接
func GetLink(url string) {
	pageStr := GetPageContent(url)
	re := regexp.MustCompile(reLink)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[1])
	}
}

//GetPhone 爬取手机号
func GetPhone(url string) {
	pageStr := GetPageContent(url)
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[0])
	}
}

//GetIDCard 获取身份证号
func GetIDCard(url string) {
	pageStr := GetPageContent(url)
	re := regexp.MustCompile(reIDCard)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result[0])
	}
}

//GetImg 爬取图片链接
func GetImg(url string) {
	pageStr := GetPageContent(url)
	fmt.Println(pageStr)
	re := regexp.MustCompile(reImage)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

func main() {
	//1. 爬邮箱
	//GetEmail("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	//2. 爬链接
	//GetLink("http://www.baidu.com/s?wd=%E8%B4%B4%E5%90%A7%20%E7%95%99%E4%B8%8B%E9%82%AE%E7%AE%B1&rsv_spt=1&rsv_iqid=0x98ace53400003985&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_dl=ib&rsv_sug2=0&inputT=5197&rsv_sug4=6345")
	//3. 爬手机号
	//GetPhone("https://www.zhaohaowang.com/")
	//4. 爬身份证号
	//GetIDCard("https://henan.qq.com/a/20171107/069413.htm")
	//5. 爬图片
	GetImg("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")

}
