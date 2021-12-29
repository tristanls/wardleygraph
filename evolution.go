package wardleygraph

import (
	"github.com/pkg/errors"
	"github.com/tristanls/sst"
)

type linkSpec struct {
	From *sst.Node
	Rel  Semantic
	To   *sst.Node
}

type Characteristic struct {
	Key         string
	Description string
}

var (
	Ubiquity    = &Characteristic{Key: "Ubiquity", Description: "Ubiquity"}
	UbiquityI   = &Characteristic{Key: "Ubiquity I", Description: "Rare"}
	UbiquityII  = &Characteristic{Key: "Ubiquity II", Description: "Slowly increasing"}
	UbiquityIII = &Characteristic{Key: "Ubiquity III", Description: "Rapidly increasing"}
	UbiquityIV  = &Characteristic{Key: "Ubiquity IV", Description: "Widespread in the applicable market / ecosystem"}

	Certainty    = &Characteristic{Key: "Certainty", Description: "Certainty"}
	CertaintyI   = &Characteristic{Key: "Certainty I", Description: "Poorly understood / exploring the unknown"}
	CertaintyII  = &Characteristic{Key: "Certainty II", Description: "Rapid increases in learning / discovery becomes refining"}
	CertaintyIII = &Characteristic{Key: "Certainty III", Description: "Rapid increases in use / increasing fit for purpose"}
	CertaintyIV  = &Characteristic{Key: "Certainty IV", Description: "Commonly understood (in terms of use)"}

	PublicationType    = &Characteristic{Key: "Publication Type", Description: "Publication Type"}
	PublicationTypeI   = &Characteristic{Key: "Publication Type I", Description: "Describe the wonder of the thing / the discovery of some marvel / a new land / an unknown frontier"}
	PublicationTypeII  = &Characteristic{Key: "Publication Type II", Description: "Focused on build / construct / awareness and learning / many models of explanation / no accepted forms / a wild west"}
	PublicationTypeIII = &Characteristic{Key: "Publication Type III", Description: "Maintenance / operations / installation / comparison between competing forms / feature analysis e.g. merits of one model over another"}
	PublicationTypeIV  = &Characteristic{Key: "Publication Type IV", Description: "Focused on use / increasingly an accepted, almost invisible component"}

	Market    = &Characteristic{Key: "Market", Description: "Market"}
	MarketI   = &Characteristic{Key: "Market I", Description: "Undefined market"}
	MarketII  = &Characteristic{Key: "Market II", Description: "Forming market / competing forms and different models of understanding"}
	MarketIII = &Characteristic{Key: "Market III", Description: "Growing market / consolidation to a few competing but accepted forms"}
	MarketIV  = &Characteristic{Key: "Market IV", Description: "Mature market / stabilized to an accepted form"}

	KnowledgeManagement    = &Characteristic{Key: "Knowledge Management", Description: "Knowledge Management"}
	KnowledgeManagementI   = &Characteristic{Key: "Knowledge Management I", Description: "Uncertain"}
	KnowledgeManagementII  = &Characteristic{Key: "Knowledge Management II", Description: "Learning on use / focused on testing prediction"}
	KnowledgeManagementIII = &Characteristic{Key: "Knowledge Management III", Description: "Learning on operation / using prediction / verification"}
	KnowledgeManagementIV  = &Characteristic{Key: "Knowledge Management IV", Description: "Known / accepted"}

	MarketPerception    = &Characteristic{Key: "Market Perception", Description: "Market Perception"}
	MarketPerceptionI   = &Characteristic{Key: "Market Perception I", Description: "Chaotic (non-linear) / Domain of the \"crazy\""}
	MarketPerceptionII  = &Characteristic{Key: "Market Perception II", Description: "Domain of \"experts\""}
	MarketPerceptionIII = &Characteristic{Key: "Market Perception III", Description: "Increasing expectation of use / Domain of \"professionals\""}
	MarketPerceptionIV  = &Characteristic{Key: "Market Perception IV", Description: "Ordered (appearing of being linear) / trivial / formula applied"}

	UserPerception    = &Characteristic{Key: "User Perception", Description: "User Perception"}
	UserPerceptionI   = &Characteristic{Key: "User Perception I", Description: "Different / confusing / exciting / surprising / dangerous"}
	UserPerceptionII  = &Characteristic{Key: "User Perception II", Description: "Leading edge / emerging / uncertainty over results"}
	UserPerceptionIII = &Characteristic{Key: "User Perception III", Description: "Increasingly common / disappointed if not used or available / feeling left behind"}
	UserPerceptionIV  = &Characteristic{Key: "User Perception IV", Description: "Standard / expected / feeling of shock if not used"}

	IndustryPerception    = &Characteristic{Key: "Industry Perception", Description: "Industry Perception"}
	IndustryPerceptionI   = &Characteristic{Key: "Industry Perception I", Description: "Future source of competitive advantage / unpredictable / unknown"}
	IndustryPerceptionII  = &Characteristic{Key: "Industry Perception II", Description: "Seen as a competitive advantage / a differential / looking for ROI and case examples"}
	IndustryPerceptionIII = &Characteristic{Key: "Industry Perception III", Description: "Advantage through implementation / features / this model is better than that"}
	IndustryPerceptionIV  = &Characteristic{Key: "Industry Perception IV", Description: "Cost of doing business / accepted / specific defined models"}

	FocusOfValue    = &Characteristic{Key: "Focus Of Value", Description: "Focus Of Value"}
	FocusOfValueI   = &Characteristic{Key: "Focus Of Value I", Description: "High future worth but immediate investment"}
	FocusOfValueII  = &Characteristic{Key: "Focus Of Value II", Description: "Seeking ways to profit and a ROI / seeking confirmation of value"}
	FocusOfValueIII = &Characteristic{Key: "Focus Of Value III", Description: "High profitability per unit / a valuable model / a feeling of understanding / focus on exploitation"}
	FocusOfValueIV  = &Characteristic{Key: "Focus Of Value IV", Description: "High volume / reducing margin / important but invisible / an essential component of something more complex"}

	Understanding    = &Characteristic{Key: "Understanding", Description: "Understanding"}
	UnderstandingI   = &Characteristic{Key: "Understanding I", Description: "Poorly understood / unpredicable"}
	UnderstandingII  = &Characteristic{Key: "Understanding II", Description: "Increasing understanding / development of measures"}
	UnderstandingIII = &Characteristic{Key: "Understanding III", Description: "Increasing education / constant refinement of needs / measures"}
	UnderstandingIV  = &Characteristic{Key: "Understanding IV", Description: "Believed to be well defined / stable / measurable"}

	Comparison    = &Characteristic{Key: "Comparison", Description: "Comparison"}
	ComparisonI   = &Characteristic{Key: "Comparison I", Description: "Constantly changing / a differential / unstable"}
	ComparisonII  = &Characteristic{Key: "Comparison II", Description: "Learning from others / testing the water / some evidential support"}
	ComparisonIII = &Characteristic{Key: "Comparison III", Description: "Competing models / feature difference / evidential support"}
	ComparisonIV  = &Characteristic{Key: "Comparison IV", Description: "Essential / any advantage is operational / accepted norm"}

	Failure    = &Characteristic{Key: "Failure", Description: "Failure"}
	FailureI   = &Characteristic{Key: "Failure I", Description: "High / tolerated / assumed to be wrong"}
	FailureII  = &Characteristic{Key: "Failure II", Description: "Moderate / unsurprising if wrong but disappointed"}
	FailureIII = &Characteristic{Key: "Failure III", Description: "Not tolerated / assumed to be in the right direction / resistance to changing"}
	FailureIV  = &Characteristic{Key: "Failure IV", Description: "Surprised by failure / focus on operational efficiency"}

	MarketAction    = &Characteristic{Key: "Market Action", Description: "Market Action"}
	MarketActionI   = &Characteristic{Key: "Market Action I", Description: "Gambling / driven by gut"}
	MarketActionII  = &Characteristic{Key: "Market Action II", Description: "Exploring a \"found\" value"}
	MarketActionIII = &Characteristic{Key: "Market Action III", Description: "Market analysis / listening to customers"}
	MarketActionIV  = &Characteristic{Key: "Market Action IV", Description: "Metric driven / build what is needed"}

	Efficiency    = &Characteristic{Key: "Efficiency", Description: "Efficiency"}
	EfficiencyI   = &Characteristic{Key: "Efficiency I", Description: "Reducing the cost of change (experimentation)"}
	EfficiencyII  = &Characteristic{Key: "Efficiency II", Description: "Reducing cost of waste (learning)"}
	EfficiencyIII = EfficiencyII
	EfficiencyIV  = &Characteristic{Key: "Efficiency IV", Description: "Reducing cost of deviation (volume)"}

	DecisionDrivers    = &Characteristic{Key: "Decision Drivers", Description: "Decision Drivers"}
	DecisionDriversI   = &Characteristic{Key: "Decision Drivers I", Description: "Heritage / culture"}
	DecisionDriversII  = &Characteristic{Key: "Decision Drivers II", Description: "Analysis & synthesis"}
	DecisionDriversIII = DecisionDriversII
	DecisionDriversIV  = &Characteristic{Key: "Decision Drivers IV", Description: "Previous experience"}
)

