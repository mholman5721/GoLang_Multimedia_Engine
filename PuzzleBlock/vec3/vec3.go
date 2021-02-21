package vec3

import "math"

// Vector3 represents the X, Y, and Z components of a 3d vector
type Vector3 struct {
	X, Y, Z float32
}

// Add returns a new vector that is the sum of the components of two other vectors
func Add(a, b Vector3) Vector3 {
	return Vector3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Mult returns a new vector that consists of the components of vector 'a' scaled by value 'b'
func Mult(a Vector3, b float32) Vector3 {
	return Vector3{a.X * b, a.Y * b, a.Z * b}
}

// Length returns the length of vector vec3
func (a Vector3) Length() float32 {
	return float32(math.Sqrt(float64(a.X*a.X + a.Y + a.Y + a.Z*a.Z)))
}

// Distance returns the distance between two vectors
func Distance(a, b Vector3) float32 {
	xDiff := a.X - b.X
	yDiff := a.Y - b.Y
	zDiff := a.Z - b.Z
	return float32(math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff + zDiff*zDiff)))
}

// DistanceSquared returns the squared distance between two vectors - useful for telling whether one thing is closer/ further than another
func DistanceSquared(a, b Vector3) float32 {
	xDiff := a.X - b.X
	yDiff := a.Y - b.Y
	zDiff := a.Z - b.Z
	return xDiff*xDiff + yDiff*yDiff + zDiff*zDiff
}

// Normalize returns a new vector in the direction of vector 'a,' but with magnitude 1
func Normalize(a Vector3) Vector3 {
	len := a.Length()
	return Vector3{a.X / len, a.Y / len, a.Z / len}
}
