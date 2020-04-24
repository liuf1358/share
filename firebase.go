package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"firebase.google.com/go/messaging"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var appIns *firebase.App

func init() {
	// Initialize another app with a different config
	opt := option.WithCredentialsFile("conf/firebase-264401-firebase-adminsdk-2zwtk-6c6c28d1e3.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	appIns = app
}

func sendFcm() {
	if appIns == nil {
		log.Panicln("failed appIns is nil")
		return
	}

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := appIns.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "dSPJ8Mcc1-k:APA91bHlZHT67tdj_BTiL43BKT6_87XxligY4m4TqUNQT46PaZLruBQHUjS8Ic-xpFkIetHOFaNAqwhDZl-k_5QHoNY8PAud6tVhCYCI-UMwLoXpGyoirr1XQXH-vvo8ffMXvVjjdyud"

	t := time.Now()
	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title:    "I'm title",
			Body:     t.String(),
			ImageURL: "",
		},
		Data: map[string]string{
			"score":  "850",
			"time":   "1222:52",
			"skipto": "12221213",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}

// verificationAuthToken
// 实际上只是解密token
// 解密后 {"auth_time":1578550741,"iss":"https://securetoken.google.com/firebase-264401","aud":"firebase-264401","exp":1578554398,
// "iat":1578550798,"sub":"7BCRxtCHLYgIXEey9Ex3yapZzkK2","uid":"7BCRxtCHLYgIXEey9Ex3yapZzkK2","firebase":{"sign_in_provider":"google.com",
// "tenant":"","identities":{"email":["18322595861@163.com"],"google.com":["101397198388277826613"]}}}
func verificationAuthToken() {

	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjUxMjRjY2JhZDVkNWZiZjNiYTJhOGI1ZWE3MTE4NDVmOGNiMjZhMzYiLCJ0eXAiOiJKV1QifQ." +
		"eyJuYW1lIjoibWVuZ2ZlaXEgcWluIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS8tMXdvYjUwdFdIM28vQUFB" +
		"QUFBQUFBQUkvQUFBQUFBQUFBQUEvQUNIaTNyZjFjREluZjhHZ2RMZWQzZUs1VzlvZFFVWDRTQS9zOTYtYy9waG90by5qcGciLCJpc3MiOiJodHRwczovL3Nl" +
		"Y3VyZXRva2VuLmdvb2dsZS5jb20vZmlyZWJhc2UtMjY0NDAxIiwiYXVkIjoiZmlyZWJhc2UtMjY0NDAxIiwiYXV0aF90aW1lIjoxNTc4NTUwNzQxLCJ1c2Vy" +
		"X2lkIjoiN0JDUnh0Q0hMWWdJWEVleTlFeDN5YXBaemtLMiIsInN1YiI6IjdCQ1J4dENITFlnSVhFZXk5RXgzeWFwWnprSzIiLCJpYXQiOjE1Nzg1NTA3OTgs" +
		"ImV4cCI6MTU3ODU1NDM5OCwiZW1haWwiOiIxODMyMjU5NTg2MUAxNjMuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRp" +
		"ZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMDEzOTcxOTgzODgyNzc4MjY2MTMiXSwiZW1haWwiOlsiMTgzMjI1OTU4NjFAMTYzLmNvbSJdfSwic2lnbl9pbl9wcm92" +
		"aWRlciI6Imdvb2dsZS5jb20ifX0.V5-Yozho5TnKyHOFe10imtiZkrhwZL2eRxeemLluDLEnZr8UsUbY9sRUcTIxGV2BBgbkWJ-XG-WL_oe8IeeW4L0bOWke" +
		"lNwMaWsL6S-7RGMNzUrWXsMEhMbXa2ypFY-PiEyRw2aAZXkaUrCKGIztZpgAJimaFAzpw12ESpsSLpB-CKTSzguJUg3Ld2bplV9k7HGSn6Sc7QzWxsWpYOFm" +
		"i5vayc02EKSzNw2ZlV7YRSEEgPJvLBzDVcV1vxbhbqq4DvIJ7feVJfoqkvBQkWtUoUx67g7MhM33JhFMuVZ-frDLeku0tEXQdS-uozC2Husl-" +
		"BzQ0JgqSLeGAEuYAwbXeg"

	ctx := context.Background()
	client, err := appIns.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}
	ss, _ := json.Marshal(token)
	log.Printf("Verified ID token: %v\n", string(ss))
}

// VerifyIDTokenAndCheckRevoked
// 检查token是否被撤销

