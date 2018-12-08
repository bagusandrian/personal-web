package tax

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/bagusandrian/mini-api/src/common"
	"github.com/julienschmidt/httprouter"
)

func (m *Module) ListTaxHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.JSONResponse) {
	resp = &common.JSONResponse{}
	data, err := m.ListTaxCode()
	if err != nil {
		log.Printf("[tax][ListTaxHandler] err get data ListTaxCode, err:%+v\n", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.ErrorString = "Failed get list tax code"
		return
	}
	resp.StatusCode = http.StatusOK
	resp.Data = data
	resp.Message = "Success"
	return
}
func (m *Module) RegisterUserHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.JSONResponse) {
	resp = &common.JSONResponse{}
	name := rgxSQLInjectionChar.ReplaceAllString(r.PostFormValue("name"), "")
	if name == "" {
		log.Printf("[tax][RegisterUserHandler] name is empty")
		resp.ErrorString = "Name is empty"
		resp.StatusCode = http.StatusBadRequest
		return
	}
	data, err := m.RegisterUser(name)
	if err != nil {
		log.Printf("[tax][RegisterUserHandler] name: %s error RegisterUser: %+v\n", name, err)
		resp.ErrorString = err.Error()
		resp.StatusCode = http.StatusInternalServerError
		return
	}
	resp.Message = "Success"
	resp.Data = data
	resp.StatusCode = http.StatusOK
	return
}

func (m *Module) GetListUserHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.JSONResponse) {
	resp = &common.JSONResponse{}
	var data []UserData
	var err error
	data, err = m.GetListUser()
	if err != nil {
		log.Printf("[tax][ListUserHandler] err get data ListUser, err:%+v\n", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.ErrorString = "Failed get list users"
		return
	}
	resp.Message = "Success"
	resp.Data = data
	resp.StatusCode = http.StatusOK
	return
}

func (m *Module) RegisterProductHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.JSONResponse) {
	data := ProductData{}
	var err error
	resp = &common.JSONResponse{}
	data.Name = rgxSQLInjectionChar.ReplaceAllString(r.PostFormValue("product_name"), "")
	data.Price, err = strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		log.Printf("[tax][RegisterProductHandler] error ParseFloat price: %s err:%+v\n", r.PostFormValue("price"), err)
		resp.ErrorString = "Invalid Price"
		resp.StatusCode = http.StatusBadRequest
		return
	}
	data.TaxCodeID, err = strconv.Atoi(r.PostFormValue("tax_code_id"))
	if err != nil {
		log.Printf("[tax][RegisterProductHandler] error Itoa tax_code_id: %s err:%+v\n", r.PostFormValue("tax_code_id"), err)
		resp.ErrorString = "Invalid tax code ID"
		resp.StatusCode = http.StatusBadRequest
		return
	}
	if data.TaxCodeID != TaxFoodID && data.TaxCodeID != TaxTobaccoID && data.TaxCodeID != TaxEntertainmentID {
		log.Printf("[tax][RegisterProductHandler] undefined tax code id")
		resp.ErrorString = "Undefined tax code ID"
		resp.StatusCode = http.StatusBadRequest
		return
	}
	data, err = m.RegisterProduct(data)
	resp.Message = "Success"
	resp.Data = data
	resp.StatusCode = http.StatusOK
	return
}

