package core

type Action interface {
	Do() (Result, error)
}

type Validable interface {
	Validate() error
}

type Result interface {
}

func Do(a Action) (Result, error) {
	if a, ok := a.(Validable); ok {
		err := a.Validate()
		if err != nil {
			return nil, err
		}
	}

	return a.Do()
}
