package tensor3

func (v *Vector) SetX(v2 Vector) {
	v.x = v2.x
}

func (v *Vector) SetY(v2 Vector) {
	v.y = v2.y
}

func (v *Vector) SetZ(v2 Vector) {
	v.z = v2.z
}

func (v *Vector) SetXY(v2 Vector) {
	v.x = v2.x
	v.y = v2.y
}

func (v *Vector) SetYZ(v2 Vector) {
	v.y = v2.y
	v.z = v2.z
}

func (v *Vector) SetXZ(v2 Vector) {
	v.x = v2.x
	v.z = v2.z
}

func (v *Vector) R() {
	v.YZX()
}

func (v *Vector) YZX() {
	v.y, v.z, v.x = v.x, v.y, v.z
}

func (v *Vector) RR() {
	v.ZXY()
}

func (v *Vector) ZXY() {
	v.z, v.x, v.y = v.x, v.y, v.z
}

func (v *Vector) YXZ() {
	v.y, v.x, v.z = v.x, v.y, v.z
}

func (v *Vector) ZYX() {
	v.z, v.y, v.x = v.x, v.y, v.z
}

func (v *Vector) XZY() {
	v.x, v.z, v.y = v.x, v.y, v.z
}

func (v Vector) DotR(v2 Vector) float64 {
	return v.DotYZX(v2)
}

func (v Vector) DotYZX(v2 Vector) float64 {
	return v.y*v2.x + v.z*v2.y + v.x*v2.z
}

func (v Vector) DotRR(v2 Vector) float64 {
	return v.DotZXY(v2)
}

func (v Vector) DotZXY(v2 Vector) float64 {
	return v.z*v2.x + v.x*v2.y + v.y*v2.z

}
func (v Vector) DotYXZ(v2 Vector) float64 {
	return v.y*v2.x + v.x*v2.y + v.z*v2.z
}

func (v Vector) DotZYX(v2 Vector) float64 {
	return v.z*v2.x + v.y*v2.y + v.x*v2.z
}

func (v Vector) DotXZY(v2 Vector) float64 {
	return v.x*v2.x + v.z*v2.y + v.y*v2.z
}

func (v *Vector) AddR(v2 Vector) {
	v.AddYZX(v2)
}

func (v *Vector) AddYZX(v2 Vector) {
	v.y += v2.x
	v.z += v2.y
	v.x += v2.z
}

func (v *Vector) AddRR(v2 Vector) {
	v.AddZXY(v2)
}

func (v *Vector) AddZXY(v2 Vector) {
	v.z += v2.x
	v.x += v2.y
	v.y += v2.z
}

func (v *Vector) AddYXZ(v2 Vector) {
	v.y += v2.x
	v.x += v2.y
	v.z += v2.z
}
func (v *Vector) AddZYX(v2 Vector) {
	v.z += v2.x
	v.y += v2.y
	v.x += v2.z
}
func (v *Vector) AddXZY(v2 Vector) {
	v.x += v2.x
	v.z += v2.y
	v.y += v2.z
}

func (v *Vector) SubtractR(v2 Vector) {
	v.SubtractYZX(v2)
}

func (v *Vector) SubtractYZX(v2 Vector) {
	v.y -= v2.x
	v.z -= v2.y
	v.x -= v2.z
}

func (v *Vector) SubtractRR(v2 Vector) {
	v.SubtractZXY(v2)
}

func (v *Vector) SubtractZXY(v2 Vector) {
	v.z -= v2.x
	v.x -= v2.y
	v.y -= v2.z
}

func (v *Vector) SubtractYXZ(v2 Vector) {
	v.y -= v2.x
	v.x -= v2.y
	v.z -= v2.z
}
func (v *Vector) SubtractZYX(v2 Vector) {
	v.z -= v2.x
	v.y -= v2.y
	v.x -= v2.z
}
func (v *Vector) SubtractXZY(v2 Vector) {
	v.x -= v2.x
	v.z -= v2.y
	v.y -= v2.z
}

func (v *Vector) CrossR(v2 Vector) {
	v.CrossYZX(v2)
}
func (v *Vector) CrossYZX(v2 Vector) {
	v.y, v.z, v.x = v.y*v2.x-v.z*v2.z, v.z*v2.y-v.x*v2.x, v.x*v2.z-v.y*v2.y
}

func (v *Vector) CrossRR(v2 Vector) {
	v.CrossZXY(v2)
}

func (v *Vector) CrossZXY(v2 Vector) {
	v.z, v.x, v.y = v.y*v2.y-v.z*v2.x, v.z*v2.z-v.x*v2.y, v.x*v2.x-v.y*v2.z
}

func (v *Vector) CrossYXZ(v2 Vector) {
	v.y, v.x, v.z = v.y*v2.y-v.z*v2.x, v.z*v2.z-v.x*v2.y, v.x*v2.x-v.y*v2.z
}

