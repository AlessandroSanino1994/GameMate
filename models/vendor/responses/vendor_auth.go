package vendorResponses

//VendorAuth represents the POSITIVE response to a registration or login request.
//
//For NEGATIVE response, please refer to errorResponses.ErrorResponse.
type VendorAuth struct {
	Type         string `json:"Type" xml:"Type" form:"Type"`
	SessionToken string `json:"SessionToken" xml:"SessionToken" form:"SessionToken"`
}

//FromToken fills the struct's data with proper definition, based on a session token.
func (receiver *VendorAuth) FromToken(SessionToken string) {
	receiver.Type = "Session Token"
	receiver.SessionToken = SessionToken
}
