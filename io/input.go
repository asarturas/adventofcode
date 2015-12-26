package io

import (
	"fmt"
	"bufio"
	"os"
)

func ReadOneInt() int {
	var a int
	fmt.Scanf("%v", &a)
	return a
}

func ReadTwoInts() (int, int) {
	var a, b int
	fmt.Scanf("%v %v", &a, &b)
	return a, b
}

func ReadThreeInts() (int, int, int) {
	var a, b, c int
	fmt.Scanf("%v %v %v", &a, &b, &c)
	return a, b, c
}

func ReadManyInts(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ints[i])
	}
	return ints
}

func ReadMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = ReadManyInts(n)
	}
	return matrix
}

func ReadWord() string {
	var a string
	fmt.Scanf("%v", &a)
	return a
}

var reader = bufio.NewReader(os.Stdin)

func ReadLine() string {
	str, _, _ := reader.ReadLine();
	return string(str);
}
