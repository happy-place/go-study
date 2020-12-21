package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main(){
	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc) // YWJjMTIzIT8kKiYoKSctPUB+

	sDec,_ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec)) // abc123!?$*&()'-=@~

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc) // YWJjMTIzIT8kKiYoKSctPUB-

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec)) // abc123!?$*&()'-=@~
}


