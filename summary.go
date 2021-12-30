package wardleygraph

import (
	"context"
	"fmt"

	arango "github.com/arangodb/go-driver"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/tristanls/sst"
)

type Summary struct {
	// Count of all components summarized, including nested ones
	All int
	// Count of components directly summarized
	Direct int
}

type summaryContainsModel struct {
	linkIDs []*string
}
type summaryLinkModel struct {
	// sst.Link's have already been allocated with node ids, etc, no need to copy memory
	// for the data that will be reused
	links  []*sst.Link
	weight float64
}
type summaryLinksModel struct {
	// constituents map indexes by component so that we can assign links
	// to the Contains link
	constituents map[string]*summaryContainsModel

	// the below maps are indexed by destination link, so that we can
	// create links outside of the summary component

	dependsOn   map[string]*summaryLinkModel
	expresses   map[string]*summaryLinkModel
	fulfilledBy map[string]*summaryLinkModel
}

type LinkSummary map[string]int

func encodeLinkName(name string, to, negate bool) string {
	if negate {
		if to {
			return "-t" + name
		}
		return "-f" + name
	} else {
		if to {
			return "+t" + name
		}
		return "+f" + name
	}
}

func decodeLinkName(name string) (string, bool, bool) {
	n := name[2:]
	if name[0:1] == "-" {
		if name[1:2] == "t" {
			return n, true, true
		}
		return n, false, true
	} else {
		if name[1:2] == "t" {
			return n, true, false
		}
		return n, false, false
	}
}

func (wg *WardleyGraph) CreateSummary(name string, constituents []*sst.Node) (*sst.Node, error) {
	sumData, sumLinks, err := wg.calculateSummary(name, constituents)
	if err != nil {
		return nil, err
	}

	fmt.Printf("SUMMARY DATA: %v\n", sumData)
	fmt.Printf("ALL OF THE LINKS:\n")
	for v, l := range sumLinks.dependsOn {
		fmt.Printf("%v : %v\n", v, *l)
	}
	for v, l := range sumLinks.expresses {
		fmt.Printf("%v : %v\n", v, *l)
	}
	for v, l := range sumLinks.fulfilledBy {
		fmt.Printf("%v : %v\n", v, *l)
	}

	var data map[string]interface{}
	err = mapstructure.Decode(sumData, &data)
	if err != nil {
		return nil, errors.Wrap(err, "wardleygraph: failed to translate struct to map")
	}

	// FIXME: Need to distinguish between creating a new component (to which we need to link all the things)
	//        and recreating an existing summary component (maybe track a hash of all constituent IDs?).
	//        Either way, to make summarizing idempotent, need symmetry breaking at this level to know
	//        if we should be adding links, if we should be adding link weights, or if we should leave
	//        links alone
	// 🤔 maybe the Contains link can include ids of summarized links in its metadata?? Yes, it should.

	// FIXME: should be a transaction
	sum, err := wg.ComponentWithData(name, data)
	if err != nil {
		return nil, errors.Wrap(err, "wardleygraph: failed to create the summary component")
	}

	// First, generate all the summary links in order to gather IDs of links
	// to be included as satellite data in the Contains link between the summary component
	// and its constituent. The satellite data is a map[string]int in order to
	// pass reflect.DeepEquals test (array order comes into play otherwise),
	// and int so that 0 (false) and 1 (true) are serialized as one character in JSON
	for k, v := range sumLinks.dependsOn {
		other, isTo, negate := decodeLinkName(k)
		var fromID, toID string
		if isTo {
			fromID = sst.MustNodeID(sum)
			toID = other
		} else {
			fromID = other
			toID = sst.MustNodeID(sum)
		}
		if negate {
			_, err = wg.sst.BlockLinkByID(fromID, string(DependsOn), toID, nil, v.weight)
		} else {
			_, err = wg.sst.CreateLinkByID(fromID, string(DependsOn), toID, nil, v.weight)
		}
		if err != nil {
			return nil, errors.Wrapf(err, "wardleygraph: failed to create DependsOn link from %v to %v", fromID, toID)
		}
	}
	for k, v := range sumLinks.expresses {
		other, _, negate := decodeLinkName(k)
		if negate {
			_, err = wg.sst.BlockLinkByID(sst.MustNodeID(sum), string(Expresses), other, nil, v.weight)
		} else {
			_, err = wg.sst.CreateLinkByID(sst.MustNodeID(sum), string(Expresses), other, nil, v.weight)
		}
		if err != nil {
			return nil, errors.Wrapf(err, "wardleygraph: failed to create Expresses link from %v to %v", sst.MustNodeID(sum), other)
		}
	}
	for k, v := range sumLinks.fulfilledBy {
		other, _, negate := decodeLinkName(k)
		if negate {
			_, err = wg.sst.BlockLinkByID(other, string(FulfilledBy), sst.MustNodeID(sum), nil, v.weight)
		} else {
			_, err = wg.sst.CreateLinkByID(other, string(FulfilledBy), sst.MustNodeID(sum), nil, v.weight)
		}
		if err != nil {
			return nil, errors.Wrapf(err, "wardleygraph: failed to create FulfilledBy link from %v to %v", other, sst.MustNodeID(sum))
		}
	}

	// Mark all the constituents as being contained by the summary component
	// Include IDs of links that were incorporated into the summary component to faciliate future updates
	for _, c := range constituents {
		var data *ContainsLinkData
		summary := sumLinks.constituents[sst.MustNodeID(c)]
		if summary != nil {
			linkSummary := LinkSummary{}
			for _, id := range summary.linkIDs {
				linkSummary[*id] = 1
			}
			data = &ContainsLinkData{
				LinkSummary: linkSummary,
			}
		}
		fmt.Printf("Contains key: %v, data: %v\n", sst.MustNodeID(c), data)
		_, err := wg.Contains(sum, c, data)
		if err != nil {
			return nil, errors.Wrapf(err, "wardleygraph: failed to create Contains link from summary %v to constituent component %v", sum.Key, c.Key)
		}
	}

	return sum, nil
}

