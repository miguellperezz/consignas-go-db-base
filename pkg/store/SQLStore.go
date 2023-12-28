package store

import ("database/sql"
		"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)


type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) StoreInterface {
	return &SQLStore{
		db: db,
	}
}

func (s *SQLStore) Read(id int) (domain.Product, error) {
	product := domain.Product{}
	err := s.db.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(&product.Id, &product.CodeValue, &product.Name, &product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *SQLStore) Create(product domain.Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?) ", product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price)
	
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLStore) Update(product domain.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?", product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price, product.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLStore) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLStore) Exists(codeValue string) bool {
	var id int
	err := s.db.QueryRow("SELECT id FROM products WHERE code_value = ?", codeValue).Scan(&id)
	if err != nil {
		return false
	}
	return true
}


