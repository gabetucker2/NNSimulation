package main

import (
	"fmt"
	"math"
)

// vectors can be *Vector2 or *Vector2Int
func SqrMagnitude(v1, v2 any) float64 {
	var x1, y1, x2, y2 float64
	if fmt.Sprintf("%T", v1) == "*main.Vector2" {
		x1 = v1.(*Vector2).x
		y1 = v1.(*Vector2).y
	} else {
		x1 = float64(v1.(*Vector2Int).x)
		y1 = float64(v1.(*Vector2Int).y)
	}
	if fmt.Sprintf("%T", v2) == "*main.Vector2" {
		x2 = v2.(*Vector2).x
		y2 = v2.(*Vector2).y
	} else {
		x2 = float64(v2.(*Vector2Int).x)
		y2 = float64(v2.(*Vector2Int).y)
	}
	xDiff := x2 - x1
	yDiff := y1 - y2
	xSqrDist := xDiff * xDiff
	ySqrDist := yDiff * yDiff
	return xSqrDist + ySqrDist
}

func m(A, B *Vector2) float64 {
	xDist := B.x - A.x
	yDist := B.y - A.y
	if xDist == 0 {
		return inf
	} else if yDist == 0 {
		return 0
	} else {
		return yDist / xDist
	}
}

func M(m float64) float64 {
	if m == 0 {
		return inf
	} else {
		return -(1 / m)
	}
}

// calculates the inverse function of fx at x for a circle
func F_Nx(x, M, r_N float64, N *Vector2) float64 {
	if M == inf {
		return N.y + r_N
	} else {
		return M*(x-N.x) + N.y
	}
}

// calculates vertex
func N_n(N, A, B *Vector2, r_N, M float64, bNotT bool) (V *Vector2) {
	mult := 0.0
	if bNotT {
		if A.y < B.y {
			mult = 1
		} else {
			mult = -1
		}
	} else {
		if A.y < B.y {
			mult = -1
		} else {
			mult = 1
		}
	}
	V_x := 0.0
	if 1+math.Pow(M,2) == 0 {
		V_x = inf
	} else {
		V_x = N.x + mult*r_N*math.Sqrt(1/(1+math.Pow(M,2)))
	}
	V = NewVector2(
		V_x,
		F_Nx(V_x, M, r_N*mult, N),
	)
	return
}

func inPolygon(polygon *Polygon, vector *Vector2) (in bool) {

	// initialize condition
	in = false

	// initialize variables
    numVertices := len(polygon.vertices)
	x := vector.x
	y := vector.y

	// iterate over pairs of adjacent vertices of the polygon (ray casting algorithm)
    j := numVertices - 1
    for i := 0; i < numVertices; i++ {
        xi := polygon.vertices[i].x
        yi := polygon.vertices[i].y
        xj := polygon.vertices[j].x
        yj := polygon.vertices[j].y
        if ((yi > y) != (yj > y)) && (x < (xj-xi)*(y-yi)/(yj-yi)+xi) {
            in = !in
        }
        j = i
    }

	// return condition
	return

}
