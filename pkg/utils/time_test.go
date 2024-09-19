package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsInHourRange(t *testing.T) {
	type args struct {
		start    time.Time
		end      time.Time
		minHours uint
		maxHours uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is in valid range",
			args: args{
				start:    time.Now(),
				end:      time.Now().Add(3 * time.Hour),
				minHours: 2,
				maxHours: 4,
			},
			want: true,
		},
		{
			name: "is below range",
			args: args{
				start:    time.Now(),
				end:      time.Now().Add(1 * time.Hour),
				minHours: 2,
				maxHours: 4,
			},
			want: false,
		},
		{
			name: "is above range",
			args: args{
				start:    time.Now(),
				end:      time.Now().Add(5 * time.Hour),
				minHours: 2,
				maxHours: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsInHourRange(tt.args.start, tt.args.end, tt.args.minHours, tt.args.maxHours), "IsInHourRange(%v, %v, %v, %v)", tt.args.start, tt.args.end, tt.args.minHours, tt.args.maxHours)
		})
	}
}

func TestIsInTheFuture(t *testing.T) {
	type args struct {
		start time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is in the future",
			args: args{
				start: time.Now().Add(15 * time.Second),
			},
			want: true,
		},
		{
			name: "is in the past",
			args: args{
				start: time.Now().Add(-15 * time.Second),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsInTheFuture(tt.args.start), "IsInTheFuture(%v)", tt.args.start)
		})
	}
}

func TestParseAndValidateTime(t *testing.T) {
	st := time.Now().Add(15 * time.Second).Round(time.Second)
	et := time.Now().Add(15 * time.Second).Add(2 * time.Hour).Round(time.Second)

	type args struct {
		startEpoch int64
		endEpoch   int64
	}
	tests := []struct {
		name      string
		args      args
		wantStart *time.Time
		wantEnd   *time.Time
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "is valid time",
			args: args{
				startEpoch: st.Unix(),
				endEpoch:   et.Unix(),
			},
			wantStart: &st,
			wantEnd:   &et,
			wantErr:   assert.NoError,
		},
		{
			name: "is invalid start time",
			args: args{
				startEpoch: time.Now().Unix(),
				endEpoch:   et.Unix(),
			},
			wantErr: assert.Error,
		},
		{
			name: "is invalid end time",
			args: args{
				startEpoch: time.Now().Unix(),
				endEpoch:   time.Now().Add(-15 * time.Second).Unix(),
			},
			wantErr: assert.Error,
		},
		{
			name: "duration is too short",
			args: args{
				startEpoch: st.Unix(),
				endEpoch:   et.Add(-1 * time.Hour).Unix(),
			},
			wantErr: assert.Error,
		},
		{
			name: "duration is too long",
			args: args{
				startEpoch: st.Unix(),
				endEpoch:   et.Add(15 * time.Hour).Unix(),
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParseAndValidateTime(tt.args.startEpoch, tt.args.endEpoch)
			if !tt.wantErr(t, err, fmt.Sprintf("ParseAndValidateTime(%v, %v)", tt.args.startEpoch, tt.args.endEpoch)) {
				return
			}
			assert.Equalf(t, tt.wantStart, got, "ParseAndValidateTime(%v, %v)", tt.args.startEpoch, tt.args.endEpoch)
			assert.Equalf(t, tt.wantEnd, got1, "ParseAndValidateTime(%v, %v)", tt.args.startEpoch, tt.args.endEpoch)
		})
	}
}

func TestTimeIsBeforeEnd(t *testing.T) {
	type args struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true",
			args: args{
				start: time.Now(),
				end:   time.Now().Add(15 * time.Second),
			},
			want: true,
		},
		{
			name: "should return false",
			args: args{
				start: time.Now().Add(15 * time.Second),
				end:   time.Now(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, TimeIsBeforeEnd(tt.args.start, tt.args.end), "TimeIsBeforeEnd(%v, %v)", tt.args.start, tt.args.end)
		})
	}
}
