package formas

type Quadrado struct{
	Lado float64
}

func (q Quadrado) Area() float64{
	return q.Lado * q.Lado
}
