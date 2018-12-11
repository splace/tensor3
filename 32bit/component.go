package tensor3
// this file contains what is different between typed versions of this package
// other files can/are just copies,(symlinks) they are generic.

// notice in this, float32, base type version, scaling is a no-op, inlined/removed by all but the most basic compiler. 

// base component type for this version of the package.
type BaseType float32

const scale = 1

func Base64(f float64) BaseType{
	return BaseType(float32(f))
}

func Base32(f float32) BaseType{
	return BaseType(f)
}


func baseScale(v BaseType) BaseType {
	return v
}

func baseUnscale(v BaseType) BaseType {
	return v
}

func vectorScale(v *Vector) {
	return
}

func vectorUnscale(v *Vector) {
	return
}

// new vector reference from float64's 
func New(x,y,z float64) *Vector{
	return &Vector{BaseType(x),BaseType(y),BaseType(z)}
}


