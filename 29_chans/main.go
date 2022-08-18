package main

import (
	"channels/pkg/conv"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	//1. Конвейер
	fmt.Println("=========================")
	fmt.Println("*** Pipeline ***")
	fmt.Println("Enter an integer or \"стоп\" to finish the program:")
	log.Println("Starting pipeline")

	wg := new(sync.WaitGroup)

	for {
		n, err := conv.ReadInt() //reading integers from input using infinite loop
		if err != nil {
			if err == conv.StopWord { //the program ends if the word "стоп" received
				break
			}
			fmt.Println(err)
		}
		intChan := conv.PassToGoroutine(n, wg) //writing an integer into a channel
		sqrChan := conv.GetSquare(intChan, wg)
		multChan := conv.GetMultiple(sqrChan, wg)
		<-multChan
	}
	wg.Wait() //waiting for all goroutines
	fmt.Println("=========================")
	time.Sleep(2 * time.Second) //a little pause between tasks

	//2.1 Graceful shutdown (using channels)
	fmt.Println("*** Graceful shutdown ***")
	fmt.Println("This app counts the squares of natural numbers:")
	log.Println("Starting app")

	var (
		quitChan     = make(chan os.Signal)
		shutdownChan = make(chan struct{})
	)

	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func(shutdownChan chan struct{}, wg *sync.WaitGroup) {
		defer wg.Done()
		log.Println("Starting goroutine")
		for i := 1; ; i++ {
			select {
			case <-shutdownChan:
				log.Println("Exiting the program")
				return
			default:
				fmt.Printf("Square of %d = %d\n", i, i*i)
				time.Sleep(3 * time.Second)
			}
		}
	}(shutdownChan, wg)

	<-quitChan // received SIGINT/SIGTERM
	log.Println("Interrupt signal received")
	close(shutdownChan)
	wg.Wait()
	log.Println("Done!")
	fmt.Println("=========================")
	time.Sleep(2 * time.Second) //a little pause between tasks

	//2.1 Graceful shutdown (using context)
	fmt.Println("*** Graceful shutdown with context ***")
	fmt.Println("This app counts the squares of natural numbers:")
	log.Println("Starting app with context")

	ctx, cancel := context.WithCancel(context.Background())
	go func(cancel context.CancelFunc) {
		exitChan := make(chan os.Signal)
		signal.Notify(exitChan, os.Interrupt)
		for {
			sig := <-exitChan
			if sig == os.Interrupt {
				log.Println("Interrupt signal received")
				cancel()
				return
			}
		}
	}(cancel)

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		log.Println("Starting goroutine with context")
		defer wg.Done()
		//loop:
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				log.Println("Exiting the program by ctx.Done")
				return
				//break loop
			default:
				<-time.After(3 * time.Second)
				fmt.Printf("Square of %d = %d\n", i, i*i)

			}
		}
	}(ctx, wg)

	wg.Wait()
	log.Println("The context app is complete!")
	fmt.Println("=========================")
}
