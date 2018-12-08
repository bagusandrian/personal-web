package tax

import (
	"log"
)

func (m *Module) ListTaxCode() (result TaxLists, err error) {
	var taxList TaxList
	var list []TaxList
	rows, err := m.queries.CheckIDTax.Query()
	if err != nil {
		log.Printf("[tax][listTaxCode] queries CheckIDTax err: %+v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&taxList.IDTax, &taxList.Type)
		if err != nil {
			log.Printf("[tax][ListTaxCode] failed scan data : %+v", err)
			return
		}
		list = append(list, taxList)
	}
	result.List = list
	return
}

func (m *Module) GetProductDetail(data ProductData) (result ProductData, err error) {
	rows, err := m.queries.GetProductByProductID.Query(data.ID)
	if err != nil {
		log.Printf("[tax][GetProductDetail] queries GetProductByProudctID err: %+v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data.ID, &data.Name, &data.Price, &data.TaxCodeID)
		if err != nil {
			log.Printf("[tax][GetProductDetail] failed scan data : %+v", err)
			return
		}
	}
	result = data
	return
}
