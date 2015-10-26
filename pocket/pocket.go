package pocket

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/browser"
)

func responseBodyAsValues(r *http.Response) (url.Values, error) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return url.Values{}, err
	}

	return url.ParseQuery(string(body))
}

// GetPocketRequestToken will get the first request token, kicks off the authentication
// process.
func GetPocketRequestToken(apiKey *string, callbackUrl string) string {
	apikeyValue := *apiKey
	resp, err := http.PostForm(
		"https://getpocket.com/v3/oauth/request",
		url.Values{"consumer_key": {apikeyValue}, "redirect_uri": {callbackUrl}},
	)

	if err != nil {
		log.Fatalf("Error getting code from Pocket: %v", err)
	}
	values, err := responseBodyAsValues(resp)

	return values.Get("code")
}

func AuthorizePocket(code string, callbackUrl string) {
	browser.OpenURL("https://getpocket.com/auth/authorize?request_token=" + code + "&redirect_uri=" + callbackUrl)
}

func GetPocketAccessToken(apiKey *string, code string, callbackUrl string) (string, string) {
	apikeyValue := *apiKey
	resp, err := http.PostForm(
		"https://getpocket.com/v3/oauth/authorize",
		url.Values{"consumer_key": {apikeyValue}, "code": {code}},
	)

	if err != nil {
		log.Fatalf("Error getting code from Pocket: %v", err)
	}
	values, err := responseBodyAsValues(resp)

	return values.Get("username"), values.Get("access_token")
}

func AddItemToPocket(apiKey *string, access_token string, tweeturl string, tweet_id int64) {
	apikeyValue := *apiKey
	_, err := http.PostForm(
		"https://getpocket.com/v3/add",
		url.Values{"consumer_key": {apikeyValue}, "access_token": {access_token}, "url": {tweeturl}, "tweet_id": {strconv.FormatInt(tweet_id, 10)}},
	)

	if err != nil {
		log.Fatalf("Error getting code from Pocket: %v", err)
	}
	//responseBodyAsValues(resp)
}
