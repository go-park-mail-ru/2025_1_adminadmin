package utils

import (
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateOrderInput(t *testing.T) {
	tests := []struct {
		name    string
		input   models.OrderInReq
		wantErr string
	}{
		{
			name: "Valid input",
			input: models.OrderInReq{
				Status:            "Ожидается",
				Address:           "г. Москва, ул. Пушкина, д. 10",
				ApartmentOrOffice: "12Б",
				Intercom:          "123",
				Entrance:          "1",
				Floor:             "3",
				CourierComment:    "Оставьте у двери",
				FinalPrice:        999,
			},
			wantErr: "",
		},
		{
			name: "Invalid status (too long)",
			input: models.OrderInReq{
				Status:            "оченьдлинноенекорректноезначениестатуса",
				Address:           "г. Москва",
				ApartmentOrOffice: "12",
				Intercom:          "123",
				Entrance:          "1",
				Floor:             "3",
				FinalPrice:        100,
			},
			wantErr: "некорректный статус (макс 20 символов)",
		},
		{
			name: "Invalid address (empty)",
			input: models.OrderInReq{
				Status:            "Ожидается",
				Address:           "",
				ApartmentOrOffice: "12",
				Intercom:          "123",
				Entrance:          "1",
				Floor:             "3",
				FinalPrice:        100,
			},
			wantErr: "некорректный адрес (макс 200 символов)",
		},
		{
			name: "Invalid apartment (contains !)",
			input: models.OrderInReq{
				Status:            "Ожидается",
				Address:           "г. Москва",
				ApartmentOrOffice: "12!",
				Intercom:          "123",
				Entrance:          "1",
				Floor:             "3",
				FinalPrice:        100,
			},
			wantErr: "некорректная квартира/офис (макс 20 символов)",
		},
		{
			name: "Invalid comment (too long)",
			input: models.OrderInReq{
				Status:            "Ожидается",
				Address:           "г. Москва",
				ApartmentOrOffice: "12",
				Intercom:          "123",
				Entrance:          "1",
				Floor:             "3",
				CourierComment:    string(make([]byte, 301)),
				FinalPrice:        100,
			},
			wantErr: "некорректный комментарий (макс 300 символов)",
		},
		{
			name: "Negative price",
			input: models.OrderInReq{
				Status:            "Ожидается",
				Address:           "г. Москва",
				ApartmentOrOffice: "12",
				Intercom:          "123",
				Entrance:          "1",
				Floor:             "3",
				FinalPrice:        -10,
			},
			wantErr: "цена не может быть отрицательной",
		},
		{
			name: "Valid status with max length",
			input: models.OrderInReq{
				Status:            "12345678901234567890", 
				Address:           "г. Москва",
				ApartmentOrOffice: "12",
				Intercom:          "123",
				Entrance:          "1",
				Floor:             "3",
				FinalPrice:        100,
			},
			wantErr: "",
		},
		{
			name: "Valid address with max length",
			input: models.OrderInReq{
				Status:            "Ожидается",
				Address:           strings.Repeat("a", maxAddressLength), 
				ApartmentOrOffice: "12",
				Intercom:          "123",
				Entrance:          "1",
				Floor:             "3",
				FinalPrice:        100,
			},
			wantErr: "",
		},
		{
			name: "Invalid address with too long value",
			input: models.OrderInReq{
				Status:            "Ожидается",
				Address:           strings.Repeat("a", maxAddressLength+1), // превышает максимальную длину
				ApartmentOrOffice: "12",
				Intercom:          "123",
				Entrance:          "1",
				Floor:             "3",
				FinalPrice:        100,
			},
			wantErr: "некорректный адрес (макс 200 символов)",
		},
		{
			name: "Empty fields",
			input: models.OrderInReq{
				Status:            "",
				Address:           "",
				ApartmentOrOffice: "",
				Intercom:          "",
				Entrance:          "",
				Floor:             "",
				FinalPrice:        100,
			},
			wantErr: "некорректный статус (макс 20 символов)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateOrderInput(&tt.input)
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
