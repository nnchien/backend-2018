package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

func main() {
	readMemStats()

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("mem size of nums: ", unsafe.Sizeof(nums))

	fmt.Printf("cap(nums): %d, len(nums): %d\n", cap(nums), len(nums))

	middle := nums[1:3]
	fmt.Println("create middle:", middle)
	fmt.Printf("cap(middle): %d, len(middle): %d\n", cap(middle), len(middle))
	fmt.Println("mem size of middle: ", unsafe.Sizeof(middle))
	readMemStats()

	middle[1] *= 2
	fmt.Println("double first elemen in middle:", middle)

	middle = append(middle, 11)
	fmt.Println("append 11 to end of middle:", middle)
	fmt.Printf("cap(middle): %d, len(middle): %d\n", cap(middle), len(middle))
	fmt.Println("mem size of middle: ", unsafe.Sizeof(middle))
	fmt.Println("nums:", nums)
	fmt.Printf("cap(nums): %d, len(nums): %d\n", cap(nums), len(nums))

	readMemStats()

	middle = append(middle, 12)
	middle = append(middle, 13)
	middle = append(middle, 14)
	middle = append(middle, 15)

	fmt.Println("append 12->15 to end of middle:", middle)
	fmt.Printf("cap(middle): %d, len(middle): %d\n", cap(middle), len(middle))
	fmt.Println("mem size of middle: ", unsafe.Sizeof(middle))
	fmt.Println("nums:", nums)
	fmt.Printf("cap(nums): %d, len(nums): %d\n", cap(nums), len(nums))

	readMemStats()

	middle = append(middle, 16)
	middle = append(middle, 17)
	fmt.Println("append 16 17 to end of middle:", middle)
	fmt.Printf("cap(middle): %d, len(middle): %d\n", cap(middle), len(middle))
	fmt.Println("mem size of middle: ", unsafe.Sizeof(middle))
	fmt.Println("nums:", nums)
	fmt.Printf("cap(nums): %d, len(nums): %d\n", cap(nums), len(nums))

	readMemStats()

	middle = append(middle, 19)
	fmt.Println("append 19 to end of middle:", middle)
	fmt.Printf("cap(middle): %d, len(middle): %d\n", cap(middle), len(middle))
	fmt.Println("mem size of middle: ", unsafe.Sizeof(middle))
	fmt.Println("nums:", nums)
	fmt.Printf("cap(nums): %d, len(nums): %d\n", cap(nums), len(nums))

	fmt.Println("append 21->100000000 to end of middle:")
	last := readMemStats()
	start := time.Now()
	for i := 21; i < 100000000; i++ {
		middle = append(middle, i)
		if time.Since(start) > (time.Duration(10) * time.Millisecond) {
			current := readMemStats()
			going := "down"
			if last-current > 0 {
				going = "up"
			}
			fmt.Printf("going %s: %d\n", going, current)
			last = current
		}
		start = time.Now()
	}
	fmt.Printf("cap(middle): %d, len(middle): %d\n", cap(middle), len(middle))
	fmt.Println("mem size of middle: ", unsafe.Sizeof(middle))
}

func readMemStats() uint32 {
	var r runtime.MemStats
	runtime.ReadMemStats(&r)

	fmt.Printf("Heap size %d\n", r.HeapAlloc)
	fmt.Printf("NumGC\t %d\n", r.NumGC)

	return r.NumGC
}
