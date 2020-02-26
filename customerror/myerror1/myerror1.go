package myerror1

type MyError1 struct {
	Desc   string
	Radius float64
}

func (e *MyError1) Error() string {
	return e.Desc
}

func (e *MyError1) GetRadius() float64 {
	return e.Radius
}

func (e *MyError1) getDesc() string {
	return e.Desc
}
