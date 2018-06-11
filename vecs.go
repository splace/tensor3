package tensor3

type Vectors []Vector

func NewVectors(cs ...BaseType) (vs Vectors) {
	vs = make(Vectors, (len(cs)+2)/3)
	for i := range vs {
		vs[i] = NewVector(cs[i*3:]...)
	}
	return
}

func (vs Vectors) Cross(v Vector) {
	vs.ForEach((*Vector).Cross, v)
}

func (vs Vectors) Add(v Vector) {
	vs.ForEach((*Vector).Add, v)
}

func (vs Vectors) Subtract(v Vector) {
	vs.ForEach((*Vector).Subtract, v)
}

func (vs Vectors) Project(v Vector) {
	vs.ForEach((*Vector).Project, v)
}

func (vs Vectors) Sum() (v Vector) {
	v.ForAll(vs, (*Vector).Add)
	return
}

func (vs Vectors) Multiply(s BaseType) {
	var multiply func(*Vector)
	multiply = func(v *Vector) {
		v.Multiply(s)
	}
	vs.ForEachNoParameter(multiply)
}

func (vs Vectors) Max() (v Vector) {
	v.Set(vs[0])
	v.ForAll(vs[1:], (*Vector).Max)
	return
}

func (vs Vectors) Min() (v Vector) {
	v.Set(vs[0])
	v.ForAll(vs[1:], (*Vector).Min)
	return
}

func (vs Vectors) Middle() (v Vector) {
	v = vs.Max()
	v.Mid(vs.Min())
	return
}

func (vs Vectors) Interpolate(v Vector, f float64) {
	f1 := Base64(1 - f)
	f2 := Base64(f)
	var interpolate func(*Vector, Vector)
	interpolate = func(v *Vector, v2 Vector) {
		v2.Multiply(f1)
		v.Multiply(f2)
		v.Add(v2)
	}
	vs.ForEach(interpolate, v)
}

func (vs Vectors) ForEach(fn func(*Vector, Vector), v Vector) {
	if !Parallel {
		vectorsApply(vs, fn, v)
	} else {
		vectorsApplyChunked(vs, fn, v)
	}
}

func vectorsApply(vs Vectors, fn func(*Vector, Vector), v Vector) {
	for i := range vs {
		fn(&vs[i], v)
	}
}

func vectorsApplyChunked(vs Vectors, fn func(*Vector, Vector), v Vector) {
	done := make(chan struct{}, 1)
	var running uint
	for chunk := range vectorsInChunks(vs, chunkSize(len(vs))) {
		running++
		go func(c Vectors) {
			vectorsApply(c, fn, v)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

// for each vector apply a function with no parameters
func (vs Vectors) ForEachNoParameter(fn func(*Vector)) {
	var inner func(*Vector, Vector)
	inner = func(v *Vector, _ Vector) {
		fn(v)
	}
	vs.ForEach(inner, Vector{})
}

func (vs Vectors) CrossAll(vs2 Vectors) {
	vs.ForAll((*Vector).Cross, vs2)
}

func (vs Vectors) AddAll(vs2 Vectors) {
	vs.ForAll((*Vector).Add, vs2)
}

func (vs Vectors) SubtractAll(vs2 Vectors) {
	vs.ForAll((*Vector).Subtract, vs2)
}

func (vs Vectors) ProjectAll(vs2 Vectors) {
	vs.ForAll((*Vector).Project, vs2)
}

func (vs Vectors) ForAll(fn func(*Vector, Vector), vs2 Vectors) {
	if !Parallel {
		vectorsApplyAll(vs, fn, vs2)
	} else {
		vectorsApplyAllChunked(vs, fn, vs2)
	}
}

func vectorsApplyAll(vs Vectors, fn func(*Vector, Vector), vs2 Vectors) {
	for i := range vs {
		fn(&vs[i], vs2[i])
	}
}

func vectorsApplyAllChunked(vs Vectors, fn func(*Vector, Vector), vs2 Vectors) {
	done := make(chan struct{}, 1)
	var running uint
	// shorten vs to use only what we have in vs2
	if len(vs) > len(vs2) {
		vs = vs[:len(vs2)]
	}
	cs := chunkSize(len(vs))
	chunks2 := vectorsInChunks(vs2, cs)
	for chunk := range vectorsInChunks(vs, cs) {
		running++
		go func(c Vectors) {
			vectorsApplyAll(c, fn, <-chunks2)
			done <- struct{}{}
		}(chunk)
	}
	for ; running > 0; running-- {
		<-done
	}
}

// search Vectors for the two Vector's that return the minimum value from the provided function
func (vs Vectors) SearchMin(toMin func(Vector, Vector) BaseType) (i, j int, value BaseType) {
	// TODO search in chunks
	value = toMin(vs[0], vs[1])
	var v1, v2 Vector
	var il, jl int = 0, 1
	for jl, v2 = range vs[2:] {
		nl := toMin(vs[0], v2)
		if nl < value {
			value, j = nl, jl+2
		}
	}

	for il, v1 = range vs[1:] {
		for jl, v2 = range vs[il+2:] {
			nl := toMin(v1, v2)
			if nl < value {
				value, i, j = nl, il+1, jl+il+2
			}
		}
	}
	return
}


func (b *BaseType) Aggregate(vs Vectors,length,stride int, fn func(*BaseType, Vectors)) {
	for vsc := range vectorSlicesInChunks(vs,chunkSize(len(vs)),length,stride,false) {
		for _,vs:=range(vsc){
			fn(b, vs)
		}
	}
}

//func(vs Vectors) Vector
//func(vs Vectors) BaseType


