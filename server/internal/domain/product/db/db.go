package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/srinivasaleti/quickbite/server/internal/database"
	"github.com/srinivasaleti/quickbite/server/internal/domain/product/model"
)

var ErrNoProductFound = errors.New("product not found")

type IProductDB interface {
	GetProducts() ([]model.Product, error)
	GetProductById(id string) (*model.Product, error)
	InsertOrUpdateCategories(categories []model.Category) ([]model.Category, error)
	InsertOrUpdateProducts(products []model.Product) ([]model.Product, error)
}

type ProductDB struct {
	DB database.DB
}

func (db *ProductDB) GetProducts() ([]model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), database.DefaultDBOperationTimeout)
	defer cancel()

	query := `
		SELECT 
			p.id, 
			p.name, 
			p.price, 
			p.image_url, 
			c.name AS category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
	`

	rows, err := db.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
			&p.ImageURL,
			&p.CategoryName,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

// InsertOrUpdateCategoriesBatch inserts or updates multiple categories in the database.
// It uses a batch to send all insert/update queries together for better performance.
// Insert happens if category name does not exist.
// Update happens if category name already exists (based on unique name constraint).
func (db *ProductDB) InsertOrUpdateCategories(categories []model.Category) ([]model.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), database.DefaultDBOperationTimeout)
	defer cancel()

	batch := &pgx.Batch{}
	for _, cat := range categories {
		sql := `
		INSERT INTO categories (name)
		VALUES (@name)
		ON CONFLICT (name) DO UPDATE
		SET name = EXCLUDED.name,
			updated_at = NOW()
		RETURNING id;
		`
		batch.Queue(sql, pgx.NamedArgs{
			"name": cat.Name,
		})
	}

	br := db.DB.SendBatch(ctx, batch)
	defer br.Close()

	for i := range categories {
		var id string
		err := br.QueryRow().Scan(&id)
		if err != nil {
			return nil, err
		}
		categories[i].ID = id
	}

	return categories, nil
}

// InsertOrUpdateProducts inserts new products or updates existing ones based on external_id.
// If a product with the same external_id already exists, it updates the name, price, image_url,
// category_id, and sets updated_at to current timestamp.
// It returns the list of products with their corresponding database-generated IDs.
func (db *ProductDB) InsertOrUpdateProducts(products []model.Product) ([]model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), database.DefaultDBOperationTimeout)
	defer cancel()

	batch := &pgx.Batch{}

	sql := `
		INSERT INTO products (name, external_id, price, image_url, category_id)
		VALUES (@name, @external_id, @price, @image_url, @category_id)
		ON CONFLICT (external_id) DO UPDATE
		SET name = EXCLUDED.name,
		    price = EXCLUDED.price,
		    image_url = EXCLUDED.image_url,
		    category_id = EXCLUDED.category_id,
		    updated_at = NOW()
		RETURNING id
	`

	for _, p := range products {
		args := pgx.NamedArgs{
			"name":        p.Name,
			"external_id": p.ExternalID,
			"price":       p.Price,
			"image_url":   p.ImageURL,
			"category_id": p.CategoryID,
		}
		batch.Queue(sql, args)
	}

	br := db.DB.SendBatch(ctx, batch)
	defer br.Close()

	for i := range products {
		var id string
		err := br.QueryRow().Scan(&id)
		if err != nil {
			return nil, err
		}
		products[i].ID = id
	}

	return products, nil
}

func (db *ProductDB) GetProductById(id string) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), database.DefaultDBOperationTimeout)
	defer cancel()

	query := `
		SELECT 
			p.id, 
			p.name, 
			p.price, 
			p.image_url, 
			c.name AS category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1
	`

	var product model.Product

	err := db.DB.QueryRow(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.ImageURL,
		&product.CategoryName,
	)
	if err == pgx.ErrNoRows {
		return nil, ErrNoProductFound
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func NewProductDB(db database.DB) IProductDB {
	return &ProductDB{
		DB: db,
	}
}
