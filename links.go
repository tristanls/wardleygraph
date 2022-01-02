package wardleygraph

import (
	"context"
	"fmt"

	arango "github.com/arangodb/go-driver"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/tristanls/sst"
)

type ContainsLinkData struct {
	LinkSummary LinkSummary `json:"linkSummary,omitempty"`
}

func validDependsOn(dependent, dependency *sst.Node) error {
	if dependent.Prefix != string(Component)+"/" {
		return dependentIsNotComponent
	}
	if dependency.Prefix != string(Component)+"/" {
		return dependencyIsNotComponent
	}
	return nil
}

func (wg *WardleyGraph) DependsOn(dependent, dependency *sst.Node) (*sst.Link, error) {
	err := validDependsOn(dependent, dependency)
	if err != nil {
		return nil, err
	}
	return wg.sst.CreateLink(dependent, string(DependsOn), dependency, nil, 1)
}

func (wg *WardleyGraph) MustDependsOn(dependent, dependency *sst.Node) *sst.Link {
	link, err := wg.DependsOn(dependent, dependency)
	if err != nil {
		panic(err)
	}
	return link
}

func (wg *WardleyGraph) DependsOnAll(dependent *sst.Node, dependencies []*sst.Node) ([]*sst.Link, error) {
	links := make([]*sst.Link, len(dependencies))
	var err error
	for i, d := range dependencies {
		links[i], err = wg.DependsOn(dependent, d)
		if err != nil {
			return nil, err
		}
	}
	return links, nil
}

func (wg *WardleyGraph) MustDependsOnAll(dependent *sst.Node, dependencies []*sst.Node) []*sst.Link {
	links, err := wg.DependsOnAll(dependent, dependencies)
	if err != nil {
		panic(err)
	}
	return links
}

func validExpressUserNeed(customer, need *sst.Node) error {
	if customer.Prefix != string(Component)+"/" {
		return customerIsNotComponent
	}
	if need.Prefix != string(UserNeed)+"/" {
		return needIsNotUserNeed
	}
	return nil
}

func (wg *WardleyGraph) ExpressUserNeed(customer, need *sst.Node) (*sst.Link, error) {
	err := validExpressUserNeed(customer, need)
	if err != nil {
		return nil, err
	}
	return wg.sst.CreateLink(customer, string(Expresses), need, nil, 1)
}

func (wg *WardleyGraph) MustExpressUserNeed(customer, need *sst.Node) *sst.Link {
	link, err := wg.ExpressUserNeed(customer, need)
	if err != nil {
		panic(err)
	}
	return link
}

func validFulfilledBy(need, dependency *sst.Node) error {
	if need.Prefix != string(UserNeed)+"/" {
		return needIsNotUserNeed
	}
	if dependency.Prefix != string(Component)+"/" {
		return dependencyIsNotComponent
	}
	return nil
}

func (wg *WardleyGraph) FulfilledBy(need, dependency *sst.Node) (*sst.Link, error) {
	err := validFulfilledBy(need, dependency)
	if err != nil {
		return nil, err
	}
	return wg.sst.CreateLink(need, string(FulfilledBy), dependency, nil, 1)
}

func (wg *WardleyGraph) MustFulfilledBy(need, dependency *sst.Node) *sst.Link {
	link, err := wg.FulfilledBy(need, dependency)
	if err != nil {
		panic(err)
	}
	return link
}

func invalidContains(container, content string) error {
	return errors.New(fmt.Sprintf("wardleygraph: %v cannot contain %v", container, content))
}

func validContains(container, content *sst.Node) error {
	switch container.Prefix {
	case string(Component) + "/":
		if content.Prefix != string(Component)+"/" {
			return invalidContains(container.Prefix, content.Prefix)
		}
	case string(Map) + "/":
		if content.Prefix != string(Component)+"/" && content.Prefix != string(UserNeed)+"/" {
			return invalidContains(container.Prefix, content.Prefix)
		}
	default:
		return invalidContains(container.Prefix, content.Prefix)
	}
	return nil
}

func (wg *WardleyGraph) Contains(container, content *sst.Node) (*sst.Link, error) {
	err := validContains(container, content)
	if err != nil {
		return nil, err
	}
	return wg.sst.CreateLink(container, string(Contains), content, nil, 1)
}

func (wg *WardleyGraph) MustContains(container, content *sst.Node) *sst.Link {
	link, err := wg.Contains(container, content)
	if err != nil {
		panic(err)
	}
	return link
}

func (wg *WardleyGraph) ContainsAll(container *sst.Node, content []*sst.Node) ([]*sst.Link, error) {
	links := make([]*sst.Link, len(content))
	var err error
	for i, c := range content {
		links[i], err = wg.Contains(container, c)
		if err != nil {
			return nil, err
		}
	}
	return links, nil
}

func (wg *WardleyGraph) MustContainsAll(container *sst.Node, content []*sst.Node) []*sst.Link {
	links, err := wg.ContainsAll(container, content)
	if err != nil {
		panic(err)
	}
	return links
}

func (wg *WardleyGraph) ContainsWithData(container, content *sst.Node, data *ContainsLinkData) (*sst.Link, error) {
	err := validContains(container, content)
	if err != nil {
		return nil, err
	}
	var d map[string]interface{}
	err = mapstructure.Decode(data, &d)
	if err != nil {
		return nil, errors.Wrapf(err, "wardleygraph: failed to create Contains link from %v to %v", container.Key, content.Key)
	}
	return wg.sst.CreateLink(container, string(Contains), content, d, 1)
}

func (wg *WardleyGraph) MustContainsWithData(container, content *sst.Node, data *ContainsLinkData) *sst.Link {
	link, err := wg.ContainsWithData(container, content, data)
	if err != nil {
		panic(err)
	}
	return link
}

func validExpressCharacteristic(component, characteristic *sst.Node) error {
	if component.Prefix != string(Component)+"/" && component.Prefix != string(UserNeed) {
		return componentIsNotComponentOrUserNeed
	}
	if characteristic.Prefix != string(EvolutionCharacteristic)+"/" {
		return characteristicIsnt
	}
	return nil
}

func (wg *WardleyGraph) ExpressCharacteristic(component, characteristic *sst.Node) (*sst.Link, error) {
	err := validExpressCharacteristic(component, characteristic)
	if err != nil {
		return nil, err
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
		return nil, errors.Wrap(err, "wardleygraph: failed to query for related characteristics")
	}
	defer cursor.Close()
	for {
		var charstc sst.Node
		_, err := cursor.ReadDocument(context.TODO(), &charstc)
		if arango.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, errors.Wrap(err, "wardleygraph: failed to read existing characteristic")
		}
		err = wg.sst.DeleteLink(component, string(Expresses), &charstc, false)
		if err != nil {
			return nil, errors.Wrapf(err, "wardleygraph: failed to remove expression of existing characteristic: %v", charstc.Key)
		}
	}
	return wg.sst.CreateLink(component, string(Expresses), characteristic, nil, 1)
}

func (wg *WardleyGraph) MustExpressCharacteristic(component, characteristic *sst.Node) *sst.Link {
	link, err := wg.ExpressCharacteristic(component, characteristic)
	if err != nil {
		panic(err)
	}
	return link
}
