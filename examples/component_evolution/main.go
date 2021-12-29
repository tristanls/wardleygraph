package main

import (
	"fmt"

	"github.com/tristanls/wardleygraph"
)

func main() {
	conf := &wardleygraph.Config{
		Name:     "wg_component_evolution",
		Password: "",
		URL:      "http://localhost:8529",
		Username: "root",
	}
	wg, err := wardleygraph.New(conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("wardley graph evolution setup complete")

	gql := wg.MustComponent("GraphQL")
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.UbiquityIII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.CertaintyIII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.PublicationTypeII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.MarketIII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.KnowledgeManagementII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.MarketPerceptionIII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.UserPerceptionII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.IndustryPerceptionIII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.FocusOfValueII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.UnderstandingIII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.ComparisonIII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.FailureIII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.MarketActionII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.EfficiencyIII.Key))
	wg.MustExpressCharacteristic(gql, wg.MustCharacteristic(wardleygraph.DecisionDriversIV.Key))

	e := wg.MustComponent("Electricity")
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.UbiquityIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.CertaintyIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.PublicationTypeIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.MarketIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.KnowledgeManagementIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.MarketPerceptionIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.UserPerceptionIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.IndustryPerceptionIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.FocusOfValueIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.UnderstandingIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.ComparisonIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.FailureIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.MarketActionIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.EfficiencyIV.Key))
	wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.DecisionDriversIV.Key))

	fmt.Println("component evolution graph setup complete")
}
