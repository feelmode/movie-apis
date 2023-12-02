package rating

import "math"

func Round(rating float32) float32 {
	return float32(math.Round(float64(rating)*10)) / 10
}
