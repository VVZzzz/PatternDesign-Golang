package observe

type Observer interface {
	Update(observable *Observable, data interface{})
}
