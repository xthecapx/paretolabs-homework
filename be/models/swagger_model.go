package models

// PaginatedUsersResponse represents the response for GetUsersHandler.
type PaginatedUsersResponse struct {
	UserList   []User `json:"userList"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	TotalCount int    `json:"total_count"`
}

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

type UserProfileResponse struct {
	ProfileData UserProfile `json:"profileData"`
}
