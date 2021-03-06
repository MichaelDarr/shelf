// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/google/uuid"
)

type Node interface {
	IsNode()
}

type User struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

func (User) IsNode() {}

type UserCreateInput struct {
	Email string `json:"email"`
}

type UserInput struct {
	ID uuid.UUID `json:"id"`
}
