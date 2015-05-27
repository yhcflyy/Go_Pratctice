package main

import "fmt"
import "os"
import "sync"
import "testing"
/*
• 不⽀支持 嵌套 (nested)、重载 (overload) 和 默认参数 (default parameter)。
• ⽆无需声明原型。
• ⽀支持不定⻓长变参。
• ⽀支持多返回值。
• ⽀支持命名返回参数。 • ⽀支持匿名函数和闭包。
*/
func main() {
	fmt.Println(testReturn(1, 2, "testReturn:%d"))
	testFunctionTyoe()
	testMultifyParam()
	testAnonymousFuc()
	testClosure()
	//testDefer()
	testDefer3()
	testPanic()
}

/*
• 函数返回多个值,
• 返回参数可看做与形参类似的局部变量,最后由 return 隐式返回。
*/
func testReturn(x, y int, s string) (a int, str string) {
	// 类型相同的相邻参数可合并。
	a = x + y
	// 多返回值必须⽤用括号。
	return a, fmt.Sprintf(s, a)
}

/*
• 函数是第一类对象,可作为参数传递
*/

type Formatfun func(str string, a, b int) string

func format(fn Formatfun, str string, a, b int) string {
	return fn(str, a, b)
}

func testFunctionTyoe() {
	s := format(func(s string, a, b int) string {
		return fmt.Sprintf(s, a, b)
	}, "%d,%d", 10, 20)
	println(s)
}

/*
• 参数可以有变参，变参本质上就是 slice。只能有一个,且必须是最后一个
*/
func multifyParam(s string, a ...int) string {
	var total int
	for _, i := range a {
		total += i
	}
	return fmt.Sprintf(s, total)
}

func testMultifyParam() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	str := multifyParam("the sum of the param is:%d", a...)
	println(str)
}

/*
  匿名函数可赋值给变量,做为结构字段,或者在 channel ⾥里传送。
*/
func testAnonymousFuc() {
	fn := func() { println("Hello World fn") }
	fn()

	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	println(fns[0](100))

	s := struct {
		fn func() string
	}{
		fn: func() string { return "Hello world struct" },
	}
	println(s.fn())

	chanFuc := make(chan func() string, 2)
	chanFuc <- func() string { return "chanFunc" }
	println((<-chanFuc)())

}

//Closure,闭包复制的是原对象指针,这就很容易解释延迟引⽤用现象
//在汇编层⾯面,test 实际返回的是 FuncVal 对象,其中包含了匿名
//函数地址、闭包对象指 针。当调⽤用匿名函数时,只需以某个寄存器传递该对象即可。

func testclosure() func() {
	x := 100
	fmt.Printf("111 x (%p) = %d\n", &x, x)
	return func() {
		fmt.Printf("222 x (%p) = %d\n", &x, x)
	}
}

func testClosure() {
	f := testclosure()
	f()
}

//延迟调⽤
//用关键字 defer ⽤用于注册延迟调⽤。这些调⽤用直到 ret 前才被执⾏行,
//通常⽤用于释放资源或错误处理。
//ret是汇编中的指令，用于将函数弹出栈

func testDefer() {
	deferTest1()
	deferTest2(0)
}

func deferTest1() error {
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}
	defer f.Close() //注册调⽤用,⽽而不是注册函数。必须提供参数,哪怕为空。
	f.WriteString("Hello, World testDefer")
	return nil
}

//多个 defer 注册,按 FILO 次序执⾏行。哪怕函数或某个延迟调⽤用发⽣生错误,这些调⽤用依旧 会被执⾏行。
func deferTest2(x int) {
	defer println("a")
	defer println("b")
	defer func() {
		fmt.Printf("%f\n", float32(100.0/x))
	}()
	defer println("c")

}

//延迟调⽤用参数在注册时求值或复制,可⽤用指针或闭包 "延迟" 读取。
func testDefer3(){
	x,y :=10,20
	defer func (i int){
		println("defer",i,y)// y 闭包引⽤,引用y的地址
	}(x)//x被复制
	
	x += 10
	y += 100
	
	println("x=",x,"y=",y)
}

//大循环使用defer导致性能下降
var lock sync.Mutex
func test() {
    lock.Lock()
    lock.Unlock()
}
func testdefer() {
    lock.Lock()
    defer lock.Unlock()
}
func BenchmarkTest(b *testing.B) {
    for i := 0; i < b.N; i++ {
test() }
}
func BenchmarkTestDefer(b *testing.B) {
    for i := 0; i < b.N; i++ {
        testdefer()
    }
}

func testPanic(){
	defer func(){
		if err:=recover();err != nil{
			println(err.(string))
		}
	}()
	panic("panic error!")
}












