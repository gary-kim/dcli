package cmds

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cobra"

	"gomod.garykim.dev/dcli/cmd"
)

func init() {
	getStatus := &cobra.Command{
		Use:   "get-status <guild-id> <username#discriminator>",
		Short: "Get users's status",
		RunE: func(command *cobra.Command, args []string) error {
			cmd.CheckArgs(2, 2, command, args)

			discord, err := discordgo.New(cmd.Token)
			if err != nil {
				fmt.Fprintf(os.Stderr, "could not connect to discord: %s\n", err)
				return err
			}
			discord.StateEnabled = true
			discord.Identify.Intents = discordgo.MakeIntent(
				discordgo.IntentsGuildPresences,
			)
			c := make(chan *discordgo.GuildMembersChunk)
			discord.AddHandler(getGuildMembersChunkHandler(c))
			err = discord.Open()
			if err != nil {
				return err
			}
			err = discord.RequestGuildMembers(args[0], strings.Split(args[1], "#")[0], 10, true)
			if err != nil {
				fmt.Fprintf(os.Stderr, "could not get guild: %s\n", err)
				return err
			}
			m := <- c

			// Find User ID
			userID := ""
			for _, user := range m.Members {
				if user.User.String() == args[1] {
					userID = user.User.ID
					break
				}
			}
			found := false
			for _, presence := range m.Presences {
				// Correct User
				if presence.User.ID == userID {
					for _, activity := range presence.Activities {
						if activity.Type == discordgo.ActivityTypeCustom {
							found = true
							tp, err := json.Marshal(activity)
							if err != nil {
								fmt.Fprintf(os.Stderr, "could not json marshal: %s\n", err)
								return err
							}
							fmt.Println(string(tp))
						}
					}
				}
			}
			if !found {
				fmt.Println("{}")
			}
			return nil
		},
	}

	cmd.Root.AddCommand(getStatus)
}

func getGuildMembersChunkHandler (c chan *discordgo.GuildMembersChunk) func(s *discordgo.Session, m *discordgo.GuildMembersChunk) {
	return func (s *discordgo.Session, m *discordgo.GuildMembersChunk) {
		c <- m
	}
}
