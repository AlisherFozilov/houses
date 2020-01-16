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

func sortBy(houses []House, compare func(a, b House) bool) []House {
	result := make([]House, len(houses))
	copy(result, houses)

	sort.Slice(result, func(i, j int) bool {
		return compare(result[i], result[j])
	})
	return result
}

func sortByPriceAsc(houses []House) []House {
	return sortBy(houses, func(a, b House) bool {
		return a.price < b.price
	})
}
func sortByPriceDesc(houses []House) []House {
	return sortBy(houses, func(a, b House) bool {
		return a.price > b.price
	})
}

func sortByDistanceFromCenterAsc(houses []House) []House {
	return sortBy(houses, func(a, b House) bool {
		return a.distanceFromCenter < b.distanceFromCenter
	})
}

func sortByDistanceFromCenterDesc(houses []House) []House {
	return sortBy(houses, func(a, b House) bool {
		return a.distanceFromCenter > b.distanceFromCenter
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
				break
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
