package mailer

import (
	"github.com/williamchang80/sea-apd/common/constants/user_role"
	"github.com/williamchang80/sea-apd/common/mailer"
	"github.com/williamchang80/sea-apd/domain/user"
	"reflect"
	"testing"
)

func TestAuthMailer_CreateMail(t *testing.T) {
	type args struct {
		user user.User
		role string
	}
	tests := []struct {
		name    string
		args    args
		want    []mailer.Mail
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				user: user.User{},
				role: user_role.ToString(user_role.MERCHANT),
			},
			want:    []mailer.Mail{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AuthMailer{}
			if got := a.CreateMail(tt.args.user, tt.args.role);
				reflect.TypeOf(tt.want) != reflect.TypeOf(got) {
				t.Errorf("NewAuthUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateMerchantProposalMailer(t *testing.T) {
	type args struct {
		user user.User
	}
	tests := []struct {
		name    string
		args    args
		want    []mailer.Mail
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				user: user.User{},
			},
			want:    []mailer.Mail{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMerchantProposalMailer(tt.args.user);
				reflect.TypeOf(tt.want) != reflect.TypeOf(got) {
				t.Errorf("NewAuthUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
