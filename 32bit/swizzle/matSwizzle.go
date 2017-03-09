package tensor3

func (m *Matrix) AddR(m2 Matrix) {
	m.AddYZX(m2)
}

func (m *Matrix) AddYZX(m2 Matrix) {
	m.y.Add(m2.x)
	m.z.Add(m2.y)
	m.x.Add(m2.z)
}

func (m *Matrix) AddRR(m2 Matrix) {
	m.AddZXY(m2)
}

func (m *Matrix) AddZXY(m2 Matrix) {
	m.z.Add(m2.x)
	m.x.Add(m2.y)
	m.y.Add(m2.z)
}

func (m *Matrix) AddYXZ(m2 Matrix) {
	m.y.Add(m2.x)
	m.x.Add(m2.y)
	m.z.Add(m2.z)
}
func (m *Matrix) AddZYX(m2 Matrix) {
	m.z.Add(m2.x)
	m.y.Add(m2.y)
	m.x.Add(m2.z)
}
func (m *Matrix) AddXZY(m2 Matrix) {
	m.x.Add(m2.x)
	m.z.Add(m2.y)
	m.y.Add(m2.z)
}

func (m *Matrix) SubtractR(m2 Matrix) {
	m.SubtractYZX(m2)
}

func (m *Matrix) SubtractYZX(m2 Matrix) {
	m.y.Subtract(m2.x)
	m.z.Subtract(m2.y)
	m.x.Subtract(m2.z)
}

func (m *Matrix) SubtractRR(m2 Matrix) {
	m.SubtractZXY(m2)
}

func (m *Matrix) SubtractZXY(m2 Matrix) {
	m.z.Subtract(m2.x)
	m.x.Subtract(m2.y)
	m.y.Subtract(m2.z)
}

func (m *Matrix) SubtractYXZ(m2 Matrix) {
	m.y.Subtract(m2.x)
	m.x.Subtract(m2.y)
	m.z.Subtract(m2.z)
}
func (m *Matrix) SubtractZYX(m2 Matrix) {
	m.z.Subtract(m2.x)
	m.y.Subtract(m2.y)
	m.x.Subtract(m2.z)
}
func (m *Matrix) SubtractXZY(m2 Matrix) {
	m.x.Subtract(m2.x)
	m.z.Subtract(m2.y)
	m.y.Subtract(m2.z)
}

func (m *Matrix) MultiplyR(s float) {
	m.MultiplyYZX(s)
}

func (m *Matrix) MultiplyYZX(s float) {
	m.y.Multiply(s)
	m.z.Multiply(s)
	m.x.Multiply(s)
}

func (m *Matrix) MultiplyRR(s float) {
	m.MultiplyZXY(s)
}

func (m *Matrix) MultiplyZXY(s float) {
	m.z.Multiply(s)
	m.x.Multiply(s)
	m.y.Multiply(s)
}

func (m *Matrix) MultiplyYXZ(s float) {
	m.y.Multiply(s)
	m.x.Multiply(s)
	m.z.Multiply(s)
}
func (m *Matrix) MultiplyZYX(s float) {
	m.z.Multiply(s)
	m.y.Multiply(s)
	m.x.Multiply(s)
}
func (m *Matrix) MultiplyXZY(s float) {
	m.x.Multiply(s)
	m.z.Multiply(s)
	m.y.Multiply(s)
}

func (m *Matrix) ProductR(m2 Matrix) {
	m.ProductYZX(m2)
}

func (m *Matrix) ProductYZX(m2 Matrix) {
	m.y.x, m.y.y, m.y.z, m.z.x, m.z.y, m.z.z, m.x.x, m.x.y, m.x.z =
		m.x.x*m2.x.x+m.x.y*m2.y.x+m.x.z*m2.z.x, m.x.x*m2.x.y+m.x.y*m2.y.y+m.x.z*m2.z.y, m.x.x*m2.x.z+m.x.y*m2.y.z+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.y.x+m.y.z*m2.z.x, m.y.x*m2.x.y+m.y.y*m2.y.y+m.y.z*m2.z.y, m.y.x*m2.x.z+m.y.y*m2.y.z+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.y.x+m.z.z*m2.z.x, m.z.x*m2.x.y+m.z.y*m2.y.y+m.z.z*m2.z.y, m.z.x*m2.x.z+m.z.y*m2.y.z+m.z.z*m2.z.z
}

