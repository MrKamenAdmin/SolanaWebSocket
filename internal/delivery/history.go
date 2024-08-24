package delivery

import "time"

type History struct {
	CaptureDate time.Time `json:"capture_date"`
	Stake       uint64    `json:"stake"`
}
