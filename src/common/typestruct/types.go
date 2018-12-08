// typestruct package to contain common structure  of external dependency
// that are commonly used by several package in this repository
// for example the ACE Product Response
package typestruct

// === ACE Proudct Response ===
type AceProdRespProduct struct {
	ProductID      int64                  `json:"id,omitempty"`
	ProductName    string                 `json:"name,omitempty"`
	URL            string                 `json:"url,omitempty"`
	URLApps        string                 `json:"url_apps,omitempty"`
	URLMobile      string                 `json:"url_mobile,omitempty"`
	ImageURL       string                 `json:"image_url,omitempty"`
	ImageURL700    string                 `json:"image_url_700,omitempty"`
	PriceFormatted string                 `json:"price,omitempty"`
	Shop           AceProdRespShop        `json:"shop,omitempty"`
	Wholesale      []AceProdRespWholesale `json:"wholesale_price"`
	CourierCount   int64                  `json:"courier_count,omitempty"`
	Condition      int64                  `json:"condition,omitempty"`
	DepartmentId   int64                  `json:"department_id,omitempty"`
	// Labels []string
	// Badges             []ProdRespBadge `json:"badge,omitempty"`
	Rating             int64               `json:"rating,omitempty"`
	CountReview        int64               `json:"count_review,omitempty"`
	SKU                string              `json:"sku,omitempty"`
	PriceOriginal      string              `json:"original_price,omitempty"`
	PriceDiscounted    string              `json:"discounted_price,omitempty"`
	Stock              int64               `json:"stock"`
	mini_apiInfo       AceProdRespmini_api `json:"mini_api,omitempty"`
	DiscountPercentage int64               `json:"discount_percentage,omitempty"`
	Returnable         bool                `json:"returnable,omitempty"`
	LockStatus         int64               `json:"lock_status,omitempty"`
	MaxOrder           int64               `json:"max_order"`
	ParentID           int64               `json:"parent_id"`
	Score              int64               `json:"score"`
	IsPreOrder         bool                `json:"is_preorder"`
	CriteriaID         int64               `json:"criteria_id"`
}

type AceProdRespShop struct {
	ShopID         int64  `json:"id,omitempty"`
	ShopName       string `json:"name,omitempty"`
	URL            string `json:"url,omitempty"`
	IsGoldMerchant bool   `json:"is_gold,omitempty"`
	Location       string `json:"location,omitempty"`
	City           string `json:"city,omitempty"`
	Reputation     string `json:"reputation,omitempty"`
	Clover         string `json:"clover,omitempty"`
}

type AceProdRespWholesale struct {
	QuantityMin int64   `json:"quantity_min,omitempty"`
	QuantityMax int64   `json:"quantity_max,omitempty"`
	Price       float64 `json:"price,omitempty"`
}

type AceProdRespmini_api struct {
	mini_apiID               int64  `json:"mini_api_id,omitempty"`
	DiscountPercentage       int64  `json:"discount_percentage,omitempty"`
	DiscountedPriceFormatted string `json:"discounted_price,omitempty"`
	OriginalPriceFormatted   string `json:"original_price,omitempty"`
	Cashback                 int64  `json:"cashback,omitempty"`
	CustomStock              int64  `json:"custom_stock,omitempty"`
	StockSoldPercentage      int64  `json:"stock_sold_percentage,omitempty"`
	StatusID                 int64  `json:"status_id,omitempty"`
	AdminStatus              int64  `json:"admin_status,omitempty"`
	ProductMapStatus         int64  `json:"product_map_status,omitempty"`
	StartDate                string `json:"start_date,omitempty"`
	EndDate                  string `json:"end_date,omitempty"`
	MaxOrder                 int64  `json:"max_order,omitempty"`
	ShortName                string `json:"short_name,omitempty"`
	mini_apiType             string `json:"mini_api_type,omitempty"`
	OriginalMaxOrder         int64  `json:"original_max_order,omitempty"`
}

// === END ACE Product Response ===

