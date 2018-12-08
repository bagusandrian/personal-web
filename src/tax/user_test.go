package tax

import (
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/bagusandrian/mini-api/src/db"
	"github.com/jmoiron/sqlx"
)

func TestRegisterUser(t *testing.T) {
	TestRegisterUserSuccess(t)
	TestQueryErrorRegisterUser(t)
	TestScanErrorRegisterUser(t)
}

func TestGetListUser(t *testing.T) {
	TestGetListUserSuccess(t)
	TestQueryErrorGetListUser(t)
	TestScanErrorGetListUser(t)
}
func TestRegisterUserSuccess(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	INSERT INTO users (.+) VALUES (.+) RETURNING id, name
	`
	rowUser := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "user test")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowUser)

	n := &Module{
		queries: &Queries{
			InsertUser: prepare(rgxQuery, sx),
		},
	}
	data, err := n.RegisterUser("user test")
	if err != nil {
		t.Errorf("[TestRegisterUserSuccess] function returns error: %+v\n", err)
	}
	if data.Name != "user test" {
		t.Errorf("[TestRegisterUserSuccess] error validation in Name. Name: %+v", data.Name)
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestRegisterUserSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestQueryErrorRegisterUser(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	INSERT INTO users (.+) VALUES (.+) RETURNING id, name
	`
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnError(errors.New("error"))

	n := &Module{
		queries: &Queries{
			InsertUser: prepare(rgxQuery, sx),
		},
	}
	_, err = n.RegisterUser("user test")
	if err == nil {
		t.Errorf("[TestQueryErrorRegisterUser] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestQueryErrorRegisterUser] there were unfulfilled expections: %+v\n", err)
	}
}
func TestScanErrorRegisterUser(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreMaster", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	INSERT INTO users (.+) VALUES (.+) RETURNING id, name
	`
	rowUser := sqlmock.NewRows([]string{"id", "name"}).AddRow("troble maker", "user test")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowUser)

	n := &Module{
		queries: &Queries{
			InsertUser: prepare(rgxQuery, sx),
		},
	}
	_, err = n.RegisterUser("user test")
	if err == nil {
		t.Errorf("[TestScanErrorRegisterUser] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestScanErrorRegisterUser] there were unfulfilled expections: %+v\n", err)
	}
}
func TestGetListUserSuccess(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, name
	FROM users
	ORDER BY id ASC
	`
	rowUser := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "user test")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowUser)

	n := &Module{
		queries: &Queries{
			GetlistUsers: prepare(rgxQuery, sx),
		},
	}
	data, err := n.GetListUser()
	if err != nil {
		t.Errorf("[TestRegisterUserSuccess] function returns error: %+v\n", err)
	}
	if data[0].Name != "user test" {
		t.Errorf("[TestRegisterUserSuccess] error validation in Name. Name: %+v", data[0].Name)
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestRegisterUserSuccess] there were unfulfilled expections: %+v\n", err)
	}
}
func TestQueryErrorGetListUser(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, name
	FROM users
	ORDER BY id ASC
	`
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnError(errors.New("error"))

	n := &Module{
		queries: &Queries{
			GetlistUsers: prepare(rgxQuery, sx),
		},
	}
	_, err = n.GetListUser()
	if err == nil {
		t.Errorf("[TestQueryErrorRegisterUser] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestQueryErrorRegisterUser] there were unfulfilled expections: %+v\n", err)
	}
}
func TestScanErrorGetListUser(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	sx := sqlx.NewDb(mockdb, "mockdb")
	db.PutMock("CoreSlave", sx)
	if err != nil {
		t.Error("failed mock db")
		return
	}

	rgxQuery := `
	SELECT id, name
	FROM users
	ORDER BY id ASC
	`
	rowUser := sqlmock.NewRows([]string{"id", "name"}).AddRow("troble maker", "user test")
	m.ExpectPrepare(rgxQuery)
	m.ExpectQuery(rgxQuery).WillReturnRows(rowUser)

	n := &Module{
		queries: &Queries{
			GetlistUsers: prepare(rgxQuery, sx),
		},
	}
	_, err = n.GetListUser()
	if err == nil {
		t.Errorf("[TestScanErrorRegisterUser] expected error but error is nil\n")
	}
	if err = m.ExpectationsWereMet(); err != nil {
		t.Errorf("[TestScanErrorRegisterUser] there were unfulfilled expections: %+v\n", err)
	}
}
