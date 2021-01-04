package dto

type Subscription struct {
	Name  string `json:"name" conform:"trim" validate:"required,max=50"`
	Email string `json:"email" conform:"trim" validate:"required,email,max=150"`
}
