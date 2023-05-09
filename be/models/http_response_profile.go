package models

type ProfileResponse struct {
	Body             ProfileBody `json:"body"`
	ConnectedAddress string      `json:"connectedAddress"`
}

type ProfileBody struct {
	ID               int64  `json:"id"`
	Address          string `json:"address"`
	Username         string `json:"username"`
	DisplayName      string `json:"displayName"`
	Bio              string `json:"bio"`
	Followers        int64  `json:"followers"`
	Following        int64  `json:"following"`
	AvatarURL        string `json:"avatarUrl"`
	IsVerifiedAvatar bool   `json:"isVerifiedAvatar"`
	RegisteredAt     int64  `json:"registeredAt"`
}
