package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[2:]
	s1[3] = 6
	s1 = append(s1, 6)
	fmt.Println("s1", s1)
	fmt.Println("s2", s2)
	fmt.Println("Hasta el 0", s1[:0])
	fmt.Println("Desde el 1", s1[1:])
	s3 := append(s1[:0], s1[1:]...)
	fmt.Println("s1 sin el primer elemento", s1)
	fmt.Println("s3", s3)
	s1 = s1[1:]
	fmt.Println("s1 sin el primer parametro nuevament ctiene los siguientes valores:", s1)
}
