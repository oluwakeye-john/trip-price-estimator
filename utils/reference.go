package utils

import (
	"math/rand"
	"strconv"
)

func GenerateReference() string {
	return "hyperdrive_trip_" + strconv.FormatInt(rand.Int63n(1000000*1000000), 10)
}
