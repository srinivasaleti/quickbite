package seeder

import (
	"embed"

	"github.com/srinivasaleti/quickbite/server/internal/database"
	productdb "github.com/srinivasaleti/quickbite/server/internal/domain/product/db"
	"github.com/srinivasaleti/quickbite/server/internal/domain/product/model"
	"github.com/srinivasaleti/quickbite/server/pkg/logger"
	"gopkg.in/yaml.v2"
)

// DataFS represent data that needed to be migrated.
//
//go:embed data
var DataFS embed.FS

// DataDir where migration data resides.
var DataDir = "data"

var ProductsFile = "data/products.yaml"
var CategoriesFile = "data/categories.yaml"

type IProductSeeder interface {
	SeedProducts() (*SeedProductsResponse, error)
}

type ProductSeeder struct {
	ProductDB productdb.IProductDB
	Logger    logger.ILogger
}

type SeedProductsResponse struct {
	Products   []model.Product
	Categories []model.Category
}

func (seeder *ProductSeeder) SeedProducts() (*SeedProductsResponse, error) {
	seeder.Logger.Info("seeding products")
	categories, err := seeder.seedCategories()
	if err != nil {
		seeder.Logger.Error(err, "Unable to seed data")
		return nil, err
	}
	products, err := seeder.seedProducts(categories)
	if err != nil {
		seeder.Logger.Error(err, "Unable to seed data")
		return nil, err
	}
	seeder.Logger.Info("successfully seeded products")
	return &SeedProductsResponse{Products: products, Categories: categories}, nil
}

func (seeder *ProductSeeder) seedCategories() ([]model.Category, error) {
	var categories []model.Category
	seeder.Logger.Info("reading categories data from file ...")
	if err := readData(CategoriesFile, &categories); err != nil {
		return nil, err
	}
	return seeder.ProductDB.InsertOrUpdateCategories(categories)
}

func (seeder *ProductSeeder) seedProducts(categories []model.Category) ([]model.Product, error) {
	var products []model.Product
	seeder.Logger.Info("reading categories data from file ...")
	if err := readData(ProductsFile, &products); err != nil {
		return nil, err
	}

	// Assign CategoryID to each product by matching CategoryName
	for i := range products {
		if products[i].CategoryName != nil {
			category := findCategoryByName(categories, *products[i].CategoryName)
			if category != nil {
				products[i].CategoryID = &category.ID
			}
		}
		products[i].PriceInCents = products[i].Price.ToCents()
	}
	return seeder.ProductDB.InsertOrUpdateProducts(products)
}

func findCategoryByName(categories []model.Category, name string) *model.Category {
	for _, c := range categories {
		if c.Name == name {
			return &c
		}
	}
	return nil
}

// readData reads the file specified by filePath and converts it to the given data type.
func readData(filePath string, dataType interface{}) error {
	file, err := DataFS.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(dataType)
	if err != nil {
		return err
	}
	return nil
}

func NewProductSeeder(logger logger.ILogger, db database.DB) IProductSeeder {
	return &ProductSeeder{
		ProductDB: productdb.NewProductDB(db),
		Logger:    logger,
	}
}
