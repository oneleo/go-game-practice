package main

import "fmt"

func main() {
	// String slice, and func call by pointer
	var a = []string{"1", "2", "3"}
	//changeSliceA(&a)
	changeSliceD(a)

	// String pointer slice, and func call by value
	s1 := "1"
	s2 := "2"
	//s3 := "3"
	var b = []*string{&s1, &s2, new(string)}
	changeSliceB(b)

	// String pointer slice, and func call by pointer
	c := make([]*string, len(b))
	copy(c, b)
	changeSliceC(&c)

	// Result
	fmt.Print("a at ", fmt.Sprintf("%p", &a), "\ta = ", a, "\n")
	fmt.Print("b at ", fmt.Sprintf("%p", &b), "\tb = ", b, "\n")
	fmt.Print("c at ", fmt.Sprintf("%p", &c), "\tc = ", c, "\n")
}

func changeSliceA(i *([]string)) {
	(*i)[0] = "3"
	*i = append(*i, "4")
}

func changeSliceB(i [](*string)) {
	*(i[0]) = "3"
	s := "4"
	i = append(i, &s)
}

func changeSliceC(i *([](*string))) {
	*((*i)[0]) = "3"
	s := "4"
	*i = append(*i, &s)
}

func changeSliceD(i []string) {
	i[0] = "3"
	i = append(i, "4")
}
