package entity_example

type EntityExampleRepository interface {
	Find() (*EntityExample, error)
	Update(c EntityExample) (*EntityExample, error)
	Create(c EntityExample) (*EntityExample, error)
	Delete(c EntityExample) error
}

type EntityExampleBusiness interface {
	Find() (*EntityExample, error)
	Update(c EntityExample) (*EntityExample, error)
	Create(c EntityExample) (*EntityExample, error)
	Delete(c EntityExample) error
}
