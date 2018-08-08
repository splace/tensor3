package tensor3
import "fmt"

// this file contains what is different between typed versions of this package
// other files can/are just copies,(symlinks) they are generic.

// base component type for this version of the package.
type BaseType int

const scaleShift=10
const scale = 1<<scaleShift

func Base64(f float64) BaseType{
	return BaseType(f*float64(scale))
}

func Base32(f float32) BaseType{
	return BaseType(f*float32(scale))
}

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

// this base type needs to be scaled for printing 
func (v BaseType) String()string{
	return fmt.Sprint(float32(v)/float32(scale))
}

func (v Vector) String()string{
	return fmt.Sprintf("{%s %s %s}",v.x,v.y,v.z)
}

// TODO scan scaled?

// new vector reference from float64's 
func New(x,y,z float64) *Vector{
	return &Vector{Base64(x),Base64(y),Base64(z)}
}



