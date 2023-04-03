package main

import (
	"fmt"
	"math"
)

type Vector2 struct {
	x float64
	y float64
}

// x and y can be floats or ints
func NewVector2(x, y any) (v2 *Vector2) {
	v2 = new(Vector2)
	if fmt.Sprintf("%T", x) == "float64" {
		v2.x = x.(float64)
		v2.y = y.(float64)
	} else {
		v2.x = float64(x.(int))
		v2.y = float64(y.(int))
	}
	return
}

type Vector2Int struct {
	x int
	y int
}

// x and y can be floats or ints
func NewVector2Int(x, y any) (v2 *Vector2Int) {
	v2 = new(Vector2Int)
	if fmt.Sprintf("%T", x) == "int" {
		v2.x = x.(int)
		v2.y = y.(int)
	} else {
		v2.x = int(math.Round(x.(float64)))
		v2.y = int(math.Round(y.(float64)))
	}
	return
}

type Color struct {
	r uint8
	g uint8
	b uint8
}

func NewColor(r, g, b uint8) (c *Color) {
	c = new(Color)
	c.r = r
	c.g = g
	c.b = b
	return
}

type Pixel struct {
	pos *Vector2Int
	col *Color
}

func NewPixel(x, y int, col *Color) (p *Pixel) {
	p = new(Pixel)
	p.pos = NewVector2Int(x, y)
	p.col = col
	return
}

type Effector struct {
	pos            *Vector2
	radius         float64
	sqrRadius      float64
	anchor         bool
	connectionIdxs []int
}

func NewEffector(x, y, radius float64, anchor bool, connectionIdxs []int) (e *Effector) {
	e = new(Effector)
	e.pos = NewVector2(x, y)
	e.radius = radius
	e.sqrRadius = radius * radius
	e.anchor = anchor
	e.connectionIdxs = connectionIdxs
	return
}

type Polygon struct {
	vertices []*Vector2
}

func MakePolygon(vertices []*Vector2) (p *Polygon) {
	p = new(Polygon)
	p.vertices = vertices
	return
}

type CaenorhabditisElegans struct {
	anchor    *Effector
	anchorIdx int
	effectors []*Effector
}

func NewCE(effectors []*Effector) (ce *CaenorhabditisElegans) {
	ce = new(CaenorhabditisElegans)
	ce.effectors = effectors
	// set ce's anchor to its anchor effector
	for i, e := range ce.effectors {
		if e.anchor {
			ce.anchor = e
			ce.anchorIdx = i
		}
	}
	return
}
