package entity

type QueryBySpuIDReq struct {
	SpuID string `json:"spu_id"`
}
type QueryByGoodsName struct {
	GoodsName string `json:"goods_name"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
}
