package model

type VisibilityType string

const (
	VISIBILITY_EVERYONE_READ VisibilityType = "read"
	VISIBILITY_EVERYONE_EDIT VisibilityType = "edit"
	VISIBILITY_PRIVATE       VisibilityType = "private"
	VISIBILITY_SUPER_ADMIN   VisibilityType = "super"
)
