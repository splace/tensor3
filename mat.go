package tensor3

type Matrix [3]Vector

// missing parameters default to zero, more than 9 are ignored
func NewMatrix(cs ...BaseType) (m Matrix) {
	switch len(cs) {
	default:
		m[2].z = baseScale(cs[8])
		fallthrough
	case 8:
		m[2].y = baseScale(cs[7])
		fallthrough
	case 7:
		m[2].x = baseScale(cs[6])
		fallthrough
	case 6:
		m[1].z = baseScale(cs[5])
		fallthrough
	case 5:
		m[1].y = baseScale(cs[4])
		fallthrough
	case 4:
		m[1].x = baseScale(cs[3])
		fallthrough
	case 3:
		m[0].z = baseScale(cs[2])
		fallthrough
	case 2:
		m[0].y = baseScale(cs[1])
		fallthrough
	case 1:
		m[0].x = baseScale(cs[0])
	case 0:
	}
	return
}

var Identity = Matrix{xAxis, yAxis, zAxis}

func (m *Matrix) Add(m2 Matrix) {
	m.ApplyToComponents((*Vector).Add, m2)
}

func (m *Matrix) Subtract(m2 Matrix) {
	m.ApplyToComponents((*Vector).Subtract, m2)
}

// component-vector-wise cross with vector
func (m *Matrix) Cross(v Vector) {
	m.ApplyToComponentsBySameVector((*Vector).Cross, v)
}

// component-vector-wise length squared
func (m Matrix) LengthLength() (v Vector) {
	v.x = m[0].LengthLength()
	v.y = m[1].LengthLength()
	v.z = m[2].LengthLength()
	return
}

// component-vector-wise dot with vector
func (m Matrix) Dot(v2 Vector) (v Vector) {
	v.x = m[0].Dot(v2)
	v.y = m[1].Dot(v2)
	v.z = m[2].Dot(v2)
	return
}

func (m Matrix) Determinant() BaseType {
	return baseUnscale(baseUnscale(m[0].x*m[1].y*m[2].z - m[0].x*m[1].z*m[2].y + m[0].y*m[1].z*m[2].x - m[0].y*m[1].x*m[2].z + m[0].z*m[1].x*m[2].y - m[0].z*m[1].y*m[2].x))
}

func (m *Matrix) Invert() {
	det := m.Determinant()
	var det2x2 func(BaseType, BaseType, BaseType, BaseType) BaseType
	det2x2 = func(a, b, c, d BaseType) BaseType {
		return baseUnscale(a*d - b*c)
	}
	m[0].x, m[0].y, m[0].z, m[1].x, m[1].y, m[1].z, m[2].x, m[2].y, m[2].z =
		det2x2(m[1].y, m[1].z, m[2].y, m[2].z), det2x2(m[0].z, m[0].y, m[2].z, m[2].y), det2x2(m[0].y, m[0].z, m[1].y, m[1].z),
		det2x2(m[1].z, m[1].x, m[2].z, m[2].x), det2x2(m[0].x, m[0].z, m[2].x, m[2].z), det2x2(m[0].z, m[0].x, m[1].z, m[1].x),
		det2x2(m[1].x, m[1].y, m[2].x, m[2].y), det2x2(m[0].y, m[0].x, m[2].y, m[2].x), det2x2(m[0].x, m[0].y, m[1].x, m[1].y)
	m.Divide(det)
}

func (m *Matrix) Multiply(s BaseType) {
	m[0].Multiply(s)
	m[1].Multiply(s)
	m[2].Multiply(s)
}

func (m *Matrix) Divide(s BaseType) {
	m[0].Divide(s)
	m[1].Divide(s)
	m[2].Divide(s)
}

func (m *Matrix) Product(m2 Matrix) {
	m[0].x, m[0].y, m[0].z, m[1].x, m[1].y, m[1].z, m[2].x, m[2].y, m[2].z =
		m[0].x*m2[0].x+m[0].y*m2[1].x+m[0].z*m2[2].x, m[0].x*m2[0].y+m[0].y*m2[1].y+m[0].z*m2[2].y, m[0].x*m2[0].z+m[0].y*m2[1].z+m[0].z*m2[2].z,
		m[1].x*m2[0].x+m[1].y*m2[1].x+m[1].z*m2[2].x, m[1].x*m2[0].y+m[1].y*m2[1].y+m[1].z*m2[2].y, m[1].x*m2[0].z+m[1].y*m2[1].z+m[1].z*m2[2].z,
		m[2].x*m2[0].x+m[2].y*m2[1].x+m[2].z*m2[2].x, m[2].x*m2[0].y+m[2].y*m2[1].y+m[2].z*m2[2].y, m[2].x*m2[0].z+m[2].y*m2[1].z+m[2].z*m2[2].z
	vectorUnscale(&m[0])
	vectorUnscale(&m[1])
	vectorUnscale(&m[2])
}

