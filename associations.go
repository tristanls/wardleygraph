package wardleygraph

import "github.com/tristanls/sst"

type Semantic string

var (
	Contains    Semantic = "contains"
	DependsOn   Semantic = "depends_on"
	EvolvedFrom Semantic = "evolved_from"
	Expresses   Semantic = "expresses"
	FulfilledBy Semantic = "fulfilled_by"
	Near        Semantic = "near"
)

var (
	associations = map[string]*sst.Association{
		string(Contains):    {Key: string(Contains), SemanticType: sst.Contains, Fwd: "contains", Bwd: "constitutes", Nfwd: "does not contain", Nbwd: "does not constitute"},
		string(DependsOn):   {Key: string(DependsOn), SemanticType: sst.Follows, Fwd: "depends on", Bwd: "is depended on by", Nfwd: "does not depend on", Nbwd: "is not depended on by"},
		string(EvolvedFrom): {Key: string(EvolvedFrom), SemanticType: sst.Follows, Fwd: "evolved from", Bwd: "evolves into", Nfwd: "did not evolve from", Nbwd: "does not evolve into"},
		string(Expresses):   {Key: string(Expresses), SemanticType: sst.Expresses, Fwd: "expresses", Bwd: "is expressed by", Nfwd: "does not express", Nbwd: "is not expressed by"},
		string(FulfilledBy): {Key: string(FulfilledBy), SemanticType: sst.Follows, Fwd: "fulfilled by", Bwd: "fulfills", Nfwd: "is not fulfilled by", Nbwd: "does not fulfill"},
		string(Near):        {Key: string(Near), SemanticType: sst.Near, Fwd: "is near", Bwd: "is near", Nfwd: "is not near", Nbwd: "is not near"},
	}
)
