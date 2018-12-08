package tax

import (
	"errors"
	"log"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/bagusandrian/mini-api/src/db"
	"github.com/jmoiron/sqlx"
)

func TestCreateTransaction(t *testing.T) {
	TestCreateTransactionSuccess(t)
	TestQueryErrorCreateTransaction(t)
	TestScanErrorCreateTransaction(t)
}
func TestGetListTransaction(t *testing.T) {
	TestGetListTransactionSuccess(t)
	TestQueryErrorGetListTransaction(t)
	TestScanErrorGetListTransaction(t)
}
func TestGetTransactionByID(t *testing.T) {
	TestGetTransactionByIDSuccess(t)
	TestQueryErrorGetTransactionByID(t)
	TestScanErrorGetTransactionByID(t)
}
func TestGetDetailTransactionByIDTransaction(t *testing.T) {
	TestGetDetailTransactionByIDTransactionSuccess(t)
	TestQueryErrorGetDetailTransactionByIDTransaction(t)
	TestScanErrorGetDetailTransactionByIDTransaction(t)
}
func TestCreateTransactionItem(t *testing.T) {
	TestCreateTransactionItemSuccess(t)
	TestScanErrorCreateTransactionItem(t)
}
func TestCalculationTotalPrice(t *testing.T) {
	TestCalculationTotalPriceSuccess(t)
	TestScanErrorCalculationTotalPrice(t)
	TestCalculationTotalPriceSuccessTax2(t)
	TestCalculationTotalPriceSuccessTax3(t)
	TestCalculationTotalPriceSuccessTax3qty101(t)
	TestCalculationTotalPriceSuccessTax0(t)
}
func TestGetListTransactionSuccess(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, user_id, transaction_date FROM transaction
	`
	rowTransaction := sqlmock.NewRows([]string{"id", "user_id", "transaction_date"}).AddRow(1, 1, "2018-09-10")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowTransaction)

	n := &Module{
		queries: &Queries{
			GetDataTransaction: prepare(rgxQuery, sx),
		},
	}
	data, err := n.GetListTransaction()
	if err != nil {
		t.Errorf("[TestGetListTransactionSuccess] function returns error: %+v\n", err)
	}
	if data[0].ID != 1 {
		t.Errorf("[TestGetListTransactionSuccess] error validation in ID. ID: %+v", data[0].ID)
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestGetListTransactionSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestQueryErrorGetListTransaction(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, user_id, transaction_date FROM transaction
	`
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnError(errors.New("error"))

	n := &Module{
		queries: &Queries{
			GetDataTransaction: prepare(rgxQuery, sx),
		},
	}
	_, err = n.GetListTransaction()
	if err == nil {
		t.Errorf("[TestQueryErrorGetListTransaction] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestQueryErrorGetListTransaction] there were unfulfilled expections: %+v\n", err)
	}
}
func TestScanErrorGetListTransaction(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, user_id, transaction_date FROM transaction
	`
	rowTransaction := sqlmock.NewRows([]string{"id", "user_id", "transaction_date"}).AddRow(1, "troble maker", "2018-09-10")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowTransaction)

	n := &Module{
		queries: &Queries{
			GetDataTransaction: prepare(rgxQuery, sx),
		},
	}
	_, err = n.GetListTransaction()
	if err == nil {
		t.Errorf("[TestScanErrorGetListTransaction] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestScanErrorGetListTransaction] there were unfulfilled expections: %+v\n", err)
	}
}
func TestCreateTransactionSuccess(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	INSERT INTO transaction (.+) RETURNING id, user_id, transaction_date
	`
	rowTransaction := sqlmock.NewRows([]string{"id", "user_id", "transaction_date"}).AddRow(1, 1, "2019-08-08")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowTransaction)

	n := &Module{
		queries: &Queries{
			InsertTransaction: prepare(rgxQuery, sx),
		},
	}
	var data DataTransaction
	data.UserID = 1
	data, err = n.CreateTransaction(data)
	if err != nil {
		t.Errorf("[TestCreateTransactionSuccess] function returns error: %+v\n", err)
	}
	log.Printf("%+v\n", data)
	if data.ID != 1 {
		t.Errorf("[TestCreateTransactionSuccess] error validation in ID. ID: %+v", data.ID)
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestCreateTransactionSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestQueryErrorCreateTransaction(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	INSERT INTO transaction (.+) RETURNING id, user_id, transaction_date
	`
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnError(errors.New("error"))

	n := &Module{
		queries: &Queries{
			InsertTransaction: prepare(rgxQuery, sx),
		},
	}
	var data DataTransaction
	data.UserID = 1
	_, err = n.CreateTransaction(data)
	if err == nil {
		t.Errorf("[TestQueryErrorCreateTransaction] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestQueryErrorCreateTransaction] there were unfulfilled expections: %+v\n", err)
	}
}
func TestScanErrorCreateTransaction(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	INSERT INTO transaction (.+) RETURNING id, user_id, transaction_date
	`
	rowTransaction := sqlmock.NewRows([]string{"id", "user_id", "transaction_date"}).AddRow(1, "Troble maker", "2019-08-08")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowTransaction)

	n := &Module{
		queries: &Queries{
			InsertTransaction: prepare(rgxQuery, sx),
		},
	}
	var data DataTransaction
	data.UserID = 1
	_, err = n.CreateTransaction(data)
	if err == nil {
		t.Errorf("[TestScanErrorCreateTransaction] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestScanErrorCreateTransaction] there were unfulfilled expections: %+v\n", err)
	}
}
func TestGetTransactionByIDSuccess(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, user_id, transaction_date FROM transaction WHERE id = (.+)
	`
	rowTransaction := sqlmock.NewRows([]string{"id", "user_id", "transaction_date"}).AddRow(1, 1, "2019-08-08")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowTransaction)

	n := &Module{
		queries: &Queries{
			GetTransactionByIDTransaction: prepare(rgxQuery, sx),
		},
	}
	var data DataTransaction
	data.UserID = 1
	data, err = n.GetTransactionByID(data)
	if err != nil {
		t.Errorf("[TestGetTransactionByIDSuccess] function returns error: %+v\n", err)
	}
	log.Printf("%+v\n", data)
	if data.ID != 1 {
		t.Errorf("[TestGetTransactionByIDSuccess] error validation in ID. ID: %+v", data.ID)
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestGetTransactionByIDSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestQueryErrorGetTransactionByID(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, user_id, transaction_date FROM transaction WHERE id = (.+)
	`
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnError(errors.New("error"))

	n := &Module{
		queries: &Queries{
			GetTransactionByIDTransaction: prepare(rgxQuery, sx),
		},
	}
	var data DataTransaction
	data.UserID = 1
	_, err = n.GetTransactionByID(data)
	if err == nil {
		t.Errorf("[TestQueryErrorGetTransactionByID] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestQueryErrorGetTransactionByID] there were unfulfilled expections: %+v\n", err)
	}
}
func TestScanErrorGetTransactionByID(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, user_id, transaction_date FROM transaction WHERE id = (.+)
	`
	rowTransaction := sqlmock.NewRows([]string{"id", "user_id", "transaction_date"}).AddRow(1, "Troble maker", "2019-08-08")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowTransaction)

	n := &Module{
		queries: &Queries{
			GetTransactionByIDTransaction: prepare(rgxQuery, sx),
		},
	}
	var data DataTransaction
	data.UserID = 1
	_, err = n.GetTransactionByID(data)
	if err == nil {
		t.Errorf("[TestScanErrorGetTransactionByID] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestScanErrorGetTransactionByID] there were unfulfilled expections: %+v\n", err)
	}
}
func TestGetDetailTransactionByIDTransactionSuccess(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT trx.id, p.id, p.name, p.price, p.tax_code_id, tx.type, trx.quantity
    FROM products p
    JOIN transaction_item trx ON trx.product_id = p.id
    JOIN tax_code tx ON tx.id = p.tax_code_id
    WHERE trx.transaction_id = (.+)
	`
	rowTransactionItem := sqlmock.NewRows([]string{"id", "product_id", "name", "price", "tax_code_id", "type", "quantity"}).
		AddRow(1, 1, "test product", 10000.00, 1, "food", 10)
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowTransactionItem)

	n := &Module{
		queries: &Queries{
			GetDetailTransactionByIDTransaction: prepare(rgxQuery, sx),
		},
	}
	var data DataTransaction
	data.ID = 1
	result, err := n.GetDetailTransactionByIDTransaction(data)
	if err != nil {
		t.Errorf("[TestGetDetailTransactionByIDTransactionSuccess] function returns error: %+v\n", err)
	}
	log.Printf("%+v\n", data)
	if result.DetailTransaction[0].ID != 1 {
		t.Errorf("[TestGetDetailTransactionByIDTransactionSuccess] error validation in ID. ID: %+v", result.DetailTransaction[0].ID)
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestGetDetailTransactionByIDTransactionSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestQueryErrorGetDetailTransactionByIDTransaction(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT trx.id, p.id, p.name, p.price, p.tax_code_id, tx.type, trx.quantity
    FROM products p
    JOIN transaction_item trx ON trx.product_id = p.id
    JOIN tax_code tx ON tx.id = p.tax_code_id
    WHERE trx.transaction_id = (.+)
	`
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnError(errors.New("error"))

	n := &Module{
		queries: &Queries{
			GetDetailTransactionByIDTransaction: prepare(rgxQuery, sx),
		},
	}
	var data DataTransaction
	data.ID = 1
	_, err = n.GetDetailTransactionByIDTransaction(data)
	if err == nil {
		t.Errorf("[TestQueryErrorGetDetailTransactionByIDTransaction] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestQueryErrorGetDetailTransactionByIDTransaction] there were unfulfilled expections: %+v\n", err)
	}
}
func TestScanErrorGetDetailTransactionByIDTransaction(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT trx.id, p.id, p.name, p.price, p.tax_code_id, tx.type, trx.quantity
    FROM products p
    JOIN transaction_item trx ON trx.product_id = p.id
    JOIN tax_code tx ON tx.id = p.tax_code_id
    WHERE trx.transaction_id = (.+)
	`
	rowTransactionItem := sqlmock.NewRows([]string{"id", "product_id", "name", "price", "tax_code_id", "type", "quantity"}).
		AddRow(1, 1, "test product", 10000.00, 1, "food", "troble Maker")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowTransactionItem)

	n := &Module{
		queries: &Queries{
			GetDetailTransactionByIDTransaction: prepare(rgxQuery, sx),
		},
	}
	var data DataTransaction
	data.ID = 1
	_, err = n.GetDetailTransactionByIDTransaction(data)
	if err == nil {
		t.Errorf("[TestScanErrorGetDetailTransactionByIDTransaction] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestScanErrorGetDetailTransactionByIDTransaction] there were unfulfilled expections: %+v\n", err)
	}
}
func TestCreateTransactionItemSuccess(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	INSERT INTO transaction_item (.+) VALUES (.+) RETURNING id, transaction_id, product_id, quantity
	`
	rowTransactionItem := sqlmock.NewRows([]string{"id", "transaction_id", "product_id", "quantity"}).
		AddRow(1, 1, 1, 10)
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowTransactionItem)

	n := &Module{
		queries: &Queries{
			InsertTransactionItem: prepare(rgxQuery, sx),
		},
	}
	var data DataTransaction
	var item DataTransactionItem

	item.ID = 1
	item.Product.ID = 1
	item.Product.Quantity = 1
	data.DetailTransaction = append(data.DetailTransaction, item)
	result, err := n.CreateTransactionItem(data)
	if err != nil {
		t.Errorf("[TestCreateTransactionItemSuccess] function returns error: %+v\n", err)
	}
	if result.DetailTransaction[0].ID != 1 {
		t.Errorf("[TestCreateTransactionItemSuccess] error validation in ID. ID: %+v", result.DetailTransaction[0].ID)
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestCreateTransactionItemSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestScanErrorCreateTransactionItem(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}
	rgxQuery := `
	INSERT INTO transaction_item (.+) VALUES (.+) RETURNING id, transaction_id, product_id, quantity
	`
	rowTransactionItem := sqlmock.NewRows([]string{"id", "transaction_id", "product_id", "quantity"}).
		AddRow(1, 1, 1, "troble maker")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowTransactionItem)

	n := &Module{
		queries: &Queries{
			InsertTransactionItem: prepare(rgxQuery, sx),
		},
	}
	var data DataTransaction
	var item DataTransactionItem

	item.ID = 1
	item.Product.ID = 1
	item.Product.Quantity = 1
	data.DetailTransaction = append(data.DetailTransaction, item)
	_, err = n.CreateTransactionItem(data)
	if err == nil {
		t.Errorf("[TestScanErrorCreateTransactionItem] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestScanErrorCreateTransactionItem] there were unfulfilled expections: %+v\n", err)
	}
}
func TestCalculationTotalPriceSuccess(t *testing.T) {
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
	rowPRoduct := sqlmock.NewRows([]string{"id", "name", "price", "tax_code_id"}).
		AddRow(1, "product test", 10000.00, 1)
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowPRoduct)

	n := &Module{
		queries: &Queries{
			GetProductByProductID: prepare(rgxQuery, sx),
		},
	}
	var transaction DataTransaction
	var item DataTransactionItem
	item.Product.ID = 1
	transaction.DetailTransaction = append(transaction.DetailTransaction, item)
	result, err := n.CalculationTotalPrice(transaction)
	if err != nil {
		t.Errorf("[TestCalculationTotalPriceSuccess] function returns error: %+v\n", err)
	}
	log.Printf("%+v\n", result)
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestCreateTransactionItemSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestScanErrorCalculationTotalPrice(t *testing.T) {
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
	var transaction DataTransaction
	var item DataTransactionItem
	item.Product.ID = 1
	transaction.DetailTransaction = append(transaction.DetailTransaction, item)
	_, err = n.CalculationTotalPrice(transaction)
	if err == nil {
		t.Errorf("[TestCalculationTotalPriceSuccess] function returns error: %+v\n", err)
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestCreateTransactionItemSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestCalculationTotalPriceSuccessTax2(t *testing.T) {
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
	rowPRoduct := sqlmock.NewRows([]string{"id", "name", "price", "tax_code_id"}).
		AddRow(1, "product test", 10000.00, 2)
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowPRoduct)

	n := &Module{
		queries: &Queries{
			GetProductByProductID: prepare(rgxQuery, sx),
		},
	}
	var transaction DataTransaction
	var item DataTransactionItem
	item.Product.ID = 1
	transaction.DetailTransaction = append(transaction.DetailTransaction, item)
	result, err := n.CalculationTotalPrice(transaction)
	if err != nil {
		t.Errorf("[TestCalculationTotalPriceSuccess] function returns error: %+v\n", err)
	}
	log.Printf("%+v\n", result)
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestCreateTransactionItemSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestCalculationTotalPriceSuccessTax3(t *testing.T) {
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
	rowPRoduct := sqlmock.NewRows([]string{"id", "name", "price", "tax_code_id"}).
		AddRow(1, "product test", 10000.00, 3)
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowPRoduct)

	n := &Module{
		queries: &Queries{
			GetProductByProductID: prepare(rgxQuery, sx),
		},
	}
	var transaction DataTransaction
	var item DataTransactionItem
	item.Product.ID = 1
	transaction.DetailTransaction = append(transaction.DetailTransaction, item)
	result, err := n.CalculationTotalPrice(transaction)
	if err != nil {
		t.Errorf("[TestCalculationTotalPriceSuccess] function returns error: %+v\n", err)
	}
	log.Printf("%+v\n", result)
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestCreateTransactionItemSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestCalculationTotalPriceSuccessTax3qty101(t *testing.T) {
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
	rowPRoduct := sqlmock.NewRows([]string{"id", "name", "price", "tax_code_id"}).
		AddRow(1, "product test", 10000.00, 3)
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowPRoduct)

	n := &Module{
		queries: &Queries{
			GetProductByProductID: prepare(rgxQuery, sx),
		},
	}
	var transaction DataTransaction
	var item DataTransactionItem
	item.Product.ID = 1
	item.Product.Quantity = 101
	transaction.DetailTransaction = append(transaction.DetailTransaction, item)
	result, err := n.CalculationTotalPrice(transaction)
	if err != nil {
		t.Errorf("[TestCalculationTotalPriceSuccess] function returns error: %+v\n", err)
	}
	log.Printf("%+v\n", result)
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestCreateTransactionItemSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestCalculationTotalPriceSuccessTax0(t *testing.T) {
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
	rowPRoduct := sqlmock.NewRows([]string{"id", "name", "price", "tax_code_id"}).
		AddRow(1, "product test", 10000.00, 0)
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowPRoduct)

	n := &Module{
		queries: &Queries{
			GetProductByProductID: prepare(rgxQuery, sx),
		},
	}
	var transaction DataTransaction
	var item DataTransactionItem
	item.Product.ID = 1
	transaction.DetailTransaction = append(transaction.DetailTransaction, item)
	result, err := n.CalculationTotalPrice(transaction)
	if err != nil {
		t.Errorf("[TestCalculationTotalPriceSuccess] function returns error: %+v\n", err)
	}
	log.Printf("%+v\n", result)
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestCreateTransactionItemSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
