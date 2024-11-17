package main

import "fmt"

func main() {
	// array
	var ages [3]int = [3]int{20, 25, 30}
	fmt.Println(ages, len(ages))

	names := [4]string{"John", "Doe", "Jane", "Doe"}
	fmt.Println(names, len(names))
	names[1] = "Smith"
	fmt.Println(names, len(names))

	//slices
	var scores = []int{100, 200, 300}

	fmt.Println(scores, len(scores))
	scores[2] = 400
	fmt.Println(scores, len(scores))
	// appending to slices
	scores = append(scores, 500)
	fmt.Println(scores, len(scores))

	//slice ranges
	range1 := names[1:3]
	fmt.Println(range1, len(range1))
	range2 := names[2:]
	fmt.Println(range2, len(range2))
	range3 := names[:3]
	fmt.Println(range3, len(range3))

	//append in slices
	range3 = append(range3, "David")
	fmt.Println(range3, len(range3))

}
