package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d
	)
	//notice how iota + 3 works  : have not hidden this golanD hint
	const (
		e = iota + 3
		f
		g
		h
	)
	const (
		i = iota + 3
		j
		k
		l = iota
	)
	const (
		m = iota
		n = iota
		o = iota
		p
		q = iota + 1
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	fmt.Printf("the values are %v \t %T", a, a)
	fmt.Printf("the values are %v \t %T", b, b)
	fmt.Printf("the values are %v \t %T", c, c)
	fmt.Printf("the values are %v \t %T", d, d)
	fmt.Printf("the values are %v \t %T", g, g)

}
