package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
    "strconv"
)

func main() {
	router := gin.Default()
	router.Use(HeaderMiddleware()/*, CheckLoginMiddleware()*/)

	router.GET("/commodity/list", func(c *gin.Context) {
        page := DEFAULT_PAGE
        if pageStr := c.Query("page"); pageStr != "" {
            page, _ = strconv.Atoi(pageStr)
        }

        pageSize := DEFAULT_PAGE_SIZE
        if pageSizeStr := c.Query("page_size"); pageSizeStr != "" {
            pageSize, _ = strconv.Atoi(pageSizeStr)
        }

		commodities := QueryCommoditiesByPage(pageSize, page)
		jsonResult := new(JsonResult)
		jsonResult.Status = http.StatusOK
		jsonResult.Data.List.Total = QueryCommodityCount()
		jsonResult.Data.List.PerPage = uint8(pageSize)
		jsonResult.Data.List.CurrentPage = uint32(page)
		jsonResult.Data.List.Data = make([]interface{}, len(commodities))
		for i, v := range commodities {
			jsonResult.Data.List.Data[i] = v
		}

		c.JSON(http.StatusOK, jsonResult)
	})

	router.Run(":8081")
}
