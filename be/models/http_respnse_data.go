package models

type HttpResponseData struct {
	Result struct {
		User              User          `json:"user"`
		Inviter           User          `json:"inviter"`
		InviterIsReferrer bool          `json:"inviterIsReferrer"`
		CollectionsOwned  []interface{} `json:"collectionsOwned"`
	} `json:"result"`
}
