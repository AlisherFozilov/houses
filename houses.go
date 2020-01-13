package houses

import (
	"sort"
)

type Rubles int64
type Km uint
type District struct {
	name string
}

type House struct {
	price              Rubles
	distanceFromCenter Km
	district           District
}

/*type HouseSorter struct {
	houses []House
	sorter func(i, j int) bool
}*/
var result []House

func sortBy(houses []House, compare func(i, j int) bool) []House {
	result = make([]House, len(houses))
	copy(result, houses)

	sort.Slice(result, compare)
	return result
}

func sortByPriceAsc(houses []House) []House {
	return sortBy(houses, func(i, j int) bool {
		return result[i].price < result[j].price
	})
}
func sortByPriceDesc(houses []House) []House {
	return sortBy(houses, func(i, j int) bool {
		return result[i].price > result[j].price
	})
}

func sortByDistanceFromCenterAsc(houses []House) []House {
	return sortBy(houses, func(i, j int) bool {
		return result[i].distanceFromCenter < result[j].distanceFromCenter
	})
}

func sortByDistanceFromCenterDesc(houses []House) []House {
	return sortBy(houses, func(i, j int) bool {
		return result[i].distanceFromCenter > result[j].distanceFromCenter
	})
}

func searchByMaxPrice(houses []House, priceLimit Rubles) []House {
	tmpResult := make([]House, 0, len(houses))

	for _, house := range houses {
		if house.price < priceLimit {
			tmpResult = append(tmpResult, house)
		}
	}
	if len(tmpResult) == 0 {
		return make([]House, 0)
	}

	result := make([]House, len(tmpResult))
	copy(result, tmpResult)
	return result
}

func searchByIntervalPrice(houses []House, lowerPriceLimit, upperPriceLimit Rubles) []House {
	tmpResult := make([]House, 0, len(houses))

	for _, house := range houses {
		if house.price < lowerPriceLimit {
			continue
		}
		if house.price <= upperPriceLimit {
			tmpResult = append(tmpResult, house)
		}
	}
	if len(tmpResult) == 0 {
		return make([]House, 0)
	}

	result := make([]House, len(tmpResult))
	copy(result, tmpResult)
	return result
}

func searchByDistrict(houses []House, requireDistrict District) []House {
	tmpResult := make([]House, 0, len(houses))

	for _, house := range houses {
		if house.district == requireDistrict {
			tmpResult = append(tmpResult, house)
		}
	}
	if len(tmpResult) == 0 {
		return make([]House, 0)
	}

	result := make([]House, len(tmpResult))
	copy(result, tmpResult)
	return result
}

func searchByDistricts(houses []House, requireDistricts []District) []House {
	tmpResult := make([]House, 0, len(houses))

	for _, house := range houses {
		for _, requireDistrict := range requireDistricts {
			if house.district == requireDistrict {
				tmpResult = append(tmpResult, house)
			}
		}
	}
	if len(tmpResult) == 0 {
		return make([]House, 0)
	}

	result := make([]House, len(tmpResult))
	copy(result, tmpResult)
	return result
}
