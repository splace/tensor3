package tensor3

//	methods with a matrix parameter can come with 'T' prefixed version, meaning the result is post transposed, or a "T" suffix meaning the parameter is transposed before the operation. adding transpose(s) is a no extra cost operation.

func (m *Matrix) TProduct(m2 Matrix) {
	m[0].x, m[1].x, m[2].x, m[0].y, m[1].y, m[2].y, m[0].z, m[1].z, m[2].z =
		m[0].x*m2[0].x+m[0].y*m2[1].x+m[0].z*m2[2].x, m[0].x*m2[0].y+m[0].y*m2[1].y+m[0].z*m2[2].y, m[0].x*m2[0].z+m[0].y*m2[1].z+m[0].z*m2[2].z,
		m[1].x*m2[0].x+m[1].y*m2[1].x+m[1].z*m2[2].x, m[1].x*m2[0].y+m[1].y*m2[1].y+m[1].z*m2[2].y, m[1].x*m2[0].z+m[1].y*m2[1].z+m[1].z*m2[2].z,
		m[2].x*m2[0].x+m[2].y*m2[1].x+m[2].z*m2[2].x, m[2].x*m2[0].y+m[2].y*m2[1].y+m[2].z*m2[2].y, m[2].x*m2[0].z+m[2].y*m2[1].z+m[2].z*m2[2].z
	vectorUnscale(&m[0])
	vectorUnscale(&m[1])
	vectorUnscale(&m[2])
}

func (m *Matrix) ProductT(m2 Matrix) {
	m[0].x, m[0].y, m[0].z, m[1].x, m[1].y, m[1].z, m[2].x, m[2].y, m[2].z =
		m[0].x*m2[0].x+m[0].y*m2[0].y+m[0].z*m2[0].z, m[0].x*m2[1].x+m[0].y*m2[1].y+m[0].z*m2[1].z, m[0].x*m2[2].x+m[0].y*m2[2].y+m[0].z*m2[2].z,
		m[1].x*m2[0].x+m[1].y*m2[0].y+m[1].z*m2[0].z, m[1].x*m2[1].x+m[1].y*m2[1].y+m[1].z*m2[1].z, m[1].x*m2[2].x+m[1].y*m2[2].y+m[1].z*m2[2].z,
		m[2].x*m2[0].x+m[2].y*m2[0].y+m[2].z*m2[0].z, m[2].x*m2[1].x+m[2].y*m2[1].y+m[2].z*m2[1].z, m[2].x*m2[2].x+m[2].y*m2[2].y+m[2].z*m2[2].z
	vectorUnscale(&m[0])
	vectorUnscale(&m[1])
	vectorUnscale(&m[2])
}

func (m *Matrix) TProductT(m2 Matrix) {
	m[0].x, m[1].x, m[2].x, m[0].y, m[1].y, m[2].y, m[0].z, m[1].z, m[2].z =
		m[0].x*m2[0].x+m[0].y*m2[0].y+m[0].z*m2[0].z, m[0].x*m2[1].x+m[0].y*m2[1].y+m[0].z*m2[1].z, m[0].x*m2[2].x+m[0].y*m2[2].y+m[0].z*m2[2].z,
		m[1].x*m2[0].x+m[1].y*m2[0].y+m[1].z*m2[0].z, m[1].x*m2[1].x+m[1].y*m2[1].y+m[1].z*m2[1].z, m[1].x*m2[2].x+m[1].y*m2[2].y+m[1].z*m2[2].z,
		m[2].x*m2[0].x+m[2].y*m2[0].y+m[2].z*m2[0].z, m[2].x*m2[1].x+m[2].y*m2[1].y+m[2].z*m2[1].z, m[2].x*m2[2].x+m[2].y*m2[2].y+m[2].z*m2[2].z
	vectorUnscale(&m[0])
	vectorUnscale(&m[1])
	vectorUnscale(&m[2])
}

func (m *Matrix) TProductRight(m2 Matrix) {
	m[0].x, m[1].x, m[2].x, m[0].y, m[1].y, m[2].y, m[0].z, m[1].z, m[2].z =
		m2[0].x*m[0].x+m2[0].y*m[1].x+m2[0].z*m[2].x, m2[0].x*m[0].y+m2[0].y*m[1].y+m2[0].z*m[2].y, m2[0].x*m[0].z+m2[0].y*m[1].z+m2[0].z*m[2].z,
		m2[1].x*m[0].x+m2[1].y*m[1].x+m2[1].z*m[2].x, m2[1].x*m[0].y+m2[1].y*m[1].y+m2[1].z*m[2].y, m2[1].x*m[0].z+m2[1].y*m[1].z+m2[1].z*m[2].z,
		m2[2].x*m[0].x+m2[2].y*m[1].x+m2[2].z*m[2].x, m2[2].x*m[0].y+m2[2].y*m[1].y+m2[2].z*m[2].y, m2[2].x*m[0].z+m2[2].y*m[1].z+m2[2].z*m[2].z
	vectorUnscale(&m[0])
	vectorUnscale(&m[1])
	vectorUnscale(&m[2])
}

func (m *Matrix) TProductRightT(m2 Matrix) {
	m[0].x, m[1].x, m[2].x, m[0].y, m[1].y, m[2].y, m[0].z, m[1].z, m[2].z =
		m2[0].x*m[0].x+m2[0].y*m[0].y+m2[0].z*m[0].z, m2[0].x*m[1].x+m2[0].y*m[1].y+m2[0].z*m[1].z, m2[0].x*m[2].x+m2[0].y*m[2].y+m2[0].z*m[2].z,
		m2[1].x*m[0].x+m2[1].y*m[0].y+m2[1].z*m[0].z, m2[1].x*m[1].x+m2[1].y*m[1].y+m2[1].z*m[1].z, m2[1].x*m[2].x+m2[1].y*m[2].y+m2[1].z*m[2].z,
		m2[2].x*m[0].x+m2[2].y*m[0].y+m2[2].z*m[0].z, m2[2].x*m[1].x+m2[2].y*m[1].y+m2[2].z*m[1].z, m2[2].x*m[2].x+m2[2].y*m[2].y+m2[2].z*m[2].z
	vectorUnscale(&m[0])
	vectorUnscale(&m[1])
	vectorUnscale(&m[2])
}

func (m *Matrix) ProductRightT(m2 Matrix) {
	m[0].x, m[0].y, m[0].z, m[1].x, m[1].y, m[1].z, m[2].x, m[2].y, m[2].z =
		m2[0].x*m[0].x+m2[0].y*m[0].y+m2[0].z*m[0].z, m2[0].x*m[1].x+m2[0].y*m[1].y+m2[0].z*m[1].z, m2[0].x*m[2].x+m2[0].y*m[2].y+m2[0].z*m[2].z,
		m2[1].x*m[0].x+m2[1].y*m[0].y+m2[1].z*m[0].z, m2[1].x*m[1].x+m2[1].y*m[1].y+m2[1].z*m[1].z, m2[1].x*m[2].x+m2[1].y*m[2].y+m2[1].z*m[2].z,
		m2[2].x*m[0].x+m2[2].y*m[0].y+m2[2].z*m[0].z, m2[2].x*m[1].x+m2[2].y*m[1].y+m2[2].z*m[1].z, m2[2].x*m[2].x+m2[2].y*m[2].y+m2[2].z*m[2].z
	vectorUnscale(&m[0])
	vectorUnscale(&m[1])
	vectorUnscale(&m[2])
}
