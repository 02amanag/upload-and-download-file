//this file include the structures with member variable to be used over project with json naming too
package model

type Example struct {
	var_string string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	UserID    int64
	ProfileId int64
}
