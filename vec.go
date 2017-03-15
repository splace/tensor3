package tensor3

type Vector struct {
	x, y, z BaseType
}

func (v Vector) Components() (BaseType, BaseType, BaseType) {
	return v.x, v.y, v.z
}

var xAxis, yAxis, zAxis = Vector{1, 0, 0}, Vector{0, 1, 0}, Vector{0, 0, 1}
var yzPlane, xzPlane, zxPlane = Vector{0, 1, 1}, Vector{1, 0, 1}, Vector{1, 1, 0}

type Axis uint

const ( 
	xAxisIndex Axis = iota 
	yAxisIndex
	zAxisIndex
)

var Axes = [3]Vector{xAxis, yAxis, zAxis}
var AxisPlanes = [3]Vector{yzPlane, xzPlane, zxPlane}

// missing components default to zero, more than 3 are ignored
func NewVector(cs ...BaseType) (v Vector) {
	switch len(cs) {
	default:
		v.z = cs[2]
		fallthrough
	case 2:
		v.y = cs[1]
		fallthrough
	case 1:
		v.x = cs[0]
	case 0:
		return
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

// components independently multipled, see Product for matrix multiplication
func (v *Vector) Multiply(s BaseType) {
	v.x *= s
	v.y *= s
	v.z *= s
}

func (v Vector) Dot(v2 Vector) BaseType {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

func (v *Vector) Cross(v2 Vector) {
	v.x, v.y, v.z = v.y*v2.z-v.z*v2.y, v.z*v2.x-v.x*v2.z, v.x*v2.y-v.y*v2.x
}

// length squared. (returning squared means this package is not dependent on math package.)
func (v Vector) LengthLength() BaseType {
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
	v.x = (v2.x + v.x) * 0.5
	v.y = (v2.y + v.y) * 0.5
	v.z = (v2.z + v.z) * 0.5
}

func (v *Vector) Interpolate(v2 Vector, f BaseType) {
	v2.Multiply(1 - f)
	v.Multiply(f)
	v.Add(v2)
}

func (v *Vector) Reduce(vs Vectors, fn func(*Vector, Vector)) {
	for _, v2 := range vs {
		fn(v, v2)
	}
}

func (v *Vector) ReduceRefs(vs VectorRefs, fn func(*Vector, Vector)) {
	for _, v2 := range vs {
		fn(v, *v2)
	}
}

// component wise multiplication. Use an axis-plane and this will project the vector to that axis.
func (v *Vector) Project(axis Vector) {
	v.x *= axis.x
	v.y *= axis.y
	v.z *= axis.z
}

// axis which vector is most aligned with. 
func (v *Vector) LongestAxis() Axis {
	var ll Vector
	ll.Set(*v)
	ll.Project(*v)
	if ll.z >ll.y {
		if ll.y>ll.x {
			return zAxisIndex
			}
		if ll.x > ll.z{
			return xAxisIndex
			}
		}
	if ll.x >ll.y {
		return xAxisIndex
	}
	return yAxisIndex
}

// vector - matrix multiplication
func (v *Vector) Product(m Matrix) {
	v.x, v.y, v.z = v.x*m.x.x+v.y*m.y.x+v.z*m.z.x, v.x*m.x.y+v.y*m.y.y+v.z*m.z.y, v.x*m.x.z+v.y*m.y.z+v.z*m.z.z
}

// vector - transposed matrix multiplication
func (v *Vector) ProductT(m Matrix) {
	v.x, v.y, v.z = v.x*m.x.x+v.y*m.x.y+v.z*m.x.z, v.x*m.y.x+v.y*m.y.y+v.z*m.y.z, v.x*m.z.x+v.y*m.z.y+v.z*m.z.z
}
