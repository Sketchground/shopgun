package shopgun

// Dealers list
type Dealers []Dealer

// Dealer type
type Dealer struct {
	ID          string      `json:"id"`
	Ern         string      `json:"ern"`
	Name        string      `json:"name"`
	URLName     string      `json:"url_name"`
	Website     string      `json:"website"`
	Logo        string      `json:"logo"`
	Color       string      `json:"color"`
	Pageflip    Pageflip    `json:"pageflip"`
	CategoryIds []string    `json:"category_ids"`
	Country     interface{} `json:"country"`
}
