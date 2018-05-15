package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("cap(nums): %d, len(nums): %d\n", cap(nums), len(nums))

	middle := nums[1:3]
	fmt.Println("create middle:", middle)
	fmt.Printf("cap(middle): %d, len(middle): %d\n", cap(middle), len(middle))

	middle[1] *= 2
	fmt.Println("double first elemen in middle:", middle)

	middle = append(middle, 11)
	fmt.Println("append 11 to end of middle:", middle)

	fmt.Println("nums:", nums)
	fmt.Printf("cap(nums): %d, len(nums): %d\n", cap(nums), len(nums))

	fmt.Printf("cap(middle): %d, len(middle): %d\n", cap(middle), len(middle))

	fmt.Print("they are: ")
	for i := 0; i < len(middle); i++ {
		fmt.Printf("%d ", middle[i])
	}
	fmt.Print("\n")

	middle = append(middle, 12)
	middle = append(middle, 13)
	middle = append(middle, 14)
	middle = append(middle, 15)

	fmt.Println("append 12 13 14 15 to end of middle:", middle)
	fmt.Printf("cap(middle): %d, len(middle): %d\n", cap(middle), len(middle))
	fmt.Println("nums:", nums)

	middle = append(middle, 16)
	middle = append(middle, 17)
	fmt.Println("append 16 17 to end of middle:", middle)
	fmt.Printf("cap(middle): %d, len(middle): %d\n", cap(middle), len(middle))
	fmt.Println("nums:", nums)

	middle = append(middle, 19)
	fmt.Println("append 19 to end of middle:", middle)
	fmt.Printf("cap(middle): %d, len(middle): %d\n", cap(middle), len(middle))
	fmt.Println("nums:", nums)
}
