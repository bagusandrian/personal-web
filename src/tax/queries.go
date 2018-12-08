package tax

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Queries struct {
	CheckIDTax                          *sqlx.Stmt
	GetProductByProductID               *sqlx.Stmt
	InsertUser                          *sqlx.Stmt
	GetlistUsers                        *sqlx.Stmt
	InsertProduct                       *sqlx.Stmt
	GetListProducts                     *sqlx.Stmt
	InsertTransaction                   *sqlx.Stmt
	InsertTransactionItem               *sqlx.Stmt
	GetTransactionByIDTransaction       *sqlx.Stmt
	GetDetailTransactionByIDTransaction *sqlx.Stmt
	GetDataTransaction                  *sqlx.Stmt
}

func prepare(query string, db *sqlx.DB) *sqlx.Stmt {
	s, err := db.Preparex(query)
	if err != nil {
		log.Println("failed to prepare query", query, err)
	}
	return s
}

func NewQueries(coreMasterDB, coreSlaveDB *sqlx.DB) *Queries {
	q := &Queries{
		CheckIDTax:                          prepare(checkIDTax, coreSlaveDB),
		GetProductByProductID:               prepare(getProductByProductID, coreSlaveDB),
		GetlistUsers:                        prepare(getlistUsers, coreSlaveDB),
		InsertUser:                          prepare(insertUser, coreMasterDB),
		InsertProduct:                       prepare(insertProduct, coreMasterDB),
		GetListProducts:                     prepare(getListProducts, coreMasterDB),
		InsertTransaction:                   prepare(insertTransaction, coreMasterDB),
		InsertTransactionItem:               prepare(insertTransactionItem, coreMasterDB),
		GetTransactionByIDTransaction:       prepare(getTransactionByIDTransaction, coreSlaveDB),
		GetDetailTransactionByIDTransaction: prepare(getDetailTransactionByIDTransaction, coreSlaveDB),
		GetDataTransaction:                  prepare(getDataTransaction, coreSlaveDB),
	}
	return q
}

const (
	checkIDTax = `
		SELECT id, type FROM tax_code ORDER BY id ASC`
	getProductByProductID = `
		SELECT id, name, price,tax_code_id FROM products WHERE id = $1`
	insertUser = `
		INSERT INTO users (name) VALUES ($1) RETURNING id, name`
	insertProduct = `
		INSERT INTO products (name, price, tax_code_id) VALUES ($1, $2, $3) RETURNING id, name, price, tax_code_id`
	getlistUsers = `
		SELECT id, name
		FROM users
		ORDER BY id ASC`
	getListProducts = `
		SELECT p.id, p.name, p.price,  p.tax_code_id, t.type
		FROM products p
		JOIN tax_code t on t.id = p.tax_code_id
		ORDER BY p.id ASC`
	insertTransaction = `
		INSERT INTO transaction (user_id) VALUES ($1) RETURNING id, user_id, transaction_date`
	insertTransactionItem = `
		INSERT INTO transaction_item (transaction_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING id, transaction_id, product_id, quantity`
	getTransactionByIDTransaction = `
		SELECT id, user_id, transaction_date FROM transaction WHERE id = $1`
	getDetailTransactionByIDTransaction = `
		SELECT trx.id, p.id, p.name, p.price, p.tax_code_id, tx.type, trx.quantity
		FROM products p
		JOIN transaction_item trx ON trx.product_id = p.id
		JOIN tax_code tx ON tx.id = p.tax_code_id
		WHERE trx.transaction_id = $1`
	getDataTransaction = `
		SELECT id, user_id, transaction_date FROM transaction`
)
