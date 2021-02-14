package main

import "fmt"

func main() {
	const constName string = "minsu" // 상수 with 타입 선언
	var variableNameOne string = "minsu" // 변수 with 타입 선언
	variableNameTwo := "minsu" // 축약된 변수 선언 with 타입 추론 only 함수 안에서만

	fmt.Println(constName)
	fmt.Println(variableNameOne)
	fmt.Println(variableNameTwo)
}