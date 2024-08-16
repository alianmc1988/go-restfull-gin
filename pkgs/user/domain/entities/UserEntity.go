package user_entities

import user_models "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/models"


type UserEntity struct {
	ID  string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserSerialized struct {
	ID  string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func NewUserEntity(id string, firstName string, lastName string, email string, password string) *UserEntity {
	return &UserEntity{
		ID: id,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: password,
	}
}

func SerializeUserEntity(u *user_models.User) *UserSerialized {
	return &UserSerialized{
		ID: u.ID,
		FirstName: u.FirstName,
		LastName: u.LastName,
		Email: u.Email,
		Password: u.Password,
	}
}