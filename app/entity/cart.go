package entity

type CartDto struct {
	ShopId         string        `json:"shopId"`         // 购物车所属商家id
	ShopName       string        `json:"shopName"`       // 购物车所属商家id购物车所属商铺名称
	ProductDtoList []*ProductDto `json:"productDtoList"` // 购物车下订单项列表
}

type AddCartDto struct {
	Member  MemberDto  `json:"member"`
	Product ProductDto `json:"product"`
}

type CartDtoList struct {
	LocalCartList []*CartDto `json:"localCartList"`
	MemberDto     *MemberDto `json:"memberDto"`
}
