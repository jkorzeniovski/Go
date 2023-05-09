package main

import (
	"fmt"
	"math"
	"time"
)

const c1 = "constant"

func cony(){
	const c2 = 500000

	const c3 = 3e20 / c2
	fmt.Println(c3)

	fmt.Println(int64(c3))

	fmt.Println(math.Sin(c2))
}

func loops(){
	for j := 1; j<=5; j++{
		fmt.Println(j)
	//	j = j+1
	}
}

func vars(){
	var a = "initial"
	var b, c int = 1, 2
	var d int
	e := "japko"
	fmt.Println(a)
	fmt.Println(b, c, b+c)
	fmt.Println(d)
	fmt.Println(e)
}

func ifs(){
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

func switches(){
	var i int = 2 
	fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("its weekend")
	default:
		fmt.Println("its weekday")
	}
}

func arrays(){
	var ar1[5]int
	ar1[0] = 3
	fmt.Println(ar1)
}

func slices(){
	sl1 := make([]string, 3)
	sl1[0] = "x"
	sl1[1] = "d"
	sl1[2] = "b"
	fmt.Println(sl1)

	// sl1 = append(sl1, "h")
	// fmt.Println(sl1) // wywala odmowe dostępu
}

func main(){
	fmt.Println("go" + "lang")
	//fmt.Println(true && true)
	vars()
	fmt.Println(c1)
	cony()
	loops()
	ifs()
	switches()
	arrays()
	slices()
}
