package main

import (
	"fmt"
	"time"
)

func main() {
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
