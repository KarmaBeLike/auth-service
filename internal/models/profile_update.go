package models

type ProfileUpdater struct {
	Email         string `json:"email" validate:"email,required"`
	Name          string `json:"name" validate:"alphaunicode,required"`
	Phone         string `json:"phone" validate:"e164,required"`
	BirthDate     string `json:"bDate" validate:"date,required"`
}