func (m *Matrix) ProductRR(m2 Matrix) {
	m.ProductZXY(m2)
}

func (m *Matrix) ProductZXY(m2 Matrix) {
	m.z.x, m.z.y, m.z.z, m.x.x, m.x.y, m.x.z, m.y.x, m.y.y, m.y.z =
		m.x.x*m2.x.x+m.x.y*m2.y.x+m.x.z*m2.z.x, m.x.x*m2.x.y+m.x.y*m2.y.y+m.x.z*m2.z.y, m.x.x*m2.x.z+m.x.y*m2.y.z+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.y.x+m.y.z*m2.z.x, m.y.x*m2.x.y+m.y.y*m2.y.y+m.y.z*m2.z.y, m.y.x*m2.x.z+m.y.y*m2.y.z+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.y.x+m.z.z*m2.z.x, m.z.x*m2.x.y+m.z.y*m2.y.y+m.z.z*m2.z.y, m.z.x*m2.x.z+m.z.y*m2.y.z+m.z.z*m2.z.z
}

func (m *Matrix) ProductYXZ(m2 Matrix) {
	m.y.x, m.y.y, m.y.z, m.x.x, m.x.y, m.x.z, m.z.x, m.z.y, m.z.z =
		m.x.x*m2.x.x+m.x.y*m2.y.x+m.x.z*m2.z.x, m.x.x*m2.x.y+m.x.y*m2.y.y+m.x.z*m2.z.y, m.x.x*m2.x.z+m.x.y*m2.y.z+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.y.x+m.y.z*m2.z.x, m.y.x*m2.x.y+m.y.y*m2.y.y+m.y.z*m2.z.y, m.y.x*m2.x.z+m.y.y*m2.y.z+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.y.x+m.z.z*m2.z.x, m.z.x*m2.x.y+m.z.y*m2.y.y+m.z.z*m2.z.y, m.z.x*m2.x.z+m.z.y*m2.y.z+m.z.z*m2.z.z
}
func (m *Matrix) ProductZYX(m2 Matrix) {
	m.z.x, m.z.y, m.z.z, m.y.x, m.y.y, m.y.z, m.x.x, m.x.y, m.x.z =
		m.x.x*m2.x.x+m.x.y*m2.y.x+m.x.z*m2.z.x, m.x.x*m2.x.y+m.x.y*m2.y.y+m.x.z*m2.z.y, m.x.x*m2.x.z+m.x.y*m2.y.z+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.y.x+m.y.z*m2.z.x, m.y.x*m2.x.y+m.y.y*m2.y.y+m.y.z*m2.z.y, m.y.x*m2.x.z+m.y.y*m2.y.z+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.y.x+m.z.z*m2.z.x, m.z.x*m2.x.y+m.z.y*m2.y.y+m.z.z*m2.z.y, m.z.x*m2.x.z+m.z.y*m2.y.z+m.z.z*m2.z.z
}
func (m *Matrix) ProductXZY(m2 Matrix) {
	m.x.x, m.x.y, m.x.z, m.z.x, m.z.y, m.z.z, m.y.x, m.y.y, m.y.z =
		m.x.x*m2.x.x+m.x.y*m2.y.x+m.x.z*m2.z.x, m.x.x*m2.x.y+m.x.y*m2.y.y+m.x.z*m2.z.y, m.x.x*m2.x.z+m.x.y*m2.y.z+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.y.x+m.y.z*m2.z.x, m.y.x*m2.x.y+m.y.y*m2.y.y+m.y.z*m2.z.y, m.y.x*m2.x.z+m.y.y*m2.y.z+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.y.x+m.z.z*m2.z.x, m.z.x*m2.x.y+m.z.y*m2.y.y+m.z.z*m2.z.y, m.z.x*m2.x.z+m.z.y*m2.y.z+m.z.z*m2.z.z
}

