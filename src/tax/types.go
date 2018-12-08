package tax

type (
	DataTransaction struct {
		ID                int64                 `json:"id"`
		UserID            int64                 `json:"user_id"`
		TransactionDate   string                `json:"transaction_date"`
		DetailTransaction []DataTransactionItem `json:"detail_transaction,omitempty"`
		TotalPrice        float64               `json:"total_price,omitempty"`
	}
	DataTransactionItem struct {
		ID            int64       `json:"transaction_item_id"`
		Product       ProductData `json:"Product"`
		Tax           float64     `json:"tax"`
		SubTotalPrice float64     `json:"sub_total_price"`
	}
	TaxLists struct {
		List []TaxList `json:"tax_list"`
	}
	TaxList struct {
		IDTax int    `json:"id"`
		Type  string `json:"type"`
	}
	UserData struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	ProductData struct {
		ID        int64   `json:"id"`
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		TaxCodeID int     `json:"tax_code"`
		TaxType   string  `json:"tax_type,omitempty"`
		Quantity  int     `json:"quantity,omitempty"`
	}
)
