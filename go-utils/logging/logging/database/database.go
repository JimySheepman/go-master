package database

type dummyDatabase struct {
}

func NewDummyDatabase() *dummyDatabase {
	return &dummyDatabase{}
}

func (d *dummyDatabase) Create() error {
	return nil
}
