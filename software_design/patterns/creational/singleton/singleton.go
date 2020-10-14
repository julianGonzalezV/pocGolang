package singleton

/// Nota: en varios lenguajes like java songleton es considerado un anti-pattern
/*
one reason: Static members can cause zommbie instances (old java version)

Main reason:

For example, and this applies to static classes also, if you unit test, or have concurrency, then the state of one request will change the state and that may cause problems, as the class calling the instance may be assuming the state is as it expected.
I think the best way to challenge the use is to think of how to handle it if your program is multi-threaded, and a simple way to do that is to unit test it, if you have several tests that run at one time.
If you find that you still need it, then use it, but realize the problems that will be encountered later.
by author of https://stackoverrun.com/es/q/3426227
*/

/* Pero en Go no existe el famoso Static, en su lugar se hace uso de package scope
para lograr un comportamient similar*/

///It is recommended start from declared-only function  when using TDD

/// This implementation is not thread safe!!!!, please see the safe one in ..(in progress)
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

// acá por defecto se incia en nil
var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
