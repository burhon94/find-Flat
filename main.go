package main

import (
	"sort"
)

func sortByPriceAsc(flats []flat) []flat {
	result := make([]flat, len(flats))
	copy(result, flats)
	sort.Slice(result, func(i, j int) bool {
		return result[i].priceToRuble < result[j].priceToRuble
	})
	return result
}
func sortByPriceDesc(flats []flat) []flat {
	result := make([]flat, len(flats))
	copy(result, flats)
	sort.Slice(result, func(i, j int) bool {
		return result[i].priceToRuble > result[j].priceToRuble
	})
	return result
}

func sortByRemoteOfCenterFar(flats []flat) []flat {
	result := make([]flat, len(flats))
	copy(result, flats)
	sort.Slice(result, func(i, j int) bool {
		return result[i].remoteOfCenterPerKm > result[j].remoteOfCenterPerKm
	})
	return result
}
func sortByRemoteOfCenterNear(flats []flat) []flat {
	result := make([]flat, len(flats))
	copy(result, flats)
	sort.Slice(result, func(i, j int) bool {
		return result[i].remoteOfCenterPerKm < result[j].remoteOfCenterPerKm
	})
	return result
}

func searchByMaxPrice(flats []flat, maxPrice int64) []flat {
	result := make([]flat, 0)
	for _, flat := range flats {
		if flat.priceToRuble <= maxPrice {
			result = append(result, flat)
		}
	}
	return result
}

func searchFromByMinPriceToByMaxPrice(flats []flat, minPrice, MaxPrice int64) []flat {
	result := make([]flat, 0)
	for _, flat := range flats {
		if flat.priceToRuble >= minPrice && flat.priceToRuble <= MaxPrice {
			result = append(result, flat)
		}
	}
	return result
}

func searchByDistrict(flats []flat, district string) []flat {
	result := make([]flat, 0)
	for _, flat := range flats {
		if flat.district == district {
			result = append(result, flat)
		}
	}
	return result
}

func searchByDistricts(flats []flat, districts [] string) []flat {
	result := make([]flat, 0)
	for _, flat := range flats {
		for _, district := range districts {
			if flat.district == district {
				result = append(result, flat)
				//continue
			}
		}
	}
	return result
}

type flat struct {
	id                  int64
	rooms               int64
	priceToRuble        int64
	squareOfLife        int64
	squereOfFlat        int64
	squreOfKitchen      int64
	remoteOfCenterPerKm int64
	floor               int64
	district            string
	typeFlat            string
}

func main() {
}
