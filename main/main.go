package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const pi = 3.14

func main() {

	circleRadius := 2
	printCircleArea(circleRadius)
	a := 2
	b := -4
	res, err := calculateRectangleArea(&a, &b)
	fmt.Println(err)
	fmt.Printf("Input area for %d and %d is %d\n", a, b, res)
	res, err = calculateRectanglePerimetr(a, b)
	fmt.Println(errors.Unwrap(err))
	fmt.Printf("Input perimetr for %d and %d is %d\n", a, b, res)

	var list = [...]string{"kek", "lol"}

	for index, value := range list {
		fmt.Printf("%s and %s\n", list[index], value)
	}

	var arr = []string{"ou"}
	arr = append(arr, "yee")
	fmt.Println(strings.Join(arr, "god"))
	fmt.Println(arr)

	for _, value := range arr {
		fmt.Printf("%s", value)
	}

	home := list[:]
	fmt.Println(home)
	reverteArray(home)
	fmt.Println("update array", home)
	
}

func printCircleArea(inputRadius int) {

	circleArea, err := calculateCircleArea(inputRadius)

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Printf("Радиус круга: %d см.\n", inputRadius)
	fmt.Printf("Площадь круга: %f32 см. кв.\n", circleArea)
}

func calculateCircleArea(inputRadius int) (float32, error) {
	if inputRadius >= 0 {
		return float32(inputRadius) * float32(inputRadius) * math.E, nil
	} else {
		return float32(0), errors.New("Radius couldn't be negative")
	}
}

func calculateRectangleArea(inputA *int, inputB *int) (int, error) {
	if (*inputA < 0) || (*inputB < 0) {
		return 0, errors.New("a, b must be positive")
	}
	return *inputA * *inputB, nil
}

func calculateRectanglePerimetr(inputA int, inputB int) (int, error) {
	if (inputA < 0) || (inputB < 0) {
		return 0, errors.New("a, b must be positive")
	}
	return (inputA + inputB) * 2, nil
}

func reverteArray(inpur []string) {

	tmp := make([]string, len(inpur))
	copy(tmp, inpur)
	fmt.Println(len(inpur), "len tmp ", len(tmp))
	var tmpIndex = 0

	for i := len(inpur) - 1; i >= 0; i-- {
		fmt.Println(i)
		inpur[i] = tmp[tmpIndex]
		tmpIndex++
	}

}
