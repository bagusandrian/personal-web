package tax

import (
	"log"
)

func (m *Module) RegisterProduct(data ProductData) (result ProductData, err error) {
	rows, err := m.queries.InsertProduct.Query(data.Name, data.Price, data.TaxCodeID)
	if err != nil {
		log.Printf("[tax][RegisterUser] queries CheckIDTax err: %+v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Name, &result.Price, &result.TaxCodeID)
		if err != nil {
			log.Printf("[tax][RegisterUser] failed scan data : %+v", err)
			return
		}
	}
	return
}

func (m *Module) ListProducts() (result []ProductData, err error) {
	var product ProductData
	rows, err := m.queries.GetListProducts.Query()
	if err != nil {
		log.Printf("[tax][ListProducts] queries GetListProducts err: %+v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.TaxCodeID, &product.TaxType)
		if err != nil {
			log.Printf("[tax][ListProducts] failed scan data : %+v", err)
			return
		}
		result = append(result, product)
	}
	return

}