func (m *Matrix) ProjectR(m2 Matrix) {
	m.ProjectYZX(m2)
}

func (m *Matrix) ProjectYZX(m2 Matrix) {
	m.y.Project(m2.x)
	m.z.Project(m2.y)
	m.x.Project(m2.z)
}

func (m *Matrix) ProjectRR(m2 Matrix) {
	m.ProjectZXY(m2)
}

func (m *Matrix) ProjectZXY(m2 Matrix) {
	m.z.Project(m2.x)
	m.x.Project(m2.y)
	m.y.Project(m2.z)
}

func (m *Matrix) ProjectYXZ(m2 Matrix) {
	m.y.Project(m2.x)
	m.x.Project(m2.y)
	m.z.Project(m2.z)
}
func (m *Matrix) ProjectZYX(m2 Matrix) {
	m.z.Project(m2.x)
	m.y.Project(m2.y)
	m.x.Project(m2.z)
}
func (m *Matrix) ProjectXZY(m2 Matrix) {
	m.x.Project(m2.x)
	m.z.Project(m2.y)
	m.y.Project(m2.z)
}

func (m *Matrix) TransposeR() {
	m.TransposeYZX()
}

func (m *Matrix) TransposeYZX() {
	m.y.y, m.y.z, m.z.z, m.z.x, m.x.x, m.x.y = m.y.x, m.z.x, m.z.y, m.x.y, m.x.z, m.y.z
}

func (m *Matrix) TransposeRR() {
	m.TransposeZXY()
}

func (m *Matrix) TransposeZXY() {
	m.y.y, m.y.z, m.z.z, m.z.x, m.x.x, m.x.y = m.y.x, m.z.x, m.z.y, m.x.y, m.x.z, m.y.z
}

func (m *Matrix) TransposeYXZ() {
	m.y.y, m.y.z, m.z.z, m.z.x, m.x.x, m.x.y = m.y.x, m.z.x, m.z.y, m.x.y, m.x.z, m.y.z
}
func (m *Matrix) TransposeZYX() {
	m.y.y, m.y.z, m.z.z, m.z.x, m.x.x, m.x.y = m.y.x, m.z.x, m.z.y, m.x.y, m.x.z, m.y.z
}
func (m *Matrix) TransposeXZY() {
	m.y.y, m.y.z, m.z.z, m.z.x, m.x.x, m.x.y = m.y.x, m.z.x, m.z.y, m.x.y, m.x.z, m.y.z
}

func (m *Matrix) ProductTR(m2 Matrix) {
	m.ProductTYZX(m2)
}

func (m *Matrix) ProductTYZX(m2 Matrix) {
	m.y.x, m.y.y, m.y.z, m.z.x, m.z.y, m.z.z, m.x.x, m.x.y, m.x.z =
		m.x.x*m2.x.x+m.x.y*m2.x.y+m.x.z*m2.x.z, m.x.x*m2.y.x+m.x.y*m2.y.y+m.x.z*m2.y.z, m.x.x*m2.z.x+m.x.y*m2.z.y+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.x.y+m.y.z*m2.x.z, m.y.x*m2.y.x+m.y.y*m2.y.y+m.y.z*m2.y.z, m.y.x*m2.z.x+m.y.y*m2.z.y+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.x.y+m.z.z*m2.x.z, m.z.x*m2.y.x+m.z.y*m2.y.y+m.z.z*m2.y.z, m.z.x*m2.z.x+m.z.y*m2.z.y+m.z.z*m2.z.z
}

func (m *Matrix) ProductTRR(m2 Matrix) {
	m.ProductTZXY(m2)
}

