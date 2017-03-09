package tensor3

func (vs Vectors) CrossR(v Vector) {
	vs.ForEach((*Vector).CrossR, v)
}

func (vs Vectors) CrossRR(v Vector) {
	vs.ForEach((*Vector).CrossRR, v)
}

func (vs Vectors) CrossYZX(v Vector) {
	vs.ForEach((*Vector).CrossYZX, v)
}

func (vs Vectors) CrossZXY(v Vector) {
	vs.ForEach((*Vector).CrossZXY, v)
}
func (vs Vectors) CrossYXZ(v Vector) {
	vs.ForEach((*Vector).CrossYXZ, v)
}

func (vs Vectors) CrossZYX(v Vector) {
	vs.ForEach((*Vector).CrossZYX, v)
}

func (vs Vectors) CrossXZY(v Vector) {
	vs.ForEach((*Vector).CrossXZY, v)
}

func (vs Vectors) ProjectR(v Vector) {
	vs.ForEach((*Vector).ProjectR, v)
}
func (vs Vectors) ProjectRR(v Vector) {
	vs.ForEach((*Vector).ProjectRR, v)
}

func (vs Vectors) ProjectYZX(v Vector) {
	vs.ForEach((*Vector).ProjectYZX, v)
}

func (vs Vectors) ProjectZXY(v Vector) {
	vs.ForEach((*Vector).ProjectZXY, v)
}

func (vs Vectors) ProjectYXZ(v Vector) {
	vs.ForEach((*Vector).ProjectYXZ, v)
}
func (vs Vectors) ProjectZYX(v Vector) {
	vs.ForEach((*Vector).ProjectZYX, v)
}
func (vs Vectors) ProjectXZY(v Vector) {
	vs.ForEach((*Vector).ProjectXZY, v)
}
