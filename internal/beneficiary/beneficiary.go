package beneficiary

import (
	"fmt"
	"time"
)

type Beneficiary struct {
	ID               int
	Name             string
	Address          string
	Email            *string
	Currency         string
	AddedDate        time.Time
	LastModifiedDate time.Time
}

func (b Beneficiary) String() string {
	return fmt.Sprintf("Beneficiary %d: %s, %s", b.ID, b.Name, *b.Email)
}
