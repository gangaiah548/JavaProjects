package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type Shape interface {
	area() float64
	//perimeter() float64
}
type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func printShape(s Shape) {
	fmt.Println("Area:", s.area())
}
func main() {
	
	var d interface{} = "hello"

    // Type assertion
    st := d.(string)
    fmt.Println(st)

    // Type assertion with a check
    st, ok := d.(string)
    fmt.Println(st, ok)

    // Type assertion that will fail
    ft, ok := d.(float64)
    fmt.Println(ft, ok)


	rect := Rectangle{5, 3}
	printShape(rect)
	fmt.Print("hell")
	keys := make([]int, 0, 2)
	for k := range keys {
		keys = append(keys, k)
	}

	//arr []int=make()
	s := []int{5, 8, 3, 2, 7, 1}
	sort.Ints(s)
	fmt.Print(keys, "sort", s)

	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Original array:", arr)

	sorted := mergeSort(arr)
	fmt.Println("Sorted array:", sorted)

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			val, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- val * 2
		}
		close(ch2)
	}()

	for val := range ch2 {
		fmt.Println(val)
	}

	ch3 := make(chan int)
	ch4 := make(chan int, 1)
	go func() {
		ch3 <- 1
		close(ch3)
	}()
	go func() {
		for {
			val, ok := <-ch3

			if !ok {
				break
			}
			ch4 <- val
		}
		close(ch4)
	}()

	for val := range ch4 {
		fmt.Println("ch4 ", val)
	}

	// ch5 := make(chan int)
	// ch6 := make(chan int)

	// go func() {
	//     ch5 <- 1
	//     <-ch6
	// }()

	// go func() {
	//     ch6 <- 1
	//     <-ch5
	// }()

	slice := []int{1, 2, 3}
	slice = append(slice, 2)
	slice = append(slice[:2], slice[2+1:]...) //index := 2 // Index of the element to remove
	//slice = append(slice[:index], slice[index+1:]...)
	ind := 2
	slice = append(slice[:ind], slice[ind+1:]...)
	mmap := make(map[int]int)
	omap := map[int]int{
		1: 1,
	}

	for k, v := range omap {
		mmap[k] = v
	}

}

// func sorting string(w string){
// w.sort(s,)
// }

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for len(a) < i && len(b) < j {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func timeLoop(ml []int, n int) time.Duration {
	var v0 = time.Now()
	for len(ml) < n {
		ml = append(ml, 1)
	}
	return time.Since(v0)
}

func charFrequncy(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}
	return freq
}

func returnPrimes(n int) []int {
	ar := []int{}
	for i := 2; i < n; i++ { // start from 2 as 0 and 1 are not prime numbers
		if isprime(i) {
			ar = append(ar, i)
		}
	}
	return ar
}

func isprime(s int) bool {
	if s <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(s))); i++ {
		if s%i == 0 {
			return false
		}
	}

	return true
}

func lengthnonrpeatedsubstring(s string) int {
	charin := make(map[rune]int)
	start, res := 0, 0
	for i, char := range s {
		if i, found := charin[char]; found && i >= start {
			start = i + 1
		}
		charin[char] = i
		if i-start+1 > res {
			res = i - start + 1
		}

	}
	return res

}

type charFreq struct {
	char byte
	freq int
}

func moveZeros(ar []int) {
	j := 0
	for i := 0; i < len(ar); i++ {
		if ar[i] != 0 {
			ar[i], ar[j] = ar[j], ar[i]
			j++
		}

	}

}

func chantest() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
		}

	}()
	close(ch1)

	func() {
		for num := range ch1 {
			ch2 <- num * 2
		}

	}()

	for res := range ch2 {
		fmt.Print(res)
	}

}

// //
func nthHighestRepeatedChar(s string, n int) (byte, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("string is empty")
	}

	freqMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freqMap[s[i]]++
	}

	type charFreq struct {
		char byte
		freq int
	}

	freqSlice := make([]charFreq, 0, len(freqMap))
	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i].freq > freqSlice[j].freq
	})
	if n <= 0 || n > len(freqSlice) {
		return 0, fmt.Errorf("n is out of range")
	}
	return 0, nil
}

type Person struct {
	name string
	ag   int
}

func structSort() {
	tp := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	sort.Slice(tp, func(i, j int) bool {
		return tp[i].ag < tp[j].ag
	})
	fmt.Println("post mod", tp)
}
