package hello

import "fmt"

// 一个用于演示如何发布自定义包的示例项目。

// SayHi 一个打招呼的函数
func SayHi(name string) {
	fmt.Printf("你好%s，我是七米。很高兴认识你。\n", name)
}
