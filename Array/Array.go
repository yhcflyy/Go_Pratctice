/*
• 数组是值类型,赋值和传参会复制整个数组,⽽而不是指针。
• 数组⻓长度必须是常量,且是类型的组成部分。[2]int 和 [3]int 是不同类型。
• ⽀支持 "=="、"!=" 操作符,因为内存总是被初始化过的。
• 指针数组 [n]*T,数组指针 *[n]T。
*/

/*
a := [3]int{1, 2}					// 未初始化元素值为 0。
b := [...]int{1, 2, 3, 4}			// 通过初始化值确定数组⻓长度。
c := [5]int{2: 100, 4:200}			// 使⽤用索引号初始化元素。
d := [...]struct {					// 可省略元素类型。
    		name string
		age uint8 
	}{
    {"user1", 10},
    {"user2", 20},					// 别忘了最后⼀一⾏行的逗号。
}
*/
package main

import "fmt"

func test(a [2]int){
	fmt.Printf(" in function a:%p\n",&a)
	a[1]=10000
}

func main(){
	a :=[2]int{}
	fmt.Printf("out function a:%p\n",&a);
	test(a)
	fmt.Println(a)
	println(len(a),cap(a))	//内置函数len和cap都返回数组的长度
}

//slice 并不是数组或数组指针。它通过内部指针和相关属性引⽤用数组⽚片段,以实现变⻓长⽅方案。
/*
struct Slice
{   					// must not move anything
byte*    array;		// actual data
uintgo   len;		// number of elements
uintgo   cap;    	// allocated number of elements                 
};
*/

/*
• 引⽤用类型。但⾃自⾝身是结构体,值拷⻉贝传递。
• 属性 len 表⽰示可⽤用元素数量,读写操作不能超过该限制。 
• 属性 cap 表⽰示最⼤大扩张容量,不能超出数组限制。
• 如果 slice == nil,那么 len、cap 结果都等于 0。
*/


