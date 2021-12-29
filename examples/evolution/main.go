package main

import (
	"fmt"

	"github.com/tristanls/wardleygraph"
)

func main() {
	conf := &wardleygraph.Config{
		Name:     "wg_evolution",
		Password: "",
		URL:      "http://localhost:8529",
		Username: "root",
	}
	_, err := wardleygraph.New(conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("wardley graph evolution setup")
	// create two components
	// link components to characteristics
}
