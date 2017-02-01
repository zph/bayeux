package bayeux

import "fmt"

func Example() {
	b := Bayeux{}
	creds := GetSalesforceCredentials()
	c := b.TopicToChannel(creds, "topicName")
	for {
		select {
		case e := <-c:
			fmt.Printf("TriggerEvent Received: %+v", e)
		}
	}
}
