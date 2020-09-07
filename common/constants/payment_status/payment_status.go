package payment_status

func GetPaymentStatus() map[string]string {
	return map[string]string{
		"ACCEPTED":   "accepted",
		"ONPROGRESS": "on progress",
		"DECLINED":   "declined",
	}
}
