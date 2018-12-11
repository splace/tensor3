package tensor3

type Vector struct {
	x, y, z Scalar
}

var xAxis, yAxis, zAxis = Vector{scale, 0, 0}, Vector{0, scale, 0}, Vector{0, 0, scale}
var yzPlane, zxPlane, xyPlane = Vector{0, scale, scale}, Vector{scale, 0, scale}, Vector{scale, scale, 0}

type Axis uint

const (
	XAxisIndex Axis = iota
	YAxisIndex
	ZAxisIndex
)

var Axes = [3]Vector{xAxis, yAxis, zAxis}
var AxisPlanes = [3]Vector{yzPlane, zxPlane, xyPlane}

// vector from component values, type BaseType, missing components default to zero, more than 3 are ignored
func NewVector(cs ...Scalar) (v Vector) {
	switch len(cs) {
	default:
		v.z = baseScale(cs[2])
		fallthrough
	case 2:
		v.y = baseScale(cs[1])
		fallthrough
	case 1:
		v.x = baseScale(cs[0])
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

// components independently multiplied, see Product for matrix multiplication
func (v *Vector) Multiply(s Scalar) {
	v.x *= s
	v.y *= s
	v.z *= s
	vectorUnscale(v)
}

// components independently divided
func (v *Vector) Divide(s Scalar) {
	vectorScale(v) //TODO half scale?
	v.x /= s
	v.y /= s
	v.z /= s
}

func (v Vector) Dot(v2 Vector) Scalar {
	return baseUnscale(v.x*v2.x + v.y*v2.y + v.z*v2.z)
}

func (v *Vector) Cross(v2 Vector) {
	v.x, v.y, v.z = v.y*v2.z-v.z*v2.y, v.z*v2.x-v.x*v2.z, v.x*v2.y-v.y*v2.x
	vectorUnscale(v)
}

// length squared. (returning squared means this package is not dependent on math package.)
func (v Vector) LengthLength() Scalar {
	return baseUnscale(v.x*v.x + v.y*v.y + v.z*v.z)
}

// distance squared. (returning squared means this package is not dependent on math package.)
func (v Vector) DistDist(v2 Vector) Scalar {
	v.Subtract(v2)
	return v.LengthLength()
}

func (v *Vector) Set(v2 Vector) {
	v.x = v2.x
	v.y = v2.y
	v.z = v2.z
}

// Max sets the referenced Vector's components to being the greater of their original value and the presented Vector's.
// (both vectors are inside a bounding box with the returned value as its upper bound.)
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

// Min sets the referenced Vector's components to being the lesser of their original value and the presented Vector's.
// (both vectors are inside a bounding box with the returned value as its lower bound.)
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
	v2.Multiply(1 - Base64(f))
	v.Multiply(Base64(f))
	v.Add(v2)
}

// three points, the projections to axis planes.
func (v Vector) AxisProjections(o Vector) (Vector, Vector, Vector) {
	return Vector{v.x, o.y, o.z}, Vector{o.x, v.y, o.z}, Vector{o.x, o.y, v.z}
}

// apply a function repeatedly to the vector, parameterised by its current value and each vector in the supplied vectors in order.
func (v *Vector) Aggregate(vs Vectors, fn func(*Vector, Vector)) {
	for _, v2 := range vs {
		fn(v, v2)
	}
}

// apply a function repeatedly to the vector, parameterised by the current value of the vector and each vector in the supplied vectors.
func (v *Vector) ForAll(vs Vectors, fn func(*Vector, Vector)) {
	if !Parallel {
		vectorApplyAll(v, fn, vs)
	} else {
		vectorApplyAllChunked(v, fn, vs)
	}
}

func vectorApplyAll(v *Vector, fn func(*Vector, Vector), vs Vectors) {
	for i := range vs {
		fn(v, vs[i])
	}
}

func vectorApplyAllChunked(v *Vector, fn func(*Vector, Vector), vs Vectors) {
	done := make(chan Vector, 1)
	var running uint
	for chunk := range vectorsInChunks(vs, chunkSize(len(vs))) {
		running++
		go func(cvs Vectors) {
			var nv Vector
			vectorApplyAll(&nv, fn, cvs)
			done <- nv
		}(chunk)
	}
	for ; running > 0; running-- {
		fn(v, <-done)
	}
}

// component wise multiplication. Use an axis-plane and this will project the vector to that axis.
func (v *Vector) Project(axis Vector) {
	v.x *= axis.x
	v.y *= axis.y
	v.z *= axis.z
	vectorUnscale(v)
}

// 2D projection, to axis aligned plane, functions.
func (v Vector) ProjectX() (Scalar, Scalar) {
	return v.y, v.z
}

func (v Vector) ProjectY() (Scalar, Scalar) {
	return v.z, v.x
}

func (v Vector) ProjectZ() (Scalar, Scalar) {
	return v.x, v.y
}

func (v *Vector) AddNewell(v1, v2 Vector) {
	v.x += (v1.y + v2.y) * (v1.z - v2.z)
	v.y += (v1.z + v2.z) * (v1.x - v2.x)
	v.z += (v1.x + v2.x) * (v1.y - v2.y)
}

// axis which vector is most aligned with.
func (v Vector) LongestAxis() Axis {
	v.Project(v)
	if v.z > v.y {
		if v.z > v.x {
			return ZAxisIndex
		}
		return XAxisIndex
	}
	if v.x > v.y {
		return XAxisIndex
	}
	return YAxisIndex
}

// axis which vector is least aligned with.
func (v Vector) ShortestAxis() Axis {
	v.Project(v)
	if v.z < v.y {
		if v.x < v.z {
			return XAxisIndex
		}
		return ZAxisIndex
	}
	if v.x > v.y {
		return YAxisIndex
	}
	return XAxisIndex
}

// vector - matrix multiplication
func (v *Vector) Product(m Matrix) {
	v.x, v.y, v.z = v.x*m[0].x+v.y*m[1].x+v.z*m[2].x, v.x*m[0].y+v.y*m[1].y+v.z*m[2].y, v.x*m[0].z+v.y*m[1].z+v.z*m[2].z
	vectorUnscale(v)
}

// vector - transposed matrix multiplication
func (v *Vector) ProductT(m Matrix) {
	v.x, v.y, v.z = v.x*m[0].x+v.y*m[0].y+v.z*m[0].z, v.x*m[1].x+v.y*m[1].y+v.z*m[1].z, v.x*m[2].x+v.y*m[2].y+v.z*m[2].z
	vectorUnscale(v)
}
