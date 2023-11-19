package model

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type CartTask struct {
	CartID    uuid.UUID `json:"cart_id"`
	Timestamp int64     `json:"timestamp"`
}

func (task *CartTask) MarshalBinary() ([]byte, error) {
	return json.Marshal(task)
}

func (task *CartTask) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &task); err != nil {
		return fmt.Errorf("cannot unmarshal binary to cart task: %w", err)
	}

	return nil
}
