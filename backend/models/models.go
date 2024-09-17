package models

import (
    "time"

    "github.com/google/uuid"
    // "gorm.io/gorm"
)

type OrganizationType string

const (
    IE  OrganizationType = "IE"
    LLC OrganizationType = "LLC"
    JSC OrganizationType = "JSC"
)

type Employee  struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    Username  string    `gorm:"size:50;unique;not null" json:"username"`
    FirstName string    `gorm:"size:50" json:"first_name"`
    LastName  string    `gorm:"size:50" json:"last_name"`
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Organization struct {
    ID          uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    Name        string          `gorm:"size:100;not null" json:"name"`
    Description string          `json:"description"`
    Type        OrganizationType `gorm:"type:organization_type" json:"type"`
    CreatedAt   time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt   time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type OrganizationResponsible struct {
    ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    OrganizationID uuid.UUID `gorm:"type:uuid;not null;index" json:"organization_id"`
    UserID         uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
}

type Tender struct {
    ID             uuid.UUID `json:"id"`
    Name           string    `json:"name"`
    Description    string    `json:"description"`
    ServiceType    string    `json:"serviceType"`
    Status         string    `json:"status"`
    OrganizationID uuid.UUID `json:"organizationId"`
    Version        int       `json:"version"`
    CreatedAt      time.Time `json:"createdAt"`
    CreatorUsername      time.Time `json:"creatorUsername"`
}

type Bid struct {
    ID          uuid.UUID `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
    TenderID    uuid.UUID `json:"tenderId"`
    AuthorType  string    `json:"authorType"`
    Feedback    string    `json:"bidFeedback"`
    Decision    string    `json:"bidDecision"`
    AuthorID    uuid.UUID `json:"authorId"`
    Version     int       `json:"version"`
    CreatedAt   time.Time `json:"createdAt"`
}

type BidReview struct {
    ID          uuid.UUID `json:"id"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"createdAt"`
    BidID       uuid.UUID `json:"bidId"`
}