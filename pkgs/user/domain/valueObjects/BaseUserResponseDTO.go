package user_value_objects



type BaseUserResponseDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}


func NewBaseUserResponseDTO(
	id string,
	firstName string,
	lastName string,
	email string,
) *BaseUserResponseDTO {
	return &BaseUserResponseDTO{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}