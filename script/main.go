package main

import (
	"context"
	"fmt"
	"os"

	traq "github.com/traPtitech/go-traq"
)

var TOKEN = os.Getenv("TOKEN")

// map の key
type mapKeys struct {
	UserName    string
	ChannelName string
}

func main() {
	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, TOKEN)

	// v, _, _ := client.ChannelApi.
	// 	GetChannels(auth).
	// 	IncludeDm(true).
	// 	Execute()
	// fmt.Printf("%#v", v)

	usersIncludeBot, r, err := client.UserApi.GetUsers(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// responce from `GetUsers`
	// fmt.Printf("%v\n", usersIncludeBot)

	// #gps/times の id
	timesId := "8ed62c7d-3f4b-41c8-a446-29edeebc36c3"

	times, r, err := client.ChannelApi.GetChannel(auth, timesId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	// fmt.Println(times)

	timesesId := times.GetChildren()
	var timeses []traq.Channel
	for _, v := range timesesId {
		channel, r, err := client.ChannelApi.GetChannel(auth, v).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannel``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		if !channel.GetArchived() {
			timeses = append(timeses, *channel)
		}
	}

	var users []traq.User
	for _, v := range usersIncludeBot {
		if !v.GetBot() {
			users = append(users, v)
		}
	}

	fmt.Println(len(users), len(timeses))

	userIdToUserName := make(map[string]string)
	for _, v := range users {
		userIdToUserName[v.GetId()] = v.GetName()
	}

	m := make(map[mapKeys]int)
	for _, channel := range timeses {
		stats, r, err := client.ChannelApi.GetChannelStats(auth, channel.GetId()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannelStats``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}

		for _, userStat := range stats.Users {
			if userName, ok := userIdToUserName[userStat.GetId()]; ok {
				m[mapKeys{UserName: userName, ChannelName: channel.Name}] += int(userStat.GetMessageCount())
				fmt.Printf("#gps/times/%s での %s の発言: %d\n", channel.Name, userName, userStat.GetMessageCount())
			}
		}
	}

	fmt.Println(m[mapKeys{UserName: "Series_205", ChannelName: "Series_205"}])
}
