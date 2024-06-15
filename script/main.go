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
	id1 string
	id2 string
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

	// fmt.Println(len(users), len(timeses))

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

	connection := make(map[mapKeys]int)
	for _, root := range timeses {
		if timesUser := timesNameToUserName[root.GetName()]; timesUser != "" {
			que := []traq.Channel{root}
			for len(que) > 0 {
				channel := que[0]
				que = que[1:]

				stats, r, err := client.ChannelApi.GetChannelStats(auth, channel.GetId()).Execute()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannelStats``: %v\n", err)
					fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
				}

				for _, userStat := range stats.GetUsers() {
					if userName, ok := userIdToUserName[userStat.GetId()]; ok && timesUser != userName {
						connection[mapKeys{timesUser, userName}] += int(userStat.GetMessageCount())
						connection[mapKeys{userName, timesUser}] += int(userStat.GetMessageCount())
						// fmt.Printf("%s での %s の発言: %d\n", channel.GetName(), userName, userStat.GetMessageCount())
					}
				}

				for _, childId := range channel.GetChildren() {
					child, r, err := client.ChannelApi.GetChannel(auth, childId).Execute()
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannel``: %v\n", err)
						fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
					}
					que = append(que, *child)
				}
			}
			// fmt.Println()
		}
	}

	fmt.Println("TRUNCATE TABLE `user_connections`;")
	fmt.Println("TRUNCATE TABLE `user_groups`;")
	fmt.Println("TRUNCATE TABLE `users`;")
	fmt.Println()

	fmt.Println("INSERT INTO `users` VALUES")
	for i, usr := range users {
		suffix := ","
		if i+1 == len(users) {
			suffix = ";"
		}
		fmt.Printf("('%s', '%s')%s\n", usr.GetName(), usr.GetDisplayName(), suffix)
	}
	fmt.Println()

	fmt.Println("INSERT INTO `user_groups` VALUES")
	for i, usr := range users {
		userDetail, r, err := client.UserApi.GetUser(auth, usr.GetId()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		for j, groupId := range userDetail.GetGroups() {
			group, r, err := client.GroupApi.GetUserGroup(auth, groupId).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when calling `GroupApi.GetUserGroup``: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			}
			suffix := ","
			if i+1 == len(users) && j+1 == len(userDetail.GetGroups()) {
				suffix = ";"
			}
			fmt.Printf("('%s', '%s')%s\n", userDetail.GetName(), group.GetName(), suffix)
		}
	}
	fmt.Println()

	fmt.Println("INSERT INTO `user_connections` VALUES")
	remain := len(connection)
	// fmt.Println(connection)
	for _, id1 := range users {
		for _, id2 := range users {
			if connect, ok := connection[mapKeys{id1.GetName(), id2.GetName()}]; ok {
				remain--
				suffix := ","
				if remain == 0 {
					suffix = ";"
				}
				fmt.Printf("('%s', '%s', '%d')%s\n", id1.GetName(), id2.GetName(), connect, suffix)
			}
		}
	}
}
