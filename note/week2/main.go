package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

func testfor() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("第", i, "次循环")
	}

	switch i := 1; i {
	case 1:
		fmt.Println("进入case 1")
		fmt.Println("i等于1")
		break

	case 2:
		fmt.Println("i等于2")
	default:
		fmt.Println("default case ")
	}

	select {
	case <-time.After(time.Second * 2):
		fmt.Println("timeout2s")
	case <-time.After(time.Second):
		fmt.Println("timeout1s")
		fmt.Println("break之后")
		break

	}

	for i := 0; i < 5; i++ {
		fmt.Printf("不使用标记，外部循环，i=%d\n", i)
		for j := 0; j < 10; j++ {
			fmt.Printf("不使用标记，内部循环，j=%d\n", j)
		}
	}

	//outter:
	//	for i := 1; i <= 3; i++ {
	//		fmt.Printf("使用标记,外部循环, i = %d\n", i)
	//		for j := 5; j <= 10; j++ {
	//			fmt.Printf("使用标记,内部循环 j = %d\n", j)
	//			break outter
	//		}
	//	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("第", i, "次循环")
	}

	// 不使用标记
	for i := 1; i <= 2; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
			if j >= 7 {
				continue
			}
			fmt.Println("不使用标记，内部循环，在continue之后执行")
		}
	}

outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环 j = %d\n", j)
			if j >= 7 {
				continue outter
			}
			fmt.Println("使用标记，内部循环，在continue之后执行")
		}
	}

}

func testArr() {
	// 仅声明
	var a [5]int
	fmt.Println("a = ", a)

	var marr [2]map[string]string
	fmt.Println("marr = ", marr)

	var b [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	var c = [5]string{"c1", "c2", "c3", "c4", "c5"}
	fmt.Println("c = ", c)

	d := [3]int{3, 2, 1}
	fmt.Println("d = ", d)

	autoLen := [...]string{"auto1", "auto2", "auto3"}
	fmt.Println("autoLen = ", autoLen)

	positionInit := [5]string{1: "position1", 3: "position3"}
	fmt.Println("positionInit = ", positionInit)

}

func testInterview() {
	a := []int{5, 4, 3, 2, 0}

	element := a[2]
	fmt.Println("element=", element)

	for i, v := range a {
		fmt.Println("Index=", i, "value =", v)
	}

	for i := range a {
		fmt.Println("only Index,Index=", a[i])
	}

	fmt.Println("len(a) =", len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println("use len() index = ", i, "value=", a[i])
	}
}

func testArrs() {
	a := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}

	layer1 := a[0]
	layer2 := a[0][1]
	element := a[0][1][1]
	fmt.Println(layer1)
	fmt.Println(layer2)
	fmt.Println(element)

	// 多维数组遍历时，需要使用嵌套for循环遍历
	for i, v := range a {
		fmt.Println("index = ", i, "value = ", v)
		for j, inner := range v {
			fmt.Println("inner, index = ", j, "value = ", inner)
		}
	}
}

type Custom struct {
	i int
}

var carr [5]*Custom = [5]*Custom{
	{6},
	{7},
	{8},
	{9},
	{10},
}

func testCarr() {
	a := [5]int{5, 4, 3, 2, 1}
	fmt.Println("before all, a = ", a)
	for i := range carr {
		fmt.Printf("in main func, carr[%d] = %p, value = %v \n", i, &carr[i], *carr[i])
	}
	printFuncParamPointer(carr)

	receiveArray(a)
	fmt.Println("after receiveArray, a = ", a)

	receiveArrayPointer(&a)
	fmt.Println("after receiveArrayPointer, a = ", a)
}

func receiveArray(param [5]int) {
	fmt.Println("in receiveArray func, before modify, param = ", param)
	param[1] = -5
	fmt.Println("in receiveArray func, after modify, param = ", param)
}

func receiveArrayPointer(param *[5]int) {
	fmt.Println("in receiveArrayPointer func, before modify, param = ", param)
	param[1] = -5
	fmt.Println("in receiveArrayPointer func, after modify, param = ", param)
}

func printFuncParamPointer(param [5]*Custom) {
	for i := range param {
		fmt.Printf("in printFuncParamPointer func, param[%d] = %p, value = %v \n", i, &param[i], *param[i])
	}
}

