package main

import (
	"bytes"
	"fmt"
	"sort"
	"sync"
	"time"

	"golang.org/x/exp/constraints"
	"github.com/emirpasic/gods/maps/linkedhashmap"
)

type userId int

func Sub[T ~int | ~float64](a T, b T) T { return a - b }

type Num interface {
	int
	int16
	int32
	int64
	int8
	float32
	float64
}
type Sam struct {
	Name string
	data string // Unexported field
}
type Employee struct {
	EmployeeId   int
	EmployeeName string
	EmployeeAdd  []string
}

func Add[T constraints.Ordered](a T, b T) T {
	return a + b
}

var sad int = 100

// fmt.Println("sahdow2",sad);
//

func main() {
	m := linkedhashmap.New()

	//input -> [2,7,11,15] target = 9
	fmt.Println(m)
	tar := []int{2, 7, 11, 15}
	//output -> [0,1]
	op := twoSum(9, tar)
	fmt.Println(op)
	test()
	pyramid()
	// numbers := [5]int{4, 2, 7, 1, 5}
	// sort.Ints(numbers[:]) ///slice
	// fmt.Println(numbers)
	// s1 := Sam{Name: "John", data: "secret1"}
	// s2 := Sam{Name: "John", data: "secret2"}

	// if cmp.Equal(s1, s2, cmpopts.IgnoreUnexported(Sam{})) {
	// 	fmt.Println("x and y are equal")
	// } else {
	// 	fmt.Println("x and y are not equal")
	// }

	// fmt.Println("arrOfNUmberAddition", arrOfNUmberAddition(1, 2, 3, 4, 5))
	// fail := func(err error) (int64, error) {
	// 	return 0, fmt.Errorf("create order %v", err)
	// }
	// fmt.Println(fail)
	// myMap := make(map[string]int)
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("Enter key-value pairs (format: key=value), type 'exit' to quit:")

	// for {
	// 	fmt.Print("Enter key=value: ")
	// 	scanner.Scan()
	// 	input := scanner.Text()
	// 	if strings.ToLower(input) == "exit" {
	// 		break
	// 	}
	// 	parts := strings.Split(input, "=")
	// 	if len(parts) != 2 {
	// 		fmt.Println("Invalid input. Use the format: key=value")
	// 		continue
	// 	}
	// 	key := parts[0]
	// 	value, err := strconv.Atoi(parts[1])
	// 	if err != nil {
	// 		fmt.Println("int onlly allwoed")
	// 		continue
	// 	}
	// 	myMap[key] = value
	// }
	// fmt.Println("myMap ", myMap)
	// go printNumbers()
	// on := [8]int{1, 2, 34, 5, 6, 7, 8}
	// onslice := on[2:5]
	// fmt.Println(onslice)
	// emp := Employee{
	// 	101,
	// 	"mm",
	// 	[]string{"add1", "add2"},
	// }
	// fmt.Println("Hello", emp)
	// fmt.Println("printbyteSlicecompare()", printbyteSlicecompare())
	// fmt.Print("Hello World")
	// var arr []int = []int{1, 2, 3, 4}
	// var mat [2][2]int = [2][2]int{{1, 2}, {3, 4}}
	// arr2 := append(arr, 55)
	// arr3 := append(arr, arr2...)
	// fmt.Println(arr2)
	// fmt.Println(arr3)
	// fmt.Println(mat)
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Print(arr[i])
	// }
	// var fib0 int = 0
	// var fib1 int = 1
	// var sad int = 66
	// for i := 1; i < 5; i++ {
	// 	var t int = fib0
	// 	fib0 = fib0 + fib1
	// 	fmt.Print(fib0, fib1)
	// 	fib1 = t
	// }
	// fmt.Println("sahdow", sad)
	// fibSeq := fibonacciIterative(5)
	// fmt.Println("Fibonacci sequence (iterative):", fibSeq)
	// empsal := make(map[string]int)
	// empsal = map[string]int{
	// 	"neha": 2000,
	// }
	// fmt.Print("map ", empsal)
	// fmt.Print("neha record ", empsal["neha"])
	// _, flag := empsal["neha"]
	// fmt.Print("neha valid ", flag)
	// fmt.Print(reverseString("Helllo World"))

	// var wg sync.WaitGroup

	// // Add the number of goroutines to wait for
	// wg.Add(2)

	// // Start two goroutines
	// go printNumbersWithWg(&wg)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("Hello from another goroutine")
	// 	time.Sleep(time.Second * 1)
	// }()

	// // Wait for both goroutines to finish
	// wg.Wait()
	// fmt.Println("All goroutines finished")

	// numch := make(chan int)
	// strch := make(chan string)
	// go sendNumbersChan(numch)
	// go sendCharsChan(strch)
	// for i := 0; i < 5; i++ {
	// 	select {
	// 	case num, ok := <-numch:
	// 		if ok {
	// 			fmt.Println(num)
	// 		} else {
	// 			numch = nil
	// 		}
	// 	case str, ok := <-strch:
	// 		if ok {
	// 			fmt.Println(str)
	// 		} else {
	// 			strch = nil
	// 		}

	// 	}

	// 	if numch == nil && strch == nil {
	// 		break
	// 	}

	// }
	// fmt.Println("All data received")
}

