package utils

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func GenerateID() string {
	u1 := uuid.Must(uuid.NewV4())

	return fmt.Sprintf("%s", u1)
}
