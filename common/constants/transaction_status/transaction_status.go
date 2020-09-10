package transaction_status

type TransactionStatus int

const (
	ON_CARTS = iota
	WAITING_PAYMENT
	WAITING_CONFIRMATION
	DECLINED
	WAITING_DELIVERY
	ACCEPTED
	OTHER
)

var TransactionStatusList = []string{
	"on carts",
	"waiting payment",
	"waiting confirmation",
	"declined",
	"waiting delivery",
	"accepted",
	"other",
}

func ToString(ts TransactionStatus) string {
	if ts < ON_CARTS || ts > OTHER {
		return ""
	}
	return TransactionStatusList[ts]
}

func ParseToEnum(src string) TransactionStatus {
	transactionStatusMap := map[string]TransactionStatus{
		"on carts":             ON_CARTS,
		"waiting payment":      WAITING_PAYMENT,
		"waiting confirmation": WAITING_CONFIRMATION,
		"declined":             DECLINED,
		"waiting delivery":     WAITING_DELIVERY,
		"accepted":             ACCEPTED,
		"other":                OTHER,
	}
	if val, exist := transactionStatusMap[src]; exist {
		return transactionStatusMap[string(val)]
	}
	return transactionStatusMap["other"]
}

func GetStatusListForTransactionHistory() []string {
	transactionHistoryStatusEnumList := []TransactionStatus{
		ACCEPTED,
		DECLINED,
		WAITING_DELIVERY,
	}
	var transactionHistoryStatusList []string
	for _, t := range transactionHistoryStatusEnumList {
		transactionHistoryStatusList =
			append(transactionHistoryStatusList, ToString(t))
	}
	return transactionHistoryStatusList
}
