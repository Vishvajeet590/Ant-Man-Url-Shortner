package entity

type User struct {
	Username string
	Id       int32
	Roles    string
	Token    string
}

func NewUser(usename, roles, token string, id int32) *User {

	return &User{
		Username: usename,
		Id:       id,
		Roles:    roles,
		Token:    token,
	}

}
