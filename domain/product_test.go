package domain

import (
	"reflect"
	"testing"
)

func TestNewProduct(t *testing.T) {
	type args struct {
		name  string
		desc  string
		price int
		image string
		stock int
	}

	tests := []struct {
		name string
		args args
		want Product
	}{
		{
			name: "success",
			args: args{
				name:  "test name",
				desc:  "description for test",
				price: 50,
				image: "mock_image",
				stock: 60,
			},
			want: Product{
				Name:        "test name",
				Description: "description for test",
				Price:       50,
				Image:       "mock_image",
				Stock:       60,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.args
			if got := NewProduct(a.name, a.desc, a.price, a.image, a.stock);
				!reflect.DeepEqual(got, tt.want) {
				t.Errorf("func NewProduct() = %#v, but want %#v", got, tt.want)
			}
		})
	}
}
