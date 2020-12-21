package test

import (
	_ "fmt"
	"testing"
	"go_study/c36_test/model"
)


/**
	使用单元测试目标：逻辑校对、性能分析
	单元测试：
	1. 编写业务代码 model/Calculator.go，至于函数名大小；
	2. 编写单元测试类 test_case/calculator_test.go
		1）注意名称为 xx_test.go格式，xx最好与业务代码相关
		2）导包 import "testing"
		3) 编写测试函数 func TestGetSum(t *testing.T){...} ，函数名称必须是 TestXxx类型，即Test之后必须是大写字母
		4）T为结构体可以获取测试相关上下文环境，最主要的是记录日志相关
			type T struct {
				common
				isParallel bool
				context    *testContext // For running tests and subtests.
			}
		5）t.Fatalf(format,args) 记录异常信息，t.Logf(format,args)记录正常信息
	3. 测试日志相关 [-v]
		1) 正确 和 错误显示日志
			go test -v [xx_test.go yy_test.go]
			=== RUN   TestGetSum
				calculator_test.go:13: model.TestGetSum(1, 2) 期望值：3，实际值：-1，测试不通过
			--- FAIL: TestGetSum (0.00s)  << 第一个异常
			=== RUN   TestGetSub
				calculator_test.go:24: model.GetSub(1, 2)测试通过
			--- PASS: TestGetSub (0.00s)  << 第二个正常
			FAIL  << 全部通过才正常
			FAIL    command-line-arguments  0.395s
			FAIL
		2）只有错误显示日志
			go test
			--- FAIL: TestGetSum (0.00s)
				calculator_test.go:42: model.TestGetSum(1, 2) 期望值：3，实际值：-1，测试不通过
			FAIL
			exit status 1
			FAIL    go_study/c36_test/test_case     0.328s
	4. 测试全部 or 测试指定案例 or 测试指定指定函数
		1）测试目录下全部测试样例
			go test
		2）测试目录下指定测试样例
			go test xx_test.go
		3) 测试指定指定函数
			go test -v -test.run TestGetSum
			=== RUN   TestGetSum
				calculator_test.go:52: model.TestGetSum(1, 2) 期望值：3，实际值：-1，测试不通过
			--- FAIL: TestGetSum (0.00s)
			FAIL
			exit status 1
			FAIL    go_study/c36_test/test_case     0.581s



 */

func TestGetMuitl(t *testing.T){
	sub := model.GetMilti(1, 2)
	expected := 0
	if sub != expected {
		t.Fatalf("model.GetMilti(1, 2) 期望值：%v，实际值：%v，测试不通过",sub,expected)
	}
	t.Logf("model.GetMilti(1, 2)测试通过")
}

func TestGetDivid(t *testing.T){
	sub := model.GetDivid(1, 2)
	expected := -1
	if sub != expected {
		t.Fatalf("GetDivid.GetSub(1, 2) 期望值：%v，实际值：%v，测试不通过",sub,expected)
	}
	t.Logf("model.GetDivid(1, 2)测试通过")
}

