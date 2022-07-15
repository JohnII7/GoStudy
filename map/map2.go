package main

import "fmt"

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key:", key, " value:", value)
	}
}
func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "BeiJing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	printMap(cityMap)

	// 删除
	delete(cityMap, "China")

	// 修改
	cityMap["USA"] = "DC"

	fmt.Println("=================")
	printMap(cityMap)

}
