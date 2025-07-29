package main

import (
	"fmt"
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
func main() {
	//testfor()
	//testArr()
	//testInterview()
	//testArrs()
	//testCarr()
	//testAppend()
	//testCopy()
	//testMap()
	testMapde()
}
