package service

import (
	"applicationDesignTest/internal/models"
	"applicationDesignTest/internal/repository/inMemory"
	"context"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	tests := []struct {
		name           string
		order          models.Order
		expectedResult bool
		expectedDays   map[time.Time]struct{}
	}{
		{
			name: "successful order",
			order: models.Order{
				HotelID:   "reddison",
				RoomID:    "lux",
				UserEmail: "test@example.com",
				From:      date(2024, 1, 1),
				To:        date(2024, 1, 3),
			},
			expectedResult: true,
			expectedDays:   nil,
		},
		{
			name: "unavailable date",
			order: models.Order{
				HotelID:   "reddison",
				RoomID:    "lux",
				UserEmail: "test@example.com",
				From:      date(2024, 1, 4),
				To:        date(2024, 1, 5),
			},
			expectedResult: false,
			expectedDays:   map[time.Time]struct{}{date(2024, 1, 5): {}},
		},
		{
			name: "partially unavailable date",
			order: models.Order{
				HotelID:   "reddison",
				RoomID:    "lux",
				UserEmail: "test@example.com",
				From:      date(2024, 1, 3),
				To:        date(2024, 1, 5),
			},
			expectedResult: false,
			expectedDays:   map[time.Time]struct{}{date(2024, 1, 5): {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			success, unavailableDays := NewOrderService(inMemory.NewOrderRepository(), inMemory.NewHotelRepository()).CreateOrder(ctx, tt.order)
			if success != tt.expectedResult {
				t.Errorf("expected result %v, got %v", tt.expectedResult, success)
			}
			if len(unavailableDays) != len(tt.expectedDays) {
				t.Errorf("expected unavailable days %v, got %v", tt.expectedDays, unavailableDays)
			}
			for day := range tt.expectedDays {
				if _, ok := unavailableDays[day]; !ok {
					t.Errorf("expected unavailable day %v", day)
				}
			}
		})
	}
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
