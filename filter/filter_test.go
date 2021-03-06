package filter_test

import (
	"reflect"
	"testing"

	"github.com/bastianrob/go-experiences/filter"
)

func TestFilter(t *testing.T) {
	intptr := func(num int) *int {
		return &num
	}
	type args struct {
		arr     interface{}
		filterf interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    interface{}
	}{
		{"Success", args{
			arr: []int{1, 2, 3, 4},
			filterf: func(entry int) bool {
				return entry == 1
			}}, false, []int{1}},
		{"Success", args{
			arr: []*int{intptr(1), intptr(2), intptr(3), intptr(4)},
			filterf: func(entry *int) bool {
				return *entry == 1
			}}, false, []*int{intptr(1)}},
		{"Failed", args{
			arr:     "[]int{1, 2, 3, 4}",
			filterf: nil}, true, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := filter.Filter(tt.args.arr, tt.args.filterf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Filter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFilterFast(b *testing.B) {
	source := [100]int{}
	for i := 0; i < len(source); i++ {
		source[i] = i + 1
	}
	isMultipliedBy3 := func(num int) bool {
		return num%3 == 0
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		filter.Filter(source, isMultipliedBy3)
	}
}