var (
	characteristicNotFound = errors.New("wardleygraph: characteristic not found")
)

func (wg *WardleyGraph) evolutionSpec() []*linkSpec {
	spec := []*linkSpec{
		{wg.stages["II"], EvolvedFrom, wg.stages["I"]},
		{wg.stages["III"], EvolvedFrom, wg.stages["II"]},
		{wg.stages["IV"], EvolvedFrom, wg.stages["III"]},

		{wg.characteristics[UbiquityIV.Key], EvolvedFrom, wg.characteristics[UbiquityIII.Key]},
		{wg.characteristics[UbiquityIII.Key], EvolvedFrom, wg.characteristics[UbiquityII.Key]},
		{wg.characteristics[UbiquityII.Key], EvolvedFrom, wg.characteristics[UbiquityI.Key]},
		{wg.characteristics[UbiquityIV.Key], Near, wg.characteristics[UbiquityIII.Key]},
		{wg.characteristics[UbiquityIII.Key], Near, wg.characteristics[UbiquityII.Key]},
		{wg.characteristics[UbiquityII.Key], Near, wg.characteristics[UbiquityI.Key]},
		{wg.characteristics[Ubiquity.Key], Contains, wg.characteristics[UbiquityI.Key]},
		{wg.characteristics[Ubiquity.Key], Contains, wg.characteristics[UbiquityII.Key]},
		{wg.characteristics[Ubiquity.Key], Contains, wg.characteristics[UbiquityIII.Key]},
		{wg.characteristics[Ubiquity.Key], Contains, wg.characteristics[UbiquityIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[UbiquityI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[UbiquityII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[UbiquityIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[UbiquityIV.Key]},

		{wg.characteristics[CertaintyIV.Key], EvolvedFrom, wg.characteristics[CertaintyIII.Key]},
		{wg.characteristics[CertaintyIII.Key], EvolvedFrom, wg.characteristics[CertaintyII.Key]},
		{wg.characteristics[CertaintyII.Key], EvolvedFrom, wg.characteristics[CertaintyI.Key]},
		{wg.characteristics[CertaintyIV.Key], Near, wg.characteristics[CertaintyIII.Key]},
		{wg.characteristics[CertaintyIII.Key], Near, wg.characteristics[CertaintyII.Key]},
		{wg.characteristics[CertaintyII.Key], Near, wg.characteristics[CertaintyI.Key]},
		{wg.characteristics[Certainty.Key], Contains, wg.characteristics[CertaintyI.Key]},
		{wg.characteristics[Certainty.Key], Contains, wg.characteristics[CertaintyII.Key]},
		{wg.characteristics[Certainty.Key], Contains, wg.characteristics[CertaintyIII.Key]},
		{wg.characteristics[Certainty.Key], Contains, wg.characteristics[CertaintyIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[CertaintyI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[CertaintyII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[CertaintyIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[CertaintyIV.Key]},

		{wg.characteristics[PublicationTypeIV.Key], EvolvedFrom, wg.characteristics[PublicationTypeIII.Key]},
		{wg.characteristics[PublicationTypeIII.Key], EvolvedFrom, wg.characteristics[PublicationTypeII.Key]},
		{wg.characteristics[PublicationTypeII.Key], EvolvedFrom, wg.characteristics[PublicationTypeI.Key]},
		{wg.characteristics[PublicationTypeIV.Key], Near, wg.characteristics[PublicationTypeIII.Key]},
		{wg.characteristics[PublicationTypeIII.Key], Near, wg.characteristics[PublicationTypeII.Key]},
		{wg.characteristics[PublicationTypeII.Key], Near, wg.characteristics[PublicationTypeI.Key]},
		{wg.characteristics[PublicationType.Key], Contains, wg.characteristics[PublicationTypeI.Key]},
		{wg.characteristics[PublicationType.Key], Contains, wg.characteristics[PublicationTypeII.Key]},
		{wg.characteristics[PublicationType.Key], Contains, wg.characteristics[PublicationTypeIII.Key]},
		{wg.characteristics[PublicationType.Key], Contains, wg.characteristics[PublicationTypeIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[PublicationTypeI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[PublicationTypeII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[PublicationTypeIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[PublicationTypeIV.Key]},

		{wg.characteristics[MarketIV.Key], EvolvedFrom, wg.characteristics[MarketIII.Key]},
		{wg.characteristics[MarketIII.Key], EvolvedFrom, wg.characteristics[MarketII.Key]},
		{wg.characteristics[MarketII.Key], EvolvedFrom, wg.characteristics[MarketI.Key]},
		{wg.characteristics[MarketIV.Key], Near, wg.characteristics[MarketIII.Key]},
		{wg.characteristics[MarketIII.Key], Near, wg.characteristics[MarketII.Key]},
		{wg.characteristics[MarketII.Key], Near, wg.characteristics[MarketI.Key]},
		{wg.characteristics[Market.Key], Contains, wg.characteristics[MarketI.Key]},
		{wg.characteristics[Market.Key], Contains, wg.characteristics[MarketII.Key]},
		{wg.characteristics[Market.Key], Contains, wg.characteristics[MarketIII.Key]},
		{wg.characteristics[Market.Key], Contains, wg.characteristics[MarketIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[MarketI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[MarketII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[MarketIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[MarketIV.Key]},

		{wg.characteristics[KnowledgeManagementIV.Key], EvolvedFrom, wg.characteristics[KnowledgeManagementIII.Key]},
		{wg.characteristics[KnowledgeManagementIII.Key], EvolvedFrom, wg.characteristics[KnowledgeManagementII.Key]},
		{wg.characteristics[KnowledgeManagementII.Key], EvolvedFrom, wg.characteristics[KnowledgeManagementI.Key]},
		{wg.characteristics[KnowledgeManagementIV.Key], Near, wg.characteristics[KnowledgeManagementIII.Key]},
		{wg.characteristics[KnowledgeManagementIII.Key], Near, wg.characteristics[KnowledgeManagementII.Key]},
		{wg.characteristics[KnowledgeManagementII.Key], Near, wg.characteristics[KnowledgeManagementI.Key]},
		{wg.characteristics[KnowledgeManagement.Key], Contains, wg.characteristics[KnowledgeManagementI.Key]},
		{wg.characteristics[KnowledgeManagement.Key], Contains, wg.characteristics[KnowledgeManagementII.Key]},
		{wg.characteristics[KnowledgeManagement.Key], Contains, wg.characteristics[KnowledgeManagementIII.Key]},
		{wg.characteristics[KnowledgeManagement.Key], Contains, wg.characteristics[KnowledgeManagementIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[KnowledgeManagementI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[KnowledgeManagementII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[KnowledgeManagementIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[KnowledgeManagementIV.Key]},

		{wg.characteristics[MarketPerceptionIV.Key], EvolvedFrom, wg.characteristics[MarketPerceptionIII.Key]},
		{wg.characteristics[MarketPerceptionIII.Key], EvolvedFrom, wg.characteristics[MarketPerceptionII.Key]},
		{wg.characteristics[MarketPerceptionII.Key], EvolvedFrom, wg.characteristics[MarketPerceptionI.Key]},
		{wg.characteristics[MarketPerceptionIV.Key], Near, wg.characteristics[MarketPerceptionIII.Key]},
		{wg.characteristics[MarketPerceptionIII.Key], Near, wg.characteristics[MarketPerceptionII.Key]},
		{wg.characteristics[MarketPerceptionII.Key], Near, wg.characteristics[MarketPerceptionI.Key]},
		{wg.characteristics[MarketPerception.Key], Contains, wg.characteristics[MarketPerceptionI.Key]},
		{wg.characteristics[MarketPerception.Key], Contains, wg.characteristics[MarketPerceptionII.Key]},
		{wg.characteristics[MarketPerception.Key], Contains, wg.characteristics[MarketPerceptionIII.Key]},
		{wg.characteristics[MarketPerception.Key], Contains, wg.characteristics[MarketPerceptionIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[MarketPerceptionI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[MarketPerceptionII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[MarketPerceptionIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[MarketPerceptionIV.Key]},

		{wg.characteristics[UserPerceptionIV.Key], EvolvedFrom, wg.characteristics[UserPerceptionIII.Key]},
		{wg.characteristics[UserPerceptionIII.Key], EvolvedFrom, wg.characteristics[UserPerceptionII.Key]},
		{wg.characteristics[UserPerceptionII.Key], EvolvedFrom, wg.characteristics[UserPerceptionI.Key]},
		{wg.characteristics[UserPerceptionIV.Key], Near, wg.characteristics[UserPerceptionIII.Key]},
		{wg.characteristics[UserPerceptionIII.Key], Near, wg.characteristics[UserPerceptionII.Key]},
		{wg.characteristics[UserPerceptionII.Key], Near, wg.characteristics[UserPerceptionI.Key]},
		{wg.characteristics[UserPerception.Key], Contains, wg.characteristics[UserPerceptionI.Key]},
		{wg.characteristics[UserPerception.Key], Contains, wg.characteristics[UserPerceptionII.Key]},
		{wg.characteristics[UserPerception.Key], Contains, wg.characteristics[UserPerceptionIII.Key]},
		{wg.characteristics[UserPerception.Key], Contains, wg.characteristics[UserPerceptionIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[UserPerceptionI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[UserPerceptionII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[UserPerceptionIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[UserPerceptionIV.Key]},

		{wg.characteristics[IndustryPerceptionIV.Key], EvolvedFrom, wg.characteristics[IndustryPerceptionIII.Key]},
		{wg.characteristics[IndustryPerceptionIII.Key], EvolvedFrom, wg.characteristics[IndustryPerceptionII.Key]},
		{wg.characteristics[IndustryPerceptionII.Key], EvolvedFrom, wg.characteristics[IndustryPerceptionI.Key]},
		{wg.characteristics[IndustryPerceptionIV.Key], Near, wg.characteristics[IndustryPerceptionIII.Key]},
		{wg.characteristics[IndustryPerceptionIII.Key], Near, wg.characteristics[IndustryPerceptionII.Key]},
		{wg.characteristics[IndustryPerceptionII.Key], Near, wg.characteristics[IndustryPerceptionI.Key]},
		{wg.characteristics[IndustryPerception.Key], Contains, wg.characteristics[IndustryPerceptionI.Key]},
		{wg.characteristics[IndustryPerception.Key], Contains, wg.characteristics[IndustryPerceptionII.Key]},
		{wg.characteristics[IndustryPerception.Key], Contains, wg.characteristics[IndustryPerceptionIII.Key]},
		{wg.characteristics[IndustryPerception.Key], Contains, wg.characteristics[IndustryPerceptionIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[IndustryPerceptionI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[IndustryPerceptionII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[IndustryPerceptionIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[IndustryPerceptionIV.Key]},

		{wg.characteristics[FocusOfValueIV.Key], EvolvedFrom, wg.characteristics[FocusOfValueIII.Key]},
		{wg.characteristics[FocusOfValueIII.Key], EvolvedFrom, wg.characteristics[FocusOfValueII.Key]},
		{wg.characteristics[FocusOfValueII.Key], EvolvedFrom, wg.characteristics[FocusOfValueI.Key]},
		{wg.characteristics[FocusOfValueIV.Key], Near, wg.characteristics[FocusOfValueIII.Key]},
		{wg.characteristics[FocusOfValueIII.Key], Near, wg.characteristics[FocusOfValueII.Key]},
		{wg.characteristics[FocusOfValueII.Key], Near, wg.characteristics[FocusOfValueI.Key]},
		{wg.characteristics[FocusOfValue.Key], Contains, wg.characteristics[FocusOfValueI.Key]},
		{wg.characteristics[FocusOfValue.Key], Contains, wg.characteristics[FocusOfValueII.Key]},
		{wg.characteristics[FocusOfValue.Key], Contains, wg.characteristics[FocusOfValueIII.Key]},
		{wg.characteristics[FocusOfValue.Key], Contains, wg.characteristics[FocusOfValueIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[FocusOfValueI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[FocusOfValueII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[FocusOfValueIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[FocusOfValueIV.Key]},

		{wg.characteristics[UnderstandingIV.Key], EvolvedFrom, wg.characteristics[UnderstandingIII.Key]},
		{wg.characteristics[UnderstandingIII.Key], EvolvedFrom, wg.characteristics[UnderstandingII.Key]},
		{wg.characteristics[UnderstandingII.Key], EvolvedFrom, wg.characteristics[UnderstandingI.Key]},
		{wg.characteristics[UnderstandingIV.Key], Near, wg.characteristics[UnderstandingIII.Key]},
		{wg.characteristics[UnderstandingIII.Key], Near, wg.characteristics[UnderstandingII.Key]},
		{wg.characteristics[UnderstandingII.Key], Near, wg.characteristics[UnderstandingI.Key]},
		{wg.characteristics[Understanding.Key], Contains, wg.characteristics[UnderstandingI.Key]},
		{wg.characteristics[Understanding.Key], Contains, wg.characteristics[UnderstandingII.Key]},
		{wg.characteristics[Understanding.Key], Contains, wg.characteristics[UnderstandingIII.Key]},
		{wg.characteristics[Understanding.Key], Contains, wg.characteristics[UnderstandingIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[UnderstandingI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[UnderstandingII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[UnderstandingIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[UnderstandingIV.Key]},

		{wg.characteristics[ComparisonIV.Key], EvolvedFrom, wg.characteristics[ComparisonIII.Key]},
		{wg.characteristics[ComparisonIII.Key], EvolvedFrom, wg.characteristics[ComparisonII.Key]},
		{wg.characteristics[ComparisonII.Key], EvolvedFrom, wg.characteristics[ComparisonI.Key]},
		{wg.characteristics[ComparisonIV.Key], Near, wg.characteristics[ComparisonIII.Key]},
		{wg.characteristics[ComparisonIII.Key], Near, wg.characteristics[ComparisonII.Key]},
		{wg.characteristics[ComparisonII.Key], Near, wg.characteristics[ComparisonI.Key]},
		{wg.characteristics[Comparison.Key], Contains, wg.characteristics[ComparisonI.Key]},
		{wg.characteristics[Comparison.Key], Contains, wg.characteristics[ComparisonII.Key]},
		{wg.characteristics[Comparison.Key], Contains, wg.characteristics[ComparisonIII.Key]},
		{wg.characteristics[Comparison.Key], Contains, wg.characteristics[ComparisonIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[ComparisonI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[ComparisonII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[ComparisonIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[ComparisonIV.Key]},

		{wg.characteristics[FailureIV.Key], EvolvedFrom, wg.characteristics[FailureIII.Key]},
		{wg.characteristics[FailureIII.Key], EvolvedFrom, wg.characteristics[FailureII.Key]},
		{wg.characteristics[FailureII.Key], EvolvedFrom, wg.characteristics[FailureI.Key]},
		{wg.characteristics[FailureIV.Key], Near, wg.characteristics[FailureIII.Key]},
		{wg.characteristics[FailureIII.Key], Near, wg.characteristics[FailureII.Key]},
		{wg.characteristics[FailureII.Key], Near, wg.characteristics[FailureI.Key]},
		{wg.characteristics[Failure.Key], Contains, wg.characteristics[FailureI.Key]},
		{wg.characteristics[Failure.Key], Contains, wg.characteristics[FailureII.Key]},
		{wg.characteristics[Failure.Key], Contains, wg.characteristics[FailureIII.Key]},
		{wg.characteristics[Failure.Key], Contains, wg.characteristics[FailureIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[FailureI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[FailureII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[FailureIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[FailureIV.Key]},

		{wg.characteristics[MarketActionIV.Key], EvolvedFrom, wg.characteristics[MarketActionIII.Key]},
		{wg.characteristics[MarketActionIII.Key], EvolvedFrom, wg.characteristics[MarketActionII.Key]},
		{wg.characteristics[MarketActionII.Key], EvolvedFrom, wg.characteristics[MarketActionI.Key]},
		{wg.characteristics[MarketActionIV.Key], Near, wg.characteristics[MarketActionIII.Key]},
		{wg.characteristics[MarketActionIII.Key], Near, wg.characteristics[MarketActionII.Key]},
		{wg.characteristics[MarketActionII.Key], Near, wg.characteristics[MarketActionI.Key]},
		{wg.characteristics[MarketAction.Key], Contains, wg.characteristics[MarketActionI.Key]},
		{wg.characteristics[MarketAction.Key], Contains, wg.characteristics[MarketActionII.Key]},
		{wg.characteristics[MarketAction.Key], Contains, wg.characteristics[MarketActionIII.Key]},
		{wg.characteristics[MarketAction.Key], Contains, wg.characteristics[MarketActionIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[MarketActionI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[MarketActionII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[MarketActionIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[MarketActionIV.Key]},

		{wg.characteristics[EfficiencyIV.Key], EvolvedFrom, wg.characteristics[EfficiencyIII.Key]},
		// {wg.characteristics[EfficiencyIII.Key], EvolvedFrom, wg.characteristics[EfficiencyII.Key]}, II == III
		{wg.characteristics[EfficiencyII.Key], EvolvedFrom, wg.characteristics[EfficiencyI.Key]},
		{wg.characteristics[EfficiencyIV.Key], Near, wg.characteristics[EfficiencyIII.Key]},
		// {wg.characteristics[EfficiencyIII.Key], Near, wg.characteristics[EfficiencyII.Key]}, II == III
		{wg.characteristics[EfficiencyII.Key], Near, wg.characteristics[EfficiencyI.Key]},
		{wg.characteristics[Efficiency.Key], Contains, wg.characteristics[EfficiencyI.Key]},
		{wg.characteristics[Efficiency.Key], Contains, wg.characteristics[EfficiencyII.Key]},
		{wg.characteristics[Efficiency.Key], Contains, wg.characteristics[EfficiencyIII.Key]},
		{wg.characteristics[Efficiency.Key], Contains, wg.characteristics[EfficiencyIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[EfficiencyI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[EfficiencyII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[EfficiencyIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[EfficiencyIV.Key]},

		{wg.characteristics[DecisionDriversIV.Key], EvolvedFrom, wg.characteristics[DecisionDriversIII.Key]},
		// {wg.characteristics[DecisionDriversIII.Key], EvolvedFrom, wg.characteristics[DecisionDriversII.Key]}, II == III
		{wg.characteristics[DecisionDriversII.Key], EvolvedFrom, wg.characteristics[DecisionDriversI.Key]},
		{wg.characteristics[DecisionDriversIV.Key], Near, wg.characteristics[DecisionDriversIII.Key]},
		// {wg.characteristics[DecisionDriversIII.Key], Near, wg.characteristics[DecisionDriversII.Key]}, II == III
		{wg.characteristics[DecisionDriversII.Key], Near, wg.characteristics[DecisionDriversI.Key]},
		{wg.characteristics[DecisionDrivers.Key], Contains, wg.characteristics[DecisionDriversI.Key]},
		{wg.characteristics[DecisionDrivers.Key], Contains, wg.characteristics[DecisionDriversII.Key]},
		{wg.characteristics[DecisionDrivers.Key], Contains, wg.characteristics[DecisionDriversIII.Key]},
		{wg.characteristics[DecisionDrivers.Key], Contains, wg.characteristics[DecisionDriversIV.Key]},
		{wg.stages["I"], Contains, wg.characteristics[DecisionDriversI.Key]},
		{wg.stages["II"], Contains, wg.characteristics[DecisionDriversII.Key]},
		{wg.stages["III"], Contains, wg.characteristics[DecisionDriversIII.Key]},
		{wg.stages["IV"], Contains, wg.characteristics[DecisionDriversIV.Key]},
	}
	return spec
}

func (wg *WardleyGraph) Characteristic(characteristic string) (*sst.Node, error) {
	char := wg.characteristics[characteristic]
	if char == nil {
		return nil, characteristicNotFound
	}
	return char, nil
}

func (wg *WardleyGraph) MustCharacteristic(characteristic string) *sst.Node {
	node, err := wg.Characteristic(characteristic)
	if err != nil {
		panic(err)
	}
	return node
}
