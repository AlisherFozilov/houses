package houses

import "fmt"

var houses = []House{
	{
		price:              Rubles(100_000),
		distanceFromCenter: Km(10),
		district:           District{name: "Rudaki"},
	},
	{
		price:              Rubles(200_000),
		distanceFromCenter: Km(5),
		district:           District{name: "Sino"},
	},
	{
		price:              Rubles(500_000),
		distanceFromCenter: Km(2),
		district:           District{name: "Sino"},
	},
	{
		price:              Rubles(300_000),
		distanceFromCenter: Km(7),
		district:           District{name: "Shohmansur"},
	},
}

func ExampleSortByPriceAsc() {
	fmt.Println(sortByPriceAsc(houses))
	// Output:[{100000 10 {Rudaki}} {200000 5 {Sino}} {300000 7 {Shohmansur}} {500000 2 {Sino}}]
}

func ExampleSortByPriceDesc() {
	fmt.Println(sortByPriceDesc(houses))
	// Output:[{500000 2 {Sino}} {300000 7 {Shohmansur}} {200000 5 {Sino}} {100000 10 {Rudaki}}]
}

func ExampleSortByDistanceFromCenterAsc() {
	fmt.Println(sortByDistanceFromCenterAsc(houses))
	// Output:[{500000 2 {Sino}} {200000 5 {Sino}} {300000 7 {Shohmansur}} {100000 10 {Rudaki}}]
}

func ExampleSortByDistanceFromCenterDesc() {
	fmt.Println(sortByDistanceFromCenterDesc(houses))
	// Output:[{100000 10 {Rudaki}} {300000 7 {Shohmansur}} {200000 5 {Sino}} {500000 2 {Sino}}]
}

func ExampleSearchByMaxPrice_ManyResults() {
	fmt.Println(searchByMaxPrice(houses, Rubles(250000)))
	// Output:[{100000 10 {Rudaki}} {200000 5 {Sino}}]
}
func ExampleSearchByMaxPrice_NoResults() {
	fmt.Println(searchByMaxPrice(houses, Rubles(0)))
	// Output:[]
}

func ExampleSearchByIntervalPrice_ManyResults() {
	fmt.Println(searchByIntervalPrice(houses, Rubles(300000), Rubles(600000)))
	// Output: [{500000 2 {Sino}} {300000 7 {Shohmansur}}]
}
func ExampleSearchByIntervalPrice_NoResults() {
	fmt.Println(searchByIntervalPrice(houses, Rubles(0), Rubles(0)))
	// Output: []
}

func ExampleSearchByDistrict_ManyResults() {
	fmt.Println(searchByDistrict(houses, District{name: "Sino"}))
	// Output:[{200000 5 {Sino}} {500000 2 {Sino}}]
}
func ExampleSearchByDistrict_NoResults() {
	fmt.Println(searchByDistrict(houses, District{name: ""}))
	// Output:[]
}

func ExampleSearchByDistricts_ManyResults() {
	fmt.Println(searchByDistricts(houses, []District{
		{name: "Rudaki"},
		{name: "Shohmansur"},
	}))
	// Output:[{100000 10 {Rudaki}} {300000 7 {Shohmansur}}]
}
func ExampleSearchByDistricts_NoResult() {
	fmt.Println(searchByDistricts(houses, []District{
		{name: ""},
		{name: ""},
	}))
	// Output:[]
}