// === Pacman Proudct ===
type PacmanProduct struct {
	ProductID                     int64   `json:"id"`
	ParentID                      int64   `json:"parent_id,omitempty"`
	DefaultChildID                int64   `json:"default_child_id,omitempty"`
	ShopID                        int64   `json:"shop_id,omitempty"`
	PacmanTimeUnix                int64   `json:"pacman_time_unix"`
	Stock                         int     `json:"stock,omitempty"`
	Status                        int     `json:"status"`
	Condition                     int     `json:"condition,omitempty"`
	Flag                          int     `json:"flag"`
	PriceLevel                    int     `json:"price_level"`
	StarRating                    int     `json:"star_rating"`
	DiscountPercentage            float64 `json:"discount_percentage,omitempty"`
	Returnable                    int     `json:"returnable"`
	CityID                        int     `json:"city_id"`
	CountView                     int     `json:"count_view"`
	CountReview                   int     `json:"count_review"`
	CountTalk                     int     `json:"count_talk"`
	CountSold                     int     `json:"count_sold,omitempty"`
	CountTransaction              int     `json:"count_transaction"`
	ServerID                      int     `json:"server_id"`
	CatalogID                     int     `json:"catalog_id,omitempty"`
	ChildCategoryID               int     `json:"child_category_id"`
	CurrencyID                    int     `json:"currency_id"`
	MenuID                        int     `json:"menu_id,omitempty"`
	Position                      int     `json:"position,omitempty"`
	ClusterID                     int     `json:"cluster_id"`
	countTransactionReject        int
	countTransactionRejectMonthly int
	countTransactionMonthly       int
	Price                         float64 `json:"price"`
	MinPrice                      float64 `json:"min_price,omitempty"`
	MaxPrice                      float64 `json:"max_price,omitempty"`
	Rating                        float64 `json:"rating"`
	BestMatchDink                 float64 `json:"dink_time,omitempty"`
	ReviewScore                   float64 `json:"review_score"`
	CashbackAmount                float64 `json:"cashback_amount,omitempty"`
	ProductScore                  float64 `json:"product_score,omitempty"`
	BestmatchScore4               float64 `json:"bestmatch_score4"`
	BestmatchScore4Fixed          float64 `json:"bestmatch_score4_fixed"`
	BestmatchScore5               float64 `json:"bestmatch_score5"`
	BestmatchScore5Fixed          float64 `json:"bestmatch_score5_fixed"`
	BestmatchScore6               float64 `json:"bestmatch_score6"`
	BestmatchScore6Fixed          float64 `json:"bestmatch_score6_fixed"`
	OriginalPrice                 float64 `json:"original_price"`
	DiscountedPrice               float64 `json:"discounted_price"`
	Gender                        float64 `json:"gender,omitempty"`
	HasCashback                   bool    `json:"has_cashback,omitempty"`
	HasDiscount                   bool    `json:"has_discount,omitempty"`
	IsPreorder                    bool    `json:"is_preorder,omitempty"`
	IsKreasiLokal                 bool    `json:"is_kreasi_lokal,omitempty"`
	IsNewMerchant                 bool    `json:"is_new_merchant,omitempty"`
	IsVariant                     bool    `json:"is_variant,omitempty"`
	IsGoldMerchant                bool    `json:"is_gm,omitempty"`
	IsFlashSale                   bool    `json:"is_flash_sale,omitempty"`
	Name                          string  `json:"name"`
	NameSort                      string  `json:"name_sort"`
	Keyword                       string  `json:"keyword"`
	CityName                      string  `json:"city_name"`
	URLPath                       string  `json:"url_path"`
	UpdateTime                    string  `json:"update_time,omitempty"`
	CreateTime                    string  `json:"create_time,omitempty"`
	Location                      string  `json:"location,omitempty"`
	ImagePath                     string  `json:"image_path,omitempty"`
	ImageName                     string  `json:"image_name,omitempty"`
	DiscountStartTime             string  `json:"discount_start_time,omitempty"`
	DiscountExpiredTime           string  `json:"discount_expired_time,omitempty"`
	DiscountStart                 string  `json:"discount_start,omitempty"`
	DiscountExpired               string  `json:"discount_expired,omitempty"`
	Sku                           string  `json:"sku,omitempty"`
	shippingList                  string
	catalogName                   string
	categoryName                  string
	shortDesc                     string
	ChildProductID                []int64                  `json:"child_ids,omitempty"`
	DepartmentID                  []int64                  `json:"department_id,omitempty"`
	TermID                        []int64                  `json:"term_id,omitempty"`
	Shipping                      []int64                  `json:"shipping,omitempty"`
	HotCurated                    []int64                  `json:"hot_curated,omitempty"`
	HotFeatured                   []int64                  `json:"hot_featured,omitempty"`
	HotlistID                     []int64                  `json:"hotlist_id,omitempty"`
	VariantsFilter                []string                 `json:"variants_filter,omitempty"`
	Wholesale                     []PacmanProductWholesale `json:"wholesale,omitempty"`
	Brand                         PacmanProductBrand       `json:"brand,omitempty"`
	Shop                          PacmanProductShop        `json:"shop,omitempty"`
	Variants                      map[string][]string      `json:"variants,omitempty"`
	mini_api                      Productmini_api          `json:"mini_api"`       // EXTRA field; used in getSellerPrdForSubmission
	IsValid                       bool                     `json:"is_valid"`       // EXTRA field; used in getSellerPrdForSubmission
	ProductStatus                 int16                    `json:"product_status"` // EXTRA field; used in getSellerPrdForSubmission
}

