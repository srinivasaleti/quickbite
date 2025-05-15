package model

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name" yaml:"name"`
}

type Product struct {
	ID           string  `json:"id" yaml:"id"`
	ExternalID   *string `json:"externalID,omitempty" yaml:"externalID"`
	Name         string  `json:"name" yaml:"name"`
	Price        float64 `json:"price" yaml:"price"`
	ImageURL     string  `json:"imageUrl" yaml:"imageUrl"`
	CategoryName *string `json:"category,omitempty" yaml:"categoryName,omitempty"`
	CategoryID   *string `json:"categoryId,omitempty" yaml:"categoryId,omitempty"`
}
