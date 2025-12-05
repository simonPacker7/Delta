package entities

type RegisterUserInput struct {
	Name     string `redis:"name" json:"name" validate:"required,min=3,max=40"`
	Email    string `redis:"email" json:"email" validate:"required,email"`
	Password string `redis:"password" json:"password" validate:"required,min=6,max=50"`
}

type LoginUserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

type UserProfile struct {
	Email string `json:"email" redis:"email"`
	Name  string `json:"name" redis:"name"`
	ID    string `json:"id" redis:"id"`
}
