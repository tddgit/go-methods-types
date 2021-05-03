package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//ioFiles()
	//stringsPackage()

}

func maps() {
	{
		var x [5]int
		x[4] = 1000
		fmt.Println(x, x[4])

		var y [5]float64
		y[0] = 98
		y[1] = 93
		y[2] = 77
		y[3] = 82
		y[4] = 83

		var total float64
		for i := 0; i < 5; i++ {
			total += y[i]
		}
		fmt.Println(total / 5)

		var total1 float64 = 0
		for _, value := range y {
			total1 += value
		}
		fmt.Println(total / float64(len(y)))

		z := [5]float64{98, 93, 77, 68, 34}
		fmt.Println(z)

	}
	{
		var x []float64
		x = make([]float64, 8)
		y := make([]int, 2)
		
		fmt.Println(x, y)
	}

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y[1])

	delete(y, 1)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"

	fmt.Println(elements["H"])

	name, ok := elements["Li"]
	if ok {
		fmt.Println(name, ok)
	}

	elems := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
	}

	if el, ok := elems["H"]; ok {
		fmt.Println(el["name"], el["state"])
	}

}

func ioFiles() {
	file, err := os.Open("log.txt")
	if err != nil {
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Println(stat.IsDir(), stat.ModTime())

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)

	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func stringsPackage() {

	fmt.Println(strings.Contains("test", "st"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "t"))
	fmt.Println(strings.HasSuffix("test", "t"))

	fmt.Println(strings.Index("test", "e"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))

	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	fmt.Println(strings.ToUpper("test"))
	fmt.Println(strings.ToLower("Test"))
}
