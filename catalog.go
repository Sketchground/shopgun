package shopgun

// Catalog type
type Catalog struct {
	ID          string        `json:"id"`
	Ern         string        `json:"ern"`
	Label       interface{}   `json:"label"`
	Background  interface{}   `json:"background"`
	RunFrom     Time          `json:"run_from"`
	RunTill     Time          `json:"run_till"`
	PageCount   int           `json:"page_count"`
	OfferCount  int           `json:"offer_count"`
	Branding    Branding      `json:"branding"`
	DealerID    string        `json:"dealer_id"`
	DealerURL   string        `json:"dealer_url"`
	StoreID     string        `json:"store_id"`
	StoreURL    string        `json:"store_url"`
	Dimensions  Dimensions    `json:"dimensions"`
	Images      Images        `json:"images"`
	CategoryIds []interface{} `json:"category_ids"`
	PdfURL      interface{}   `json:"pdf_url"`
}

// Branding type
type Branding struct {
	Name           string      `json:"name"`
	URLName        string      `json:"url_name"`
	Website        string      `json:"website"`
	Logo           string      `json:"logo"`
	LogoBackground interface{} `json:"logo_background"`
	Color          string      `json:"color"`
	Pageflip       Pageflip    `json:"pageflip"`
}

// Pageflip type
type Pageflip struct {
	Logo  string `json:"logo"`
	Color string `json:"color"`
}

// Dimensions type
type Dimensions struct {
	Height float64 `json:"height"`
	Width  int     `json:"width"`
}

// Images type
type Images struct {
	Thumb string `json:"thumb"`
	View  string `json:"view"`
	Zoom  string `json:"zoom"`
}
