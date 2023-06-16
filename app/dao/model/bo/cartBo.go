package bo

type CartBo struct {
	ShopId        string       `json:"shopId"`        // 购物车所属商家id
	ShopName      string       `json:"shopName"`      // 购物车所属商铺名称
	ProductBoList []*ProductBo `json:"productBoList"` // 购物车下订单项列表
}