func (m *Matrix) ProductTZXY(m2 Matrix) {
	m.z.x, m.z.y, m.z.z, m.x.x, m.x.y, m.x.z, m.y.x, m.y.y, m.y.z =
		m.x.x*m2.x.x+m.x.y*m2.x.y+m.x.z*m2.x.z, m.x.x*m2.y.x+m.x.y*m2.y.y+m.x.z*m2.y.z, m.x.x*m2.z.x+m.x.y*m2.z.y+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.x.y+m.y.z*m2.x.z, m.y.x*m2.y.x+m.y.y*m2.y.y+m.y.z*m2.y.z, m.y.x*m2.z.x+m.y.y*m2.z.y+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.x.y+m.z.z*m2.x.z, m.z.x*m2.y.x+m.z.y*m2.y.y+m.z.z*m2.y.z, m.z.x*m2.z.x+m.z.y*m2.z.y+m.z.z*m2.z.z
}

func (m *Matrix) ProductTYXZ(m2 Matrix) {
	m.y.x, m.y.y, m.y.z, m.x.x, m.x.y, m.x.z, m.z.x, m.z.y, m.z.z =
		m.x.x*m2.x.x+m.x.y*m2.x.y+m.x.z*m2.x.z, m.x.x*m2.y.x+m.x.y*m2.y.y+m.x.z*m2.y.z, m.x.x*m2.z.x+m.x.y*m2.z.y+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.x.y+m.y.z*m2.x.z, m.y.x*m2.y.x+m.y.y*m2.y.y+m.y.z*m2.y.z, m.y.x*m2.z.x+m.y.y*m2.z.y+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.x.y+m.z.z*m2.x.z, m.z.x*m2.y.x+m.z.y*m2.y.y+m.z.z*m2.y.z, m.z.x*m2.z.x+m.z.y*m2.z.y+m.z.z*m2.z.z
}
func (m *Matrix) ProductTZYX(m2 Matrix) {
	m.z.x, m.z.y, m.z.z, m.y.x, m.y.y, m.y.z, m.x.x, m.x.y, m.x.z =
		m.x.x*m2.x.x+m.x.y*m2.x.y+m.x.z*m2.x.z, m.x.x*m2.y.x+m.x.y*m2.y.y+m.x.z*m2.y.z, m.x.x*m2.z.x+m.x.y*m2.z.y+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.x.y+m.y.z*m2.x.z, m.y.x*m2.y.x+m.y.y*m2.y.y+m.y.z*m2.y.z, m.y.x*m2.z.x+m.y.y*m2.z.y+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.x.y+m.z.z*m2.x.z, m.z.x*m2.y.x+m.z.y*m2.y.y+m.z.z*m2.y.z, m.z.x*m2.z.x+m.z.y*m2.z.y+m.z.z*m2.z.z
}
func (m *Matrix) ProductTXZY(m2 Matrix) {
	m.x.x, m.x.y, m.x.z, m.z.x, m.z.y, m.z.z, m.y.x, m.y.y, m.y.z =
		m.x.x*m2.x.x+m.x.y*m2.x.y+m.x.z*m2.x.z, m.x.x*m2.y.x+m.x.y*m2.y.y+m.x.z*m2.y.z, m.x.x*m2.z.x+m.x.y*m2.z.y+m.x.z*m2.z.z,
		m.y.x*m2.x.x+m.y.y*m2.x.y+m.y.z*m2.x.z, m.y.x*m2.y.x+m.y.y*m2.y.y+m.y.z*m2.y.z, m.y.x*m2.z.x+m.y.y*m2.z.y+m.y.z*m2.z.z,
		m.z.x*m2.x.x+m.z.y*m2.x.y+m.z.z*m2.x.z, m.z.x*m2.y.x+m.z.y*m2.y.y+m.z.z*m2.y.z, m.z.x*m2.z.x+m.z.y*m2.z.y+m.z.z*m2.z.z
}
