package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"context"

	"github.com/judwhite/go-svc/svc"
)

var curDir string

// program implements svc.Service
type program struct {
	wg   sync.WaitGroup
	quit chan struct{}

	CancelFunc context.CancelFunc
}

func (p *program) Init(env svc.Environment) error {
	log.Printf("is win service? %v\n", env.IsWindowsService())

	return nil
}

func (p *program) Start() error {
	// The Start method must not block, or Windows may assume your service failed
	// to start. Launch a Goroutine here to do something interesting/blocking.
	// start crawling
	ctx, cancel := context.WithCancel(context.Background())
	p.CancelFunc = cancel

	go run(ctx)

	return nil
}

func (p *program) Stop() error {
	// The Stop method is invoked by stopping the Windows service, or by pressing Ctrl+C on the console.
	// This method may block, but it's a good idea to finish quickly or your process may be killed by
	// Windows during a shutdown/reboot. As a general rule you shouldn't rely on graceful shutdown.

	log.Println("Stopping...")
	p.CancelFunc()

	return nil
}

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

	prg := &program{}

	// Call svc.Run to start your program/service.
	if err := svc.Run(prg); err != nil {
		log.Fatal(err)
	}
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
