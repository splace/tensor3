package tensor3

// this file contains what is different between typed versions of this package
// other files can/are just copies,(symlinks) they are generic.

// notice in this Scalar scaling is a no-op, inlined/removed by all but the most basic compiler.

// base component type for this version of the package.
type Scalar float64

const scale = 1

func Base64(f float64) Scalar {
	return Scalar(f)
}

func Base32(f float32) Scalar {
	return Scalar(float64(f))
}

func baseScale(v Scalar) Scalar {
	return v
}

func baseUnscale(v Scalar) Scalar {
	return v
}

func vectorScale(v *Vector) {
	return
}

func vectorUnscale(v *Vector) {
	return
}

// new vector reference from float64's
func New(x, y, z float64) *Vector {
	return &Vector{Scalar(x), Scalar(y), Scalar(z)}
}/* benchmark: "Component" hal3 Tue 11 Dec 16:20:49 GMT 2018 go version go1.11.2 linux/amd64
FAIL	_/run/media/simon/6a5530c2-1442-4e9b-b35f-3db0c9a6984c/home/simon/Dropbox/github/working/tensor3 [build failed]
Tue 11 Dec 16:20:49 GMT 2018
*/

