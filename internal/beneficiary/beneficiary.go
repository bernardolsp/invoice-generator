package beneficiary

import "time"

type Beneficiary struct {
	ID               int
	Name             string
	Address          string
	Email            *string
	Currency         string
	AddedDate        time.Time
	LastModifiedDate time.Time
}
