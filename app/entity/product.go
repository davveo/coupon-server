package entity

type ProductDto struct {
	SkuId            string  `json:"skuId"`
	ProductTitle     string  `json:"productTitle"`
	ProductName      string  `json:"productName"`
	ShopId           string  `json:"shopId"`
	ShopName         string  `json:"shopName"`
	OriginPrice      float64 `json:"originPrice"`
	ActualPrice      float64 `json:"actualPrice"`
	PurchaseQuantity int64   `json:"purchaseQuantity"`
	StockQuantity    int64   `json:"stockQuantity"`
	SkuCategoryName  string  `json:"skuCategoryName"`
	ProductModelNo   string  `json:"productModelNo"`
	ProductSpec      string  `json:"productSpec"`
}
