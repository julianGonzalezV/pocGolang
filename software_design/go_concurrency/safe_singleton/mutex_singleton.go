package safe_singleton

/// Note lo limpio que queda la implementacion con Mutex

import "sync"

type singletonM struct {
	count int
	sync.RWMutex// acá usamos sync.RWMutex en lugar de sync.Mutex, debido a que 
	// es un bloqueo de lectura y escritura, en la lectura no habrá espera siempre y cuando 
	// no exista una operación de escritura ..así multiples goRoutines podrán leer al mismo tiempo 
	// Note que puede estár superando el performance de la implementación con channels
}

var instanceM singletonM

func GetInstanceM() *singletonM {
	return &instanceM
}
func (s *singletonM) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}
func (s *singletonM) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}
