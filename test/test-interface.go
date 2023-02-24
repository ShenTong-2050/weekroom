package main

type people struct {}

type ChineseSpeak interface {
	SayChinese()	string
	ForeignSpeak
}

type ForeignSpeak interface {
	SayEnglish()	string
}

func (people) SayChinese() string {
	return "你好,世界!"
}

func (people) SayEnglish() string {
	return "hello world"
}

type data struct {

}

func (data) testString() string {
	return "匿名接口变量"
}

// 定义匿名接口变量
type node struct {
	data interface{				// 匿名接口类型
		testString() string
	}
}

// Test 超级接口变量隐式转换子集接口
func Test(f ForeignSpeak) {
	println(f.SayEnglish())
}

func main() {

	/*var p people

	var x ChineseSpeak = &p

	fmt.Printf("the p ptr : %p, value:%v\n",&p,p)

	fmt.Printf("the x ptr : %p, value:%v\n",&x,x)

	fmt.Println(x.SayEnglish())

	fmt.Println(x.SayChinese())

	// 隐式转换为子集接口
	var f ForeignSpeak
	f = &p
	Test(f)*/

	// 定义匿名接口类型
	var anonymity interface{
		testString()	string
	} = data{}

	n := node{data: anonymity}

	println(n.data.testString())


}
