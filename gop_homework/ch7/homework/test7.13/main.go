//练习 7.13： 为Expr增加一个String方法来打印美观的语法树。当再一次解析的时候，检查它的结果是否生成相同的语法树。
package main

import "fmt"

type Expr interface {
	fmt.Stringer
}
