package main

import "fmt"
import "regexp"

const (
	QUERY_TYPE_ID   = "field_cid"
	QUERY_TYPE_NAME = "field_cname"
	SELECT_SQL_BASE = `SELECT
  id,
  name,
  type,
  agency_type,
  link_url,
  promotion_image,
  price,
  commission_rate,
  creation_time,
  promotion_time_type,
  promotion_time_start,
  promotion_time_end,
  biz_id,
  status
FROM cps_commodity`
	COUNT_SQL_BASE = "select count(id) from cps_commodity"
)

type CommodityQuery struct {
	QueryType    string
	QueryContent string
	RangeStart   uint32
	RangeEnd     uint32
	BizId        uint16
}

func (query *CommodityQuery) Validate() bool {
	if len(query.QueryContent) > 0 && query.QueryType == QUERY_TYPE_ID {
		matched, err := regexp.MatchString("^[1-9][1-9]+$", query)
		if err == nil || !matched {
			return false
		}
	}

	if query.RangeEnd < query.RangeStart {
		return false
	}

	return true
}

func (query *CommodityQuery) BuildListSql() (sql string) {
	arr := []string{}
	if query.BizId > 0 {
		arr = append(arr, "biz_id="+query.BizId)
	}

	return nil
}

type Commodity struct {
	Id                 uint64  `json:"cid"`
	Name               string  `json:"cname"`
	Type               uint8   `json:"ctype"`
	LinkUrl            string  `json:"link_url"`
	AgencyType         uint8   `json:"agency_type"`
	PromotionImage     string  `json:"promotion_image"`
	Price              float32 `json:"price"`
	CommissionRate     float32 `json:"commission_rate"`
	CreationTime       uint32  `json:"creation_time"`
	PromotionTimeType  uint8   `json:"promotion_time_type"`
	PromotionTimeStart uint32  `json:"promotion_time_range_start"`
	PromotionTimeEnd   uint32  `json:"promotion_time_range_end"`
	BizId              uint16  `json:"biz_id"`
	Status             uint8   `json:"status"`
}

func QueryCommodityCount() uint32 {
	var count uint32 = 0
	query := "select count(id) from cps_commodity"
	row := db.QueryRow(query)
	row.Scan(&count)

	return count
}

func QueryCommoditiesByPage(pageSize, page int) []*Commodity {
	offset := pageSize * (page - 1)
	query := "select id, name, type, agency_type, link_url, promotion_image, price, " +
		"commission_rate, creation_time, promotion_time_type, promotion_time_start, " +
		"promotion_time_end, biz_id, status from cps_commodity limit ?,?"
	stmt, err := db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil
	}

	rows, err := stmt.Query(offset, pageSize)
	fmt.Printf("offset : %v, limit : %v\n", offset, pageSize)
	if err != nil {
		return nil
	}

	var commodities []*Commodity
	for rows.Next() {
		commodity := &Commodity{}
		rows.Scan(&commodity.Id, &commodity.Name, &commodity.Type,
			&commodity.AgencyType, &commodity.LinkUrl, &commodity.PromotionImage,
			&commodity.Price, &commodity.CommissionRate, &commodity.CreationTime,
			&commodity.PromotionTimeType, &commodity.PromotionTimeStart,
			&commodity.PromotionTimeEnd, &commodity.BizId, &commodity.Status)
		fmt.Printf("commodity : %+v\n", commodity)
		commodities = append(commodities, commodity)
	}

	return commodities
}

func QueryCommodityById(id uint64) *Commodity {
	query := "select id, name, type, agency_type, link_url, promotion_image, price, " +
		"commission_rate, promotion_time_type, promotion_time_start, promotion_time_end, " +
		"biz_id, status from cps_commodity where id=?"
	stmt, err := db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil
	}

	var commodity Commodity = Commodity{}
	stmt.QueryRow(id).Scan(&commodity.Id, &commodity.Name, &commodity.Type,
		&commodity.AgencyType, &commodity.LinkUrl, &commodity.PromotionImage,
		&commodity.Price, &commodity.CommissionRate, &commodity.PromotionTimeType,
		&commodity.PromotionTimeStart, &commodity.PromotionTimeEnd, &commodity.BizId)

	return &commodity
}
