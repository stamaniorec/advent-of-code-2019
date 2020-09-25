package vector

import "math"

type Vector struct {
	X, Y, Z int
}

func AddVectors(v, u Vector) Vector {
	return Vector{
		X: v.X + u.X,
		Y: v.Y + u.Y,
		Z: v.Z + u.Z,
	}
}

func GetAbsoluteVector(vec Vector) Vector {
	return Vector{
		X: int(math.Abs(float64(vec.X))),
		Y: int(math.Abs(float64(vec.Y))),
		Z: int(math.Abs(float64(vec.Z))),
	}
}
