package usecase

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/02amanag/upload-and-download-file/internal/model"
	"github.com/dgrijalva/jwt-go"
)

//ExtractToken ...
func (u *UsecaseStruct) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

//VerifyToken ...
func (u *UsecaseStruct) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := u.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil //secret key
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (u *UsecaseStruct) ExtractTokenMetadata(r *http.Request) (*model.AccessDetails, error) {
	token, err := u.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userid, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		profileid, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["profile_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &model.AccessDetails{
			UserID:    userid,
			ProfileId: profileid,
		}, nil
	}
	return nil, err
}

func (u *UsecaseStruct) CreateToken(userID int64, Profile_id int) (*model.TokenDetails, error) {

	td := &model.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 120).Unix()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["profile_id"] = Profile_id
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}