func (m *Module) GetListProductHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.JSONResponse) {
	resp = &common.JSONResponse{}
	var data []ProductData
	var err error
	data, err = m.ListProducts()
	if err != nil {
		log.Printf("[tax][GetListProductHandler] err get data ListProducts, err:%+v\n", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.ErrorString = "Failed get list products"
		return
	}
	resp.Message = "Success"
	resp.Data = data
	resp.StatusCode = http.StatusOK
	return
}

func (m *Module) CreateTransactionHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.JSONResponse) {
	resp = &common.JSONResponse{}
	var data DataTransaction
	var itemTransaction DataTransactionItem
	var itemTransactions []DataTransactionItem
	var err error
	data.UserID, err = strconv.ParseInt(r.PostFormValue("user_id"), 10, 64)
	if err != nil || data.UserID == 0 {
		log.Printf("[tax][CreateTransactionHandler] undefined user_id: %s err: %+v\n", r.PostFormValue("user_id"), err)
		resp.StatusCode = http.StatusInternalServerError
		resp.ErrorString = "Undefined user_id"
		return
	}
	transactionItem := r.PostFormValue("transaction_item")
	stringSlice := strings.Split(transactionItem, ",")
	for _, s := range stringSlice {
		splitIDProductQuantity := strings.Split(s, ":")
		for i, j := range splitIDProductQuantity {
			if i < 1 {
				itemTransaction.Product.ID, err = strconv.ParseInt(j, 10, 64)
				if err != nil {
					log.Printf("[tax][CreateTransactionHandler] params: %s error parseInt for ProductID err: %+v\n", transactionItem, err)
					resp.ErrorString = "transaction item is undefined"
					resp.Message = "ERROR"
					resp.StatusCode = http.StatusBadRequest
					return
				}
				continue
			}
			itemTransaction.Product.Quantity, err = strconv.Atoi(j)
			if err != nil {
				log.Printf("[tax][CreateTransactionHandler] params: %s error Itoa for quantity err: %+v\n", transactionItem, err)
				resp.ErrorString = "transaction item is undefined"
				resp.Message = "ERROR"
				resp.StatusCode = http.StatusBadRequest
				return
			}
		}
		itemTransactions = append(itemTransactions, itemTransaction)
	}
	data.DetailTransaction = itemTransactions

	data, err = m.CreateTransaction(data)
	if err != nil {
		log.Printf("[tax][CreateTransactionHandler] params: %+v CreateTransaction Err:%+v\n", data, err)
		resp.ErrorString = "Internal Server error"
		resp.Message = "ERROR"
		resp.StatusCode = http.StatusInternalServerError
		return
	}
	data, err = m.CreateTransactionItem(data)
	if err != nil {
		log.Printf("[tax][CreateTransactionHandler] params: %+v CreateTransactionItem Err:%+v\n", data, err)
		resp.ErrorString = "Internal Server error"
		resp.Message = "ERROR"
		resp.StatusCode = http.StatusInternalServerError
		return
	}
	data, err = m.CalculationTotalPrice(data)
	resp.Message = "Success"
	resp.Data = data
	resp.StatusCode = http.StatusOK
	return
}
func (m *Module) GetListTransactionHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.JSONResponse) {
	resp = &common.JSONResponse{}
	data, err := m.GetListTransaction()
	if err != nil {
		log.Printf("[tax][GetListTransactionHandler] err get data GetListTransaction, err:%+v\n", err)
		resp.StatusCode = http.StatusInternalServerError
		resp.ErrorString = "Failed get list transaction"
		return
	}
	resp.StatusCode = http.StatusOK
	resp.Data = data
	resp.Message = "Success"
	return
}
func (m *Module) GetDetailTransactionHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) (resp *common.JSONResponse) {
	resp = &common.JSONResponse{}
	var data DataTransaction
	var err error
	data.ID, err = strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		log.Printf("[tax][RegisterProductHandler] error ParseFloat transaction_id: %s err:%+v\n", r.PostFormValue("transaction_id"), err)
		resp.ErrorString = "Invalid transaction_id"
		resp.StatusCode = http.StatusBadRequest
		return
	}
	data, err = m.GetTransactionByID(data)
	if err != nil {
		log.Printf("[tax][CreateTransactionHandler] params: %+v GetTransactionBYID Err:%+v\n", data, err)
		resp.ErrorString = "Internal Server error"
		resp.Message = "ERROR"
		resp.StatusCode = http.StatusInternalServerError
		return
	}
	data, err = m.GetDetailTransactionByIDTransaction(data)
	if err != nil {
		log.Printf("[tax][CreateTransactionHandler] params: %+v GetDetailTransactionByIDTransaction Err:%+v\n", data, err)
		resp.ErrorString = "Internal Server error"
		resp.Message = "ERROR"
		resp.StatusCode = http.StatusInternalServerError
		return
	}
	data, err = m.CalculationTotalPrice(data)
	resp.Message = "Success"
	resp.Data = data
	resp.StatusCode = http.StatusOK
	return
}
