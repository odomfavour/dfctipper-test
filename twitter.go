package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func _main() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file")
		return
	}

	fmt.Println("Go-Twitter Bot v0.01")
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	client, err := getClient(&creds)
	if err != nil {
		log.Error("Error getting Twitter Client")
		log.Error(err)
		return
	}

	user, _, err := client.Users.Lookup(&twitter.UserLookupParams{
		ScreenName: []string{"deficonnect"},
	})

	log.Errorf("\n\nInformation of user\n%+v\n\n\n", user[0].ID)

	if err != nil {
		panic(err)
	}

	tweets, _, err := client.Statuses.Retweets(1466430599909027846, &twitter.StatusRetweetsParams{
		Count: 1,
	})

	if err != nil {
		log.Error(err)
		return
	}

	for _, t := range tweets {
		log.Error("%+v\n\n\n", t.User.ScreenName)
	}

	log.Error("%d", len(tweets))

}

// getClient is a helper function that will return a twitter client
// that we can subsequently use to send tweets, or to stream new tweets
// this will take in a pointer to a Credential struct which will contain
// everything needed to authenticate and return a pointer to a twitter Client
// or an error
func getClient(creds *Credentials) (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Errorf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}
