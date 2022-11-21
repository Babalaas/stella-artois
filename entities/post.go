package entities

type Post struct {
	ID            uint   `json:"id"`
	User_ID       uint   `json:"user_id"`
	Collection_ID uint   `json:"collection_id"`
	Caption       string `json:"caption"`
	Location      string `json:"location"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Image         string `json:"image"`
	Image2        string `json:"image2"`
}
