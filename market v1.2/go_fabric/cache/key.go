package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey = "rank"  //排名
)

func ProductViewKey(id uint) string{
	return fmt.Sprintf("view:product:%s",strconv.Itoa(int(id)))
}