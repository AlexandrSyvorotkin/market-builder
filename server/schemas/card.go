package schemas

type CardSchema struct {
	UserId uint16 `json:"userId"`
	Title string `json:"title"`
	Image string `json:"image"`
}
