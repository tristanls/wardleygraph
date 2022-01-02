package main

import (
	"fmt"

	"github.com/tristanls/wardleygraph"
)

func main() {
	conf := &wardleygraph.Config{
		Name:     "wg_movement",
		Password: "",
		URL:      "http://localhost:8529",
		Username: "root",
	}
	wg, err := wardleygraph.New(conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("wardley graph evolution setup complete")

	comp := wg.MustComponent("component")

	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.UbiquityI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.CertaintyI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.PublicationTypeI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.MarketI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.KnowledgeManagementI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.MarketPerceptionI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.UserPerceptionI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.IndustryPerceptionI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.FocusOfValueI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.UnderstandingI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.ComparisonI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.FailureI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.MarketActionI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.EfficiencyI.Key))
	wg.MustExpressCharacteristic(comp, wg.MustCharacteristic(wardleygraph.DecisionDriversI.Key))
}
