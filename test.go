package main

import (
	"fmt"
	"reflect"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("^[1-9][0-9]+$", "01232325")
	fmt.Printf("Matched : %v\n", matched)

	str := `SELECT
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
	fmt.Println(reflect.TypeOf(str))
}
