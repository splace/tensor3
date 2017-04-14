package tensor3

type Matrix struct {
	x, y, z Vector
}

func (m Matrix) Components() (Vector, Vector, Vector) {
	return m.x, m.y, m.z
}

// missing components default to zero, more than 9 are ignored
func NewMatrix(cs ...BaseType) (m Matrix) {
	switch len(cs) {
	case 9:
		m.z.z = cs[8]
		fallthrough
	case 8:
		m.z.y = cs[7]
		fallthrough
	case 7:
		m.z.x = cs[6]
		fallthrough
	case 6:
		m.y.z = cs[5]
		fallthrough
	case 5:
		m.y.y = cs[4]
		fallthrough
	case 4:
		m.y.x = cs[3]
		fallthrough
	case 3:
		m.x.z = cs[2]
		fallthrough
	case 2:
		m.x.y = cs[1]
		fallthrough
	case 1:
		m.x.x = cs[0]
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
func (m *Matrix) LengthLength() (v Vector) {
	v.x = m.x.LengthLength()
	v.y = m.y.LengthLength()
	v.z = m.z.LengthLength()
	return
}

// component-vector-wise dot with vector
func (m Matrix) Dot(v2 Vector) (v Vector) {
	v.x = m.x.Dot(v2)
	v.y = m.y.Dot(v2)
	v.z = m.z.Dot(v2)
	return
}

func (m Matrix) Determinant() BaseType {
	return m.x.x*m.y.y*m.z.z - m.x.x*m.y.z*m.z.y + m.x.y*m.y.z*m.z.x - m.x.y*m.y.x*m.z.z + m.x.z*m.y.x*m.z.y - m.x.z*m.y.y*m.z.x
}

func (m *Matrix) Invert() {
	det := m.Determinant()
	var det2x2 func(BaseType, BaseType, BaseType, BaseType) BaseType
	det2x2 = func(a, b, c, d BaseType) BaseType {
		return a*d - b*c
	}
	m.x.x, m.x.y, m.x.z, m.y.x, m.y.y, m.y.z, m.z.x, m.z.y, m.z.z =
		det2x2(m.y.y, m.y.z, m.z.y, m.z.z), det2x2(m.x.z, m.x.y, m.z.z, m.z.y), det2x2(m.x.y, m.x.z, m.y.y, m.y.z),
		det2x2(m.y.z, m.y.x, m.z.z, m.z.x), det2x2(m.x.x, m.x.z, m.z.x, m.z.z), det2x2(m.x.z, m.x.x, m.y.z, m.y.x),
		det2x2(m.y.x, m.y.y, m.z.x, m.z.y), det2x2(m.x.y, m.x.x, m.z.y, m.z.x), det2x2(m.x.x, m.x.y, m.y.x, m.y.y)
	m.Multiply(1 / det)
}

func (m *Matrix) Multiply(s BaseType) {
	m.x.Multiply(s)
	m.y.Multiply(s)
	m.z.Multiply(s)
}

func (m *Matrix) Product(m2 Matrix) {
	m.x.x, m.x.y, m.x.z, m.y.x, m.y.y, m.y.z, m.z.x, m.z.y, m.z.z =
		m.x.x*m2.x.x+m.x.y*m2.y.x+m.x.z*m2.z.x, m.x.x*m2.x.y+m.x.y*m2.y.y+m.x.z*m2.z.y, m.x.x*m2.x.z+m.x.y*m2.y.z+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.y.x+m.y.z*m2.z.x, m.y.x*m2.x.y+m.y.y*m2.y.y+m.y.z*m2.z.y, m.y.x*m2.x.z+m.y.y*m2.y.z+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.y.x+m.z.z*m2.z.x, m.z.x*m2.x.y+m.z.y*m2.y.y+m.z.z*m2.z.y, m.z.x*m2.x.z+m.z.y*m2.y.z+m.z.z*m2.z.z
}

func (m *Matrix) ProductRight(m2 Matrix) {
	m.x.x, m.x.y, m.x.z, m.y.x, m.y.y, m.y.z, m.z.x, m.z.y, m.z.z =
		m2.x.x*m.x.x+m2.x.y*m.y.x+m2.x.z*m.z.x, m2.x.x*m.x.y+m2.x.y*m.y.y+m2.x.z*m.z.y, m2.x.x*m.x.z+m2.x.y*m.y.z+m2.x.z*m.z.z,
		m2.y.x*m.x.x+m2.y.y*m.y.x+m2.y.z*m.z.x, m2.y.x*m.x.y+m2.y.y*m.y.y+m2.y.z*m.z.y, m2.y.x*m.x.z+m2.y.y*m.y.z+m2.y.z*m.z.z,
		m2.z.x*m.x.x+m2.z.y*m.y.x+m2.z.z*m.z.x, m2.z.x*m.x.y+m2.z.y*m.y.y+m2.z.z*m.z.y, m2.z.x*m.x.z+m2.z.y*m.y.z+m2.z.z*m.z.z
}

func (m *Matrix) Transpose() {
	m.x.y, m.x.z, m.y.z, m.y.x, m.z.x, m.z.y = m.y.x, m.z.x, m.z.y, m.x.y, m.x.z, m.y.z
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

func (m *Matrix) Reduce(ms Matrices, fn func(*Matrix, Matrix)) {
	for _, m2 := range ms {
		fn(m, m2)
	}
}

// apply a function to each of the matrices 3 vector components, parameterised by the existing component and the corresponding component of the passed matrix. 
func (m *Matrix) ApplyToComponents(fn func(*Vector, Vector), m2 Matrix) {
	if !ParallelComponents {
		fn(&m.x, m2.x)
		fn(&m.y, m2.y)
		fn(&m.z, m2.z)
	} else {
		done := make(chan struct{}, 1)
		go func() {
			fn(&m.x, m2.x)
			done <- struct{}{}
		}()
		go func() {
			fn(&m.y, m2.y)
			done <- struct{}{}
		}()
		go func() {
			fn(&m.z, m2.z)
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
		fn(&m.x, v)
		fn(&m.y, v)
		fn(&m.z, v)
	} else {
		done := make(chan struct{}, 1)
		go func() {
			fn(&m.x, v)
			done <- struct{}{}
		}()
		go func() {
			fn(&m.y, v)
			done <- struct{}{}
		}()
		go func() {
			fn(&m.z, v)
			done <- struct{}{}
		}()
		<-done
		<-done
		<-done
	}
}

// apply a function to a matrices first vector component, parameterised by the existing component and the passed vector. 
func (m *Matrix) applyX(fn func(*Vector, Vector), v Vector) {
	fn(&m.x, v)
}

// apply a function to a matrices second vector component, parameterised by the existing component and the passed vector. 
func (m *Matrix) applyY(fn func(*Vector, Vector), v Vector) {
	fn(&m.y, v)
}

// apply a function to a matrices third vector component, parameterised by the existing component and the passed vector. 
func (m *Matrix) applyZ(fn func(*Vector, Vector), v Vector) {
	fn(&m.z, v)
}
