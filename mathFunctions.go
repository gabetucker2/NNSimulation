package main

func SqrDistance(v1, v2 *Vector2) float32 {
	// square distance comparison is far more optimized than dist comparison since we aren't taking the square root of 2
	xSqrDist := (v1.X-v2.X) * (v1.X-v2.X)
	ySqrDist := (v1.Y-v2.Y) * (v1.Y-v2.Y)
	return float32(xSqrDist) + float32(ySqrDist)
}
