package tensor3

// component type
type BaseType float32

const scale = 1

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

func New(x,y,z float32) *Vector{
	return &Vector{BaseType(x),BaseType(y),BaseType(z)}
}


