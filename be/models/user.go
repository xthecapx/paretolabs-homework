package models

type User struct {
	FID         int    `json:"fid"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	PFP         struct {
		URL      string `json:"url"`
		Verified bool   `json:"verified"`
	} `json:"pfp"`
	Profile struct {
		Bio struct {
			Text     string   `json:"text"`
			Mentions []string `json:"mentions"`
		} `json:"bio"`
		Location struct {
			PlaceID     string `json:"placeId"`
			Description string `json:"description"`
		} `json:"location"`
	} `json:"profile"`
	FollowerCount    int    `json:"followerCount"`
	FollowingCount   int    `json:"followingCount"`
	ReferrerUsername string `json:"referrerUsername"`
	ViewerContext    struct {
		Following          bool `json:"following"`
		FollowedBy         bool `json:"followedBy"`
		CanSendDirectCasts bool `json:"canSendDirectCasts"`
	} `json:"viewerContext"`
}