func fibonacciIterative(n int) []int {
	if n < 1 {
		return []int{}
	}

	fibseq := make([]int, n)
	fibseq[0] = 0
	if n > 1 {
		fibseq[1] = 1
	}
	for i := 2; i < n; i++ {
		fibseq[i] = fibseq[i-1] + fibseq[i-2]
	}

	return fibseq
}

func reverseString(inp string) string {
	runes := []rune(inp)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}

}
func printNumbersWithWg(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
}

func printbyteSlicecompare() int {
	x1 := []byte{'c', 'o', 'd', 'i', 'g'}
	x2 := []byte{'n', 'i', 'n', 'j', 'a', 's'}
	op := bytes.Compare(x1, x2)
	return op
}

func arrOfNUmberAddition(val ...int) int {
	s := 0
	for i := range val {
		s += val[i]
	}
	return s
}

func sendNumbersChan(ch chan int) {
	for i := 0; i < 6; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func sendCharsChan(ch chan string) {
	for _, c := range []string{"a", "b", "c", "d"} {
		ch <- c
		time.Sleep(time.Millisecond * 600)
	}
	close(ch)
}

func sorChars(s string) string {
	var runes []rune = []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}
func reverseStringPreserveSpaces(s string) string {
	n := len(s)
	res := make([]byte, n)
	space := make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			space = append(space, i)
		}
	}

	j := 0
	for i := n - 1; i >= 0; i-- {
		if len(space) > 0 && i == space[len(space)-1] {
			res[j] = ' '
			space = space[:len(space)-1]
		} else {
			res[j] = s[i]
		}
		j++
	}

	return string(res)
}
func twoSum(t int, ar []int) []int {
	//var rv int
	//var id int
	//f := 1

	for i := 0; i < len(ar)-1; i++ {
		//rv = i
		for j := i + 1; j < len(ar); j++ {
			//id = j
			//fmt.Println("match case",ar[i],ar[j])
			if ar[i]+ar[j] == t {
				//f = 0
				tar := []int{i, j}
				return tar
			}

		}

	}

	return nil

}

func twoSumd(t int, ar []int) []int {
	//var rv int
	//var id int
	//f := 1
	mp := make(map[int]int)
	for i := 0; i < len(ar)-1; i++ {

		if ar[i] < t {
			value, exists := mp[t-ar[i]]
			if exists {
				tr := []int{i, value}
				return tr
			} else {
				mp[ar[i]] = i
			}
		}
	}

	return nil

}

//input -> [2,7,11,15] target = 9

//output -> [0,1]
