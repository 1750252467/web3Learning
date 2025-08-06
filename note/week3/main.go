package main

import (
	"context"
	"fmt"
	"time"
)

// func main() {
// 	waitOnCond := func(ctx context.Context, cond *sync.Cond, conditionMet func() bool) error {
// 		stopf := context.AfterFunc(ctx, func() {
// 			// We need to acquire cond.L here to be sure that the Broadcast
// 			// below won't occur before the call to Wait, which would result
// 			// in a missed signal (and deadlock).
// 			cond.L.Lock()
// 			defer cond.L.Unlock()

// 			// If multiple goroutines are waiting on cond simultaneously,
// 			// we need to make sure we wake up exactly this one.
// 			// That means that we need to Broadcast to all of the goroutines,
// 			// which will wake them all up.
// 			//
// 			// If there are N concurrent calls to waitOnCond, each of the goroutines
// 			// will spuriously wake up O(N) other goroutines that aren't ready yet,
// 			// so this will cause the overall CPU cost to be O(N²).
// 			cond.Broadcast()
// 		})
// 		defer stopf()

// 		// Since the wakeups are using Broadcast instead of Signal, this call to
// 		// Wait may unblock due to some other goroutine's context being canceled,
// 		// so to be sure that ctx is actually canceled we need to check it in a loop.
// 		for !conditionMet() {
// 			cond.Wait()
// 			if ctx.Err() != nil {
// 				return ctx.Err()
// 			}
// 		}

// 		return nil
// 	}

// 	cond := sync.NewCond(new(sync.Mutex))

// 	var wg sync.WaitGroup
// 	for i := 0; i < 4; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()

// 			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
// 			defer cancel()

// 			cond.L.Lock()
// 			defer cond.L.Unlock()

// 			err := waitOnCond(ctx, cond, func() bool { return false })
// 			fmt.Println(err)
// 		}()
// 	}
// 	wg.Wait()

// }
// const (
// 	KEY = "trace_id"
// )

// func NewRequestID() string {
// 	return strings.Replace(uuid.New().String(), "-", "", -1)
// }

// func NewContextWithTraceID() context.Context {
// 	ctx := context.WithValue(context.Background(), KEY, NewRequestID())
// 	return ctx
// }

// func PrintLog(ctx context.Context, message string) {
// 	fmt.Printf("%s|info|trace_id=%s|%s", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, KEY), message)
// }

// func GetContextValue(ctx context.Context, k string) string {
// 	v, ok := ctx.Value(k).(string)
// 	if !ok {
// 		return ""
// 	}
// 	return v
// }

// func ProcessEnter(ctx context.Context) {
// 	PrintLog(ctx, "Golang梦工厂")
// }

//	func main() {
//		ProcessEnter(NewContextWithTraceID())
//	}
//
// }
// const (
//
//	KEY = "trace_id"
//
// )
// func main() {
// 	HttpHandler()
// }

// func NewContextWithTimeout() (context.Context, context.CancelFunc) {
// 	return context.WithTimeout(context.Background(), 3*time.Second)
// }

// func HttpHandler() {
// 	ctx, cancel := NewContextWithTimeout()
// 	defer cancel()
// 	deal(ctx)
// }

//	func deal(ctx context.Context) {
//		for i := 0; i < 10; i++ {
//			time.Sleep(1 * time.Second)
//			select {
//			case <-ctx.Done():
//				fmt.Println(ctx.Err())
//				return
//			default:
//				fmt.Printf("deal time is %d\n", i)
//			}
//		}
//	}
// func main() {
// 	HttpHandler1()
// }

// func NewContextWithTimeout1() (context.Context, context.CancelFunc) {
// 	return context.WithTimeout(context.Background(), 3*time.Second)
// }

// func HttpHandler1() {
// 	ctx, cancel := NewContextWithTimeout1()
// 	defer cancel()
// 	deal1(ctx, cancel)
// }

//	func deal1(ctx context.Context, cancel context.CancelFunc) {
//		for i := 0; i < 10; i++ {
//			time.Sleep(1 * time.Second)
//			select {
//			case <-ctx.Done():
//				fmt.Println(ctx.Err())
//				return
//			default:
//				fmt.Printf("deal time is %d\n", i)
//				cancel()
//			}
//		}
//	}
func main() {

	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	go Speak(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func Speak(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println("speak end")
			return
		default:
			fmt.Println("speak")

		}

	}
}
