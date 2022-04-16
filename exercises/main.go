package main

import (
	"cloud_primordial_training_camp/exercises/module1"
	"fmt"
)

func main() {
	fmt.Println("start 1.1")
	handleStr := []string{"I", "am", "stupid", "and", "weak"}
	tmpRes := module1.HandleString(handleStr)
	fmt.Println(tmpRes)

	res := module1.HandleSliceString(handleStr)
	fmt.Println(res)

	result := module1.HandleSliceStringTwo(handleStr)
	fmt.Println("two:", result)
	fmt.Println("end 1.1")
}
