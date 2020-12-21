package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func absolutePath(name string) string {
	baseDir,_ := os.Getwd()
	filePath := fmt.Sprintf("%s/src/go_study/c33_file/resources/%s",baseDir,name)
	return filePath
}

func openFile(name string) *os.File{
	filePath := absolutePath(name)
	file,err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return file
}

func closeFile(file *os.File){
	if file != nil {
		err := file.Close()
		if err != nil {
			panic(err)
		}else{
			fmt.Printf("%s关闭成功\n",file.Name())
		}
	}
}

// 带缓冲读取
func bufferRead(file *os.File){
	bufferReader := bufio.NewReader(file)
	for {
		readString, err := bufferReader.ReadString('\n') // 读到换行符就结束
		if err == io.EOF { // 读取都文末就退出
			break
		}
		fmt.Printf(string(readString))
	}
}

func createAndClose(){
	var file *os.File
	file = openFile("1.txt")
	defer closeFile(file)
	bufferRead(file)
}

/*
	一次性读取整个文件到内存，然后作为字节数组整体输出，适合读取小文件
	由于没有显式打开文件，因此也无需关闭文件，文件打开与关闭函数自己解决
 */
func readAll(){
	filePath := absolutePath("1.txt")
	file, err := ioutil.ReadFile(filePath)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(file))
}

func writeString(){
	filePath := absolutePath("2.txt")
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666) // 只读 + 创建
	defer closeFile(file)
	if err != nil {
		panic(err)
	}
	for i:=0;i<5;i++{
		file.WriteString("hello, world\n")
	}
}

func bufferWriter(){
	filePath := absolutePath("3.txt")
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666) // 只读 + 创建
	defer closeFile(file)
	if err != nil {
		panic(err)
	}
	bufferWriter := bufio.NewWriter(file)
	for i:=0;i<5;i++{
		bufferWriter.WriteString("hello, world\n")
	}
	bufferWriter.Flush()
}

func truncWriter(){
	filePath := absolutePath("3.txt")
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0666) // 只读 + 覆盖
	defer closeFile(file)
	if err != nil {
		panic(err)
	}
	bufferWriter := bufio.NewWriter(file)
	for i:=0;i<3;i++{
		bufferWriter.WriteString("hello, hi\n")
	}
	bufferWriter.Flush()
}

func appendWriter(){
	filePath := absolutePath("3.txt")
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666) // 只读 + 追加
	defer closeFile(file)
	if err != nil {
		panic(err)
	}
	bufferWriter := bufio.NewWriter(file)
	for i:=0;i<3;i++{
		bufferWriter.WriteString("goodby 吃饭去\n")
	}
	bufferWriter.Flush()
}

func readAppend(){
	filePath := absolutePath("3.txt")
	file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666) // 读写 + 追加
	defer closeFile(file)
	if err != nil {
		panic(err)
	}
	// 读取
	bufferReader := bufio.NewReader(file)
	for{
		readString, err := bufferReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf(readString)
	}

	// 追加
	bufferWriter := bufio.NewWriter(file)
	for i:=0;i<3;i++{
		bufferWriter.WriteString("ohh\n")
	}
	bufferWriter.Flush()
}

func copyFile(){
	sourcePath := absolutePath("3.txt")
	sourceFile, err := os.OpenFile(sourcePath, os.O_RDONLY, 0666)
	defer closeFile(sourceFile)
	if err != nil {
		panic(err)
	}
	bufReader := bufio.NewReader(sourceFile)

	destPath := absolutePath("4.txt")
	destFile, err := os.OpenFile(destPath, os.O_WRONLY | os.O_CREATE, 0666)
	defer closeFile(destFile)
	if err != nil {
		panic(err)
	}

	bufWriter := bufio.NewWriter(destFile)

	for{
		readString, err := bufReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		bufWriter.WriteString(readString)
	}
	bufWriter.Flush()

}

