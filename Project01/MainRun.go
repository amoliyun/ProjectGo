package main

import "fmt"

// PersonInfo 是人员信息
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func ifFuncDeclar(x int) int {
	if 0 == x {
		return 1
	} else {
		return 2
	}
}

func main() {
	fmt.Println("开始了开始了")

	fmt.Println("数组")

	var name [5]string
	name[0] = "11"

	values := [3]int{3, 4, 5}

	fmt.Println("values之前:", values)

	slice1 := values[:]

	slice2 := make([]int, 1, 2)

	slice2 = append(slice2, 444)
	slice1 = append(slice1, 444)

	slice1[0] = 111
	slice1[2] = 11111

	fmt.Println("slice2:", slice2)
	fmt.Println("slice1:", slice1)
	fmt.Println("values之后:", values)

	copy(slice2, slice1)
	fmt.Println("slice2拷贝后:", slice2)

	copy(slice1, slice2)
	fmt.Println("slice1拷贝后:", slice1)

	fmt.Println("")
	fmt.Println("map")
	var persons map[string]PersonInfo
	persons = make(map[string]PersonInfo)

	persons["111"] = PersonInfo{"111", "asd", "asdasd"}
	persons["222"] = PersonInfo{"222", "asd", "asdasd"}

	person, ok := persons["111"]
	if ok {
		fmt.Println(person)
	}

	myPersons := map[string]PersonInfo{
		"111": PersonInfo{"111", "asdasdasd", "323232323"},
	}
	fmt.Println("myPersons:", myPersons)

	fmt.Println("")
	fmt.Println("if")
	fmt.Println("ifFuncDeclar：", ifFuncDeclar(2))
}
