package unit_test

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetByID(id int) (*User, error)
}

func GetUser(repo UserRepository, id int) (*User, error) {
	return repo.GetByID(id)
}
