package tensor3

func (m *Matrices) AddR(m2 Matrix) {
	m.ForEach((*Matrix).AddR, m2)
}

func (m *Matrices) AddYZX(m2 Matrix) {
	m.ForEach((*Matrix).AddYZX, m2)
}

func (m *Matrices) AddRR(m2 Matrix) {
	m.ForEach((*Matrix).AddRR, m2)
}

func (m *Matrices) AddZXY(m2 Matrix) {
	m.ForEach((*Matrix).AddZXY, m2)
}

func (m *Matrices) AddYXZ(m2 Matrix) {
	m.ForEach((*Matrix).AddYXZ, m2)
}

func (m *Matrices) AddZYX(m2 Matrix) {
	m.ForEach((*Matrix).AddZYX, m2)
}

func (m *Matrices) AddXZY(m2 Matrix) {
	m.ForEach((*Matrix).AddXZY, m2)
}

func (m *Matrices) SubtractR(m2 Matrix) {
	m.ForEach((*Matrix).SubtractR, m2)
}

func (m *Matrices) SubtractYZX(m2 Matrix) {
	m.ForEach((*Matrix).SubtractYZX, m2)
}

func (m *Matrices) SubtractRR(m2 Matrix) {
	m.ForEach((*Matrix).SubtractRR, m2)
}

func (m *Matrices) SubtractZXY(m2 Matrix) {
	m.ForEach((*Matrix).SubtractZXY, m2)
}

func (m *Matrices) SubtractYXZ(m2 Matrix) {
	m.ForEach((*Matrix).SubtractYXZ, m2)
}

func (m *Matrices) SubtractZYX(m2 Matrix) {
	m.ForEach((*Matrix).SubtractZYX, m2)
}

func (m *Matrices) SubtractXZY(m2 Matrix) {
	m.ForEach((*Matrix).SubtractXZY, m2)
}

func (m *Matrices) ProjectR(m2 Matrix) {
	m.ForEach((*Matrix).ProjectR, m2)
}

func (m *Matrices) ProjectYZX(m2 Matrix) {
	m.ForEach((*Matrix).ProjectYZX, m2)
}

func (m *Matrices) ProjectRR(m2 Matrix) {
	m.ForEach((*Matrix).ProjectRR, m2)
}

func (m *Matrices) ProjectZXY(m2 Matrix) {
	m.ForEach((*Matrix).ProjectZXY, m2)
}

func (m *Matrices) ProjectYXZ(m2 Matrix) {
	m.ForEach((*Matrix).ProjectYXZ, m2)
}

func (m *Matrices) ProjectZYX(m2 Matrix) {
	m.ForEach((*Matrix).ProjectZYX, m2)
}

func (m *Matrices) ProjectXZY(m2 Matrix) {
	m.ForEach((*Matrix).ProjectXZY, m2)
}

func (m *Matrices) ProductR(m2 Matrix) {
	m.ForEach((*Matrix).ProductR, m2)
}

func (m *Matrices) ProductYZX(m2 Matrix) {
	m.ForEach((*Matrix).ProductYZX, m2)
}

func (m *Matrices) ProductRR(m2 Matrix) {
	m.ForEach((*Matrix).ProductRR, m2)
}

func (m *Matrices) ProductZXY(m2 Matrix) {
	m.ForEach((*Matrix).ProductZXY, m2)
}

func (m *Matrices) ProductYXZ(m2 Matrix) {
	m.ForEach((*Matrix).ProductYXZ, m2)
}

func (m *Matrices) ProductZYX(m2 Matrix) {
	m.ForEach((*Matrix).ProductZYX, m2)
}

func (m *Matrices) ProductXZY(m2 Matrix) {
	m.ForEach((*Matrix).ProductXZY, m2)
}

func (m *Matrices) ProductTR(m2 Matrix) {
	m.ForEach((*Matrix).ProductTR, m2)
}

func (m *Matrices) ProductTYZX(m2 Matrix) {
	m.ForEach((*Matrix).ProductTYZX, m2)
}

func (m *Matrices) ProductTRR(m2 Matrix) {
	m.ForEach((*Matrix).ProductTRR, m2)
}

func (m *Matrices) ProductTZXY(m2 Matrix) {
	m.ForEach((*Matrix).ProductTZXY, m2)
}

func (m *Matrices) ProductTYXZ(m2 Matrix) {
	m.ForEach((*Matrix).ProductTYXZ, m2)
}

func (m *Matrices) ProductTZYX(m2 Matrix) {
	m.ForEach((*Matrix).ProductTZYX, m2)
}

func (m *Matrices) ProductTXZY(m2 Matrix) {
	m.ForEach((*Matrix).ProductTXZY, m2)
}
