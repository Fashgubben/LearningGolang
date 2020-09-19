package main

import (
	"fmt"
	"math"

	"github.com/fashgubben/LearningGolang/online_course/03_packages/strutil"
)

func main() {

	// Use math
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(1.6))
	fmt.Println(math.Sqrt(16))

	fmt.Println(strutil.Reverse("Hello"))
}
