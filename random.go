package raru

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func RandomID() (int, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(math.MaxUint16))
	if err != nil {
		return 0, fmt.Errorf("Random generator error: %s", err)
	}
	return 31337 + int(r.Int64()), nil
}
