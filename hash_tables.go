package main

import "fmt"

func main() {
	iMap := make(map[string]int)
	iMap["k1"] = 12
	iMap["k2"] = 13

	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
	}

	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	fmt.Println("anotherMap:", anotherMap)

	_, ok := iMap["doesItExist"]
	if ok {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Does not exist!")
	}

	for key, value := range anotherMap {
		fmt.Println(key, ":", value)
	}
	fmt.Println(iMap)

}
