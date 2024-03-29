package model

import "time"

type LogPayload struct {
	UserId    int     `json:"user_id"`
	Total     float32 `json:"total"`
	Title     string  `json:"title"`
	Meta      Meta    `json:"meta"`
	Completed bool    `json:"completed"`
}

type Meta struct {
	Logins       []Login     `json:"logins"`
	PhoneNumbers PhoneNumber `json:"phone_numbers"`
}

type Login struct {
	Time time.Time `json:"time"`
	IP   string    `json:"ip"`
}

type PhoneNumber struct {
	Home   string `json:"home"`
	Mobile string `json:"mobile"`
}
