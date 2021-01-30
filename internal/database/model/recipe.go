package model

type Recipe struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Calories    int    `json:"calories"`
	Time        int    `json:"time"`
	ImageUrl    string `json:"imageUrl"`
	Author      string `json:"author"`
}
