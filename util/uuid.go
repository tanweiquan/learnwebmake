package util

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func GetUUID() string {
	id := uuid.NewV1().String()
	t0 := id[0:8]
	t1 := id[9:13]
	t2 := id[14:18]
	t3 := id[19:]
	return fmt.Sprintf("%s-%s-%s-%s", t2, t1, t3, t0)

}
