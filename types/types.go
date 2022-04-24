package types

type ImgData struct {
	ID int `json:"id"`
	Title string `json:"title"`
	URL string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
	AlbumID int `json:"albumId"`
}

type UserData struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Address struct {
		Street string `json:"street"`
		Suite string `json:"suite"`
		City string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs string `json:"bs"`
	} `json:"company"`
}
type PostData struct {
	UserId int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}
type CommentData struct {
	PostId int `json:"postId"`
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Body string `json:"body"`
}
