package user_models

import (
	user_value_objects "github.com/alianmc1988/go-restfull-gin/pkgs/user/domain/valueObjects"
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID  	uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" primary_key:"true"`
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
}

func (u *User) Serialize () *UserSerialized {
	return &UserSerialized{
		ID: u.ID.UUID.String(),
		FirstName: u.FirstName,
		LastName: u.LastName,
		Email: u.Email,
	}
}

func NewUser(u *user_value_objects.CreateUserDTO) *User {
	return &User{
		ID: uuid.UUID{},
		FirstName: u.FirstName,
		LastName: u.LastName,
		Email: u.Email,
		Password: hashPassword(u.Password),
	}
}

func hashPassword(p string) string {
	hash, err:= bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func (u *User) CheckPassword(p string) bool {
	err:= bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	return err == nil
}