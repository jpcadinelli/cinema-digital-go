package models

import "github.com/google/uuid"

type DropdownUUID struct {
	Label string    `json:"label"`
	Value uuid.UUID `json:"value"`
}
