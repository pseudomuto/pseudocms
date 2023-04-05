package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Model is a base model that can be embedded in others to include common fields
// like ID, timestamps, etc.
type Model struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

// GetID satisfies Identifiable for all models that embed this struct.
func (m Model) GetID() uuid.UUID { return m.ID }
