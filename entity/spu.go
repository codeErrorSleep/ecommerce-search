package entity

type Spu struct {
	SpuID       int      `json:"spu_id"`
	SpuType     int      `json:"spu_type"`
	GoodsName   string   `json:"goods_name"`
	Price       float64  `json:"price"`
	VisitNum    int      `json:"visit_num"`
	GoodsTag    []string `json:"goods_tag"`
	SaleStatus  int      `json:"sale_status"`
	AppraiseNum int      `json:"appraise_num"`
	IsDeleted   int      `json:"is_deleted"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}