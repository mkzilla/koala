package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "github.com/mkzilla/koala/pkg/types/config"
)

func Page(c *gin.Context) {
	if len(c.Query(PageSize)) > 0 {
		pz, _ := strconv.Atoi(c.Query(PageSize))
		c.Set(PageSize, pz)
	} else {
		c.Set(PageSize, 10)
	}
	if len(c.Query(PageNumber)) > 0 {
		pn, _ := strconv.Atoi(c.Query(PageNumber))
		c.Set(PageNumber, pn)
	} else {
		c.Set(PageNumber, 0)
	}
	c.Set(PageOffset, c.GetInt(PageSize)*c.GetInt(PageNumber))
}
