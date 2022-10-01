package distance

import "testing"

func TestDistanceOnTheEarth(t *testing.T) {
	type args struct {
		from Coordinate
		to   Coordinate
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test",
			args: args{
				from: Coordinate{
					Longitude: 139.767052,
					Latitude:  35.681167,
				},
				to: Coordinate{
					Longitude: 139.767052,
					Latitude:  35.681167,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistanceOnTheEarth(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("DistanceOnTheEarth() = %v, want %v", got, tt.want)
			}
		})
	}
}
