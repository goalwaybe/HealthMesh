package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

func GenWareTotal() string {
	timeStr := time.Now().Format("20060102150405")
	randStr := rand.Intn(9000) + 1000
	return fmt.Sprintf("zy%s%d", timeStr, randStr)
}
