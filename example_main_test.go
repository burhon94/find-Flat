package main

import "fmt"

var flats = [] flat{
	{
		1,
		4,
		4_355_900,
		39,
		47,
		5,
		50,
		4,
		"Domodedova",
		"New",
	},
	{
		2,
		2,
		1_755_900,
		26,
		29,
		3,
		42,
		2,
		"Mtishi",
		"Second",
	},
	{
		3,
		1,
		1_355_900,
		23,
		26,
		3,
		75,
		5,
		"Sheremetevo",
		"Second",
	},
	{
		4,
		3,
		3_655_900,
		23,
		28,
		5,
		51,
		4,
		"Domodedova",
		"Second",
	},
}

func ExampleSortByPriceAsc() {
	sortedByPriceAsc := sortByPriceAsc(flats)
	fmt.Println(sortedByPriceAsc)
	//Output: [{3 1 1355900 23 26 3 75 5 Sheremetevo Second} {2 2 1755900 26 29 3 42 2 Mtishi Second} {4 3 3655900 23 28 5 51 4 Domodedova Second} {1 4 4355900 39 47 5 50 4 Domodedova New}]
}
func ExampleSortByPriceDesc() {
	sortedByDesc := sortByPriceDesc(flats)
	fmt.Println(sortedByDesc)
	//Output: [{1 4 4355900 39 47 5 50 4 Domodedova New} {4 3 3655900 23 28 5 51 4 Domodedova Second} {2 2 1755900 26 29 3 42 2 Mtishi Second} {3 1 1355900 23 26 3 75 5 Sheremetevo Second}]
}

func ExampleSortByRemoteOfCenterFar() {
	sortedByRemoteFarOfCenter := sortByRemoteOfCenterFar(flats)
	fmt.Println(sortedByRemoteFarOfCenter)
	//Output: [{3 1 1355900 23 26 3 75 5 Sheremetevo Second} {4 3 3655900 23 28 5 51 4 Domodedova Second} {1 4 4355900 39 47 5 50 4 Domodedova New} {2 2 1755900 26 29 3 42 2 Mtishi Second}]
}
func ExampleSortByRemoteOfCenterNear() {
	sortedByRemoteNearOfCenter := sortByRemoteOfCenterNear(flats)
	fmt.Println(sortedByRemoteNearOfCenter)
	//Output: [{2 2 1755900 26 29 3 42 2 Mtishi Second} {1 4 4355900 39 47 5 50 4 Domodedova New} {4 3 3655900 23 28 5 51 4 Domodedova Second} {3 1 1355900 23 26 3 75 5 Sheremetevo Second}]
}

func ExampleSearchByMaxPriceNoResult() {
	result := searchByMaxPrice(flats, 500_000)
	fmt.Println(result)
	//Output: []
}
func ExampleSearchByMaxPriceOne() {
	result := searchByMaxPrice(flats, 1_500_000)
	fmt.Println(result)
	//Output: [{3 1 1355900 23 26 3 75 5 Sheremetevo Second}]
}
func ExampleSearchByMaxPriceManyResult() {
	result := searchByMaxPrice(flats, 2_000_000)
	result = sortByPriceAsc(result)
	fmt.Println(result)
	//Output: [{3 1 1355900 23 26 3 75 5 Sheremetevo Second} {2 2 1755900 26 29 3 42 2 Mtishi Second}]
}

func ExampleSearchFromMinPriceToMaxPriceNoResult() {
	result := searchFromByMinPriceToByMaxPrice(flats, 1_000_000, 1_200_000)
	fmt.Println(result)
	//Output: []
}
func ExampleSearchFromMinPriceToMaxPriceOneResult() {
	result := searchFromByMinPriceToByMaxPrice(flats, 1_000_000, 1_500_000)
	fmt.Println(result)
	//Output: [{3 1 1355900 23 26 3 75 5 Sheremetevo Second}]
}
func ExampleSearchFromMinPriceToMaxPriceManyResult() {
	result := searchFromByMinPriceToByMaxPrice(flats, 1_000_000, 2_500_000)
	result = sortByPriceDesc(result)
	fmt.Println(result)
	//Output: [{2 2 1755900 26 29 3 42 2 Mtishi Second} {3 1 1355900 23 26 3 75 5 Sheremetevo Second}]
}

func ExampleSearchByDistrictNoResult() {
	result := searchByDistrict(flats, "Vnukova")
	fmt.Println(result)
	//Output: []
}
func ExampleSearchByDistrictOne() {
	result := searchByDistrict(flats, "Sheremetevo")
	fmt.Println(result)
	//Output: [{3 1 1355900 23 26 3 75 5 Sheremetevo Second}]
}
func ExampleSearchByDistrictManyResult() {
	result := searchByDistrict(flats, "Domodedova")
	fmt.Println(result)
	//Output: [{1 4 4355900 39 47 5 50 4 Domodedova New} {4 3 3655900 23 28 5 51 4 Domodedova Second}]
}

func ExampleSearchByDistrictsNoResult() {
	result := searchByDistricts(flats, []string{"Vnukova", "Izmailova"})
	fmt.Println(result)
	//Output: []
}
func ExampleSearchByDistrictsOne() {
	result := searchByDistricts(flats, []string{"Sheremetevo", "Izmailova"})
	fmt.Println(result)
	//Output: [{3 1 1355900 23 26 3 75 5 Sheremetevo Second}]
}
func ExampleSearchByDistrictsManyResult() {
	result := searchByDistricts(flats, []string{"Sheremetevo", "Domodedova"})
	fmt.Println(result)
	//Output: [{1 4 4355900 39 47 5 50 4 Domodedova New} {3 1 1355900 23 26 3 75 5 Sheremetevo Second} {4 3 3655900 23 28 5 51 4 Domodedova Second}]
}
