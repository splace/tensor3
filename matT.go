package tensor3

//	methods with a matrix parameter can come with 'T' prefixed version, meaning the result is post transposed, or a "T" suffix meaning the parameter is transposed before the operation. adding transpose(s) is a no cost operation.

func (m *Matrix) TProduct(m2 Matrix) {
	m.x.x, m.y.x, m.z.x, m.x.y, m.y.y, m.z.y, m.x.z, m.y.z, m.z.z =
		m.x.x*m2.x.x+m.x.y*m2.y.x+m.x.z*m2.z.x, m.x.x*m2.x.y+m.x.y*m2.y.y+m.x.z*m2.z.y, m.x.x*m2.x.z+m.x.y*m2.y.z+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.y.x+m.y.z*m2.z.x, m.y.x*m2.x.y+m.y.y*m2.y.y+m.y.z*m2.z.y, m.y.x*m2.x.z+m.y.y*m2.y.z+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.y.x+m.z.z*m2.z.x, m.z.x*m2.x.y+m.z.y*m2.y.y+m.z.z*m2.z.y, m.z.x*m2.x.z+m.z.y*m2.y.z+m.z.z*m2.z.z
	vectorUnscale(&m.x)
	vectorUnscale(&m.y)
	vectorUnscale(&m.z)
}

func (m *Matrix) ProductT(m2 Matrix) {
	m.x.x, m.x.y, m.x.z, m.y.x, m.y.y, m.y.z, m.z.x, m.z.y, m.z.z =
		m.x.x*m2.x.x+m.x.y*m2.x.y+m.x.z*m2.x.z, m.x.x*m2.y.x+m.x.y*m2.y.y+m.x.z*m2.y.z, m.x.x*m2.z.x+m.x.y*m2.z.y+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.x.y+m.y.z*m2.x.z, m.y.x*m2.y.x+m.y.y*m2.y.y+m.y.z*m2.y.z, m.y.x*m2.z.x+m.y.y*m2.z.y+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.x.y+m.z.z*m2.x.z, m.z.x*m2.y.x+m.z.y*m2.y.y+m.z.z*m2.y.z, m.z.x*m2.z.x+m.z.y*m2.z.y+m.z.z*m2.z.z
	vectorUnscale(&m.x)
	vectorUnscale(&m.y)
	vectorUnscale(&m.z)
}

func (m *Matrix) TProductT(m2 Matrix) {
	m.x.x, m.y.x, m.z.x, m.x.y, m.y.y, m.z.y, m.x.z, m.y.z, m.z.z =
		m.x.x*m2.x.x+m.x.y*m2.x.y+m.x.z*m2.x.z, m.x.x*m2.y.x+m.x.y*m2.y.y+m.x.z*m2.y.z, m.x.x*m2.z.x+m.x.y*m2.z.y+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.x.y+m.y.z*m2.x.z, m.y.x*m2.y.x+m.y.y*m2.y.y+m.y.z*m2.y.z, m.y.x*m2.z.x+m.y.y*m2.z.y+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.x.y+m.z.z*m2.x.z, m.z.x*m2.y.x+m.z.y*m2.y.y+m.z.z*m2.y.z, m.z.x*m2.z.x+m.z.y*m2.z.y+m.z.z*m2.z.z
	vectorUnscale(&m.x)
	vectorUnscale(&m.y)
	vectorUnscale(&m.z)
}

func (m *Matrix) TProductRight(m2 Matrix) {
	m.x.x, m.y.x, m.z.x, m.x.y, m.y.y, m.z.y, m.x.z, m.y.z, m.z.z =
		m2.x.x*m.x.x+m2.x.y*m.y.x+m2.x.z*m.z.x, m2.x.x*m.x.y+m2.x.y*m.y.y+m2.x.z*m.z.y, m2.x.x*m.x.z+m2.x.y*m.y.z+m2.x.z*m.z.z,
		m2.y.x*m.x.x+m2.y.y*m.y.x+m2.y.z*m.z.x, m2.y.x*m.x.y+m2.y.y*m.y.y+m2.y.z*m.z.y, m2.y.x*m.x.z+m2.y.y*m.y.z+m2.y.z*m.z.z,
		m2.z.x*m.x.x+m2.z.y*m.y.x+m2.z.z*m.z.x, m2.z.x*m.x.y+m2.z.y*m.y.y+m2.z.z*m.z.y, m2.z.x*m.x.z+m2.z.y*m.y.z+m2.z.z*m.z.z
	vectorUnscale(&m.x)
	vectorUnscale(&m.y)
	vectorUnscale(&m.z)
}

func (m *Matrix) TProductRightT(m2 Matrix) {
	m.x.x, m.y.x, m.z.x, m.x.y, m.y.y, m.z.y, m.x.z, m.y.z, m.z.z =
		m2.x.x*m.x.x+m2.x.y*m.x.y+m2.x.z*m.x.z, m2.x.x*m.y.x+m2.x.y*m.y.y+m2.x.z*m.y.z, m2.x.x*m.z.x+m2.x.y*m.z.y+m2.x.z*m.z.z,
		m2.y.x*m.x.x+m2.y.y*m.x.y+m2.y.z*m.x.z, m2.y.x*m.y.x+m2.y.y*m.y.y+m2.y.z*m.y.z, m2.y.x*m.z.x+m2.y.y*m.z.y+m2.y.z*m.z.z,
		m2.z.x*m.x.x+m2.z.y*m.x.y+m2.z.z*m.x.z, m2.z.x*m.y.x+m2.z.y*m.y.y+m2.z.z*m.y.z, m2.z.x*m.z.x+m2.z.y*m.z.y+m2.z.z*m.z.z
	vectorUnscale(&m.x)
	vectorUnscale(&m.y)
	vectorUnscale(&m.z)
}

func (m *Matrix) ProductRightT(m2 Matrix) {
	m.x.x, m.x.y, m.x.z, m.y.x, m.y.y, m.y.z, m.z.x, m.z.y, m.z.z =
		m2.x.x*m.x.x+m2.x.y*m.x.y+m2.x.z*m.x.z, m2.x.x*m.y.x+m2.x.y*m.y.y+m2.x.z*m.y.z, m2.x.x*m.z.x+m2.x.y*m.z.y+m2.x.z*m.z.z,
		m2.y.x*m.x.x+m2.y.y*m.x.y+m2.y.z*m.x.z, m2.y.x*m.y.x+m2.y.y*m.y.y+m2.y.z*m.y.z, m2.y.x*m.z.x+m2.y.y*m.z.y+m2.y.z*m.z.z,
		m2.z.x*m.x.x+m2.z.y*m.x.y+m2.z.z*m.x.z, m2.z.x*m.y.x+m2.z.y*m.y.y+m2.z.z*m.y.z, m2.z.x*m.z.x+m2.z.y*m.z.y+m2.z.z*m.z.z
	vectorUnscale(&m.x)
	vectorUnscale(&m.y)
	vectorUnscale(&m.z)
}
