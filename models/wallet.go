package models

import "time"

type Wallet struct {
	OwnedBy    string
	IsActive   bool
	Balance    float64
	EnabledAt  []time.Time
	DisabledAt []time.Time
	Transaction
}
