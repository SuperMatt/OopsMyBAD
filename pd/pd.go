package pd

import (
	"encoding/base64"
	"fmt"

	pagerduty "github.com/PagerDuty/go-pagerduty"
)

func Event(apiKey, action, name, message, source string) (*pagerduty.V2EventResponse, error) {
	//This is where I'm going to trigger the alert!

	if message == "" {
		message = "empty message"
	}
	fmt.Println("Sending Alert")
	fmt.Println("Action:", action)
	fmt.Println("Name:", name)
	fmt.Println("Message:", message)

	dedupKey := base64.StdEncoding.EncodeToString([]byte(name))

	e := pagerduty.V2Event{
		RoutingKey: apiKey,
		Action:     action,
		DedupKey:   dedupKey,
		Payload: &pagerduty.V2Payload{
			Summary:  message,
			Source:   source,
			Severity: "critical",
			Details: map[string]interface{}{
				"Details":  message,
				"Souce":    source,
				"Name":     name,
				"DedupKey": dedupKey,
			},
		},
	}
	return pagerduty.ManageEvent(e)
}
