package consinesimilarity

import "math"

func Cal(v1, v2 []float64) float64 {
	dotProduct, normV1, normV2 := .0, .0, .0
	for i := range v1 {
		dotProduct += v1[i] * v2[i]
		normV1 += v1[i] * v1[i]
		normV2 += v2[i] * v2[i]
	}
	return dotProduct / (float64(math.Sqrt(float64(normV1))) * float64(math.Sqrt(float64(normV2))))
}