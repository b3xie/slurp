// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateChannel struct {
	Name   string `json:"name"`
	Owner  string `json:"owner"`
	Public bool   `json:"public"`
}

type Login struct {
	Username string `json:"username"`
	Secret   string `json:"secret"`
}

type Message struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Mutation struct {
}

type NewUser struct {
	Username string `json:"username"`
	Secret   string `json:"secret"`
}

type Query struct {
}

type User struct {
	Username string `json:"username"`
	Secret   string `json:"secret"`
}
