package singleton_pattern

type Singleton interface {
	AddOne() int
	GetCount() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		return new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count += 1
	return s.count
}

func (s *singleton) GetCount() int {
	return s.count
}
