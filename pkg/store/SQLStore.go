package store

import ("database/sql"
		"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)


type SQLStore struct {
	db *sql.DB
}

func NewMySQLStore(db *sql.DB) StoreInterface {
	return &SQLStore{
		db: db,
	}
}

func (s *SQLStore) Read(id int) (domain.Product, error) {
	var product domain.Product
	query := "SELECT * FROM products WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *SQLStore) Create(product domain.Product) error {
	query := "INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?); "
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	
	res, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
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


