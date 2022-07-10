package main

import "fmt"

func main() {

	// friends
	var friends [3]string = [3]string{"Ali", "Veli", "Ayşe"}

	// distances
	distances := [5]float32{.1, .2, 10.0, 20.2, 30.5}

	//data buffer with five bytes
	bytes := [5]byte{0x0, 0x1, 0x2, 0x3, 0x4}

	// currency ratios
	ratios := [1]float64{0.02}

	//website alives
	var alives [4]bool = [4]bool{true, false, false, true}

	// empty array
	emptyArray := [0]byte{}

	// print arrayys with their types the arrays with their types.
	fmt.Printf("names    : %#v\n", friends)
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("data     : %#v\n", bytes)
	fmt.Printf("ratios   : %#v\n", ratios)
	fmt.Printf("alives   : %#v\n", alives)
	fmt.Printf("zero     : %#v\n", emptyArray)

	//types
	fmt.Printf("the type of friends : %T \n", friends)
	fmt.Printf("the type of distances : %T \n", distances)
	fmt.Printf("the type of bytes : %T \n", bytes)
	fmt.Printf("the type of ratios : %T \n", ratios)
	fmt.Printf("the type of alives : %T \n", alives)
	fmt.Printf("the type of empyarray : %T \n", emptyArray)

	// print just elements of array
	fmt.Println(alives)
	fmt.Println(emptyArray)
	fmt.Println(ratios)
	fmt.Println(bytes)
	fmt.Println(distances)
	fmt.Println(friends)
}
