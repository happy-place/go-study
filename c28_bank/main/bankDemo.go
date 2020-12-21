package main

import (
	"fmt"
	m "go_study/c28_bank/model"
)

func front(){
	acc := &m.Account{"8888","6868",100.23}
	toStop := false
	for;!toStop;{
		var opt int
		var pwd string
		var money float64
		fmt.Println("请选择操作:\n\t1)查询\n\t2)存款\n\t3)取款\n\t4)退出")
		fmt.Scanf("%v",&opt)
		switch opt {
		case 1:
			fmt.Printf("请输入密码：")
			fmt.Scanf("%s",&pwd)
			acc.Query(pwd)
		case 2:
			fmt.Printf("请输入密码和存款金额（格式：密码 金额）:")
			fmt.Scanf("%s %f",&pwd,&money)
			acc.Deposite(money,pwd)
		case 3:
			fmt.Printf("请输入密码和取款金额（格式：密码 金额）:")
			fmt.Scanf("%s %f",&pwd,&money)
			acc.WithDraw(money,pwd)
		case 4:
			fmt.Println("退出~")
			toStop = true
		}
	}
}


func main(){

	front()

}
