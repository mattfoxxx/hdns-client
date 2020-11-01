package main

import (
	"encoding/json"
	"fmt"
	hdnsclient "github.com/mattfoxxx/hdns-client"
	"os"
)

func main() {
	c := hdnsclient.NewClient(os.Getenv("HDNS_API_TOKEN"))
	res, err := c.GetRecords("5c95wmRRiFSdwNswRDcMuG", nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("An Error has occured: %s", err))
		os.Exit(1)
	}
	result, err := json.Marshal(res)
	if err != nil {
		fmt.Println(fmt.Sprintf("An Error has occured: %s", err))
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("Result:\n%s", result))

	res2, err := c.GetRecord("8ce7818ffd302808066d3732ffbb26a4")
	if err != nil {
		fmt.Println(fmt.Sprintf("An Error has occured: %s", err))
		os.Exit(1)
	}
	result2, err := json.Marshal(res2)
	if err != nil {
		fmt.Println(fmt.Sprintf("An Error has occured: %s", err))
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("Result:\n%s", result2))
}
