package handler

import (
	"fmt"
	"net/http"
)

type Post struct{}

func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a post")
}

func (p *Post) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all posts")
}

func (p *Post) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a post by ID")
}

func (p *Post) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a post by ID")
}

func (p *Post) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a post by ID")
}
