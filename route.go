package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{{1, "title 1", "text 1"}}
}

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(result)

}

func addPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "applicaation/json")
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}

	post.ID = len(posts) + 1
	posts = append(posts, post)
	result, err := json.Marshal(post)
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}
