package models

type Admin struct {
	User
	pages map[string]string
}

func NewAdmin(firstName, lastName, birthDate string) (*Admin, error) {
	user, err := NewUser(firstName, lastName, birthDate)
	if err != nil {
		return nil, err
	}
	return &Admin{
		User:  *user,
		pages: map[string]string{},
	}, nil
}
