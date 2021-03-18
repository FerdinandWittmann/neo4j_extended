package neo4j_extended

import (
	"errors"
	"fmt"
)

//AddMerge adds a create Cypher
func (r *NeoRequest) AddMerge(n *NeoNode) (err error) {
	err = nodeOp("MERGE", n, r)
	if err != nil {
		return err
	}
	return nil
}

//AddCreate adds a create Cypher
func (r *NeoRequest) AddCreate(n *NeoNode) (err error) {
	err = nodeOp("CREATE", n, r)
	if err != nil {
		return err
	}
	return nil
}

//AddMatch adds a create Cypher
func (r *NeoRequest) AddMatch(n *NeoNode) (err error) {
	err = nodeOp("MATCH", n, r)
	if err != nil {
		return err
	}
	return nil
}

//CheckName checks if a name exisits and returns one that doesnt, always new extended with last character recursive n times
func (r *NeoRequest) CheckName(name string) (availableName string) {
	_, ok := r.names[name]
	if !ok {
		return name
	} else {
		return r.CheckName(name + name[len(name)-1:])
	}
}

//SaveReturn saves a node to be returned
func (r *NeoRequest) SaveReturn(obj interface{}) {
	switch o := obj.(type) {
	case *NeoNode:
		(*r).returns = append((*r).returns, (*o).Name)
	case *NeoRelation:
		(*r).returns = append((*r).returns, (*o).Name)
	default:
		fmt.Println("@SaveReturn: InvalidType")
	}
}

//AddReturnSimple takes list of node names
func (r *NeoRequest) AddReturnSimple(list []string) (err error) {
	if len(list) == 0 {
		return errors.New("@AddReturnAll: Nothing to return")
	}
	cypher := "RETURN "
	for _, val := range list {
		cypher += val + ", "
		if _, ok := r.names[val]; !ok {
			return errors.New("@AddReturn: node is not in query ")
		}
	}
	cypher = trailingComma(cypher)
	r.multiCypher = append(r.multiCypher, cypher)
	return nil
}

//AddReturnAll returns all specificied relations and nodes
func (r *NeoRequest) AddReturnAll() (err error) {
	if len(r.names) == 0 {
		return errors.New("@AddReturnAll: Nothing to return")
	}
	cypher := "RETURN "
	for key, _ := range r.names {
		cypher += key + ", "
	}
	cypher = trailingComma(cypher)
	r.multiCypher = append(r.multiCypher, cypher)
	return nil
}

//AddReturns adds return statement with all saved returns
func (r *NeoRequest) AddReturns() (err error) {
	if len((*r).returns) == 0 {
		return errors.New("@AddReturn: no return fields")
	}
	cypher := "RETURN "
	for _, name := range (*r).returns {
		if _, ok := r.names[name]; !ok {
			return errors.New("@AddReturn: node is not in query ")
		}
		cypher += name + ", "
	}
	cypher = trailingComma(cypher)
	r.multiCypher = append(r.multiCypher, cypher)
	return nil
}

//AddReturn adds a return statemnet
func (r *NeoRequest) AddReturn(nodes *[]*NeoNode) (err error) {
	if nodes == nil {
		return errors.New("@AddReturn: no return fields")
	}
	cypher := "RETURN "
	for _, node := range *nodes {
		if node == nil {
			return errors.New("@AddReturn: nil pointer")
		}
		if _, ok := r.names[(*node).Name]; !ok {
			return errors.New("@AddReturn: node is not in query ")
		}
		cypher += (*node).Name + ", "
	}
	cypher = trailingComma(cypher)
	r.multiCypher = append(r.multiCypher, cypher)
	return nil
}
