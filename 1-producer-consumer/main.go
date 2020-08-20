//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweetsChannel chan<- *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			tweetsChannel <- tweet
			close(tweetsChannel)
			return
		}

		tweetsChannel <- tweet
	}
}

func consumer(tweetsChannel <-chan *Tweet) {
	for t := range tweetsChannel {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	tweetsChannel := make(chan *Tweet)

	// Producer
	go producer(stream, tweetsChannel)

	// Consumer
	consumer(tweetsChannel)

	fmt.Printf("Process took %s\n", time.Since(start))
}
