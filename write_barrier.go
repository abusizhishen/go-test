package main

type Test struct {
	Val *int
}

var a,b Test

func main() {
	var c int
	simple(&c)
}

func simple(c *int)  {
	b.Val = c
	a.Val = b.Val
	b.Val = nil
}
