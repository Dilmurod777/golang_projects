package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
	"time"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	_ = os.Setenv("SLACK_BOT_TOKEN", "xapp-1-A045KDF3H3K-4162265010007-fc623fe81653b48f84281e77a1f10015fd27f4a323c6c05bc51d0bc0769f01b9")
	_ = os.Setenv("SLACK_APP_TOKEN", "xoxb-4179267313380-4176752420131-3rS6U6eWjOxrt5UA6yjHXt3y")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My year of birth is <year>", &slacker.CommandDefinition{
		Description: "Year of birth calculator",
		//Examples: ["my year of birth is 2020"],
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("Error while parsing param")
			}

			currentYear, _, _ := time.Now().Date()
			age := currentYear - yob
			r := fmt.Sprintf("Your age is %d", age)
			err = response.Reply(r)
			if err != nil {
				fmt.Println("Error while sending reply")
			}
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
