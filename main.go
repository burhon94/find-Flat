package main

import (
	"sort"
)

type flat struct {
	id                  int64
	rooms               int64
	priceToRuble        int64
	squareOfLife        int64
	squareOfFlat        int64
	squareOfKitchen     int64
	remoteOfCenterPerKm int64
	floor               int64
	district            string
	typeFlat            string
}

func sortByPricesAscAndDesc(flats []flat) (resultForAsc, resultForDesc []flat) {
	resultForAsc = make([]flat, len(flats))
	copy(resultForAsc, flats)
	resultForDesc = make([]flat, len(flats))
	copy(resultForDesc, flats)
	sort.Slice(resultForAsc, func(i, j int) bool {
		return resultForAsc[i].priceToRuble < resultForAsc[j].priceToRuble
	})
	sort.Slice(resultForDesc, func(i, j int) bool {
		return resultForDesc[i].priceToRuble > resultForDesc[j].priceToRuble
	})
	return
}

func sortByRemoteOfCenterNearAndFar(flats []flat) (resultForNear, resultForFar []flat) {
	resultForNear = make([]flat, len(flats))
	copy(resultForNear, flats)
	resultForFar = make([]flat, len(flats))
	copy(resultForFar, flats)
	sort.Slice(resultForNear, func(i, j int) bool {
		return resultForNear[i].remoteOfCenterPerKm < resultForNear[j].remoteOfCenterPerKm
	})
	sort.Slice(resultForFar, func(i, j int) bool {
		return resultForFar[i].remoteOfCenterPerKm > resultForFar[j].remoteOfCenterPerKm
	})
	return
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
			}
		}
	}
	return result
}

func main() {
}