func testAppend() {
	s3 := []int{}
	fmt.Println("s3 = ", s3)

	// append函数追加元素
	s3 = append(s3)
	s3 = append(s3, 1)
	s3 = append(s3, 2, 3)
	fmt.Println("s3 = ", s3)

	s4 := []int{1, 2, 4, 5}
	s4 = append(s4[:2], append([]int{3}, s4[2:]...)...)
	fmt.Println("s4 = ", s4)

	s5 := []int{1, 2, 3, 5, 4}
	s5 = append(s5[:3], s5[4:]...)
	fmt.Println("s5 = ", s5)
}

func testCopy() {
	src1 := []int{1, 2, 3}
	dst1 := make([]int, 4, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3, 3)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)

	copy(dst1, src1)
	copy(dst2, src2)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)
}

func testMap() {
	var m1 map[string]string
	fmt.Println("m1 length:", len(m1))

	m2 := make(map[string]string)
	fmt.Println("m2 length:", len(m2))
	fmt.Println("m2 =", m2)

	m3 := make(map[string]string, 10)
	fmt.Println("m3 length:", len(m3))
	fmt.Println("m3 =", m3)

	m4 := map[string]string{}
	fmt.Println("m4 length:", len(m4))
	fmt.Println("m4 =", m4)

	m5 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println("m5 length:", len(m5))
	fmt.Println("m5 =", m5)
}

func testMapde() {
	m := make(map[string]int, 10)

	m["1"] = int(1)
	m["2"] = int(2)
	m["3"] = int(3)
	m["4"] = int(4)
	m["5"] = int(5)
	m["6"] = int(6)

	// 获取元素
	value1 := m["1"]
	fmt.Println("m[\"1\"] =", value1)

	value1, exist := m["1"]
	fmt.Println("m[\"1\"] =", value1, ", exist =", exist)

	valueUnexist, exist := m["10"]
	fmt.Println("m[\"10\"] =", valueUnexist, ", exist =", exist)

	// 修改值
	fmt.Println("before modify, m[\"2\"] =", m["2"])
	m["2"] = 20
	fmt.Println("after modify, m[\"2\"] =", m["2"])

	// 获取map的长度
	fmt.Println("before add, len(m) =", len(m))
	m["10"] = 10
	fmt.Println("after add, len(m) =", len(m))

	// 遍历map集合main
	for key, value := range m {
		fmt.Println("iterate map, m[", key, "] =", value)
	}

	// 使用内置函数删除指定的key
	_, exist_10 := m["10"]
	fmt.Println("before delete, exist 10: ", exist_10)
	delete(m, "10")
	_, exist_10 = m["10"]
	fmt.Println("after delete, exist 10: ", exist_10)

	// 在遍历时，删除map中的key
	for key := range m {
		fmt.Println("iterate map, will delete key:", key)
		delete(m, key)
	}
	fmt.Println("m = ", m)
}

