package tax

import (
	"log"
)

func (m *Module) GetListTransaction() (result []DataTransaction, err error) {
	a := DataTransaction{}
	rows, err := m.queries.GetDataTransaction.Query()
	if err != nil {
		log.Printf("[tax][GetListTransaction] queries GetDataTransaction err: %+v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&a.ID, &a.UserID, &a.TransactionDate)
		if err != nil {
			log.Printf("[tax][CreateTransaction] failed scan data : %+v", err)
			return
		}
		result = append(result, a)
	}
	return
}

func (m *Module) CreateTransaction(data DataTransaction) (result DataTransaction, err error) {
	result = data
	rows, err := m.queries.InsertTransaction.Query(result.UserID)
	if err != nil {
		log.Printf("[tax][CreateTransaction] queries InsertTransaction err: %+v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result.ID, &result.UserID, &result.TransactionDate)
		if err != nil {
			log.Printf("[tax][CreateTransaction] failed scan data : %+v", err)
			return
		}
	}
	return
}

func (m *Module) GetTransactionByID(data DataTransaction) (result DataTransaction, err error) {
	result = data
	rows, err := m.queries.GetTransactionByIDTransaction.Query(result.ID)
	if err != nil {
		log.Printf("[tax][GetTransactionBYID] queries GetTransactionByIDTransaction err: %+v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result.ID, &result.UserID, &result.TransactionDate)
		if err != nil {
			log.Printf("[tax][GetTransactionBYID] failed scan data : %+v", err)
			return
		}
	}
	return
}

func (m *Module) GetDetailTransactionByIDTransaction(data DataTransaction) (result DataTransaction, err error) {
	result = data
	var item DataTransactionItem
	rows, err := m.queries.GetDetailTransactionByIDTransaction.Query(result.ID)
	if err != nil {
		log.Printf("[tax][GetTransactionBYID] queries GetTransactionByIDTransaction err: %+v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.Product.ID, &item.Product.Name, &item.Product.Price, &item.Product.TaxCodeID, &item.Product.TaxType, &item.Product.Quantity)
		if err != nil {
			log.Printf("[tax][GetTransactionBYID] failed scan data : %+v", err)
			return
		}
		result.DetailTransaction = append(result.DetailTransaction, item)
	}
	return

}

func (m *Module) CreateTransactionItem(data DataTransaction) (result DataTransaction, err error) {
	result = data
	for i, item := range result.DetailTransaction {
		rows, err := m.queries.InsertTransactionItem.Query(result.ID, item.Product.ID, item.Product.Quantity)
		if err != nil {
			log.Printf("[tax][CreateTransactionItem] queries InsertTransactionItem err: %+v\n", err)
			return result, err
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&item.ID, &data.ID, &item.Product.ID, &item.Product.Quantity)
			if err != nil {
				log.Printf("[tax][CreateTransactionItem] failed scan data : %+v", err)
				return result, err
			}
		}
		result.DetailTransaction[i] = item
	}
	return
}
func (m *Module) CalculationTotalPrice(data DataTransaction) (result DataTransaction, err error) {
	var totalPrice float64
	totalPrice = 0
	result = data
	for i, item := range result.DetailTransaction {
		item.Product, err = m.GetProductDetail(item.Product)
		if err != nil {
			log.Printf("[tax][CalculationTotalPrice] productID: %d error GetProductDetail err:%+v\n", item.Product.ID, err)
			return
		}
		item = m.GenerateTaxProduct(item)
		result.DetailTransaction[i] = item
		totalPrice = totalPrice + item.SubTotalPrice
	}
	result.TotalPrice = totalPrice
	return
}

func (m *Module) GenerateTaxProduct(data DataTransactionItem) (result DataTransactionItem) {
	if data.Product.TaxCodeID == TaxFoodID {
		data.Tax = float64(data.Product.Quantity * int(data.Product.Price) * 10 / 100)
		data.SubTotalPrice = data.Tax + float64(int(data.Product.Price)*data.Product.Quantity)
	} else if data.Product.TaxCodeID == TaxTobaccoID {
		data.Tax = float64(data.Product.Quantity*int(data.Product.Price)*2/100) + 10
		data.SubTotalPrice = data.Tax + float64(int(data.Product.Price)*data.Product.Quantity)
	} else if data.Product.TaxCodeID == TaxEntertainmentID {
		if data.Product.Quantity > 100 {
			data.Tax = float64((data.Product.Quantity - 100) * int(data.Product.Price) * 1 / 100)
			data.SubTotalPrice = data.Tax + float64(int(data.Product.Price)*data.Product.Quantity)
		}
		data.Tax = 0
		data.SubTotalPrice = data.Tax + float64(int(data.Product.Price)*data.Product.Quantity)
	}

	result = data
	return
}
