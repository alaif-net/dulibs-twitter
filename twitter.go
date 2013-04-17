package dulibs

import (
	"encoding/json"
	"github.com/mrjones/oauth"
	"io/ioutil"
)

type Twitter struct {
	consumer    *oauth.Consumer
	accessToken *oauth.AccessToken
}

//Initializer
func NewTwitter(consumerKey, consumerSecret string) *Twitter {
	twitter := new(Twitter)
	twitter.consumer = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	return twitter
}

/* EXPORT FUNCS */
//set accesstoken
func (tw *Twitter) SetAccessToken(accessToken, accessSecret string) {
	tw.accessToken = &oauth.AccessToken{accessToken, accessSecret}
}

//Post status
func (tw *Twitter) PostStatus(message string) (interface{}, error) {
	res, err := tw.post(
		"https://api.twitter.com/1.1/statuses/update.json",
		map[string]string{"status": message})

	return res, err
}

//Post DM
func (tw *Twitter) PostDm(name, text string) (interface{}, error) {
	res, err := tw.post(
		"https://api.twitter.com/1.1/direct_messages/new.json",
		map[string]string{"screen_name": name, "text": text})

	return res, err
}

/* LOCAL FUNCS */
//post method
func (tw *Twitter) post(url string, params map[string]string) (interface{}, error) {
	res, err := tw.consumer.Post(url, params, tw.accessToken)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(body, &result)
	return result, err
}
