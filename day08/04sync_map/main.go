package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Go内置的Map不是并发安全的

var m = make(map[string]int)

var lock sync.Mutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 21; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			key := strconv.Itoa(n)
// 			lock.Lock()
// 			set(key, n)
// 			lock.Unlock()
// 			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

var m2 = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			value, _ := m2.Load(key)                // 必须使用sync.Map内置的Store方法设置键值对
			fmt.Printf("k=:%v,v:=%v\n", key, value) // 必须使用sync.Map提供的Load方法根据key取值
			wg.Done()
		}(i)
	}
	wg.Wait()
}
