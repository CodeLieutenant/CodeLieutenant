package dto

type Contact struct {
	Name    string `json:"name" conform:"trim" validate:"required,alphanumericunicodespace,max=50"`
	Email   string `json:"email" conform:"trim,email" validate:"required,email,max=150"`
	Subject string `json:"subject" conform:"trim" validate:"required,alphanumericunicodespace,min=3,max=150"`
	Message string `json:"message" conform:"trim" validate:"required,min=3,max=1000"`
}
