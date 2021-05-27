package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

const Pi = 3.14

var c, python, java bool
var i, j int = 1, 2
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func add(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(xx, nn, lim float64) float64 {
	if vv := math.Pow(xx, nn); vv < lim {
		return vv
	}
	return lim
}
func pow2(x1, n1, lim float64) float64 {
	if v1 := math.Pow(x1, n1); v1 < lim {
		return v1
	} else {
		fmt.Printf("%g >= %g\n", v1, lim)
	}
	// can't use v here, though
	return lim
}

type Vertex struct {
	X int
	Y int
}

var (
	v123 = Vertex{1, 2}  // has type Vertex
	v234 = Vertex{X: 1}  // Y:0 is implicit
	v345 = Vertex{}      // X:0 and Y:0
	p123 = &Vertex{1, 2} // has type *Vertex
)
var pow5 = []int{1, 2, 4, 8, 16, 32, 64, 128}

type Vertex2 struct {
	Lat, Long float64
}

var mm2 map[string]Vertex2

type Vertex3 struct {
	Lat, Long float64
}

var mw1 = map[string]Vertex3{
	"Bell Labs": Vertex3{
		40.68433, -74.39967,
	},
	"Google": Vertex3{
		37.42202, -122.08408,
	},
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Println("Hello, World!!!")
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(9))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i int
	fmt.Println(i, c, python, java)
	var m, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(m, j, k, c, python, java)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var q int
	var f float64
	var w bool
	var s string
	fmt.Printf("%v %v %v %q\n", q, f, w, s)
	var g, h int = 3, 4
	var v float64 = math.Sqrt(float64(g*g + h*h))
	var z uint = uint(v)
	fmt.Println(g, h, z)
	u := 43 // change me!
	fmt.Printf("u относится к типу %T\n", u)
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	l := time.Now()
	switch {
	case l.Hour() < 12:
		fmt.Println("Good morning!")
	case l.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	fmt.Println("world")
	defer fmt.Println("hello")

	fmt.Println("counting")

	for i1 := 0; i1 < 10; i1++ {
		defer fmt.Println(i1)
	}
	fmt.Println("done")
	i2, j := 42, 2701

	p2 := &i2        // point to i
	fmt.Println(*p2) // read i through the pointer
	*p2 = 21         // set i through the pointer
	fmt.Println(i2)  // see the new value of i

	p2 = &j        // point to j
	*p2 = *p2 / 37 // divide j through the pointer
	fmt.Println(j) // see the new value of j
	fmt.Println(Vertex{1, 2})
	v1 := Vertex{1, 2}
	v1.X = 4
	fmt.Println(v1.X)
	v22 := Vertex{1, 2}
	p22 := &v22
	p22.X = 1e9
	fmt.Println(v22)
	fmt.Println(v123, p123, v234, v345)
	var aa [2]string
	aa[0] = "Hello"
	aa[1] = "World"
	fmt.Println(aa[0], aa[1])
	fmt.Println(aa)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	aq := names[0:2]
	bq := names[1:3]
	fmt.Println(aq, bq)

	bq[0] = "XXX"
	fmt.Println(aq, bq)
	fmt.Println(names)

	qe := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(qe)

	re := []bool{true, false, true, true, false, true}
	fmt.Println(re)

	se := []struct {
		ie int
		be bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(se)
	sw := []int{2, 3, 5, 7, 11, 13}

	sw = sw[1:4]
	fmt.Println(sw)

	sw = sw[:2]
	fmt.Println(sw)

	sw = sw[1:]
	fmt.Println(sw)
	sa := []int{2, 3, 5, 7, 11, 13}
	printSlice(sa)

	// Slice the slice to give it zero length.
	sa = sa[:0]
	printSlice(sa)

	// Extend its length.
	sa = sa[:4]
	printSlice(sa)

	// Drop its first two values.
	sa = sa[2:]
	printSlice(sa)
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i3 := 0; i3 < len(board); i3++ {
		fmt.Printf("%s\n", strings.Join(board[i3], " "))
	}
	var st []int
	printSlice(st)

	st = append(st, 0)
	printSlice(st)

	st = append(st, 1)
	printSlice(st)

	st = append(st, 2, 3, 4)
	printSlice(st)
	for ie, ve := range pow5 {
		fmt.Printf("2**%d = %d\n", ie, ve)
	}
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	mm2 = make(map[string]Vertex2)
	mm2["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(mm2["Bell Labs"])
	fmt.Println(mw1)
	md := make(map[string]int)

	md["Answer"] = 42
	fmt.Println("The value:", md["Answer"])

	md["Answer"] = 48
	fmt.Println("The value:", md["Answer"])

	delete(md, "Answer")
	fmt.Println("The value:", md["Answer"])

	vx, ok := md["Answer"]
	fmt.Println("The value:", vx, "Present?", ok)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func printSlice1(st []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(st), cap(st), st)
}

func printSlice(sa []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(sa), cap(sa), sa)
	var sd []int
	fmt.Println(sd, len(sd), cap(sd))
	if sd == nil {
		fmt.Println("nil!")
	}

}
