package main

import (
	"context"
	"fmt"
	"os"

	"github.com/samber/lo"
	"github.com/schollz/progressbar/v3"
	"github.com/traPtitech/go-traq"
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

	var users []*traq.UserDetail
	{
		usersIncludingBot, r, err := client.UserApi.GetUsers(auth).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUsers``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		usersExcludingBot := lo.Filter(usersIncludingBot, func(u traq.User, _ int) bool { return !u.Bot })

		bar := progressbar.Default(int64(len(usersExcludingBot)), "Getting user details")
		users = lo.Map(usersExcludingBot, func(u traq.User, _ int) *traq.UserDetail {
			ud, _ := lo.Must2(client.UserApi.GetUser(auth, u.Id).Execute())
			lo.Must0(bar.Add(1))
			return ud
		})
	}

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
	bar := progressbar.Default(int64(len(timesesId)), "Getting channel details")
	for _, v := range timesesId {
		channel, r, err := client.ChannelApi.GetChannel(auth, v).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ChannelApi.GetChannel``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}
		if !channel.GetArchived() {
			timeses = append(timeses, *channel)
		}
		lo.Must0(bar.Add(1))
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
			if home := usr.GetHomeChannel(); home != "" {
				timesNameToUserName[timesIdToTimesName[home]] = usr.GetName()
				// fmt.Println(userDetail.GetName(), timesIdToTimesName[home])
			} else {
				// fmt.Println(usr.GetName())
			}
		}
	}

	connection := make(map[mapKeys]int)
	bar = progressbar.Default(int64(len(timeses)), "Getting channel stats")
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
		lo.Must0(bar.Add(1))
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

	groups, _ := lo.Must2(client.GroupApi.GetUserGroups(auth).Execute())
	groupIDToInfo := make(map[string]*traq.UserGroup, len(groups))
	for _, g := range groups {
		groupIDToInfo[g.Id] = &g
	}

	fmt.Println("INSERT INTO `user_groups` VALUES")
	for i, usr := range users {
		for j, groupId := range usr.GetGroups() {
			group := groupIDToInfo[groupId]
			suffix := ","
			if i+1 == len(users) && j+1 == len(usr.GetGroups()) {
				suffix = ";"
			}
			fmt.Printf("('%s', '%s')%s\n", usr.GetName(), group.GetName(), suffix)
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
