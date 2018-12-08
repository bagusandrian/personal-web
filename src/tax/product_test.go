package tax

import (
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/bagusandrian/mini-api/src/db"
	"github.com/jmoiron/sqlx"
)

func TestRegisterProduct(t *testing.T) {
	TestRegisterProductSuccess(t)
	TestQueryErrorRegisterProduct(t)
	TestScanErrorRegisterProduct(t)
}
func TestListProducts(t *testing.T) {
	TestListProductsSuccess(t)
	TestQueryErrorListProducts(t)
	TestScanErrorListProducts(t)
}
func TestRegisterProductSuccess(t *testing.T) {
	var data ProductData
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	INSERT INTO products (.+) VALUES (.+) RETURNING id, name, price, tax_code_id
	`
	rowProduct := sqlmock.NewRows([]string{"id", "name", "price", "tax_code_id"}).AddRow(1, "test mock", 10000.00, TaxFoodID)
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowProduct)

	n := &Module{
		queries: &Queries{
			InsertProduct: prepare(rgxQuery, sx),
		},
	}
	data.Name = "test mock"
	data.Price = 10000.00
	data.TaxCodeID = 1
	data, err = n.RegisterProduct(data)
	if err != nil {
		t.Errorf("[TestRegisterProductSuccess] function returns error: %+v\n", err)
	}
	if data.TaxCodeID != TaxFoodID {
		t.Errorf("[TestRegisterProductSuccess] error validation in TaxCodeID. TaxCodeID: %+v", data.TaxCodeID)
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestRegisterProductSuccess] there were unfulfilled expections: %+v\n", err)
	}

}
func TestQueryErrorRegisterProduct(t *testing.T) {
	var data ProductData
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	INSERT INTO products (.+) VALUES (.+) RETURNING id, name, price, tax_code_id
	`
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnError(errors.New("error"))

	n := &Module{
		queries: &Queries{
			InsertProduct: prepare(rgxQuery, sx),
		},
	}
	data.Name = "test mock"
	data.Price = 10000.00
	data.TaxCodeID = 1
	data, err = n.RegisterProduct(data)
	if err == nil {
		t.Errorf("[TestQueryErrorRegisterProduct] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestQueryErrorRegisterProduct] there were unfulfilled expections: %+v\n", err)
	}

}
func TestScanErrorRegisterProduct(t *testing.T) {
	var data ProductData
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	INSERT INTO products (.+) VALUES (.+) RETURNING id, name, price, tax_code_id
	`
	rowProduct := sqlmock.NewRows([]string{"id", "name", "price", "tax_code_id"}).AddRow(1, "test mock", "troble maker", TaxFoodID)
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowProduct)

	n := &Module{
		queries: &Queries{
			InsertProduct: prepare(rgxQuery, sx),
		},
	}
	data.Name = "test mock"
	data.Price = 10000.00
	data.TaxCodeID = 1
	data, err = n.RegisterProduct(data)
	if err == nil {
		t.Errorf("[TestNoDBRegisterProduct] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestScanErrorRegisterProduct] there were unfulfilled expections: %+v\n", err)
	}

}
func TestListProductsSuccess(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT p.id, p.name, p.price,  p.tax_code_id, t.type
	FROM products p
	JOIN tax_code t on t.id = p.tax_code_id
	ORDER BY p.id ASC
	`
	rowProduct := sqlmock.NewRows([]string{"id", "name", "price", "tax_code_id", "food"}).AddRow(1, "test mock", 10000.00, TaxFoodID, "food")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowProduct)

	n := &Module{
		queries: &Queries{
			GetListProducts: prepare(rgxQuery, sx),
		},
	}

	_, err = n.ListProducts()
	if err != nil {
		t.Errorf("[TestListProductsSuccess] function returns error: %+v\n", err)
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestListProductsSuccess] there were unfulfilled expections: %+v\n", err)
	}

}
func TestQueryErrorListProducts(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT p.id, p.name, p.price,  p.tax_code_id, t.type
	FROM products p
	JOIN tax_code t on t.id = p.tax_code_id
	ORDER BY p.id ASC
	`
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnError(errors.New("error"))

	n := &Module{
		queries: &Queries{
			GetListProducts: prepare(rgxQuery, sx),
		},
	}

	_, err = n.ListProducts()
	if err == nil {
		t.Errorf("[TestQueryErrorListProducts] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestQueryErrorListProducts] there were unfulfilled expections: %+v\n", err)
	}

}
func TestScanErrorListProducts(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT p.id, p.name, p.price,  p.tax_code_id, t.type
	FROM products p
	JOIN tax_code t on t.id = p.tax_code_id
	ORDER BY p.id ASC
	`
	rowProduct := sqlmock.NewRows([]string{"id", "name", "price", "tax_code_id", "food"}).AddRow("troble maker", "test mock", 10000.00, TaxFoodID, "food")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowProduct)

	n := &Module{
		queries: &Queries{
			GetListProducts: prepare(rgxQuery, sx),
		},
	}

	_, err = n.ListProducts()
	if err == nil {
		t.Errorf("[TestScanErrorListProducts] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestScanErrorListProducts] there were unfulfilled expections: %+v\n", err)
	}

}
