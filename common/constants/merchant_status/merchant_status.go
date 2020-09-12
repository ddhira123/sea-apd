package merchant_status

type MerchantApprovalStatus int

const (
	WAITING MerchantApprovalStatus = iota
	DECLINED
	ACCEPTED
	OTHER
)

var MerchantStatusList = []string{
	"waiting",
	"declined",
	"accepted",
	"other",
}

func ToString(ms MerchantApprovalStatus) string {
	if ms < WAITING || ms > ACCEPTED {
		return ""
	}
	return MerchantStatusList[ms]
}

func ParseToEnum(src string) MerchantApprovalStatus {
	transactionStatusMap := map[string]MerchantApprovalStatus{
		"waiting":  WAITING,
		"declined": DECLINED,
		"accepted": ACCEPTED,
		"other":    OTHER,
	}
	if val, exist := transactionStatusMap[src]; exist {
		return val
	}
	return transactionStatusMap["other"]
}
