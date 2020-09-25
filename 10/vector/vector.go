package vector

import "math"

type Vector struct {
	X, Y float64
}

func GetPositionVector(x, y int) Vector {
	return Vector{
		X: float64(x),
		Y: float64(y),
	}
}

func GetVectorRelativeTo(origin, v Vector) Vector {
	return Vector{
		X: v.X - origin.X,
		Y: origin.Y - v.Y,
	}
}

func GetScaledDownVector(v Vector) Vector {
	maxComponent := math.Max(math.Abs(v.X), math.Abs(v.Y))
	return Vector{
		X: v.X / maxComponent,
		Y: v.Y / maxComponent,
	}
}

func GetMagnitude(v Vector) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func GetAngle(origin, v Vector) float64 {
	radians := math.Atan2(origin.Y-v.Y, v.X-origin.X)
	if radians <= 0 {
		radians = 2*math.Pi + radians
	}

	return radians
}

func ToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
