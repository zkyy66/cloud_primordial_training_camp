/**
 * @Date 2022/4/16
 * @Name 1.1
 * @VariableName
**/
/**
给定一个字符串数组
[“I”,“am”,“stupid”,“and”,“weak”]
用 for 循环遍历该数组并修改为
[“I”,“am”,“smart”,“and”,“strong”]
*/
package module1

import "fmt"

const SMART = "smart"
const STRONG = "strong"

//for循环
func HandleString(str []string) []string {
	if len(str) < 0 {
		return str
	}
	for index, val := range str {
		if val == "stupid" {
			str[index] = SMART
		}
		if val == "weak" {
			str[index] = STRONG
		}
	}
	return str
}

func HandleSliceString(str []string) []string {
	if len(str) < 0 {
		return str
	}
	frontSlice := str[0:2]
	frontSlice = append(frontSlice, "smart")
	str[4] = STRONG
	backSlice := str[len(str)-2:]
	frontSlice = append(frontSlice, backSlice...)
	return frontSlice
}
func HandleSliceStringTwo(str []string) []string {
	for index, c := range str {
		if string(c) == "stupid" {
			str[index] = "smart"
		} else if string(c) == "weak" {
			str[index] = "strong"
		}
	}
	fmt.Println("HandleSliceStringTwo")
	return str
}
