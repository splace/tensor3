package tensor3

import "testing"
import "fmt"

func TestVecSwizzleR(t *testing.T) {
	v := Vector{1, 2, 3}
	v.R()
	if fmt.Sprint(v) != "{3 1 2}" {
		t.Error(v)
	}
}

func TestVecSwizzleYZX_ZXY(t *testing.T) {
	v := Vector{1, 2, 3}
	v.YZX()
	v.ZXY()
	if fmt.Sprint(v) != "{1 2 3}" {
		t.Error(v)
	}
}

func TestVecSwizzleYZX_ZXY_ZYX_XZY(t *testing.T) {
	v := Vector{1, 2, 3}
	v.YXZ()
	v.XZY()
	v.ZYX()
	v.XZY()
	if fmt.Sprint(v) != "{1 2 3}" {
		t.Error(v)
	}
}

func TestVecSwizzleDotR(t *testing.T) {
	v := Vector{1, 2, 3}
	if v.DotR(Vector{5, 6, 4}) != 32.0 {
		t.Error(v)
	}
}

func TestVecSwizzleDotYZX(t *testing.T) {
	v := Vector{1, 2, 3}
	if v.DotYZX(Vector{5, 6, 4}) != 32.0 {
		t.Error(v)
	}
}

/*

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
*/
