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
	UserA string
	UserB string
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

	timesIdToTimesName := make(map[string]string)
	for _, v := range timeses {
		timesIdToTimesName[v.GetId()] = v.GetName()
	}

	timesNameToUserName := make(map[string]string)
	for _, channel := range timeses {
		timesNameToUserName[channel.GetName()] = ""
	}
	for _, usr := range users {
		if _, ok := timesNameToUserName[usr.GetName()]; ok {
			timesNameToUserName[usr.GetName()] = usr.GetName()
		} else {
			userDetail, r, err := client.UserApi.GetUser(auth, usr.GetId()).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			}
			if home := userDetail.GetHomeChannel(); home != "" {
				timesNameToUserName[timesIdToTimesName[home]] = usr.GetName()
				// fmt.Println(userDetail.GetName(), timesIdToTimesName[home])
			} else {
				// fmt.Println(usr.GetName())
			}
		}
	}
	cnt := 0
	for _, channel := range timeses {
		if timesNameToUserName[channel.GetName()] == "" && !channel.GetArchived() {
			fmt.Println(channel.GetName())
			cnt++
		}
	}
	fmt.Println(cnt)

	m := make(map[mapKeys]int)
	for _, channel := range timeses {
		if timesNameToUserName[channel.GetName()] != "" {
			stats, r, err := client.ChannelApi.GetChannelStats(auth, channel.GetId()).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannelStats``: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			}

			for _, userStat := range stats.Users {
				if userName, ok := userIdToUserName[userStat.GetId()]; ok {
					m[mapKeys{timesNameToUserName[channel.GetName()], userName}] += int(userStat.GetMessageCount())
					m[mapKeys{userName, timesNameToUserName[channel.GetName()]}] += int(userStat.GetMessageCount())
					fmt.Printf("%s のtimesでの %s の発言: %d\n", timesNameToUserName[channel.GetName()], userName, userStat.GetMessageCount())
				}
			}
		}
	}

}
