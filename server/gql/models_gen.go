// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

type Programmer struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Picture *string  `json:"picture"`
	Company string   `json:"company"`
	Skills  []*Skill `json:"skills"`
}

type Skill struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Icon       *string `json:"icon"`
	Importance int     `json:"importance"`
}