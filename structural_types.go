package main

import (
	"fmt"
	"image/color"
)

func swapAndProduct(x, y, product *int) {
	if *x > *y {
		*x, *y = *y, *x

	}
	*product = *x * *y
}

func swapAndProduct2(x, y int) (int, int, int) {
	if x > y {
		x, y = y, x
	}
	product := x * y
	return x, y, product
}

type composer struct {
	name      string
	birthYear int
}

func inflate(numbers []int, factor int) {
	for i := range numbers {
		numbers[i] *= factor
	}
}

func inflate2(numbers []int, factor int) {
	for _, number := range numbers {
		number *= factor
	}
}

type rectangle struct {
	x0, y0, x1, y1 int
	fill           color.RGBA
}

func resizeRect(rect *rectangle, width, height int) {
	(*rect).x1 += width
	rect.y1 += height
}
func main() {
	fmt.Println("=====================================================")

	var buffer [30]byte
	var grid [3][3]int
	grid[1][0], grid[1][1], grid[1][2] = 8, 6, 2

	grid2 := [3][3]int{{4, 3}, {8, 6, 2}}

	fmt.Println(grid, grid2, buffer)

	cities := [...]string{"Shnghai", "Mumbai", "Istanbul", "Beijing"}
	cities[len(cities)-1] = "Karachi"

	fmt.Println("Type    Len Contents")
	fmt.Printf("buffer -> %-8T %2d %v\n", buffer, len(buffer), buffer)
	fmt.Printf("cities -> %-8T %d %v\n", cities, len(cities), cities)
	fmt.Printf("grid -> %-8T %d %v\n", grid, len(grid), grid)
	fmt.Printf("grid2 -> %-8T %d %[1]v\n", grid2, len(grid2))
	fmt.Printf("grid2 -> %-8T %d %v\n", grid2, len(grid2), grid2)
	
	rect := rectangle{4, 8, 20, 10, color.RGBA{0xFF, 0, 0, 0xFF}}
	fmt.Println(rect)
	resizeRect(&rect, 10, 10)
	fmt.Println(rect)

	grades := []int{87, 55, 43, 71, 60, 43, 32, 19}
	inflate(grades, 3)
	fmt.Println(grades)

	inflate2(grades, 3)
	fmt.Println(grades)

	antonio := composer{
		birthYear: 1707,
		name:      "Antonio Teixeira",
	}
	agnes := new(composer)
	agnes.name, agnes.birthYear = "Agnes Zimmermann", 1845

	julia := &composer{}
	julia.name, julia.birthYear = "Julia Ward Howe", 1819

	augusta := &composer{"Augusta Holmes", 1847}

	fmt.Println("antonio -> ", antonio)
	fmt.Println("other composers -> ", agnes, julia, augusta)

	z := 37
	pi := &z
	ppi := &pi
	fmt.Println(z, pi, ppi)
	fmt.Println(z, *pi, **ppi)

	**ppi++
	fmt.Println(z, pi, ppi)

	i := 9
	j := 5
	product := 0
	swapAndProduct(&i, &j, &product)
	fmt.Println(i, j, product)

	i, j, product = swapAndProduct2(i, j)
	fmt.Println(i, j, product)

}
