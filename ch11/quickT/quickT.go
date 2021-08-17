package quickT

type Point2D struct {
	X, Y int
}

func Add(x1, x2 Point2D) Point2D {
	temp := Point2D{}
	temp.X = x1.X + x2.X
	temp.Y = x1.Y + x2.Y
	return temp
}
