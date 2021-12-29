package wardleygraph

import (
	"context"

	arango "github.com/arangodb/go-driver"
	"github.com/pkg/errors"
	"github.com/tristanls/sst"
)

type NodeType string

var (
	Component               NodeType = "Component"
	EvolutionCharacteristic NodeType = "EvolutionCharacteristic"
	StageOfEvolution        NodeType = "StageOfEvolution"
	UserNeed                NodeType = "UserNeed"
)

var (
	characteristicIsnt                = errors.New("wardleygraph: characteristic is not an EvolutionCharacteristic")
	componentIsNotComponentOrUserNeed = errors.New("wardleygraph: component is neither Component nor UserNeed")
	customerIsNotComponent            = errors.New("wardleygraph: customer is not a Component")
	dependentIsNotComponent           = errors.New("wardleygraph: dependent is not a Component")
	dependencyIsNotComponent          = errors.New("wardleygraph: dependency is not a Component")
	needIsNotUserNeed                 = errors.New("wardleygraph: need is not a UserNeed")
)

func (wg *WardleyGraph) Component(name string) (*sst.Node, error) {
	return wg.sst.CreateNode(string(Component), name, nil, 1)
}

func (wg *WardleyGraph) MustComponent(name string) *sst.Node {
	node, err := wg.Component(name)
	if err != nil {
		panic(err)
	}
	return node
}

func (wg *WardleyGraph) UserNeed(name string, customer *sst.Node) (*sst.Node, error) {
	if customer.Prefix != string(Component)+"/" {
		return nil, customerIsNotComponent
	}
	// FIXME: should be a transaction
	un, err := wg.sst.CreateNode(string(UserNeed), name, nil, 1)
	if err != nil {
		return nil, errors.Wrap(err, "wardleygraph: failed to create UserNeed node")
	}
	err = wg.sst.CreateLink(customer, string(Expresses), un, 1)
	if err != nil {
		return nil, errors.Wrap(err, "wardleygraph: failed to create Expresses link between customer and user need")
	}
	return un, nil
}

func (wg *WardleyGraph) MustUserNeed(name string, customer *sst.Node) *sst.Node {
	node, err := wg.UserNeed(name, customer)
	if err != nil {
		panic(err)
	}
	return node
}

func (wg *WardleyGraph) DependsOn(dependent, dependency *sst.Node) error {
	if dependent.Prefix != string(Component)+"/" {
		return dependentIsNotComponent
	}
	if dependency.Prefix != string(Component)+"/" {
		return dependencyIsNotComponent
	}
	return wg.sst.CreateLink(dependent, string(DependsOn), dependency, 1)
}

func (wg *WardleyGraph) MustDependsOn(dependent, dependency *sst.Node) {
	err := wg.DependsOn(dependent, dependency)
	if err != nil {
		panic(err)
	}
}

func (wg *WardleyGraph) ExpressUserNeed(customer, need *sst.Node) error {
	if customer.Prefix != string(Component)+"/" {
		return customerIsNotComponent
	}
	if need.Prefix != string(UserNeed)+"/" {
		return needIsNotUserNeed
	}
	return wg.sst.CreateLink(customer, string(Expresses), need, 1)
}

func (wg *WardleyGraph) MustExpressUserNeed(customer, need *sst.Node) {
	err := wg.ExpressUserNeed(customer, need)
	if err != nil {
		panic(err)
	}
}

func (wg *WardleyGraph) FulfilledBy(need, dependency *sst.Node) error {
	if need.Prefix != string(UserNeed)+"/" {
		return needIsNotUserNeed
	}
	if dependency.Prefix != string(Component)+"/" {
		return dependencyIsNotComponent
	}
	return wg.sst.CreateLink(need, string(FulfilledBy), dependency, 1)
}

func (wg *WardleyGraph) MustFulfilledBy(need, dependency *sst.Node) {
	err := wg.FulfilledBy(need, dependency)
	if err != nil {
		panic(err)
	}
}

func (wg *WardleyGraph) ExpressCharacteristic(component, characteristic *sst.Node) error {
	if component.Prefix != string(Component)+"/" && component.Prefix != string(UserNeed) {
		return componentIsNotComponentOrUserNeed
	}
	if characteristic.Prefix != string(EvolutionCharacteristic)+"/" {
		return characteristicIsnt
	}
	// FIXME: should be a transaction

	// Find if another characteristic in the same category is already expressed by this component.
	// For example, if we are adding Ubiquity II, this will find if the component currently
	// expresses Ubiquity IV.
	qs := `
	LET categories = (
		FOR v, l, p
			IN 1
			INBOUND @new_charstc_id
			Contains
			FILTER v.Prefix == @charstc_prefix
				AND (LENGTH(p.edges) > 0 && p.edges[0].semantics == 'contains')
			RETURN v._id
	)
	FOR v, l, p
		IN 2
		OUTBOUND @comp_id
		Expresses, INBOUND Contains
		PRUNE (LENGTH(p.edges) > 0 && p.edges[0].semantics != 'expresses')
			OR p.vertices[1]._id == @new_charstc_id
		FILTER v.Prefix == @charstc_prefix
		FILTER p.vertices[2]._id IN categories
		RETURN p.vertices[1]
	`
	vars := map[string]interface{}{
		"charstc_prefix": string(EvolutionCharacteristic) + "/",
		"comp_id":        string(Component) + "/" + component.Key,
		"new_charstc_id": string(EvolutionCharacteristic) + "/" + characteristic.Key,
	}
	cursor, err := wg.sst.Query(context.TODO(), qs, vars)
	if err != nil {
		return errors.Wrap(err, "wardleygraph: failed to query for related characteristics")
	}
	defer cursor.Close()
	for {
		var charstc sst.Node
		_, err := cursor.ReadDocument(context.TODO(), &charstc)
		if arango.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return errors.Wrap(err, "wardleygraph: failed to read existing characteristic")
		}
		err = wg.sst.DeleteLink(component, string(Expresses), &charstc)
		if err != nil {
			return errors.Wrapf(err, "wardleygraph: failed to remove expression of existing characteristic: %v", charstc.Key)
		}
	}
	return wg.sst.CreateLink(component, string(Expresses), characteristic, 1)
}

func (wg *WardleyGraph) MustExpressCharacteristic(component, characteristic *sst.Node) {
	err := wg.ExpressCharacteristic(component, characteristic)
	if err != nil {
		panic(err)
	}
}
