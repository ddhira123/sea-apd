package converter

import (
	merchant2 "github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/dto/request/merchant"
)

func ConvertUpdateMerchantRequestToEntity(request merchant.UpdateMerchantRequest) merchant2.Merchant {
	return merchant2.Merchant{
		Name:    request.Name,
		Brand:   request.Brand,
		Address: request.Address,
	}
}
