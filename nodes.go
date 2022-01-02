package wardleygraph

import (
	"github.com/pkg/errors"
	"github.com/tristanls/sst"
)

type NodeType string

var (
	Component               NodeType = "Component"
	EvolutionCharacteristic NodeType = "EvolutionCharacteristic"
	Map                     NodeType = "Map"
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

type NodeData struct {
	Summary Summary
}

func (wg *WardleyGraph) Component(name string) (*sst.Node, error) {
	return wg.ComponentWithData(name, nil)
}

func (wg *WardleyGraph) MustComponent(name string) *sst.Node {
	node, err := wg.Component(name)
	if err != nil {
		panic(err)
	}
	return node
}

func (wg *WardleyGraph) ComponentWithData(name string, data map[string]interface{}) (*sst.Node, error) {
	return wg.sst.CreateNode(string(Component), name, data, 1)
}

func (wg *WardleyGraph) MustComponentWithData(name string, data map[string]interface{}) *sst.Node {
	node, err := wg.ComponentWithData(name, data)
	if err != nil {
		panic(err)
	}
	return node
}

func (wg *WardleyGraph) Map(name string) (*sst.Node, error) {
	return wg.MapWithData(name, nil)
}

func (wg *WardleyGraph) MustMap(name string) *sst.Node {
	node, err := wg.Map(name)
	if err != nil {
		panic(err)
	}
	return node
}

func (wg *WardleyGraph) MapWithData(name string, data map[string]interface{}) (*sst.Node, error) {
	return wg.sst.CreateNode(string(Map), name, data, 1)
}

func (wg *WardleyGraph) MustMapWithData(name string, data map[string]interface{}) *sst.Node {
	node, err := wg.MapWithData(name, data)
	if err != nil {
		panic(err)
	}
	return node
}

func (wg *WardleyGraph) UserNeed(name string, customer *sst.Node) (*sst.Node, *sst.Link, error) {
	if customer.Prefix != string(Component)+"/" {
		return nil, nil, customerIsNotComponent
	}
	// FIXME: should be a transaction
	un, err := wg.sst.CreateNode(string(UserNeed), name, nil, 1)
	if err != nil {
		return nil, nil, errors.Wrap(err, "wardleygraph: failed to create UserNeed node")
	}
	link, err := wg.sst.CreateLink(customer, string(Expresses), un, nil, 1)
	if err != nil {
		return nil, nil, errors.Wrap(err, "wardleygraph: failed to create Expresses link between customer and user need")
	}
	return un, link, nil
}

func (wg *WardleyGraph) MustUserNeed(name string, customer *sst.Node) (*sst.Node, *sst.Link) {
	node, link, err := wg.UserNeed(name, customer)
	if err != nil {
		panic(err)
	}
	return node, link
}
