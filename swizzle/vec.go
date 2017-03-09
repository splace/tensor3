package tensor3

type Vector struct {
	x, y, z float64
}

var XAxis, YAxis, ZAxis = Vector{1, 0, 0}, Vector{0, 1, 0}, Vector{0, 0, 1}
var Axes = [3]Vector{XAxis, YAxis, ZAxis}

func NewVector(cs ...float64) (v Vector) {
	switch len(cs) {
	case 3:
		v.z = cs[2]
		fallthrough
	case 2:
		v.y = cs[1]
		fallthrough
	case 1:
		v.x = cs[0]
	}
	return
}

func (v *Vector) Add(v2 Vector) {
	v.x += v2.x
	v.y += v2.y
	v.z += v2.z
}

func (v *Vector) Subtract(v2 Vector) {
	v.x -= v2.x
	v.y -= v2.y
	v.z -= v2.z
}

func (v *Vector) Multiply(s float64) {
	v.x *= s
	v.y *= s
	v.z *= s
}

func (v Vector) Dot(v2 Vector) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

func (v *Vector) Cross(v2 Vector) {
	v.x, v.y, v.z = v.y*v2.z-v.z*v2.y, v.z*v2.x-v.x*v2.z, v.x*v2.y-v.y*v2.x
}

func (v Vector) LengthLength() float64 {
	return v.Dot(v)
}

func (v *Vector) Set(v2 Vector) {
	v.x = v2.x
	v.y = v2.y
	v.z = v2.z
}

func (v *Vector) Max(v2 Vector) {
	if v2.x > v.x {
		v.x = v2.x
	}
	if v2.y > v.y {
		v.y = v2.y
	}
	if v2.z > v.z {
		v.z = v2.z
	}
}

func (v *Vector) Min(v2 Vector) {
	if v2.x < v.x {
		v.x = v2.x
	}
	if v2.y < v.y {
		v.y = v2.y
	}
	if v2.z < v.z {
		v.z = v2.z
	}
}

func (v *Vector) Mid(v2 Vector) {
	v.x = (v2.x + v.x) / 2
	v.y = (v2.y + v.y) / 2
	v.z = (v2.z + v.z) / 2
}

func (v *Vector) Interpolate(v2 Vector, f float64) {
	v2.Multiply(1 - f)
	v.Multiply(f)
	v.Add(v2)
}

func interpolater(f float64) func(*Vector, Vector) {
	return func(v *Vector, v2 Vector) {
		v.Interpolate(v2, f)
	}
}

func (v *Vector) Reduce(vs Vectors, fn func(*Vector, Vector)) {
	for _, v2 := range vs {
		fn(v, v2)
	}
}

func (v *Vector) Project(axis Vector) {
	v.x *= axis.x
	v.y *= axis.y
	v.z *= axis.z
}

func (v *Vector) Product(m Matrix) {
	v.x, v.y, v.z = v.x*m.x.x+v.y*m.y.x+v.z*m.z.x, v.x*m.x.y+v.y*m.y.y+v.z*m.z.y, v.x*m.x.z+v.y*m.y.z+v.z*m.z.z
}

func (v *Vector) ProductT(m Matrix) {
	v.x, v.y, v.z = v.x*m.x.x+v.y*m.x.y+v.z*m.x.z, v.x*m.y.x+v.y*m.y.y+v.z*m.y.z, v.x*m.z.x+v.y*m.z.y+v.z*m.z.z
}
