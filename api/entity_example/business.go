package entity_example

type EntityExampleBusinessImpl struct {
	entityExampleRepository EntityExampleRepository
}

func NewEntityExampleBusinessImpl(entityExampleRepository EntityExampleRepository) *EntityExampleBusinessImpl {
	return &EntityExampleBusinessImpl{entityExampleRepository}
}

func (b EntityExampleBusinessImpl) Find() (*EntityExample, error) {
	// TODO impl this
	return nil, nil
}

func (b EntityExampleBusinessImpl) Update(c EntityExample) (*EntityExample, error) {
	// TODO impl this
	return nil, nil
}

func (b EntityExampleBusinessImpl) Create(c EntityExample) (*EntityExample, error) {
	// TODO impl this
	return nil, nil
}

func (b EntityExampleBusinessImpl) Delete(c EntityExample) error {
	// TODO impl this
	return nil
}
