package vector

type Vector struct {
	X, Y int
}

func AddVectors(v, u Vector) Vector {
	return Vector{
		v.X + u.X,
		v.Y + u.Y,
	}
}
