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

type Joint struct {
	pos            *Vector2
	size           float32
	sqrSize        float32
	anchor         bool
	connectionIdxs []int
}

func NewJoint(x, y int, size float32, anchor bool, connectionIdxs []int) (j *Joint) {
	j = new(Joint)
	j.pos = NewVector2(x, y)
	j.size = size
	j.sqrSize = size * size
	j.anchor = anchor
	j.connectionIdxs = connectionIdxs
	return
}

type CaenorhabditisElegans struct {
	anchor *Joint
	joints []*Joint
}

func NewCE(joints []*Joint) (ce *CaenorhabditisElegans) {
	ce = new(CaenorhabditisElegans)
	ce.joints = joints
	// set ce's anchor to its anchor joint
	for _, j := range ce.joints {
		if j.anchor {
			ce.anchor = j
		}
	}
	return
}
