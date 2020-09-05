package transaction_status

func GetTransactionStatus() map[string]string {
	return map[string]string{
		"ACCEPTED":   "accepted",
		"ONPROGRESS": "on progress",
		"DECLINED":   "declined",
	}
}

func GetRequiredStatus() []string {
	transactionStatus := GetTransactionStatus()
	return []string{transactionStatus["ACCEPTED"],
		transactionStatus["DECLINED"],
		transactionStatus["ONPROGRESS"]}

}
