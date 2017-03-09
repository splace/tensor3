package tensor3

type Matrix struct {
	x, y, z Vector
}

// missing components default to zero, more than 9 are ignored
func NewMatrix(cs ...float) (m Matrix) {
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

var Identity = Matrix{XAxis, YAxis, ZAxis}

func (m *Matrix) Add(m2 Matrix) {
	m.x.Add(m2.x)
	m.y.Add(m2.y)
	m.z.Add(m2.z)
}

func (m *Matrix) Subtract(m2 Matrix) {
	m.x.Subtract(m2.x)
	m.y.Subtract(m2.y)
	m.z.Subtract(m2.z)
}

func (m *Matrix) Multiply(s float) {
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

func (m *Matrix) ProductT(m2 Matrix) {
	m.x.x, m.x.y, m.x.z, m.y.x, m.y.y, m.y.z, m.z.x, m.z.y, m.z.z =
		m.x.x*m2.x.x+m.x.y*m2.x.y+m.x.z*m2.x.z, m.x.x*m2.y.x+m.x.y*m2.y.y+m.x.z*m2.y.z, m.x.x*m2.z.x+m.x.y*m2.z.y+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.x.y+m.y.z*m2.x.z, m.y.x*m2.y.x+m.y.y*m2.y.y+m.y.z*m2.y.z, m.y.x*m2.z.x+m.y.y*m2.z.y+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.x.y+m.z.z*m2.x.z, m.z.x*m2.y.x+m.z.y*m2.y.y+m.z.z*m2.y.z, m.z.x*m2.z.x+m.z.y*m2.z.y+m.z.z*m2.z.z
}

func (m *Matrix) ProductRightT(m2 Matrix) {
	m.x.x, m.x.y, m.x.z, m.y.x, m.y.y, m.y.z, m.z.x, m.z.y, m.z.z =
		m2.x.x*m.x.x+m2.x.y*m.x.y+m2.x.z*m.x.z, m2.x.x*m.y.x+m2.x.y*m.y.y+m2.x.z*m.y.z, m2.x.x*m.z.x+m2.x.y*m.z.y+m2.x.z*m.z.z,
		m2.y.x*m.x.x+m2.y.y*m.x.y+m2.y.z*m.x.z, m2.y.x*m.y.x+m2.y.y*m.y.y+m2.y.z*m.y.z, m2.y.x*m.z.x+m2.y.y*m.z.y+m2.y.z*m.z.z,
		m2.z.x*m.x.x+m2.z.y*m.x.y+m2.z.z*m.x.z, m2.z.x*m.y.x+m2.z.y*m.y.y+m2.z.z*m.y.z, m2.z.x*m.z.x+m2.z.y*m.z.y+m2.z.z*m.z.z
}

func (m *Matrix) Project(m2 Matrix) {
	m.x.Project(m2.x)
	m.y.Project(m2.y)
	m.z.Project(m2.z)
}

func (m *Matrix) Set(m2 Matrix) {
	m.x.Set(m2.x)
	m.y.Set(m2.y)
	m.z.Set(m2.z)
}

func (m *Matrix) Max(m2 Matrix) {
	m.x.Max(m2.x)
	m.y.Max(m2.y)
	m.z.Max(m2.z)
}

func (m *Matrix) Min(m2 Matrix) {
	m.x.Min(m2.x)
	m.y.Min(m2.y)
	m.z.Min(m2.z)
}

func (m *Matrix) Reduce(ms Matrices, fn func(*Matrix, Matrix)) {
	for _, m2 := range ms {
		fn(m, m2)
	}
}

// apply a vector(vector) function to a component of a matrix
func (m *Matrix) applyX(fn func(*Vector, Vector), v Vector) {
	fn(&m.x, v)
}

func (m *Matrix) applyY(fn func(*Vector, Vector), v Vector) {
	fn(&m.y, v)
}

func (m *Matrix) applyZ(fn func(*Vector, Vector), v Vector) {
	fn(&m.z, v)
}
