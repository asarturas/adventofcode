package solution

type Light struct {
	Brightness int
}

func (this Light) IsOn() bool {
	return this.Brightness > 0;
}

func (this Light) IsOff() bool {
	return this.Brightness == 0;
}
