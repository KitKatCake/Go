package main

import (
	"log"
	"context"
	"time"
	"os"
)

var logg *log.Logger

func someHandler()  {
	ctx,cancel := context.WithCancel(context.Background())

	go doStuff(ctx)

	time.Sleep(10*time.Second)

	log.Println("cancel")


	cancel()


}

func doStuff(ctx context.Context)  {
	for{
		time.Sleep(1*time.Second)

		select {
		case <-ctx.Done():
			log.Println("done")
		return
		default:
			log.Println("work")
		}
	}
}

func main()  {
	logg = log.New(os.Stdout,"",log.Ltime)

	someHandler()

	logg.Printf("down")

}
