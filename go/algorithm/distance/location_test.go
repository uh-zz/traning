package distance

import (
	"reflect"
	"testing"
)

func TestGeoLocation(t *testing.T) {
	type args struct {
		keyword string
	}
	tests := []struct {
		name string
		args args
		want []Location
	}{
		{
			name: "test",
			args: args{
				keyword: "東京",
			},
			want: []Location{
				{
					City:     "東京都",
					CityKana: "トウキョウト",
					Town:     "千代田区",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GeoLocation(tt.args.keyword); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeoLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