func (v *Vector) CrossZYX(v2 Vector) {
	v.z, v.y, v.x = v.y*v2.y-v.z*v2.x, v.z*v2.z-v.x*v2.y, v.x*v2.x-v.y*v2.z
}

func (v *Vector) CrossXZY(v2 Vector) {
	v.x, v.z, v.y = v.y*v2.y-v.z*v2.x, v.z*v2.z-v.x*v2.y, v.x*v2.x-v.y*v2.z
}

func (v *Vector) ProjectR(v2 Vector) {
	v.ProjectYZX(v2)
}

func (v *Vector) ProjectYZX(v2 Vector) {
	v.y *= v2.x
	v.z *= v2.y
	v.x *= v2.z
}

func (v *Vector) ProjectRR(v2 Vector) {
	v.ProjectZXY(v2)
}

func (v *Vector) ProjectZXY(v2 Vector) {
	v.z *= v2.x
	v.x *= v2.y
	v.y *= v2.z
}

func (v *Vector) ProjectYXZ(v2 Vector) {
	v.y *= v2.x
	v.x *= v2.y
	v.z *= v2.z
}
func (v *Vector) ProjectZYX(v2 Vector) {
	v.z *= v2.x
	v.y *= v2.y
	v.x *= v2.z
}
func (v *Vector) ProjectXZY(v2 Vector) {
	v.x *= v2.x
	v.z *= v2.y
	v.y *= v2.z
}

func (v *Vector) ProductR(m Matrix) {
	v.ProductYZX(m)
}

func (v *Vector) ProductYZX(m Matrix) {
	v.y, v.z, v.x = v.x*m.x.x+v.y*m.y.x+v.z*m.z.x, v.x*m.x.y+v.y*m.y.y+v.z*m.z.y, v.x*m.x.z+v.y*m.y.z+v.z*m.z.z
}

func (v *Vector) ProductRR(m Matrix) {
	v.ProductZXY(m)
}

func (v *Vector) ProductZXY(m Matrix) {
	v.z, v.x, v.y = v.x*m.x.x+v.y*m.y.x+v.z*m.z.x, v.x*m.x.y+v.y*m.y.y+v.z*m.z.y, v.x*m.x.z+v.y*m.y.z+v.z*m.z.z
}

func (v *Vector) ProductYXZ(m Matrix) {
	v.y, v.x, v.z = v.x*m.x.x+v.y*m.y.x+v.z*m.z.x, v.x*m.x.y+v.y*m.y.y+v.z*m.z.y, v.x*m.x.z+v.y*m.y.z+v.z*m.z.z
}

func (v *Vector) ProductZYX(m Matrix) {
	v.z, v.y, v.x = v.x*m.x.x+v.y*m.y.x+v.z*m.z.x, v.x*m.x.y+v.y*m.y.y+v.z*m.z.y, v.x*m.x.z+v.y*m.y.z+v.z*m.z.z
}
func (v *Vector) ProductXZY(m Matrix) {
	v.x, v.z, v.y = v.x*m.x.x+v.y*m.y.x+v.z*m.z.x, v.x*m.x.y+v.y*m.y.y+v.z*m.z.y, v.x*m.x.z+v.y*m.y.z+v.z*m.z.z
}

func (v *Vector) ProductTR(m Matrix) {
	v.ProductTYZX(m)
}

func (v *Vector) ProductTYZX(m Matrix) {
	v.y, v.z, v.x = v.x*m.x.x+v.y*m.x.y+v.z*m.x.z, v.x*m.y.x+v.y*m.y.y+v.z*m.y.z, v.x*m.z.x+v.y*m.z.y+v.z*m.z.z
}

func (v *Vector) ProductTRR(m Matrix) {
	v.ProductTZXY(m)
}

func (v *Vector) ProductTZXY(m Matrix) {
	v.z, v.x, v.y = v.x*m.x.x+v.y*m.x.y+v.z*m.x.z, v.x*m.y.x+v.y*m.y.y+v.z*m.y.z, v.x*m.z.x+v.y*m.z.y+v.z*m.z.z
}

func (v *Vector) ProductTYXZ(m Matrix) {
	v.y, v.x, v.z = v.x*m.x.x+v.y*m.x.y+v.z*m.x.z, v.x*m.y.x+v.y*m.y.y+v.z*m.y.z, v.x*m.z.x+v.y*m.z.y+v.z*m.z.z
}

func (v *Vector) ProductTZYX(m Matrix) {
	v.z, v.y, v.x = v.x*m.x.x+v.y*m.x.y+v.z*m.x.z, v.x*m.y.x+v.y*m.y.y+v.z*m.y.z, v.x*m.z.x+v.y*m.z.y+v.z*m.z.z
}
func (v *Vector) ProductTXZY(m Matrix) {
	v.x, v.z, v.y = v.x*m.x.x+v.y*m.x.y+v.z*m.x.z, v.x*m.y.x+v.y*m.y.y+v.z*m.y.z, v.x*m.z.x+v.y*m.z.y+v.z*m.z.z
}
