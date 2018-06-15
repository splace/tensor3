package tensor3

// component type
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

// new vector reference from float32's 
func New(x,y,z float64) *Vector{
	return &Vector{BaseType(x),BaseType(y),BaseType(z)}
}


