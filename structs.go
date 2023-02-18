package main

type Vector2 struct {
	X int
	Y int
}

func NewVector2(x, y int) (v2 *Vector2) {
	v2 = new(Vector2)
	v2.X = x
	v2.Y = y
	return
}

type Vector3 struct {
	X int
	Y int
	Z int
}

func NewVector3(x, y, z int) (v3 *Vector3) {
	v3 = new(Vector3)
	v3.X = x
	v3.Y = y
	v3.Z = z
	return
}

type Color struct {
	R uint8
	G uint8
	B uint8
}

func NewColor(r, g, b uint8) (c *Color) {
	c = new(Color)
	c.R = r
	c.G = g
	c.B = b
	return
}

type Pixel struct {
	pos *Vector2
	col *Color
}

func NewPixel(x, y int, col *Color) (p *Pixel) {
	p = new(Pixel)
	p.pos = NewVector2(x, y)
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

func NewEffector(x, y int, radius float64, anchor bool, connectionIdxs []int) (e *Effector) {
	e = new(Effector)
	e.pos = NewVector2(x, y)
	e.radius = radius
	e.sqrRadius = radius * radius
	e.anchor = anchor
	e.connectionIdxs = connectionIdxs
	return
}

type Polygon struct {
	points []*Vector2
}

func MakePolygon(points []*Vector2) (p *Polygon) {
	p = new(Polygon)
	p.points = points
	return
}

type CaenorhabditisElegans struct {
	anchor    *Effector
	effectors []*Effector
}

func NewCE(effectors []*Effector) (ce *CaenorhabditisElegans) {
	ce = new(CaenorhabditisElegans)
	ce.effectors = effectors
	// set ce's anchor to its anchor effector
	for _, j := range ce.effectors {
		if j.anchor {
			ce.anchor = j
		}
	}
	return
}
