package main

import (
	"fmt"

	"github.com/tristanls/sst"
	"github.com/tristanls/wardleygraph"
)

func main() {
	conf := &wardleygraph.Config{
		Name:     "wg_im_a_map",
		Password: "",
		URL:      "http://localhost:8529",
		Username: "root",
	}
	wg, err := wardleygraph.New(conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("wardley graph evolution setup complete")

	business := wg.MustComponent("Business")
	public := wg.MustComponent("Public")
	profit, _ := wg.MustUserNeed("Profit", business)
	thirstForTea, _ := wg.MustUserNeed("Thirst for Tea", public)
	cupOfTea := wg.MustComponent("Cup of Tea")
	wg.MustFulfilledBy(thirstForTea, cupOfTea)
	wg.MustFulfilledBy(profit, cupOfTea)
	staff := wg.MustComponent("Staff")
	tea := wg.MustComponent("Tea")
	cup := wg.MustComponent("Cup")
	hotWater := wg.MustComponent("Hot Water")
	wg.MustDependsOnAll(cupOfTea, []*sst.Node{staff, tea, cup, hotWater})
	kettle := wg.MustComponent("Kettle")
	water := wg.MustComponent("Water")
	wg.MustDependsOnAll(hotWater, []*sst.Node{kettle, water})
	power := wg.MustComponent("Power")
	wg.MustDependsOn(kettle, power)
	fmt.Println("dependencies created")

	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.UbiquityIV.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.CertaintyIV.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.PublicationTypeII.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.MarketIII.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.KnowledgeManagementIII.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.MarketPerceptionIII.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.UserPerceptionIV.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.IndustryPerceptionIV.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.FocusOfValueIII.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.UnderstandingIV.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.ComparisonIII.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.FailureIV.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.MarketActionIII.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.EfficiencyIII.Key))
	wg.MustExpressCharacteristic(cupOfTea, wg.MustCharacteristic(wardleygraph.DecisionDriversIII.Key))

	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.UbiquityIII.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.CertaintyIII.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.PublicationTypeIII.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.MarketIII.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.KnowledgeManagementIV.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.MarketPerceptionIV.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.UserPerceptionIII.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.IndustryPerceptionIV.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.FocusOfValueIII.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.UnderstandingIV.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.ComparisonIII.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.FailureIII.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.MarketActionIII.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.EfficiencyIII.Key))
	wg.MustExpressCharacteristic(staff, wg.MustCharacteristic(wardleygraph.DecisionDriversIII.Key))

	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.UbiquityIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.CertaintyIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.PublicationTypeIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.MarketIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.KnowledgeManagementIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.MarketPerceptionIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.UserPerceptionIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.IndustryPerceptionIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.FocusOfValueIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.UnderstandingIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.ComparisonIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.FailureIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.MarketActionIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.EfficiencyIV.Key))
	wg.MustExpressCharacteristic(cup, wg.MustCharacteristic(wardleygraph.DecisionDriversIV.Key))

	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.UbiquityIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.CertaintyIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.PublicationTypeIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.MarketIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.KnowledgeManagementIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.MarketPerceptionIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.UserPerceptionIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.IndustryPerceptionIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.FocusOfValueIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.UnderstandingIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.ComparisonIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.FailureIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.MarketActionIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.EfficiencyIV.Key))
	wg.MustExpressCharacteristic(tea, wg.MustCharacteristic(wardleygraph.DecisionDriversIV.Key))

	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.UbiquityIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.CertaintyIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.PublicationTypeIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.MarketIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.KnowledgeManagementIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.MarketPerceptionIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.UserPerceptionIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.IndustryPerceptionIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.FocusOfValueIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.UnderstandingIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.ComparisonIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.FailureIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.MarketActionIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.EfficiencyIV.Key))
	wg.MustExpressCharacteristic(hotWater, wg.MustCharacteristic(wardleygraph.DecisionDriversIV.Key))

	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.UbiquityIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.CertaintyIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.PublicationTypeIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.MarketIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.KnowledgeManagementIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.MarketPerceptionIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.UserPerceptionIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.IndustryPerceptionIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.FocusOfValueIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.UnderstandingIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.ComparisonIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.FailureIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.MarketActionIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.EfficiencyIV.Key))
	wg.MustExpressCharacteristic(water, wg.MustCharacteristic(wardleygraph.DecisionDriversIV.Key))

	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.UbiquityII.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.CertaintyII.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.PublicationTypeIV.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.MarketII.Key))
	// wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.KnowledgeManagementIV.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.MarketPerceptionII.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.UserPerceptionII.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.IndustryPerceptionII.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.FocusOfValueI.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.UnderstandingII.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.ComparisonII.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.FailureII.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.MarketActionII.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.EfficiencyII.Key))
	wg.MustExpressCharacteristic(kettle, wg.MustCharacteristic(wardleygraph.DecisionDriversII.Key))

	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.UbiquityIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.CertaintyIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.PublicationTypeIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.MarketIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.KnowledgeManagementIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.MarketPerceptionIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.UserPerceptionIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.IndustryPerceptionIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.FocusOfValueIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.UnderstandingIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.ComparisonIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.FailureIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.MarketActionIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.EfficiencyIV.Key))
	wg.MustExpressCharacteristic(power, wg.MustCharacteristic(wardleygraph.DecisionDriversIV.Key))

	fmt.Println("component characteristics configured")

	m := wg.MustMap("I'm a map!")
	wg.MustContainsAll(m, []*sst.Node{
		business, public,
		profit, thirstForTea,
		cupOfTea, staff, tea, cup, hotWater, kettle, water, power,
	})

	m2 := wg.MustMap("I'm a smaller map")
	wg.MustContainsAll(m2, []*sst.Node{
		public,
		thirstForTea,
		cupOfTea, tea, cup, hotWater, kettle, water,
	})

	summary := wg.MustCreateSummary("Tea Shop", []*sst.Node{cupOfTea, staff, tea, cup, hotWater, kettle, water, power})
	m3 := wg.MustMap("Map with summary")
	wg.MustContainsAll(m3, []*sst.Node{
		business, public,
		profit, thirstForTea,
		summary,
	})

	m4 := wg.MustMap("Map with summary and detail")
	wg.MustContainsAll(m4, []*sst.Node{
		business, public,
		profit, thirstForTea,
		summary, hotWater, kettle, water,
	})
}

// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.UbiquityIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.CertaintyIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.PublicationTypeIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.MarketIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.KnowledgeManagementIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.MarketPerceptionIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.UserPerceptionIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.IndustryPerceptionIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.FocusOfValueIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.UnderstandingIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.ComparisonIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.FailureIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.MarketActionIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.EfficiencyIV.Key))
// wg.MustExpressCharacteristic(e, wg.MustCharacteristic(wardleygraph.DecisionDriversIV.Key))
