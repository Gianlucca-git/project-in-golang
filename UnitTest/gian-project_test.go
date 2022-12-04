package UnitTest

import (
	"IMPORTS/model/dto"
	"IMPORTS/repository"
	"IMPORTS/service"
	"testing"
)

func Test_serviceStruct_OrderList(t *testing.T) {
	type fields struct {
		ReplaceManager repository.ReplaceManager
		Utilities      service.Utilities
	}
	type args struct {
		list *dto.ClassifiedList
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantList []int
		wantErr  bool
	}{
		{
			name: "Test 01",
			args: args{
				list: &dto.ClassifiedList{
					Unclassified: []int{3, 5, 5, 6, 8, 3, 4, 4, 7, 7, 1, 1, 2},
				},
			},
			wantList: []int{1, 2, 3, 4, 5, 6, 7, 8, 5, 3, 4, 7, 1},
		}, {
			name: "Test 02",
			args: args{
				list: &dto.ClassifiedList{
					Unclassified: []int{1, 8, 3, 8, 56, 0, -12, 6, 6, 13, -2},
				},
			},
			wantList: []int{-12, -2, 0, 1, 3, 6, 8, 13, 56, 8, 6},
		}, {
			name: "Test 03",
			args: args{
				list: &dto.ClassifiedList{
					Unclassified: []int{1, 5, 6, 5, 9, 1, 5, 4, 4, 4, -15, 0, 0, 5},
				},
			},
			wantList: []int{-15, 0, 1, 4, 5, 6, 9, 5, 1, 5, 4, 4, 0, 5},
		}, {
			name: "Test 04 null list",
			args: args{
				list: &dto.ClassifiedList{
					Unclassified: []int{},
				},
			},
			wantList: []int{},
		}, {
			name: "Test 05 limit exceeded ",
			args: args{
				list: &dto.ClassifiedList{
					Unclassified: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 101},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &service.ServiceStruct{
				ReplaceManager: tt.fields.ReplaceManager,
				Utilities:      tt.fields.Utilities,
			}
			if err := sm.OrderList(tt.args.list); (err != nil) != tt.wantErr {
				t.Errorf("OrderList() error = %v, wantErr %v", err, tt.wantErr)
			}
			equal := func() bool {
				for i, v := range tt.args.list.Classified {
					if v != tt.wantList[i] {
						return false
					}
				}
				return true
			}()
			if !equal {
				t.Errorf("List Classified = %v is not equal to  wantList %v", tt.args.list.Classified, tt.wantList)
			}
		})
	}
}
