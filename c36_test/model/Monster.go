package model

import (
	"fmt"
	_ "io"
	"bufio"
	"os"
	"encoding/json"
	"log"
	"io/ioutil"
)
type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func absloutePath(name string) string{
	wd, _ := os.Getwd()
	return fmt.Sprintf("%s/../resources/%s",wd,name)

}

// 直接进行流对拷
func (m *Monster)Store(name string) error {
	mBytes, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Marshal object error: %s\n",err)
		return err
	}

	path := absloutePath(name)
	file, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("open file error: %s\n",err)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.Write(mBytes) // 直接将[]byte，以string 方式落盘
	if err != nil {
		log.Fatalf("write file error: %s\n",err)
		return err
	}
	writer.Flush() // bufio.NewWrite 必须使用 Flush()
	return nil
}

func (m *Monster)ReadStore(name string) (*Monster,error){
	path := absloutePath(name)
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("open file error: %s\n",err)
		return nil,err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	mBytes, _, err := reader.ReadLine() // 直接将string读取为 []byte
	if err != nil {
		log.Fatalf(" read string error: %s\n",err)
		return nil,err
	}

	var mm Monster
	err = json.Unmarshal(mBytes, &mm)
	if err != nil {
		log.Fatalf("Unmarshal string error: %s\n",err)
		return nil,err
	}
	//log.Println(mm)
	return &mm,nil
}

// 使用 ioutil 实现流对拷
func (m *Monster)Store2(name string) error{
	mBytes, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("json.Marshal(m) error：%v",err)
		return err
	}

	path := absloutePath(name)
	err = ioutil.WriteFile(path, mBytes, 0666)
	if err != nil {
		log.Fatalf("ioutil.WriteFile error：%v",err)
		return err
	}
	return nil
}

func (m *Monster)ReadStore2(name string) (*Monster,error){
	path := absloutePath(name)
	mBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("ioutil.ReadFile error：%v",err)
		return nil,err
	}
	var monster *Monster
	json.Unmarshal(mBytes,monster)
	return monster,nil
}