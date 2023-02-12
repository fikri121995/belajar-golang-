package belajargolanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func(){
		time.Sleep(2 * time.Second)
		channel <- "eko kurniawan khannedy"
		fmt.Println("selesai mengirim data ke channel")
	}()

	data := <- channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}