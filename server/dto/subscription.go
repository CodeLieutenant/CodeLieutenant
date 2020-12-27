package dto

type Subscription struct {
	Email string `json:"email" conform:"trim" validate:"required,email,max=150"`
}
