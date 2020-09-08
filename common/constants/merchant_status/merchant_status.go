package merchant_status

func GetMerchantStatus() map[string]string {
	return map[string]string{
		"WAITING":  "waiting",
		"APPROVED": "approved",
		"DECLINED": "declined",
	}
}
