package tensor3
import "fmt"

// component type
type BaseType int

const scaleShift=10
const scale = 1<<scaleShift

// TODO could use base two scale

func baseScale(v BaseType) BaseType{
	return v<<scaleShift
}

func baseUnscale(v BaseType) BaseType{
	return v>>scaleShift
}

func vectorScale(v *Vector){
	v.x<<=scaleShift
	v.y<<=scaleShift
	v.z<<=scaleShift
	return
}

func vectorUnscale(v *Vector){
	v.x>>=scaleShift
	v.y>>=scaleShift
	v.z>>=scaleShift
	return
}


func (v BaseType) String()string{
	return fmt.Sprint(float64(v)/float64(scale))
}

func (v Vector) String()string{
	return fmt.Sprintf("{%v %v %v}",v.x,v.y,v.z)
}

func (v Matrix) String()string{
	return fmt.Sprintf("{%v %v %v}",v.x,v.y,v.z)
}

