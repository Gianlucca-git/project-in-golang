package UnitTest

import (
	"IMPORTS/model/dto"
	"IMPORTS/service"
	"fmt"
	"testing"
)

func TestServiceStruct_ValidatedRequestBalance(t *testing.T) {

	sm := &service.ServiceStruct{}

	type args struct {
		request *dto.BalanceRequest
	}
	tests := []struct {
		name             string
		args             args
		wantMessageError string
	}{
		{
			name: "invalid mount in Months",
			args: args{request: &dto.BalanceRequest{
				Months: []string{"Enero", "Error"},
				Sales:  []int{1, 2},
				Bills:  []int{1, 2},
			},
			},
			wantMessageError: fmt.Sprintf("%s (value = %s)", sm.Constants("InvalidMonth"), "Error"),
		},
		{
			name: "invalid length in a list",
			args: args{request: &dto.BalanceRequest{
				Months: []string{"Enero", "Febrero"},
				Sales:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 101},
				Bills:  []int{1, 2},
			},
			},
			wantMessageError: sm.Constants("InvalidLengthList"),
		},
		{
			name: "the list Sale is not the same length",
			args: args{request: &dto.BalanceRequest{
				Months: []string{"Enero", "Febrero"},
				Sales:  []int{1, 2, 3},
				Bills:  []int{1, 2},
			},
			},
			wantMessageError: sm.Constants("InvalidLengths"),
		},
		{
			name: "integer not positive",
			args: args{request: &dto.BalanceRequest{
				Months: []string{"Enero", "Marzo", "abril"},
				Sales:  []int{0, -22, 341},
				Bills:  []int{1, 2, 0},
			},
			},
			wantMessageError: fmt.Sprintf("%s (value = %d)", sm.Constants("InvalidNumber"), -22),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sm.ValidatedRequestBalance(tt.args.request); tt.wantMessageError != err.Error() {
				t.Errorf("ValidatedRequestBalance() error = [%v], wantMessageError = [%v]", err.Error(), tt.wantMessageError)
			}
		})
	}
}
