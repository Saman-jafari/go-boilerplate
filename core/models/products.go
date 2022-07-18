package models

type Product struct {
	Id                    string `json:"id"`
	ProductFamilyId       string `json:"product_family_id"`
	ProductNr             string `json:"product_nr"`
	Name                  string `json:"name"`
	Shortname             string `json:"shortname"`
	Segment               string `json:"segment"`
	Manufacturer          string `json:"manufacturer"`
	ManufacturerArticleNr string `json:"manufacturer_article_nr"`
	Status                string `json:"status"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
	Descriptions          string `json:"descriptions"`
	Articles              string `json:"articles"`
	ArticlesCount         string `json:"articles_count"`
	ProductFamily         string `json:"productFamily"`
	ReferenceKey          string `json:"referenceKey"`
}

func GetProducts() (Product, error) {
	return Product{}, nil
}