func fileStat(name string){
	path := absolutePath(name)
	fileInfo, err := os.Stat(path)
	if err == nil { // 文件存在
		fmt.Printf("filename: %v \n",fileInfo.Name())
		fmt.Printf("file is dir: %v \n",fileInfo.IsDir())
		fmt.Printf("file mode: %v \n",fileInfo.Mode())
		fmt.Printf("file mod time: %v \n",fileInfo.ModTime())
		fmt.Printf("file size: %v \n",fileInfo.Size())
		fmt.Printf("file sys: %v \n",fileInfo.Sys())
	}else{
		fmt.Printf("%s not exist\n",path)
	}

	/**
	filename: 4.txt
	file is dir: false
	file mode: -rw-r--r--
	file mod time: 2020-11-20 08:53:15.2266301 +0800 CST
	file size: 114
	file sys: &{16777220 33188 1 51545669 501 20 0 [0 0 0 0] {1605833596 503792601} {1605833595 226630100} {1605833595 226630100} {1605833595 226530812} 114 8 4096 0 0 0 [0 0]}

	*/
}

func fileExist(name string) (bool,error){
	path := absolutePath(name)
	_, err := os.Stat(path)
	if err == nil { // 文件存在
		return true,nil
	}
	if os.IsNotExist(err) { // 文件不存在
		return false,err
	}
	return false,err // 其他未知异常
}

func testExist(){
	exist, err := fileExist("4.txt")
	if err != nil {
		fmt.Printf("4.txt exist %v\n",exist)
	}else{
		fmt.Printf("4.txt exist %v\n",exist)
	}
}

func copy(source string,dest string)(int64,error){
	sourcePath := absolutePath(source)

	exist, err := fileExist(sourcePath)
	if !exist{
		panic(err)
	}

	sourceFile,err := os.OpenFile(sourcePath,os.O_RDONLY,0666)
	defer closeFile(sourceFile)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(sourceFile)

	destPath := absolutePath(dest)
	destFile,err := os.OpenFile(destPath,os.O_WRONLY | os.O_CREATE,0666)
	defer closeFile(destFile)
	writer := bufio.NewWriter(destFile)
	written, err := io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("写入字节数：%v\n",written)
	return written,err
}

type CharCount struct {
	charCount int
	intCount int
	spaceCount int
	delimiterCount int
	otherCount int
}

func countChar(){
	counter := CharCount{}
	filePath := absolutePath("5.txt")
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	defer closeFile(file)
	if err != nil {
		fmt.Printf("%s 打开异常(%v)",filePath,err)
	}
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		//for _,ch := range readString {
		//	//fmt.Println(string(ch))
		//	switch {
		//	case ch >= 'a' && ch <= 'z':
		//		fallthrough
		//	case ch >= 'A' && ch <= 'Z':
		//		counter.charCount ++
		//	case ch >= '0' && ch <= '9':
		//		counter.intCount ++
		//	case ch == ' ' || ch == '\t':
		//		counter.spaceCount ++
		//	case ch == '\n':
		//		counter.delimiterCount ++
		//	default:
		//		counter.otherCount ++
		//	}
		//}
		rString := []rune(readString) // 中文存在必须转换称为 []rune() 然后使用经典for循环遍历
		for i:=0;i<len(rString);i++{
			ch := rString[i]
			switch {
			case ch >= 'a' && ch <= 'z':
				fallthrough
			case ch >= 'A' && ch <= 'Z':
				counter.charCount ++
			case ch >= '0' && ch <= '9':
				counter.intCount ++
			case ch == ' ' || ch == '\t':
				counter.spaceCount ++
			case ch == '\n':
				counter.delimiterCount ++
			default:
				counter.otherCount ++
			}
		}

		if err == io.EOF {
			break
		}
	}
	fmt.Printf("%+v\n",counter)
}

func main() {
	//createAndClose()
	//readAll()
	//writeString()
	//bufferWriter()
	//truncWriter()
	//appendWriter()
	//readAppend()
	//copyFile()
	//fileStat("4.txt")
	//testExist()
	//copy("4.txt","5.txt")
	countChar()
}
