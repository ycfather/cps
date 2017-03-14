package common

type Commodity struct {
	Id                      uint64  `json:"cid"`
	Name                    string  `json:"cname"`
	Type                    uint8   `json:"ctype"`
	AgencyType              uint8   `json:"agency_type"`
	LinkUrl                 string  `json:"link_url"`
	PromotionImage          string  `json:"promotion_image"`
	Price                   float32 `json:"price"`
	CommissionRate          float32 `json:"commission_rate"`
	CreationTime            uint32  `json:"creation_time"`
	PromotionTimeType       uint8   `json:"promotion_time_type"`
	PromotionTimeRangeStart uint32  `json:"promotion_time_range_start"`
	PromotionTimeRangeEnd   uint32  `json:"promotion_time_range_end"`
	BizId                   uint16  `json:"biz_id"`
	Status                  uint8   `json:"status"`
}
