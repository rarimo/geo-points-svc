package referralid

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		nullifier string
		index     uint64
		want      string
	}{
		{
			name:      "Valid nullifier with index 0",
			nullifier: "2184ae1f990d26aa5bb84d54dc945ac3cce569cd828269802f0fa5c5c28f30a7",
			want:      "6xM70VgX4eh",
		},
		{
			name:      "Valid nullifier with index 1",
			nullifier: "2184ae1f990d26aa5bb84d54dc945ac3cce569cd828269802f0fa5c5c28f30a7",
			index:     1,
			want:      "eLHv3hj5txB",
		},
		{
			name:      "Valid nullifier with index 258",
			nullifier: "2184ae1f990d26aa5bb84d54dc945ac3cce569cd828269802f0fa5c5c28f30a7",
			index:     258,
			want:      "1hhJaHQB13G",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.nullifier, tt.index); got != tt.want {
				t.Errorf("New() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestNewMany(t *testing.T) {
	tests := []struct {
		name      string
		nullifier string
		index     uint64
		count     uint64
		want      []string
	}{
		{
			name:      "Valid nullifier for basic balance creation",
			nullifier: "2184ae1f990d26aa5bb84d54dc945ac3cce569cd828269802f0fa5c5c28f30a7",
			count:     5,
			want:      []string{"6xM70VgX4eh", "eLHv3hj5txB", "8Mu12YhyDVQ", "4l3LwW9p77V", "bLnCgkUOPWT"},
		},
		{
			name:      "Valid nullifier for start from non-zero index",
			nullifier: "2184ae1f990d26aa5bb84d54dc945ac3cce569cd828269802f0fa5c5c28f30a7",
			index:     2,
			count:     3,
			want:      []string{"8Mu12YhyDVQ", "4l3LwW9p77V", "bLnCgkUOPWT"},
		},
		{
			name:      "Valid nullifier, no count",
			nullifier: "2184ae1f990d26aa5bb84d54dc945ac3cce569cd828269802f0fa5c5c28f30a7",
			index:     8,
		},
	}

	equal := func(a, b []string) bool {
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMany(tt.nullifier, tt.count, tt.index)
			if !equal(got, tt.want) {
				t.Errorf("NewMany() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestDailyQuestions(t *testing.T) {
	tests := []struct {
		name       string
		operation  string
		want       bool
		loc        *time.Location
		timeNow    time.Time
		startAt    time.Time
		newStartAt time.Time
	}{
		{
			name:      "Create question past",
			operation: "create",
			want:      false,
			loc:       time.UTC,
			timeNow:   time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			startAt:   time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "Create question today",
			operation: "create",
			want:      true,
			loc:       time.UTC,
			timeNow:   time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			startAt:   time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "Create question future",
			operation: "create",
			want:      true,
			loc:       time.UTC,
			timeNow:   time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			startAt:   time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
		},

		//EDIT
		{
			name:       "Edit question from past to past",
			operation:  "edit",
			want:       false,
			loc:        time.UTC,
			timeNow:    time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			startAt:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			newStartAt: time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name:       "Edit question from past to feature",
			operation:  "edit",
			want:       false,
			loc:        time.UTC,
			timeNow:    time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			startAt:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			newStartAt: time.Date(2025, time.January, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			name:       "Edit question from past to today",
			operation:  "edit",
			want:       false,
			loc:        time.UTC,
			timeNow:    time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			startAt:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			newStartAt: time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			name:       "Edit question from future to past",
			operation:  "edit",
			want:       false,
			loc:        time.UTC,
			timeNow:    time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			startAt:    time.Date(2024, time.January, 20, 0, 0, 0, 0, time.UTC),
			newStartAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:       "Edit question from future to today",
			operation:  "edit",
			want:       true,
			loc:        time.UTC,
			timeNow:    time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			startAt:    time.Date(2024, time.January, 20, 0, 0, 0, 0, time.UTC),
			newStartAt: time.Date(2025, time.January, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			name:       "Edit question from future to future",
			operation:  "edit",
			want:       true,
			loc:        time.UTC,
			timeNow:    time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			startAt:    time.Date(2024, time.January, 20, 0, 0, 0, 0, time.UTC),
			newStartAt: time.Date(2025, time.January, 30, 0, 0, 0, 0, time.UTC),
		},
		{
			name:       "Edit question from today to future",
			operation:  "edit",
			want:       false,
			loc:        time.UTC,
			timeNow:    time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			startAt:    time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			newStartAt: time.Date(2024, time.January, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			name:       "Edit question from today to past",
			operation:  "edit",
			want:       false,
			loc:        time.UTC,
			timeNow:    time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			startAt:    time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			newStartAt: time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name:       "Edit question from today to today",
			operation:  "edit",
			want:       false,
			loc:        time.UTC,
			timeNow:    time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			startAt:    time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
			newStartAt: time.Date(2024, time.January, 10, 0, 0, 0, 0, time.UTC),
		},

		//DELETE
		{
			name:      "Delete question in a past",
			operation: "delete",
			want:      false,
			loc:       time.UTC,
			timeNow:   time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			startAt:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "Delete question in a future",
			operation: "delete",
			want:      true,
			loc:       time.UTC,
			timeNow:   time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			startAt:   time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "Delete question today",
			operation: "delete",
			want:      true,
			loc:       time.UTC,
			timeNow:   time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			startAt:   time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.operation {
			case "edit":
				var condition bool
				got := CheckOpportunityChange(tt.timeNow, tt.startAt, tt.loc)
				if got == false {
					condition = false
				}
				got = CheckOpportunityChange(tt.timeNow, tt.newStartAt, tt.loc)
				if got == tt.want {
					condition = false
				}
				if tt.want != condition {
					t.Errorf("CREATE CheckOpportunityChange() = %v, want %v", condition, tt.want)
				}
			case "create":
				got := CheckOpportunityChange(tt.timeNow, tt.startAt, tt.loc)
				if got != tt.want {
					t.Errorf("CREATE CheckOpportunityChange() = %v, want %v", got, tt.want)
				}
			case "delete":
				got := CheckOpportunityChange(tt.timeNow, tt.startAt, tt.loc)
				if got != tt.want {
					t.Errorf("DELETE CheckOpportunityChange() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