type PacmanProductWholesale struct {
	QuantityMin int     `json:"quantity_min"`
	QuantityMax int     `json:"quantity_max"`
	Price       float64 `json:"price"`
}

type PacmanProductShop struct {
	ID                  int64   `json:"id"`
	Name                string  `json:"name"`
	Domain              string  `json:"domain"`
	GoldMerchantExpired string  `json:"gold_merchant_expired,omitempty"`
	ShopScore           int     `json:"shop_score,omitempty"`
	ShopReputationScore int     `json:"shop_reputation,omitempty"`
	TempShopScore       int     `json:"-"`
	IsOfficial          bool    `json:"is_official,omitempty"`
	IsRegular           bool    `json:"is_regular,omitempty"`
	CancellationRate    float64 `json:"rate_cancel,omitempty"`
}

type PacmanProductBrand struct {
	ID   int64  `json:"ID"`
	Name string `json:"name,omitempty"`
}

// === END Pacman Product ===

type Productmini_api struct {
	mini_apiID         int64                   `db:"mini_api_id" json:"mini_api_id"`
	StartDate          string                  `json:"start_date"`
	EndDate            string                  `json:"end_date"`
	DiscountPercentage float64                 `db:"discount_percentage" json:"discount_percentage"`
	OriginalPrice      int64                   `db:"original_price" json:"original_price"`
	DiscountedPrice    int64                   `db:"discounted_price" json:"discounted_price"`
	Criteria           Productmini_apiCriteria `json:"criteria"`
}

type Productmini_apiCriteria struct {
	CriteriaID            int64   `db:"criteria_id" json:"criteria_id"`
	PriceMin              int64   `db:"min_price" json:"price_min"`
	PriceMax              int64   `db:"max_price" json:"price_max"`
	StockMin              int64   `db:"min_stock" json:"stock_min"`
	RatingMin             int64   `db:"minimum_rating" json:"rating_min"`
	RatingMax             int64   `db:"maximum_rating" json:"rating_max"`
	CashbackMin           int64   `db:"minimum_cashback" json:"cashback_minimum"`
	CashbackMax           int64   `db:"maximum_cashback" json:"cashback_maximum"`
	DiscountPercentageMin float64 `db:"minimum_discount_percent" json:"discount_percentage_min"`
	DiscountPercentageMax float64 `db:"maximum_discount_percent" json:"discount_percentage_max"`
	SubmissonMax          int64   `db:"max_submission" json:"submission_max"`
	SubmissionCount       int64   `db:"submitted" json:"submission_count"`
	ExcludePreOrder       bool    `db:"exclude_preorder" json:"exclude_preorder"`
	ExcludeWholesale      bool    `db:"exclude_wholesale" json:"exclude_wholesale"`
	MinOrderMin           int64   `db:"min_order_min" json:"min_order_min"`
	MinOrderMax           int64   `db:"min_order_max" json:"min_order_max"`
}
