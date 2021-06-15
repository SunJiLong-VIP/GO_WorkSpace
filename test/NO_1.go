package main

import (
	"fmt"
	"math"
	"strings"
	"userinfo"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("test  go  code")
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	userinfo.Login()
	userinfo.Deleteuser()
	userinfo.Selectuser()
	v := `
		用户登录中
		输入用户名
		输入密码
		输入验证码
		登录成功
		删除用户成功
		查询用户成功`
	fmt.Println(strings.Split(v, ""))       //字符串分割
	fmt.Println(strings.Contains(v, "密码"))  //字符串是否包含指定内容
	fmt.Println(strings.HasPrefix(v, "用户")) //字符串是否以指定内容开头
	fmt.Println(strings.HasSuffix(v, "成功")) //字符串是否以指定内容结束
	fmt.Println(strings.Index(v, "用户"))     //指定内容首次出现的位置
	fmt.Println(strings.LastIndex(v, "用户")) //指定内容最后出现的位置
	v1 := []string{"请", "输", "入", "验", "证", "码"}
	fmt.Println(strings.Join(v1, "-")) //join 连接

	v2 := [10]int{4, 5, 8, 6, 2, 1, 7, 9, 0, 3}
	fmt.Println(v2)
	for i := 0; i < len(v2)-1; i++ {
		for j := 0; j < len(v2)-1-i; j++ {
			if v2[j] > v2[j+1] {
				v2[j], v2[j+1] = v2[j+1], v2[j]
			}
		}
	}
	fmt.Println(v2)

	var t []int
	for i := 0; i < 10; i++ {
		t = append(t, i)
		fmt.Printf("%v,len:%d,cap:%d,ptr:%p\n", t, len(t), cap(t), t)
		//fmt.Println(make([]int, i))
	}
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	fmt.Printf("%T", a)

	var vv = make(map[string]int, 5)
	vv["张三"] = 100
	vv["李四"] = 90
	fmt.Printf(" vv:%#v\n", vv)
	value, ok := vv["二狗子"]
	//	value, ok := vv["李四"]
	if ok {
		fmt.Println("李四:", value, "is  true")
	} else {
		fmt.Println("查无此人！")
	}

}