func VerifyIDTokenAndCheckRevoked() {
	idToken := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjUxMjRjY2JhZDVkNWZiZjNiYTJhOGI1ZWE3MTE4NDVmOGNiMjZhMzYiLCJ0eXAiOiJKV1QifQ." +
		"eyJuYW1lIjoibWVuZ2ZlaXEgcWluIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS8tMXdvYjUwdFdIM28vQUFB" +
		"QUFBQUFBQUkvQUFBQUFBQUFBQUEvQUNIaTNyZjFjREluZjhHZ2RMZWQzZUs1VzlvZFFVWDRTQS9zOTYtYy9waG90by5qcGciLCJpc3MiOiJodHRwczovL3Nl" +
		"Y3VyZXRva2VuLmdvb2dsZS5jb20vZmlyZWJhc2UtMjY0NDAxIiwiYXVkIjoiZmlyZWJhc2UtMjY0NDAxIiwiYXV0aF90aW1lIjoxNTc4NTUwNzQxLCJ1c2Vy" +
		"X2lkIjoiN0JDUnh0Q0hMWWdJWEVleTlFeDN5YXBaemtLMiIsInN1YiI6IjdCQ1J4dENITFlnSVhFZXk5RXgzeWFwWnprSzIiLCJpYXQiOjE1Nzg1NTA3OTgs" +
		"ImV4cCI6MTU3ODU1NDM5OCwiZW1haWwiOiIxODMyMjU5NTg2MUAxNjMuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRp" +
		"ZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMDEzOTcxOTgzODgyNzc4MjY2MTMiXSwiZW1haWwiOlsiMTgzMjI1OTU4NjFAMTYzLmNvbSJdfSwic2lnbl9pbl9wcm92" +
		"aWRlciI6Imdvb2dsZS5jb20ifX0.V5-Yozho5TnKyHOFe10imtiZkrhwZL2eRxeemLluDLEnZr8UsUbY9sRUcTIxGV2BBgbkWJ-XG-WL_oe8IeeW4L0bOWke" +
		"lNwMaWsL6S-7RGMNzUrWXsMEhMbXa2ypFY-PiEyRw2aAZXkaUrCKGIztZpgAJimaFAzpw12ESpsSLpB-CKTSzguJUg3Ld2bplV9k7HGSn6Sc7QzWxsWpYOFm" +
		"i5vayc02EKSzNw2ZlV7YRSEEgPJvLBzDVcV1vxbhbqq4DvIJ7feVJfoqkvBQkWtUoUx67g7MhM33JhFMuVZ-frDLeku0tEXQdS-uozC2Husl-" +
		"BzQ0JgqSLeGAEuYAwbXeg"

	ctx := context.Background()
	client, err := appIns.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	token, err := client.VerifyIDTokenAndCheckRevoked(ctx, idToken)
	if err != nil {
		log.Printf("VerifyIDTokenAndCheckRevoked err: %v\n", err)

		if err.Error() == "ID token has been revoked" {
			// Token is revoked. Inform the user to reauthenticate or signOut() the user.
		} else {
			// Token is invalid
		}
	}
	//ss, _ := json.Marshal(token)
	//log.Printf("Verified ID token: %v\n", string(ss))
	log.Printf("Verified ID token: %#v\n", token)
}

// 撤销令牌
func RevokeRefreshTokens() {
	uid := "ZDILymPJEGafV8jQYWMtbz7RyfB3"
	ctx := context.Background()
	client, err := appIns.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	if err := client.RevokeRefreshTokens(ctx, uid); err != nil {
		log.Fatalf("error revoking tokens for user: %v, %v\n", uid, err)
	}
	// accessing the user's TokenValidAfter
	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	ss, _ := json.Marshal(u)
	log.Printf("[the user]:%v", string(ss))

	timestamp := u.TokensValidAfterMillis / 1000
	log.Printf("the refresh tokens were revoked at: %d (UTC seconds) ", timestamp)
}

func getUserInfoById() {
	uid := "ZDILymPJEGafV8jQYWMtbz7RyfB3"
	ctx := context.Background()
	client, err := appIns.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	// accessing the user's TokenValidAfter
	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	ss, _ := json.Marshal(u)
	log.Printf("[the user]:%v", string(ss))
}

func main() {
	//getUserInfoById()
	RevokeRefreshTokens()
	//VerifyIDTokenAndCheckRevoked()
	//verificationAuthToken()
	//sendFcm()
}
