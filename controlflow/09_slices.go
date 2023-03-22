package main

import "fmt"

func main() {

	//we use composite literals for slices
	//Composite literals are used to construct the values for arrays, structs, slices, and maps.
	//Each time they are evaluated, it will create new value.
	//They consist of the type of the literal followed by a brace-bound list of elements.

	//x is a SLICE
	x := []int{0, 1, 2, 3, 4}
	//loop through slice
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
	fmt.Println("using range")

	//loop through using range, index, value
	for i, v := range x {
		fmt.Println(i, v)
	}

	//slicing a slice
	fmt.Println("using slice")
	fmt.Println(x[:])
	fmt.Println(x[:len(x)])

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i:], "\t ", x[:i])
	}

	//Appending a slice
	// we use append which is a function : func
	//append(slice []T, elements ...T) []T
	//where T is a placeholder for any given type. You can't actually write a function in Go where the type T is determined by the caller.
	//That's why append is built in: it needs support from the compiler.
	//its under
	fmt.Println("Before append: \n ", x)
	x = append(x, 10, 11, 12)
	y := []int{13, 14, 15}
	x = append(x, y...) // check how ...y is used

	fmt.Println("After append: \n ", x)

	//deleting from the slice  , delete the 3rd element
	x = append(x[:3], x[4:]...)
	fmt.Println("After deleting the 3rd element : \n ", x)

	//slices's underlying data structure is arrays
	//So , evertime you add , delete elements the runtime is creating new arrays & scraping old ones .
	//this requires processing . So in case you know the capacity you need to work with , you can use
	//make function : make ([]T, length , capacity)
	z := make([]int, 10, 12)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z = append(z, 1) // z[10]
	z[1] = 2
	z[2] = 3
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	// till this time processing is saved since new arrays are not made or deleted
	z = append(z, 2) //z[11]
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	// after this new array will be created N old destroyed because capacity is over, note length is +1 , capacity is *2
	z = append(z, 3) //z[11]
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
}
