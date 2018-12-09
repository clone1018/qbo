package main

import (
	"context"
	"fmt"
	"github.com/clone1018/qbo"
	"golang.org/x/oauth2"
	"time"
)

const (
	accessToken  = ""
	refreshToken = ""
	realmID      = ""
)

type TokenSource struct {
	AccessToken  string
	RefreshToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		Expiry:       time.Now().Add(time.Duration(time.Hour)),
	}
	return token, nil
}

func main() {
	tokenSource := &TokenSource{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := qbo.NewClient(oauthClient, realmID)

	ctx := context.TODO()

	entries, _, err := client.Query.Query(ctx, "select * from JournalEntry")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", entries)
}
