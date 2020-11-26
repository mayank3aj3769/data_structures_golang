//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

import "fmt"

func merge(a []int, b []int) []int {

	var r = make([]int, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {

		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}

	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r

}

//Mergesort Perform mergesort on a slice of ints
func Mergesort(items []int) []int {

	if len(items) < 2 {
		return items

	}

	var middle = len(items) / 2
	var a = Mergesort(items[:middle])
	var b = Mergesort(items[middle:])
	return merge(a, b)

}
func main() {

	fmt.Println("Enter array ")
	var arr = make([]int, 10)
	var res = make([]int, 10)
	for i := 0; i < 10; i++ {

		fmt.Scanf("%d", &arr[i])
	}
	res = Mergesort(arr)
	fmt.Println("%q", res)

}