func (m *Matrix) ProductRight(m2 Matrix) {
	m[0].x, m[0].y, m[0].z, m[1].x, m[1].y, m[1].z, m[2].x, m[2].y, m[2].z =
		m2[0].x*m[0].x+m2[0].y*m[1].x+m2[0].z*m[2].x, m2[0].x*m[0].y+m2[0].y*m[1].y+m2[0].z*m[2].y, m2[0].x*m[0].z+m2[0].y*m[1].z+m2[0].z*m[2].z,
		m2[1].x*m[0].x+m2[1].y*m[1].x+m2[1].z*m[2].x, m2[1].x*m[0].y+m2[1].y*m[1].y+m2[1].z*m[2].y, m2[1].x*m[0].z+m2[1].y*m[1].z+m2[1].z*m[2].z,
		m2[2].x*m[0].x+m2[2].y*m[1].x+m2[2].z*m[2].x, m2[2].x*m[0].y+m2[2].y*m[1].y+m2[2].z*m[2].y, m2[2].x*m[0].z+m2[2].y*m[1].z+m2[2].z*m[2].z
	vectorUnscale(&m[0])
	vectorUnscale(&m[1])
	vectorUnscale(&m[2])
}

func (m *Matrix) Transpose() {
	m[0].y, m[0].z, m[1].z, m[1].x, m[2].x, m[2].y = m[1].x, m[2].x, m[2].y, m[0].y, m[0].z, m[1].z
}

func (m *Matrix) Project(m2 Matrix) {
	m.ApplyToComponents((*Vector).Project, m2)
}

func (m *Matrix) Set(m2 Matrix) {
	m.ApplyToComponents((*Vector).Set, m2)
}

func (m *Matrix) Max(m2 Matrix) {
	m.ApplyToComponents((*Vector).Max, m2)
}

func (m *Matrix) Min(m2 Matrix) {
	m.ApplyToComponents((*Vector).Min, m2)
}

func (m *Matrix) Aggregate(ms Matrices, fn func(*Matrix, Matrix)) {
	for _, m2 := range ms {
		fn(m, m2)
	}
}

func (vs Vectors) Product(m Matrix) {
	m.ForEach((*Vector).Product, vs)
}

func (vs Vectors) ProductT(m Matrix) {
	m.ForEach((*Vector).ProductT, vs)
}

func (m Matrix) ForEach(fn func(*Vector, Matrix), vs Vectors) {
	if !Parallel {
		matrixApply(vs, fn, m)
	} else {
		if Hints.ChunkSizeFixed {
			matrixApplyChunked(vs, fn, m)
		} else {
			matrixApplyChunked(vs, fn, m)
		}
	}
}

func matrixApply(vs Vectors, fn func(*Vector, Matrix), m Matrix) {
	for i := range vs {
		fn(&vs[i], m)
	}
}

func matrixApplyChunked(vs Vectors, fn func(*Vector, Matrix), m Matrix) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range vectorsInChunks(vs,chunkSize(len(vs))) {
		running++
		go func(c Vectors) {
			matrixApply(c, fn, m)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

// apply a function to each of the matrices 3 vector components, parameterised by the existing component and the corresponding component of the passed matrix.
func (m *Matrix) ApplyToComponents(fn func(*Vector, Vector), m2 Matrix) {
	if !ParallelComponents {
		fn(&m[0], m2[0])
		fn(&m[1], m2[1])
		fn(&m[2], m2[2])
	} else {
		done := make(chan struct{}, 1)
		go func() {
			fn(&m[0], m2[0])
			done <- struct{}{}
		}()
		go func() {
			fn(&m[1], m2[1])
			done <- struct{}{}
		}()
		go func() {
			fn(&m[2], m2[2])
			done <- struct{}{}
		}()
		<-done
		<-done
		<-done
	}
}

// apply a function to each of the matrices 3 vector components, parameterised by the existing component and the corresponding component of the identity matrix, (aligned axis vectors).
func (m *Matrix) ApplyToComponentsByAxes(fn func(*Vector, Vector)) {
	m.ApplyToComponents(fn, Matrix{xAxis, yAxis, zAxis})
}

// apply a function to each of the matrices 3 vector components, parameterised by the existing component and the passed vector.
func (m *Matrix) ApplyToComponentsBySameVector(fn func(*Vector, Vector), v Vector) {
	if !ParallelComponents {
		fn(&m[0], v)
		fn(&m[1], v)
		fn(&m[2], v)
	} else {
		done := make(chan struct{}, 1)
		go func() {
			fn(&m[0], v)
			done <- struct{}{}
		}()
		go func() {
			fn(&m[1], v)
			done <- struct{}{}
		}()
		go func() {
			fn(&m[2], v)
			done <- struct{}{}
		}()
		<-done
		<-done
		<-done
	}
}

// apply a function to a matrices first vector component, parameterised by the existing component and the passed vector.
func (m *Matrix) applyX(fn func(*Vector, Vector), v Vector) {
	fn(&m[0], v)
}

// apply a function to a matrices second vector component, parameterised by the existing component and the passed vector.
func (m *Matrix) applyY(fn func(*Vector, Vector), v Vector) {
	fn(&m[1], v)
}

// apply a function to a matrices third vector component, parameterised by the existing component and the passed vector.
func (m *Matrix) applyZ(fn func(*Vector, Vector), v Vector) {
	fn(&m[2], v)
}

