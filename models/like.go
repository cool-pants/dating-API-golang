package models

type Like struct {
	Id    int64 `json:"id"`
	Liker int64 `json:"who_likes"`
	Likee int64 `json:"who_is_liked"`
}