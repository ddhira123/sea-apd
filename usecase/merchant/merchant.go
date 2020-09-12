package merchant

import (
	"github.com/williamchang80/sea-apd/common/constants/mailer_type"
	"github.com/williamchang80/sea-apd/common/constants/merchant_status"
	"github.com/williamchang80/sea-apd/common/constants/user_role"
	"github.com/williamchang80/sea-apd/common/mailer"
	"github.com/williamchang80/sea-apd/common/mailer/factory"
	"github.com/williamchang80/sea-apd/domain/merchant"
	user "github.com/williamchang80/sea-apd/domain/user"
	request "github.com/williamchang80/sea-apd/dto/request/merchant"
	"github.com/williamchang80/sea-apd/dto/request/merchant/converter"
	user2 "github.com/williamchang80/sea-apd/dto/request/user"
)

type MerchantUsecase struct {
	mc      merchant.MerchantRepository
	usecase user.UserUsecase
}

type NotifyAdminMailer struct {
}

func NewMerchantUsecase(m merchant.MerchantRepository, usecase user.
UserUsecase) merchant.MerchantUsecase {
	mc := MerchantUsecase{mc: m, usecase: usecase}
	return mc
}

func ConvertMerchantRequestToEntity(m request.MerchantRequest) merchant.Merchant {
	return merchant.Merchant{
		Name:     m.Name,
		Balance:  0,
		UserId:   m.UserId,
		Brand:    m.Brand,
		Address:  m.Address,
		Approval: merchant_status.ToString(merchant_status.WAITING),
	}
}

func (m MerchantUsecase) UpdateMerchantBalance(request request.UpdateMerchantBalanceRequest) error {
	if err := m.mc.UpdateMerchantBalance(request.Amount, request.MerchantId); err != nil {
		return err
	}
	return nil
}

func (m MerchantUsecase) GetMerchantBalance(merchantId string) (int, error) {
	balance, err := m.mc.GetMerchantBalance(merchantId)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (m MerchantUsecase) RegisterMerchant(request request.MerchantRequest) error {
	merch := ConvertMerchantRequestToEntity(request)
	r, err := m.mc.RegisterMerchant(merch)
	if err != nil {
		return err
	}
	u, err := m.usecase.GetUserById(r.UserId)
	if err != nil {
		return err
	}
	notifyAdminOnMerchantRegister(*u)
	return nil
}

func notifyAdminOnMerchantRegister(u user.User) error {
	mail := factory.CreateMailerFactory(mailer_type.AUTH)
	mails := mail.CreateMail(u, user_role.ToString(user_role.MERCHANT))
	err := mailer.SendEmail(mails)
	if err != nil {
		return err
	}
	return nil
}

func (m MerchantUsecase) GetMerchants() ([]merchant.Merchant, error) {
	mh, err := m.mc.GetMerchants()
	if err != nil {
		return nil, err
	}
	return mh, nil
}

func (m MerchantUsecase) GetMerchantById(merchantId string) (*merchant.Merchant, error) {
	mh, err := m.mc.GetMerchantById(merchantId)
	if err != nil || mh == nil {
		return nil, err
	}
	return mh, nil
}

func (m MerchantUsecase) UpdateMerchantApprovalStatus(request request.
UpdateMerchantApprovalStatusRequest) error {
	if err := m.mc.UpdateMerchantApprovalStatus(request.MerchantId,
		merchant_status.ToString(request.Status)); err != nil {
		return err
	}
	if request.Status == merchant_status.ACCEPTED {
		merch, _ := m.mc.GetMerchantById(request.MerchantId)
		updateRequest := user2.UpdateUserRoleRequest{Role: user_role.MERCHANT,
			UserId: merch.UserId}
		if err := m.usecase.UpdateUserRole(updateRequest); err != nil {
			return err
		}
	}
	return nil
}

func (m MerchantUsecase) UpdateMerchant(r request.UpdateMerchantRequest) error {
	merchant := converter.ConvertUpdateMerchantRequestToEntity(r)
	if err := m.mc.UpdateMerchant(r.MerchantId, merchant); err != nil {
		return err
	}
	return nil
}
