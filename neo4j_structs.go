package neo4j_extended

import "errors"

//NeoRequest helper struct for building neo4j Requests
type NeoRequest struct {
	multiCypher []string
	params      map[string]interface{}
	names       map[string]struct{}
}

//NewNeoRequest default constructor
func NewNeoRequest() NeoRequest {
	neoReq := NeoRequest{
		multiCypher: []string{},
		params:      make(map[string]interface{}),
		names:       make(map[string]struct{}),
	}
	return neoReq
}

//NeoField name value parameter tuple
type NeoField struct {
	Name string
	Val  interface{}
}

//NeoRelation relation
type NeoRelation struct {
	Name      string
	Label     string
	Fields    []NeoField
	NextNode  *NeoNode
	Direction int
}

//NeoNode name label variable tuple
type NeoNode struct {
	Name    string
	Label   string
	Fields  []NeoField
	NextRel *NeoRelation
}

//toCypher returns the cypher representation of the node-rel LinkedList and adds params to req
func (node *NeoNode) toCypher(req *NeoRequest) (cypher string, err error) {
	cypher = "(" + node.Name + ":" + node.Label + " "
	fieldsstring, err := getFieldsCypher(req, node.Fields)
	if err != nil {
		return "", err
	}
	cypher += fieldsstring + ")"

	if rel := node.NextRel; rel != nil {
		beginArrow := "-"
		endArrow := "->"
		if rel.Direction == -1 {
			beginArrow = "<-"
			endArrow = "-"
		}
		cypher += beginArrow + "[" + rel.Name + ":" + rel.Label
		fieldsstring, err := getFieldsCypher(req, rel.Fields)
		if err != nil {
			return "", err
		}
		cypher += fieldsstring + "]" + endArrow
		if rel.NextNode == nil {
			return "", errors.New("@toCpyher: relation points to nil")
		}
		nextNodeCypher, err := rel.NextNode.toCypher(req)
		if err != nil {
			return "", err
		}
		cypher += nextNodeCypher
	}
	return cypher, nil
}

//NewNeoNode creats the Node and does some checks if it is valid
func (req *NeoRequest) NewNeoNode(name string, label string, fields []NeoField) (node NeoNode, err error) {
	if _, ok := req.names[name]; ok {
		return NeoNode{}, errors.New("@NewNeoNode: name is already declared")
	}
	node = NeoNode{
		Name:    name,
		Label:   label,
		Fields:  fields,
		NextRel: nil,
	}
	return node, nil
}

//AddRelation NeoNode attaches a Node to a relation
func (node *NeoNode) AddRelation(name string, label string, fields []NeoField, direction int, req *NeoRequest) (rel *NeoRelation, err error) {
	if _, ok := req.names[name]; ok {
		return &NeoRelation{}, errors.New("@AddRelation: name is already declared")
	}
	node.NextRel = &NeoRelation{
		Name:      name,
		Label:     label,
		Fields:    fields,
		Direction: direction,
		NextNode:  nil,
	}
	return node.NextRel, nil
}

//AddNode attaches a Node to a relation
func (rel *NeoRelation) AddNode(name string, label string, fields []NeoField, req *NeoRequest) (node *NeoNode, err error) {
	if _, ok := req.names[name]; ok {
		return &NeoNode{}, errors.New("@AddNode: name is already declared")
	}
	rel.NextNode = &NeoNode{
		Name:    name,
		Label:   label,
		Fields:  fields,
		NextRel: nil,
	}
	return rel.NextNode, nil
}