func testMapde2() {
	m := make(map[string]int)
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(2)

	go func() {
		for {
			lock.Lock()
			m["a"]++
			lock.Unlock()
		}
	}()

	go func() {
		for {
			lock.Lock()
			m["a"]++
			fmt.Println(m["a"])
			lock.Unlock()
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("timeout, stopping")
	}
}

func testfor2() {
	array := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	slice := [][]int{{1, 2}, {3}}
	// 只拿到行的索引
	for index := range array {
		// array[index]类型是一维数组
		fmt.Println(reflect.TypeOf(array[index]))
		fmt.Printf("array -- index=%d, value=%v\n", index, array[index])
	}

	for index := range slice {
		// slice[index]类型是一维数组
		fmt.Println(reflect.TypeOf(slice[index]))
		fmt.Printf("slice -- index=%d, value=%v\n", index, slice[index])
	}

	// 拿到行索引和该行的数据
	fmt.Println("print array element")
	for row_index, row_value := range array {
		fmt.Println(row_index, reflect.TypeOf(row_value), row_value)
	}

	fmt.Println("print array slice")
	for row_index, row_value := range slice {
		fmt.Println(row_index, reflect.TypeOf(row_value), row_value)
	}

	// 双重遍历，拿到每个元素的值
	for row_index, row_value := range array {
		for col_index, col_value := range row_value {
			fmt.Printf("array[%d][%d]=%d ", row_index, col_index, col_value)
		}
		fmt.Println()
	}
	for row_index, row_value := range slice {
		for col_index, col_value := range row_value {
			fmt.Printf("slice[%d][%d]=%d ", row_index, col_index, col_value)
		}
		fmt.Println()
	}
}

func addData(ch chan int) {
	size := cap(ch)
	for i := 0; i < size; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func main1() {
	ch := make(chan int, 10)

	go addData(ch)

	for i := range ch {
		fmt.Println(i)
	}
}

func main2() {
	hash := map[string]int{
		"a": 1,
		"f": 2,
		"z": 3,
		"c": 4,
	}

	for key := range hash {
		fmt.Printf("key=%s, value=%d\n", key, hash[key])
	}

	for key, value := range hash {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
}

type PayMethod interface {
	Account
	Pay(amount int) bool
}

type Account interface {
	GetBalance() int
}

// CreditCard 结构体实现 PaymentMethod 接口
type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Printf("信用卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("信用卡支付失败: 超出额度")
	return false
}

func (c *CreditCard) GetBalance() int {
	return c.balance
}

// DebitCard 结构体实现 PaymentMethod 接口
type DebitCard struct {
	balance int
}

func (d *DebitCard) Pay(amount int) bool {
	if d.balance >= amount {
		d.balance -= amount
		fmt.Printf("借记卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("借记卡支付失败: 余额不足")
	return false
}

func (d *DebitCard) GetBalance() int {
	return d.balance
}

// 使用 PaymentMethod 接口的函数
func purchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Printf("购买成功，剩余余额: %d\n", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}

func main3() {
	creditCard := &CreditCard{balance: 0, limit: 1000}
	debitCard := &DebitCard{balance: 500}

	fmt.Println("使用信用卡购买:")
	purchaseItem(creditCard, 800)

	fmt.Println("\n使用借记卡购买:")
	purchaseItem(debitCard, 300)

	fmt.Println("\n再次使用借记卡购买:")
	purchaseItem(debitCard, 300)
}

func (c *CreditCard) Pay1(amout int) {
	if c.balance < amout {
		fmt.Println("余额不足")
		return
	}
	c.balance -= amout
}

func anyParam(param interface{}) {
	fmt.Println("param: ", param)
}

func main4() {
	c := CreditCard{balance: 100, limit: 1000}
	c.Pay1(200)
	var a PayMethod = &c
	fmt.Println("a.Pay: ", a)

	var b interface{} = &c
	fmt.Println("b: ", b)

	anyParam(c)
	anyParam(1)
	anyParam("123")
	anyParam(a)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main6() {
	go func() {
		fmt.Println("run goroutine in closure")
	}()
	go func(string) {
	}("gorouine: closure params")
	go say("in goroutine: world")
	say("hello")
}

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 获取当前计数
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

type UnsafeCounter struct {
	count int
}

// 增加计数
func (c *UnsafeCounter) Increment() {
	c.count += 1
}

// 获取当前计数
func (c *UnsafeCounter) GetCount() int {
	return c.count
}

func main7() {
	counter := SafeCounter{}

	// 启动100个goroutine同时增加计数
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	// 等待一段时间确保所有goroutine完成
	time.Sleep(time.Second)

	// 输出最终计数
	fmt.Printf("Final count: %d\n", counter.GetCount())
}
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

func sendOnly(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("发送 %d\n", i)
	}
	close(ch)
}

func main9() {

	ch := make(chan int, 3)

	go sendOnly(ch)

	go receiveOnly(ch)

	timeout := time.After(2 * time.Second)

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main10() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}
	}()
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1:
			fmt.Printf("receive %d from channel 1\n", x)
		case y := <-ch2:
			fmt.Printf("receive %d from channel 2\n", y)
		case z := <-ch3:
			fmt.Printf("receive %d from channel 3\n", z)
		}
	}
}
func main() {
	//testfor()
	//testArr()
	//testInterview()
	//testArrs()
	//testCarr()
	//testAppend()
	//testCopy()
	//testMap()
	//testMapde2()
	//testfor2()
	//main1()
	//main2()
	main10()
}
