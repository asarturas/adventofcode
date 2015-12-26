package solution

type Action interface {
	Do(current bool) bool
}

type TurnOn struct {}

func (this TurnOn) Do(current bool) bool {
	return true;
}

type TurnOff struct {}

func (this TurnOff) Do(current bool) bool {
	return false;
}

type Toggle struct {}

func (this Toggle) Do(current bool) bool {
	return !current;
}

