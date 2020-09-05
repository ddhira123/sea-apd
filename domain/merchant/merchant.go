package merchant

import (
	"github.com/williamchang80/sea-apd/domain"
)

type Merchant struct {
	domain.Base
	name string
}
