package tensor3
import "fmt"

// component type
type BaseType int

const scale=1000

func baseScale(v BaseType) BaseType{
	return v*scale
}

func baseUnscale(v BaseType) BaseType{
	return v/scale
}

func vectorScale(v *Vector){
	v.x*=scale
	v.y*=scale
	v.z*=scale
	return
}

func vectorUnscale(v *Vector){
	v.x/=scale
	v.y/=scale
	v.z/=scale
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

