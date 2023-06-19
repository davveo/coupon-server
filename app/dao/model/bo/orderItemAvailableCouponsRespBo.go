package bo

type OrderItemAvailableCouponsRespBo struct {
	GetAvailableCouponSelectListBo
	SkuId       string // 商品/服务SKU ID
	SkuCategory string //商品/服务分类
}
