package main

import (
	"fmt"
)

type arrayOps interface {
	Search(element int) int
	Sort() //inplace
}

type IntArray struct {
	arr    []int
	length int
}

func (ar IntArray) Search(el int) int {
	for i := 0; i < ar.length; i++ {
		if ar.arr[i] == el {
			return i
		}
	}
	return -1
}

func (ar IntArray) Sort() {
	for i := 0; i < ar.length; i++ {
		for j := i + 1; j < ar.length-1; j++ {
			if ar.arr[i] > ar.arr[j] {
				temp := ar.arr[i]
				ar.arr[i] = ar.arr[j]
				ar.arr[j] = temp
			}
		}
	}
}

func IsPrime(val int) bool {
	for i := 2; i <= val/2; i++ {
		if val%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	myint := IntArray{[]int{11, 22, 24}, 3}
	fmt.Println(myint.Search(24))
	myint.Sort()
	fmt.Println(myint.arr)
	//utils.Func1()
	//utils.Func2()
	//utilsnested.FuncNested()
	//a, _ := strconv.Atoi("10")
	//fmt.Println("Is 25 Prime ", IsPrime(25))
	//fmt.Println("Is 31 Prime", IsPrime(31), Triangle(2, 3, 3))
}