func (wg *WardleyGraph) MustCreateSummary(name string, constituents []*sst.Node) *sst.Node {
	sum, err := wg.CreateSummary(name, constituents)
	if err != nil {
		panic(err)
	}
	return sum
}

func (wg *WardleyGraph) calculateSummary(name string, constituents []*sst.Node) (*NodeData, *summaryLinksModel, error) {
	// Aggregate constituent specific aggregates to place into the summary component
	sumData := NodeData{
		Summary: Summary{
			All:    len(constituents),
			Direct: len(constituents),
		},
	}
	constituentIDs := []string{}
	constituentMap := make(map[string]*sst.Node)
	for _, c := range constituents {
		if c.Prefix != string(Component)+"/" {
			return nil, nil, errors.New(fmt.Sprintf("wardleygraph: cannot summarize non-component: %v", c.Key))
		}
		if c.Data != nil {
			var data NodeData
			err := mapstructure.Decode(c.Data, &data)
			if err != nil {
				return nil, nil, errors.Wrapf(err, "wardleygraph: failed to decode component data: %v", c.Key)
			}
			sumData.Summary.All += data.Summary.All
		}
		id := sst.MustNodeID(c)
		constituentIDs = append(constituentIDs, id)
		constituentMap[id] = c
	}

	// Find all the links from the constituents reaching beyond the summary component
	// and place them onto the summary component itself. Link weight is the sum of duplicates.
	// Which links have been incorporated is tracked on the Contains link between the
	// summary component and the constituent.
	// - Component Expresses EvolutionaryCharacteristic (OUTBOUND only)
	// - Component DependsOn Component (ANY direction)
	// - Component -FulfilledBy UserNeed (INBOUND only)
	qs := `
	FOR comp IN @constituents
		FOR v, l, p
			IN 1
			OUTBOUND comp
			Expresses, ANY Follows
			FILTER (p.vertices[1].Prefix == @charstc_prefix)
				OR (p.vertices[1].Prefix == @comp_prefix)
				OR (p.vertices[1].Prefix == @un_prefix)
			FILTER p.vertices[1]._id NOT IN @constituents
			RETURN l
	`
	vars := map[string]interface{}{
		"charstc_prefix": EvolutionCharacteristic + "/",
		"comp_prefix":    Component + "/",
		"constituents":   constituentIDs,
		"un_prefix":      UserNeed + "/",
	}
	cursor, err := wg.sst.Query(context.TODO(), qs, vars)
	if err != nil {
		return nil, nil, errors.Wrap(err, "wardleygraph: failed to query for constituent component characteristics")
	}
	defer cursor.Close()
	sumLinks := summaryLinksModel{
		constituents: make(map[string]*summaryContainsModel),
		dependsOn:    make(map[string]*summaryLinkModel),
		expresses:    make(map[string]*summaryLinkModel),
		fulfilledBy:  make(map[string]*summaryLinkModel),
	}
	for {
		var link sst.Link
		_, err := cursor.ReadDocument(context.TODO(), &link)
		if arango.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, nil, errors.Wrapf(err, "wardleygraph: failed to read existing link to component characteristic")
		}
		negate := sst.MustLinkKeyNegated(link.Key)
		linkID, err := wg.sst.LinkID(&link)
		if err != nil {
			return nil, nil, err
		}
		switch link.SID {
		case string(DependsOn): // ANY direction links
			// query does not return links between constituents, so if one end is a constituent, the other is not, and vice versa

			var name, comp string
			if constituentMap[link.From] == nil {
				name = encodeLinkName(link.From, false, negate)
				comp = link.To
			} else {
				name = encodeLinkName(link.To, true, negate)
				comp = link.From
			}
			dependsOn := sumLinks.dependsOn[name]
			if dependsOn == nil {
				sumLinks.dependsOn[name] = &summaryLinkModel{
					links:  []*sst.Link{&link},
					weight: link.Weight,
				}
			} else {
				dependsOn.links = append(dependsOn.links, &link)
				dependsOn.weight += link.Weight
			}
			constituents := sumLinks.constituents[comp]
			if constituents == nil {
				sumLinks.constituents[comp] = &summaryContainsModel{
					linkIDs: []*string{&linkID},
				}
			} else {
				constituents.linkIDs = append(constituents.linkIDs, &linkID)
			}
		case string(Expresses): // OUTBOUND links only
			name := encodeLinkName(link.To, true, negate)
			comp := link.From
			expresses := sumLinks.expresses[name]
			if expresses == nil {
				sumLinks.expresses[name] = &summaryLinkModel{
					links:  []*sst.Link{&link},
					weight: link.Weight,
				}
			} else {
				expresses.links = append(expresses.links, &link)
				expresses.weight += link.Weight
			}
			constituents := sumLinks.constituents[comp]
			if constituents == nil {
				sumLinks.constituents[comp] = &summaryContainsModel{
					linkIDs: []*string{&linkID},
				}
			} else {
				constituents.linkIDs = append(constituents.linkIDs, &linkID)
			}
		case string(FulfilledBy): // INBOUND links only
			name := encodeLinkName(link.From, false, negate)
			comp := link.To
			fulfilledBy := sumLinks.fulfilledBy[name]
			if fulfilledBy == nil {
				sumLinks.fulfilledBy[name] = &summaryLinkModel{
					links:  []*sst.Link{&link},
					weight: link.Weight,
				}
			} else {
				fulfilledBy.links = append(fulfilledBy.links, &link)
				fulfilledBy.weight += link.Weight
			}
			constituents := sumLinks.constituents[comp]
			if constituents == nil {
				sumLinks.constituents[comp] = &summaryContainsModel{
					linkIDs: []*string{&linkID},
				}
			} else {
				constituents.linkIDs = append(constituents.linkIDs, &linkID)
			}
		default:
			return nil, nil, errors.New(fmt.Sprintf("wardleygraph: unexpected association type: %v", link.SID))
		}
	}

	return &sumData, &sumLinks, nil
}
