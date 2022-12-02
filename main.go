package main

import (
	"fmt"
	"context"
	"log"
	"os"
	"github.com/shomali11/slacker"
	"github.com/joho/godotenv"
	"strconv"
)


func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main () {
	err := godotenv.Load(".env")
	if err!= nil {
		log.Fatal("Could not load environment tokens")
	}

	slackbotToken := os.Getenv("SLACK_BOT_TOKEN")
	slackappToken := os.Getenv("SLACK_APP_TOKEN")

	bot := slacker.NewClient(slackbotToken, slackappToken)

	go printCommandEvents(bot.CommandEvents())

	examples := []string{}
	examples = append(examples, "my yob is 2022")

	bot.Command("my YOB is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples: examples,
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err!=nil{
				println("error")
			}
			age:= 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	error := bot.Listen(ctx)
	if err != nil{
		log.Fatal(error)
	}

}