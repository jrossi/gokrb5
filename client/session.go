package client

import (
	"github.com/jcmturner/gokrb5/types"
	"time"
)

type Session struct {
	AuthTime             time.Time
	EndTime              time.Time
	RenewTill            time.Time
	TGT                  types.Ticket
	SessionKey           types.EncryptionKey
	SessionKeyExpiration time.Time
}
