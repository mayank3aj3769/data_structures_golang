package sorts

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Enter array ")
	var arr = make([]int, 10)
	var res = make([]int, 10)
	for i := 0; i < 10; i++ {

		fmt.Scanf("%d", &arr[i])
	}
	res = quicksort(arr)
	fmt.Println("%q", res)

}

func quicksort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = quicksort(lowPart)
	highPart = quicksort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}
