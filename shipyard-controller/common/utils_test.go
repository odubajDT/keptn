package common

import (
	"github.com/benbjohnson/clock"
	"github.com/keptn/go-utils/pkg/common/timeutils"
	"reflect"
	"testing"
	"time"
)

func TestParseTimestamp(t *testing.T) {
	correctISO8601Timestamp := "2020-01-02T15:04:05.000Z"

	timeObj, _ := time.Parse(timeutils.KeptnTimeFormatISO8601, correctISO8601Timestamp)

	mockClock := clock.NewMock()

	type args struct {
		ts       string
		theClock clock.Clock
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "compatible timestamp provided",
			args: args{
				ts:       correctISO8601Timestamp,
				theClock: nil,
			},
			want: timeObj,
		},
		{
			name: "incompatible timestamp provided - return now",
			args: args{
				ts:       "invalid",
				theClock: mockClock,
			},
			want: mockClock.Now().UTC(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseTimestamp(tt.args.ts, tt.args.theClock); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
