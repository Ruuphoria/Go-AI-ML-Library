package util

import (
	"math/rand"
	"time"
)

func randomFloat(r *rand.Rand, min, max float64) float64 {
	return r.Float64()*(max-min) + min
}

func randomFloat32(r *rand.Rand, min, ma