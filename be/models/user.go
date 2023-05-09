package models

type User struct {
	FID              int           `json:"fid"`
	Username         string        `json:"username"`
	DisplayName      string        `json:"displayName"`
	PFP              PFP           `json:"pfp"`
	Profile          Profile       `json:"profile"`
	FollowerCount    int           `json:"followerCount"`
	FollowingCount   int           `json:"followingCount"`
	ReferrerUsername string        `json:"referrerUsername"`
	ViewerContext    ViewerContext `json:"viewerContext"`
}

type PFP struct {
	URL      string `json:"url"`
	Verified bool   `json:"verified"`
}

type Profile struct {
	Bio      Bio      `json:"bio"`
	Location Location `json:"location"`
}

type Bio struct {
	Text     string   `json:"text"`
	Mentions []string `json:"mentions"`
}

type Location struct {
	PlaceID     string `json:"placeId"`
	Description string `json:"description"`
}

type ViewerContext struct {
	Following          bool `json:"following"`
	FollowedBy         bool `json:"followedBy"`
	CanSendDirectCasts bool `json:"canSendDirectCasts"`
}

type UserProfile struct {
	UserInformation    User        `json:"userInformation"`
	Address            string      `json:"address"`
	ProfileInformation ProfileBody `json:"profileInformation"`
}
