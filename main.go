package main

import (
	"bot_beefy/config"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	paginator "github.com/topi314/dgo-paginator"
)

func main() {
	c := config.GetConfig()

	dg, err := discordgo.New("Bot " + c.Discord.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	manager := paginator.NewManager()
	dg.AddHandler(manager.OnInteractionCreate)

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID || m.Content != "!usdc" {
			return
		}

		vaults := getVaults()
		usdVaults := filterVaults(vaults, "USD")

		pages := map[int]*discordgo.MessageEmbed{}

		for i, vault := range usdVaults {
			if i%10 == 0 {
				pages[i/10] = &discordgo.MessageEmbed{
					Title: "USD Vaults",
				}
			}

			tokens := ""
			for _, token := range vault.Assets {
				tokens += token + "\n"
			}

			if !vault.IsCLM {
				pages[i/10].Description += fmt.Sprintf("[%s](https://app.beefy.com/vault/%s) - **%.2f%%**\n", strings.Join(vault.Assets, "-"), vault.ID, vault.APY*100)
			} else {
				pages[i/10].Description += fmt.Sprintf("[%s](https://app.beefy.com/vault/%s) - CLM - **%.2f%%**\n", strings.Join(vault.Assets, "-"), vault.ID, vault.APY*100)
			}
		}

		if len(pages) == 0 {
			fmt.Println("No vaults found")
			return
		}

		if err = manager.CreateMessage(s, m.ChannelID, &paginator.Paginator{
			PageFunc: func(page int, embed *discordgo.MessageEmbed) {
				embed.Title = pages[page].Title
				embed.Description = pages[page].Description
			},
			MaxPages:        len(pages),
			Expiry:          time.Now(),
			ExpiryLastUsage: true,
		}); err != nil {
			fmt.Println(err)
		}
	})

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID || m.Content != "!btc" {
			return
		}

		vaults := getVaults()
		btcVaults := filterVaults(vaults, "BTC")

		pages := map[int]*discordgo.MessageEmbed{}

		for i, vault := range btcVaults {
			if i%10 == 0 {
				pages[i/10] = &discordgo.MessageEmbed{
					Title: "BTC Vaults",
				}
			}

			tokens := ""
			for _, token := range vault.Assets {
				tokens += token + "\n"
			}

			if !vault.IsCLM {
				pages[i/10].Description += fmt.Sprintf("[%s](https://app.beefy.com/vault/%s) - **%.2f%%**\n", strings.Join(vault.Assets, "-"), vault.ID, vault.APY*100)
			} else {
				pages[i/10].Description += fmt.Sprintf("[%s](https://app.beefy.com/vault/%s) - CLM - **%.2f%%**\n", strings.Join(vault.Assets, "-"), vault.ID, vault.APY*100)
			}
		}

		if err = manager.CreateMessage(s, m.ChannelID, &paginator.Paginator{
			PageFunc: func(page int, embed *discordgo.MessageEmbed) {
				embed.Title = pages[page].Title
				embed.Description = pages[page].Description
			},
			MaxPages:        len(pages),
			Expiry:          time.Now(),
			ExpiryLastUsage: true,
		}); err != nil {
			fmt.Println(err)
		}
	})

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID || m.Content != "!eth" {
			return
		}

		vaults := getVaults()
		ethVaults := filterVaults(vaults, "ETH")

		pages := map[int]*discordgo.MessageEmbed{}

		for i, vault := range ethVaults {
			if i%10 == 0 {
				pages[i/10] = &discordgo.MessageEmbed{
					Title: "ETH Vaults",
				}
			}

			tokens := ""
			for _, token := range vault.Assets {
				tokens += token + "\n"
			}

			if !vault.IsCLM {
				pages[i/10].Description += fmt.Sprintf("[%s](https://app.beefy.com/vault/%s) - **%.2f%%**\n", strings.Join(vault.Assets, "-"), vault.ID, vault.APY*100)
			} else {
				pages[i/10].Description += fmt.Sprintf("[%s](https://app.beefy.com/vault/%s) - CLM - **%.2f%%**\n", strings.Join(vault.Assets, "-"), vault.ID, vault.APY*100)
			}
		}

		if err = manager.CreateMessage(s, m.ChannelID, &paginator.Paginator{
			PageFunc: func(page int, embed *discordgo.MessageEmbed) {
				embed.Title = pages[page].Title
				embed.Description = pages[page].Description
			},
			MaxPages:        len(pages),
			Expiry:          time.Now(),
			ExpiryLastUsage: true,
		}); err != nil {
			fmt.Println(err)
		}
	})

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID || m.Content != "!bnb" {
			return
		}

		vaults := getVaults()
		usdVaults := filterVaults(vaults, "BNB")

		pages := map[int]*discordgo.MessageEmbed{}

		for i, vault := range usdVaults {
			if i%10 == 0 {
				pages[i/10] = &discordgo.MessageEmbed{
					Title: "BNB Vaults",
				}
			}

			tokens := ""
			for _, token := range vault.Assets {
				tokens += token + "\n"
			}

			pages[i/10].Description += fmt.Sprintf("[%s](https://app.beefy.com/vault/%s) - **%.2f%%**\n", strings.Join(vault.Assets, "-"), vault.ID, vault.APY*100)
		}

		if err = manager.CreateMessage(s, m.ChannelID, &paginator.Paginator{
			PageFunc: func(page int, embed *discordgo.MessageEmbed) {
				embed.Title = pages[page].Title
				embed.Description = pages[page].Description
			},
			MaxPages:        len(pages),
			Expiry:          time.Now(),
			ExpiryLastUsage: true,
		}); err != nil {
			fmt.Println(err)
		}
	})

	if err = dg.Open(); err != nil {
		fmt.Println("error opening connection: ", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
