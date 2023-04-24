package util

import (
	"fmt"
	"time"
)

func GenerateAID(accountId int8) string {
	timestamp := time.Now().UnixMilli()
	return fmt.Sprintf("%s_%d%d", "ADID", timestamp, accountId)
}
