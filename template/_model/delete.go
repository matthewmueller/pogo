package model

type delete struct {
}

func (d *delete) Generate() (string, error) {
	return "", nil
}

type deleteBy struct {
}

func (d *deleteBy) Generate() (string, error) {
	return "", nil
}

type deleteMany struct {
}

func (d *deleteMany) Generate() (string, error) {
	return "", nil
}
