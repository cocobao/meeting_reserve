package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var weekDayS = []string{"一", "二", "三", "四", "五", "六", "日"}

func GetDate(ctx *gin.Context) {
	now := time.Now()
	result := []map[string]string{}

	for i := 0; i < 6; i++ {
		s := fmt.Sprintf("%d月%d日", now.Month(), now.Day())

		switch i {
		case 0:
			result = append(result, map[string]string{
				"今天": s,
			})
		case 1:
			result = append(result, map[string]string{
				"明天": s,
			})
		case 2:
			result = append(result, map[string]string{
				"后台": s,
			})
		default:
			result = append(result, map[string]string{
				"星期" + weekDayS[now.Weekday()]: s,
			})
		}

		now = now.AddDate(0, 0, 1)
	}

	ctx.JSON(200, result)
}
