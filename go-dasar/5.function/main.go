package main

import (
	rand "math/rand"
	randV2 "math/rand/v2"
	time "time"
)

var randomizer = rand.New(rand.NewSource((time.Now().Unix())))

func main() {

	var randomValue int

	randomValue = randomWithRangeV1(2, 10)
	println("random number : ", randomValue)

	randomValue = randomWithRangeV1(2, 10)
	println("random number : ", randomValue)

	randomValue = randomWithRangeV2(2, 10)
	println("random number : ", randomValue)

}

func randomWithRangeV1(min, max int) int {
	// var value = randomizer.Int()%(max-min+1) + min
	// var value = randomizer.Int()%(max-min+1) + min
	// return value
	return randomizer.Intn(max-min+1) + min

}

func randomWithRangeV2(min, max int) int {
	// Calculate the size of the range
	rangeSize := max - min + 1

	// Use rand.IntN to get a value in [0, rangeSize)
	// Then add min to shift it to the correct range [min, max]
	var value = randV2.IntN(rangeSize) + min
	return value
}
