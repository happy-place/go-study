package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main(){
	match,_ := regexp.MatchString("p([a-z]+)ch","peach")
	fmt.Println(match) // true 是否具有匹配项

	r,_ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach")) // true 是否具有匹配项

	fmt.Println(r.FindString("peach punch")) // peach 从左往右找一个匹配项
	fmt.Println(r.FindStringIndex("peach punch")) // [0 5] 第一个匹配项的起止下标
	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea] 返回第一个匹配项的单词，以及匹配字符串

	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3] 返回第一个匹配项的单词，以及匹配字符串的起止下标
	fmt.Println(r.FindAllString("peach punch pinch",-1)) // [peach punch pinch] 找出全部匹配项

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch punch",-1)) // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]，找出全部匹配项，以及匹配单词和匹配字符串起止下标
	fmt.Println(r.FindAllString("peach punch punch",2)) // [peach punch] 找出头两个匹配项
	fmt.Println(r.Match([]byte("peach"))) // 字节数组匹配

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r) // p([a-z]+)ch

	fmt.Println(r.ReplaceAllString("a peach","<fruit>")) // a <fruit> 替换
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in,bytes.ToUpper)
	fmt.Println(string(out)) // a PEACH 对匹配项使用目标函数

}
