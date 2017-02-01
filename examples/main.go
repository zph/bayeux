package main

import (
	"fmt"

	bay "github.com/zph/bayeux"
)

func Example() {
	b := bay.Bayeux{}
	creds := bay.GetSalesforceCredentials()
	c := b.TopicToChannel(creds, "topicName")
	for {
		select {
		case e := <-c:
			fmt.Printf("TriggerEvent Received: %+v", e)
		}
	}
}

func main() {
	Example()
}
