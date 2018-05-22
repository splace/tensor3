package tensor3

// component type
type BaseType float64

const scale = 1

func Base64(f float64) BaseType{
	return BaseType(f)
}

func Base32(f float32) BaseType{
	return BaseType(float64(f))
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


