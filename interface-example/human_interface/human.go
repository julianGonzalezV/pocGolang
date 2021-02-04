package human_interface

type Human struct {
	Name string
}

/// NamePrinter
func (h Human) NamePrinter() string {
	return h.Name + "printer"
}
