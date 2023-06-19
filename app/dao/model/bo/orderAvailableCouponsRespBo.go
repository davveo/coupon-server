package bo

// OrderAvailableCouponsRespBo 查询订单可用优惠券结果类
type OrderAvailableCouponsRespBo struct {
	GetAvailableCouponSelectListBo `json:"getAvailableCouponSelectListBo"`
	OrderId                        string `json:"orderId"`
	OrderNo                        string `json:"orderNo"`
}
