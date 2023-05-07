package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	totalmem "github.com/pbnjay/memory"
)

// =========================================
// THIS IS ONLY A TEMPLATE FOR NOW
// =========================================
// The main function of the bot.
// We will make it as clean as possible later.

func main() {

	// Automatically set GOMAXPROCS to the number of your CPU cores.
	// Increase performance by allowing Golang to use multiple processors.
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs) // Sets the GOMAXPROCS value
	totalMem := fmt.Sprintf("Available OS Memory: %v MB | %v GB", (totalmem.TotalMemory() / Megabyte), (totalmem.TotalMemory() / Gigabyte))
	fmt.Println()
	totalcpu := fmt.Sprintf("Available CPUs: %v", numCPUs)

	fmt.Println(totalcpu)
	fmt.Println(totalMem)

	// run http server
	// go proxyServer()
	fmt.Println("HTTP server runs on port " + httpPort)

	// Create the logs folder
	osFS.RemoveAll("./logs/")
	createLogFolder := osFS.MkdirAll("./logs/", 0777)
	if createLogFolder != nil {
		fmt.Println(" [ERROR] ", createLogFolder)
	}
	fmt.Println(` [DONE] New "logs" folder has been created. \n >> `, createLogFolder)

	createDBFolder := osFS.MkdirAll("./db/", 0777)
	if createDBFolder != nil {
		fmt.Println(" [ERROR] ", createDBFolder)
	}
	fmt.Println(` [DONE] New "db" folder has been created. \n >> `, createDBFolder)

	// Create the ./cache/ folder
	osFS.RemoveAll("./cache/")
	createCacheFolder := osFS.MkdirAll("./cache/", 0777)
	if createCacheFolder != nil {
		fmt.Println(" [ERROR] ", createCacheFolder)
	}
	fmt.Println(` [DONE] New "cache" folder has been created. \n >> `, createCacheFolder)

	// Get the latest sticker list
	fmt.Println(" Processing sticker list. Please wait...")
	getStickers, err := normalclient.Get("https://0ms.run/stickers/stickers.txt")
	if err != nil {
		fmt.Println(" [getStickers] ", err)

		if len(universalLogs) >= universalLogsLimit {
			universalLogs = nil
		} else {
			universalLogs = append(universalLogs, fmt.Sprintf("\n%v", err))
		}

		return
	}

	bodyStickers, err := io.ReadAll(bufio.NewReader(getStickers.Body))
	if err != nil {
		fmt.Println(" [bodyStickers] ", err)

		if len(universalLogs) >= universalLogsLimit {
			universalLogs = nil
		} else {
			universalLogs = append(universalLogs, fmt.Sprintf("\n%v", err))
		}

		return
	}

	newstickerlist := strings.Split(string(bodyStickers), "\n")
	stickerList = append(stickerList, newstickerlist...)
	newstickerlist = nil
	fmt.Println(" Successfully fetched the sticker list.")

	// support for openai
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + discordBotToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	// dg.AddHandler(emojiReactions)
	// dg.AddHandler(getUserInfo)
	// dg.AddHandler(openBilling)
	// dg.AddHandler(openAI)

	// Just like the ping pong example, we only care about receiving message
	// events in this example.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}
