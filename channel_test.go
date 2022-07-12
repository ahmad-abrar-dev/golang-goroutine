package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// the channel commonly used for sending data into the outside go routine itself
// it can be in the function itself or other goroutine
func TestCreateChannel(t *testing.T) {
	channelName := make(chan string)

	// to send data into channel
	// channelName <- "Zhun"
	// to get data from channel
	// data := <-channel
	// or we can also use it as parameter
	// fmt.Println(<-channelName)

	// make sure to do the close to channel
	// or your memory will get leaking
	// and to make sure the close() will be execute use defer
	// when creating channel make sure there's have an action to send or recieve the channel
	// close(channelName)
	defer close(channelName)

	//create anonymous function in go
	go func() {
		time.Sleep(2 * time.Second)
		channelName <- "Zhun Weky"
		fmt.Println("finished sending data into channel")
	}() // the block code () to determinate the golang to execute

	data := <-channelName
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}
