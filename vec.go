package tensor3

type Vector struct {
	x, y, z BaseType
}

func (v Vector) Components() (BaseType, BaseType, BaseType) {
	return v.x, v.y, v.z
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

// missing components default to zero, more than 3 are ignored
func NewVector(cs ...BaseType) (v Vector) {
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
func (v *Vector) Multiply(s BaseType) {
	v.x *= s
	v.y *= s
	v.z *= s
}

// components independently divided
func (v *Vector) Divide(s BaseType) {
	v.x /= s
	v.y /= s
	v.z /= s
}

func (v Vector) Dot(v2 Vector) BaseType {
	return baseUnscale(v.x*v2.x + v.y*v2.y + v.z*v2.z)
}

func (v *Vector) Cross(v2 Vector) {
	v.x, v.y, v.z = v.y*v2.z-v.z*v2.y, v.z*v2.x-v.x*v2.z, v.x*v2.y-v.y*v2.x
	vectorUnscale(v)
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
	v.x = (v2.x + v.x) / 2
	v.y = (v2.y + v.y) / 2
	v.z = (v2.z + v.z) / 2
}

func (v *Vector) Interpolate(v2 Vector, f BaseType) {
	v2.Multiply(1 - f)
	v.Multiply(f)
	v.Add(v2)
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
		if Hints.ChunkSizeFixed {
			vectorApplyAllChunked(v, fn, vs, Hints.DefaultChunkSize)
		} else {
			cs := uint(len(vs)) / (Hints.Threads + 1)
			if cs < Hints.DefaultChunkSize {
				cs = Hints.DefaultChunkSize
			}
			vectorApplyAllChunked(v, fn, vs, cs)
		}
	}
}

func vectorApplyAll(v *Vector, fn func(*Vector, Vector), vs Vectors) {
	for i := range vs {
		fn(v, vs[i])
	}
}

func vectorApplyAllChunked(v *Vector, fn func(*Vector, Vector), vs Vectors, chunkSize uint) {
	done := make(chan Vector, 1)
	var running uint
	for chunk := range vectorsInChunks(vs) {
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
	v.x, v.y, v.z = v.x*m.x.x+v.y*m.y.x+v.z*m.z.x, v.x*m.x.y+v.y*m.y.y+v.z*m.z.y, v.x*m.x.z+v.y*m.y.z+v.z*m.z.z
	vectorUnscale(v)
}

// vector - transposed matrix multiplication
func (v *Vector) ProductT(m Matrix) {
	v.x, v.y, v.z = v.x*m.x.x+v.y*m.x.y+v.z*m.x.z, v.x*m.y.x+v.y*m.y.y+v.z*m.y.z, v.x*m.z.x+v.y*m.z.y+v.z*m.z.z
	vectorUnscale(v)
}
