// Package wardleygraph provides facilities for encoding Wardley Maps
// as a Semantic Spacetime Graph
package wardleygraph

import (
	"github.com/pkg/errors"
	"github.com/tristanls/sst"
)

type Config struct {
	// Name is the name of WardleyGraph
	Name string
	// Password to ArangoDB
	Password string
	// URL to ArangoDB
	URL string
	// Username to ArangoDB
	Username string
}

type WardleyGraph struct {
	characteristics map[string]*sst.Node
	sst             *sst.SST
	stages          map[string]*sst.Node
}

func New(config *Config) (*WardleyGraph, error) {
	sstConf := &sst.Config{
		Associations: associations,
		Name:         config.Name,
		NodeCollections: []string{
			string(Component),
			string(EvolutionCharacteristic),
			string(StageOfEvolution),
			string(UserNeed),
		},
		Password: config.Password,
		URL:      config.URL,
		Username: config.Username,
	}
	s, err := sst.NewSST(sstConf)
	if err != nil {
		return nil, errors.Wrap(err, "wardleygraph: failed to create Wardley Graph Semantic Spacetime")
	}
	wg := &WardleyGraph{
		characteristics: make(map[string]*sst.Node),
		sst:             s,
		stages:          make(map[string]*sst.Node),
	}

	err = wg.ConfigureEvolution()
	if err != nil {
		return nil, errors.Wrap(err, "wardleygraph: failed to configure evolution")
	}

	return wg, nil
}

// ConfigureEvolution creates the portion of the graph representing stages of evolution,
// characteristic, and general property which components will express.
func (wg *WardleyGraph) ConfigureEvolution() (err error) {
	for _, stage := range []string{"I", "II", "III", "IV"} {
		wg.stages[stage], err = wg.sst.CreateNode(string(StageOfEvolution), stage, nil, 0)
		if err != nil {
			return err
		}
	}
	for _, charstc := range []*Characteristic{
		Ubiquity, Certainty, PublicationType, Market, KnowledgeManagement,
		MarketPerception, UserPerception, IndustryPerception, FocusOfValue,
		Understanding, Comparison, Failure, MarketAction, Efficiency, DecisionDrivers,

		UbiquityI, UbiquityII, UbiquityIII, UbiquityIV,
		CertaintyI, CertaintyII, CertaintyIII, CertaintyIV,
		PublicationTypeI, PublicationTypeII, PublicationTypeIII, PublicationTypeIV,
		MarketI, MarketII, MarketIII, MarketIV,
		KnowledgeManagementI, KnowledgeManagementII, KnowledgeManagementIII, KnowledgeManagementIV,
		MarketPerceptionI, MarketPerceptionII, MarketPerceptionIII, MarketPerceptionIV,
		UserPerceptionI, UserPerceptionII, UserPerceptionIII, UserPerceptionIV,
		IndustryPerceptionI, IndustryPerceptionII, IndustryPerceptionIII, IndustryPerceptionIV,
		FocusOfValueI, FocusOfValueII, FocusOfValueIII, FocusOfValueIV,
		UnderstandingI, UnderstandingII, UnderstandingIII, UnderstandingIV,
		ComparisonI, ComparisonII, ComparisonIII, ComparisonIV,
		FailureI, FailureII, FailureIII, FailureIV,
		MarketActionI, MarketActionII, MarketActionIII, MarketActionIV,
		EfficiencyI, EfficiencyII, EfficiencyIII, EfficiencyIV,
		DecisionDriversI, DecisionDriversII, DecisionDriversIII, DecisionDriversIV,
	} {
		wg.characteristics[charstc.Key], err = wg.sst.CreateNode(string(EvolutionCharacteristic), charstc.Key, map[string]interface{}{
			"description": charstc.Description,
		}, 0)
		if err != nil {
			return err
		}
	}

	for _, link := range wg.evolutionSpec() {
		err = wg.sst.CreateLink(link.From, string(link.Rel), link.To, 1)
		if err != nil {
			return err
		}
	}

	return nil
}
