package model

/*
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-04-27 08:31
 */

type AccessToken struct {
	AccessToken string      `json:"access_token"`
	TokenType   string      `json:"type_type"`
	ExpiresIn   int64       `json:"expires_in"`
	Data        interface{} `json:"data"`
}
