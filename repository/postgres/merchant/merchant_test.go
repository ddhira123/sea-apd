package merchant

import (
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/common/constants/merchant_status"
	domain "github.com/williamchang80/sea-apd/domain/merchant"
	mock_psql "github.com/williamchang80/sea-apd/mocks/postgres"
	"reflect"
	"testing"
)

func TestNewMerchantRepository(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want domain.MerchantRepository
	}{
		{
			name: "success with null value on db",
			args: args{
				db: nil,
			},
			want: &MerchantRepository{
				db: nil,
			},
		},
		{
			name: "success with value on db",
			args: args{
				db: db,
			},
			want: &MerchantRepository{
				db: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMerchantRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerchantRepository_GetMerchantBalance(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		merchantId string
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "failed with not matched query",
			args: args{
				merchantId: "",
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := MerchantRepository{
				db: tt.initMock(),
			}
			balance, err := pr.GetMerchantBalance(tt.args.merchantId)
			if err != nil && !tt.wantErr {
				t.Errorf("MerchantRepository.GetMerchantBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(balance, tt.want) {
				t.Errorf("MerchantRepository.Find() = %v, want %v", balance, tt.want)
			}
		})
	}
}

func TestMerchantRepository_UpdateMerchantBalance(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		merchantId string
		amount     int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "failed with not matched query",
			args: args{
				merchantId: "",
				amount:     0,
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
		{
			name: "failed",
			args: args{
				merchantId: "1",
				amount:     10000,
			},
			want:    10000,
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := MerchantRepository{
				db: tt.initMock(),
			}
			err := pr.UpdateMerchantBalance(tt.args.amount, tt.args.merchantId)
			if err != nil && !tt.wantErr {
				t.Errorf("MerchantRepository.UpdateMerchantBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMerchantRepository_GetMerchantById(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		merchantId string
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "failed with not matched query",
			args: args{
				merchantId: "",
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
		{
			name: "failed",
			args: args{
				merchantId: "1",
			},
			want:    10000,
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := MerchantRepository{
				db: tt.initMock(),
			}
			_, err := pr.GetMerchantById(tt.args.merchantId)
			if err != nil && !tt.wantErr {
				t.Errorf("MerchantRepository.GetMerchantById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMerchantRepository_GetMerchantByUserId(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		userId string
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "failed with not matched query",
			args: args{
				userId: "",
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
		{
			name: "success",
			args: args{
				userId: "1",
			},
			want:    10000,
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := MerchantRepository{
				db: tt.initMock(),
			}
			_, err := pr.GetMerchantByUserId(tt.args.userId)
			if err != nil && !tt.wantErr {
				t.Errorf("MerchantRepository.GetMerchantByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMerchantRepository_GetMerchants(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	tests := []struct {
		name     string
		want     []domain.Merchant
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name:    "success",
			want:    []domain.Merchant{},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := MerchantRepository{
				db: tt.initMock(),
			}
			_, err := pr.GetMerchants()
			if err != nil && !tt.wantErr {
				t.Errorf("MerchantRepository.GetMerchants() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMerchantRepository_RegisterMerchant(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		merchant domain.Merchant
	}
	tests := []struct {
		name     string
		want     domain.Merchant
		args     args
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name:    "failed",
			want:    domain.Merchant{},
			args:    args{merchant: domain.Merchant{}},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
		{
			name: "success",
			want: domain.Merchant{},
			args: args{merchant: domain.Merchant{
				Address: "test",
			}},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := MerchantRepository{
				db: tt.initMock(),
			}
			_, err := pr.RegisterMerchant(tt.args.merchant)
			if err != nil && !tt.wantErr {
				t.Errorf("MerchantRepository.RegisterMerchant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMerchantRepository_UpdateMerchant(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		merchant   domain.Merchant
		merchantId string
	}
	tests := []struct {
		name     string
		want     domain.Merchant
		args     args
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name:    "failed",
			want:    domain.Merchant{},
			args:    args{merchant: domain.Merchant{}, merchantId: ""},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
		{
			name: "success",
			want: domain.Merchant{},
			args: args{merchant: domain.Merchant{
				Address: "test",
			},
				merchantId: "1"},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := MerchantRepository{
				db: tt.initMock(),
			}
			err := pr.UpdateMerchant(tt.args.merchantId, tt.args.merchant)
			if err != nil && !tt.wantErr {
				t.Errorf("MerchantRepository.UpdateMerchant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMerchantRepository_UpdateMerchantApprovalStatus(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		status     string
		merchantId string
	}
	tests := []struct {
		name     string
		want     domain.Merchant
		args     args
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "failed",
			want: domain.Merchant{},
			args: args{status: merchant_status.ToString(merchant_status.DECLINED),
				merchantId: ""},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
		{
			name: "success",
			want: domain.Merchant{},
			args: args{status: merchant_status.ToString(merchant_status.ACCEPTED),
				merchantId: "1"},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := MerchantRepository{
				db: tt.initMock(),
			}
			err := pr.UpdateMerchantApprovalStatus(tt.args.merchantId, tt.args.status)
			if err != nil && !tt.wantErr {
				t.Errorf("MerchantRepository.UpdateMerchantApprovalStatus() error = %v, " +
					"wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
