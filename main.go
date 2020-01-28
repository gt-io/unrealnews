package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"context"
)

var curDir string

func main() {

	////////////////////////////////////////////////////
	// init log...
	curDir, _ = os.Getwd()
	log.Println("log path : ", curDir+"/unrealbot.log")
	fpLog, err := os.OpenFile(curDir+"/unrealbot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fpLog.Close()
	log.SetOutput(io.MultiWriter(fpLog, os.Stdout))

	// create context for cancel func
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start crawling
	run(ctx)
}

func run(ctx context.Context) {
	// load unreal version from  .version file
	version := loadUnrealVersion(curDir + "/.version")

	log.Println("Start Crawling...", version)

	bot, err := initBot(curDir + "/conf.json")
	if err != nil {
		log.Fatal(err)
	}

	cacheFlush()

	ticker := time.Tick(time.Second * 30)

	// send to zulip test msg
	bot.sendBot("start crawling..." + version)

	for {
		select {
		case <-ticker:
			datas, err := GetPageData(version)
			if err != nil {
				continue
			}
			for _, v := range datas {
				// check cache exist
				if !cacheExist(v.URL) {

					// send to bot
					bot.sendBot(fmt.Sprintf("**[%s](%s)** %s", v.Title, v.URL, v.Desc))

					// add cache
					cacheRegister(v.URL)
				}
			}

			if cacheSize() > MaxContainerSize {
				cacheFlush()
			}

		case <-ctx.Done():
			log.Println("Stop Crawling...")
			return
		}
	}
}
