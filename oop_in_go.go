package main

import (
	"fmt"
	"image/color"
	"io"
	"strings"
	"unicode"
)

type ColoredPoint struct {
	color.Color
	x, y int
}

type SomeWriter interface {
	Write([]byte) (int, error)
}

type Count int

func (count *Count) Increment()   { *count++ }
func (count *Count) Decrement()   { *count-- }
func (count *Count) IsZero() bool { return *count == 0 }

type StringMap map[string]string
type FloatChan chan float64

type RuneForRuneFunc func(rune) rune

type Part struct {
	Id   int
	Name string
}

func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}

func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}

func (part Part) String() string {
	return fmt.Sprintf("%d %q", part.Id, part.Name)
}

func (part Part) HasPrefix(prefix string) (bool, string) {
	return strings.HasPrefix(part.Name, prefix), part.String()
}

type Item struct {
	id       string
	price    float64
	quantity int
}

func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}

type SpecialItem struct {
	Item
	catalogId int
}

func (specialItem *SpecialItem) Cost() float64 {
	return (specialItem.price * float64(specialItem.quantity)) * 1000
}

type LuxuryItem struct {
	Item
	markup float64
}

func (item *LuxuryItem) Cost() float64 {
	return item.Item.Cost() * item.markup
}

type Place struct {
	latitude, longitude float64
	Name                string
}

type Exchanger interface {
	Exchange()
}

type StringPair struct {
	first, second string
}

func (pair *StringPair) Exchange() {
	pair.first, pair.second = pair.second, pair.first
}

type Point [2]int

func (point *Point) Exchange() {
	point[0], point[1] = point[1], point[0]
}

func (pair StringPair) String() string {
	return fmt.Sprintf("%q + %q", pair.first, pair.second)
}

func exchageThese(exchangers ...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

func (pair *StringPair) Read(data []byte) (n int, err error) {
	if pair.first == "" && pair.second == "" {
		return 0, io.EOF
	}

	if pair.first != "" {
		n = copy(data, pair.first)
		pair.first = pair.first[n:]
	}

	if n < len(data) && pair.second != "" {
		m := copy(data[n:], pair.second)
		pair.second = pair.second[m:]
		n += m
	}

	return n, nil
}

type LowerCaser interface {
	LowerCase()
}

type UpperCaser interface {
	UpperCase()
}

type LowerUpperCaser interface {
	LowerCaser
	UpperCaser
}

type FixCaser interface {
	FixCase()
}

type ChangeCaser interface {
	LowerUpperCaser
	FixCaser
}

//func (part *Part) FixCase() {
//	part.Name = fixCase(part.Name)
//}

//func fixCase(s string) string {
//	var chars []rune
//	upper := true
//	for _, char := range chars {
//
//	}
//
//
//	return s
//
//}

type Integer int

type Person struct {
	Title     string
	Forenames []string
	Surname   string
}

type Author1 struct {
	Names    Person
	Title    []string
	YearBorn int
}

type Author2 struct {
	Person,
	Title []string,
	YearBorn int
}

type Tasks struct {
	slice []string
	Count
}

func (tasks *Tasks) Add(task string) {
	tasks.slice = append(tasks.slice, task)
	tasks.Increment()
}

func (tasks *Tasks) Tally() int {
	return int(tasks.Count)
}

func main() {
	fmt.Println("=========================================================")
	//=========================================================================

	author1 := Author1{
		Person{
			"Mr",
			[]string{"Robert", "Louis", "Balfour"},
			"Stevenson",
		},
		[]string{"Kidnapped", "Treasure Island"},
		1850,
	}

	fmt.Println(author1)

	author1.Names.Title = ""
	author1.Names.Forenames = []string{"Oscar", "Fingal", "O'Flahertie"}
	author1.Names.Surname = "Wilde"

	author1.Title = []string{
		"The picture of Dorian Gray",
	}
	author1.YearBorn += 4

	fmt.Println(author1)

	pointsStruct := []struct{ x, y int }{{4, 6}, {}, {-7, 11}, {15, 17}, {14, -8}}

	for _, point := range pointsStruct {
		fmt.Printf("(%d, %d) ", point.x, point.y)
	}

	fmt.Println()

	var points = [][2]int{{4, 6}, {}, {-7, 11}, {15, 17}, {14, -8}}

	for _, point := range points {
		fmt.Printf("(%d, %d) ", point[0], point[1])
	}

	fmt.Println()

	const size = 16

	robert := &StringPair{
		"Robert L.",
		"Stevenson",
	}

	david := &StringPair{
		"David",
		"Balfour",
	}

	fmt.Println(robert, david)

	jekyll := StringPair{
		"Henry",
		"Jekyll",
	}
	hyde := StringPair{
		"Edward",
		"Hyde",
	}

	point := Point{5, -3}

	fmt.Println("Before", jekyll, hyde, point)

	jekyll.Exchange()
	hyde.Exchange()
	point.Exchange()

	fmt.Println("After #1", jekyll, hyde, point)

	exchageThese(&jekyll, &hyde, &point)

	fmt.Println("After #2", jekyll, hyde, point)

	someString := StringPair{
		"First String",
		"Second String",
	}

	fmt.Println("someString ->", someString)

	someString.Exchange()

	fmt.Println("someString ->", someString)

	special := SpecialItem{
		Item{
			"Green",
			3,
			5,
		},
		207,
	}

	fmt.Println(special.Cost(), special.Item.Cost(), special.catalogId,
		special.price)

	part := Part{
		5,
		"wrench",
	}

	part.UpperCase()
	fmt.Printf("%+v \n", part)
	fmt.Println(part)
	part.LowerCase()
	fmt.Printf("%+v \n", part)
	fmt.Println(part)
	part.Id += 11

	hasPrefix, toString := part.HasPrefix("w")
	fmt.Println("hasPrefix, toString ->", hasPrefix, toString)

	asStringV := Part.String
	sv := asStringV(part)

	fmt.Println("sv->", sv)

	var count Count
	count.Increment()
	count.Decrement()
	fmt.Println(count.IsZero(), count)

	var i Count = 7
	i++
	fmt.Println(i)
	sm := make(StringMap)
	sm["key1"] = "value1"
	sm["key2"] = "value2"

	fmt.Println(sm)
	fc := make(FloatChan, 1)
	fc <- 2.2977889999
	fmt.Println(<-fc)

	var removePunctuation RuneForRuneFunc

	phrases := []string{"Day; dusk, and night.", "All day long"}

	removePunctuation = func(char rune) rune {
		if unicode.Is(unicode.Terminal_Punctuation, char) {
			return -1
		}

		return char
	}

	processPhrases := func(phrases []string, function RuneForRuneFunc) {
		for _, phrase := range phrases {
			fmt.Println(strings.Map(function, phrase))
		}
	}

	processPhrases(phrases, removePunctuation)

}
