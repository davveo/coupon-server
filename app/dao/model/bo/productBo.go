package bo

import "github.com/shopspring/decimal"

type ProductBo struct {
	SkuId          string          `json:"skuId"`          // 商品id
	Title          string          `json:"title"`          // 商品标题
	GoodsName      string          `json:"goodsName"`      // 商品名称
	ShopId         string          `json:"shopId"`         // 商家id
	ShopName       string          `json:"shopName"`       // 商铺名称
	Price          decimal.Decimal `json:"price"`          // 商品价格
	CostPrice      decimal.Decimal `json:"costPrice"`      // 优惠后价格
	Quantity       int64           `json:"quantity"`       // 购买数量
	StockCount     int64           `json:"stockCount"`     // 库存数量
	SkuCategory    string          `json:"skuCategory"`    // 商品/服务分类
	ProductModelNo string          `json:"productModelNo"` // 商品型号
	Specification  string          `json:"specification"`  // 商品规格
	ImageUrl       string          `json:"imageUrl"`       // 商品图片地址
}
