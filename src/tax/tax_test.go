package tax

import (
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/bagusandrian/mini-api/src/db"
	"github.com/jmoiron/sqlx"
)

func TestListTaxCode(t *testing.T) {
	TestListTaxCodeSuccess(t)
	TestQueryErrorListTaxCode(t)
	TestScanErrorListTaxCode(t)
}
func TestListGetProductDetail(t *testing.T) {
	TestListGetProductDetailSuccess(t)
	TestQueryErrorListGetProductDetail(t)
	TestScanErrorListGetProductDetail(t)
}
func TestListTaxCodeSuccess(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, type FROM tax_code ORDER BY id ASC
	`
	rowTax := sqlmock.NewRows([]string{"id", "type"}).AddRow(1, "food")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowTax)

	n := &Module{
		queries: &Queries{
			CheckIDTax: prepare(rgxQuery, sx),
		},
	}
	data, err := n.ListTaxCode()
	if err != nil {
		t.Errorf("[TestListTaxCodeSuccess] function returns error: %+v\n", err)
	}
	if data.List[0].Type != "food" {
		t.Errorf("[TestListTaxCodeSuccess] error validation in Name. Name: %+v", data.List[0].Type)
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestListTaxCodeSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestQueryErrorListTaxCode(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, type FROM tax_code ORDER BY id ASC
	`
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnError(errors.New("error"))

	n := &Module{
		queries: &Queries{
			CheckIDTax: prepare(rgxQuery, sx),
		},
	}
	_, err = n.ListTaxCode()
	if err == nil {
		t.Errorf("[TestQueryErrorListTaxCode] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestQueryErrorListTaxCode] there were unfulfilled expections: %+v\n", err)
	}
}
func TestScanErrorListTaxCode(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, type FROM tax_code ORDER BY id ASC
	`
	rowTax := sqlmock.NewRows([]string{"id", "type"}).AddRow("troble maker", "food")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowTax)

	n := &Module{
		queries: &Queries{
			CheckIDTax: prepare(rgxQuery, sx),
		},
	}
	_, err = n.ListTaxCode()
	if err == nil {
		t.Errorf("[TestScanErrorListTaxCode] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestScanErrorListTaxCode] there were unfulfilled expections: %+v\n", err)
	}
}
func TestListGetProductDetailSuccess(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, name, price,tax_code_id FROM products WHERE id = (.+)
	`
	rowPRoduct := sqlmock.NewRows([]string{"id", "name", "price", "tax_code_id"}).AddRow(1, "product test", 10000.00, 1)
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowPRoduct)

	n := &Module{
		queries: &Queries{
			GetProductByProductID: prepare(rgxQuery, sx),
		},
	}
	var data ProductData
	data.ID = 1
	_, err = n.GetProductDetail(data)
	if err != nil {
		t.Errorf("[TestListTaxCodeSuccess] function returns error: %+v\n", err)
	}
	if data.ID != 1 {
		t.Errorf("[TestListTaxCodeSuccess] error validation in ID. ID: %+v", data.ID)
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestListTaxCodeSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestQueryErrorListGetProductDetail(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, name, price,tax_code_id FROM products WHERE id = (.+)
	`
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnError(errors.New("error"))

	n := &Module{
		queries: &Queries{
			GetProductByProductID: prepare(rgxQuery, sx),
		},
	}
	var data ProductData
	data.ID = 1
	_, err = n.GetProductDetail(data)
	if err == nil {
		t.Errorf("[TestQueryErrorListTaxCode] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestQueryErrorListTaxCode] there were unfulfilled expections: %+v\n", err)
	}
}
func TestScanErrorListGetProductDetail(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, name, price,tax_code_id FROM products WHERE id = (.+)
	`
	rowPRoduct := sqlmock.NewRows([]string{"id", "name", "price", "tax_code_id"}).AddRow(1, "product test", 10000.00, "troble maker")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowPRoduct)

	n := &Module{
		queries: &Queries{
			GetProductByProductID: prepare(rgxQuery, sx),
		},
	}
	var data ProductData
	data.ID = 1
	_, err = n.GetProductDetail(data)
	if err == nil {
		t.Errorf("[TestScanErrorListTaxCode] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestScanErrorListTaxCode] there were unfulfilled expections: %+v\n", err)
	}
}
