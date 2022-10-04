package main

import (
	"testing"
	"time"
)

func TestEvent_HasIntersection(t *testing.T) {
	type fields struct {
		ID        int
		Name      string
		Organizer string
		Start     time.Time
		End       time.Time
	}
	type args struct {
		otherEvent Event
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "case 1",
			fields: fields{
				Start: time.Date(2022, 10, 4, 11, 0, 0, 0, time.UTC),
				End:   time.Date(2022, 10, 4, 14, 0, 0, 0, time.UTC),
			},
			args: args{
				otherEvent: Event{
					Start: time.Date(2022, 10, 4, 13, 0, 0, 0, time.UTC),
					End:   time.Date(2022, 10, 4, 15, 0, 0, 0, time.UTC),
				},
			},
			want: true,
		},
		{
			name: "case 2",
			fields: fields{
				Start: time.Date(2022, 10, 4, 13, 0, 0, 0, time.UTC),
				End:   time.Date(2022, 10, 4, 15, 0, 0, 0, time.UTC),
			},
			args: args{
				otherEvent: Event{
					Start: time.Date(2022, 10, 4, 11, 0, 0, 0, time.UTC),
					End:   time.Date(2022, 10, 4, 14, 0, 0, 0, time.UTC),
				},
			},
			want: true,
		},
		{
			name: "case 2",
			fields: fields{
				Start: time.Date(2022, 10, 4, 13, 0, 0, 0, time.UTC),
				End:   time.Date(2022, 10, 4, 15, 0, 0, 0, time.UTC),
			},
			args: args{
				otherEvent: Event{
					Start: time.Date(2022, 10, 4, 10, 0, 0, 0, time.UTC),
					End:   time.Date(2022, 10, 4, 12, 0, 0, 0, time.UTC),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Event{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Organizer: tt.fields.Organizer,
				Start:     tt.fields.Start,
				End:       tt.fields.End,
			}
			if got := e.HasIntersection(tt.args.otherEvent); got != tt.want {
				t.Errorf("HasIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
