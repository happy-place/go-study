package main

import "fmt"
import ss "strings" // 别名

var p = fmt.Println // 提取函数

func main(){
	p("Contains:", ss.Contains("test","es"))
	p("Count:", ss.Count("test","t"))
	p("HasPrefix:", ss.HasPrefix("test","te"))
	p("HasSuffix:", ss.HasSuffix("test","st"))
	p("Index:", ss.Index("test","e"))
	p("Join:", ss.Join([]string{"a","b","c"},"-"))
	p("Repeat:", ss.Repeat("a",5))
	p("Replace-All:", ss.Replace("foo","o","0",-1))
	p("Replace:", ss.Replace("foo","o","0",1))
	p("Split:", ss.Split("a-b-c-d-e","-"))
	p("ToLower:", ss.ToLower("TEST"))
	p("ToUpper:", ss.ToUpper("test"))
	p()
	p("Len:",len("hello"))
	p("Char:","hello"[0])
}

/**
Contains: true
Count: 2
HasPrefix: true
HasSuffix: true
Index: 1
Join: a-b-c
Repeat: aaaaa
Replace-All: f00
Replace: f0o
Split: [a b c d e]
ToLower: test
ToUpper: TEST

Len: 5
Char: 104
*/