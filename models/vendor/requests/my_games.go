package vendorRequests

import (
	"errors"

	"github.com/labstack/echo"
)

//MyGames represents a request to obtain a VendorGameList for the vendor.
type MyGames struct {
	Type         string `json:"Type" xml:"Type" form:"Type"`
	API_Token    string `json:"API_Token" xml:"API_Token" form:"API_Token"`
	SessionToken string `json:"SessionToken" xml:"SessionToken" form:"SessionToken"`
}

//FromForm creates a valid struct based on form data submitted, or returns error.
//
// Does not check for the validity of the items inside the struct (e.g. tokens)
func (receiver *MyGames) FromForm(c echo.Context) error {
	var err error
	receiver.Type = c.FormValue("Type")
	receiver.API_Token = c.FormValue("API_Token")
	receiver.SessionToken = c.FormValue("SessionToken")
	if receiver.Type != "MyGames" || receiver.API_Token == "" || receiver.SessionToken == "" {
		err = errors.New("Invalid Form Submitted")
	}

	return err
}
