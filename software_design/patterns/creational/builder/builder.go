package builder

/*
Abstract complex creations so that object creation is separated from the object
user
Create an object step by step by filling its fields and creating the embedded
objects
Reuse the object creation algorithm between many objects
*/

type BuildProcess interface {
	// Note that in various functions the same type BuildProcess
	// in order to chaing several steps
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	/// GetVehicle return the vehicle instance from the builder
	Build() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}
