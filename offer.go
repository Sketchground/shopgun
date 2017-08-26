package shopgun

// Offers list of offers
type Offers []Offer

// Offer type
type Offer struct {
	ID          string `json:"id"`
	Ern         string `json:"ern"`
	Heading     string `json:"heading"`
	Description string `json:"description"`
	CatalogPage int    `json:"catalog_page"`
	Pricing     `json:"pricing"`
	Quantity    `json:"quantity"`
	Images      Images `json:"images"`
	Links       struct {
		Webshop interface{} `json:"webshop"`
	} `json:"links"`
	RunFrom     string   `json:"run_from"`
	RunTill     string   `json:"run_till"`
	DealerURL   string   `json:"dealer_url"`
	StoreURL    string   `json:"store_url"`
	CatalogURL  string   `json:"catalog_url"`
	DealerID    string   `json:"dealer_id"`
	StoreID     string   `json:"store_id"`
	CatalogID   string   `json:"catalog_id"`
	CategoryIds []string `json:"category_ids"`
	Branding    Branding `json:"branding"`
}

// Pricing type
type Pricing struct {
	Price    int    `json:"price"`
	PrePrice int    `json:"pre_price"`
	Currency string `json:"currency"`
}

// Quantity type
type Quantity struct {
	Unit   Unit  `json:"unit"`
	Size   Range `json:"size"`
	Pieces Range `json:"pieces"`
}

// Unit type
type Unit struct {
	Symbol string `json:"symbol"`
	Si     Si     `json:"si"`
}

// Si type
type Si struct {
	Symbol string `json:"symbol"`
	Factor int    `json:"factor"`
}

// Range type
type Range struct {
	From float64 `json:"from"`
	To   float64 `json:"to"`
}
