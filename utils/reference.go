package utils

import (
	"math/rand"
	"strconv"
)

func GenerateReference() string {
	return "trip_" + strconv.FormatInt(rand.Int63n(100000*100000), 10)
}
