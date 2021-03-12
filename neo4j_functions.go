package neo4j_extended

import "errors"

//AddCreate adds a create Cypher
func (r *NeoRequest) AddCreate(n NeoNode) (err error) {
	cypherkey := "CREATE"
	cypher, err := n.toCypher(r)
	if err != nil {
		return err
	}
	cypher = cypherkey + cypher
	r.multiCypher = append(r.multiCypher, cypher)
	return nil
}

//AddReturn adds a return statemnet
func (r *NeoRequest) AddReturn(nodes *[]NeoNode) (err error) {
	if nodes == nil {
		return errors.New("@AddReturn: no return fields")
	}
	cypher := "RETURN"
	for _, node := range *nodes {
		if _, ok := r.names[node.Name]; !ok {
			return errors.New("@AddReturn: node is not in query ")
		}
		cypher += node.Name + ", "
	}
	cypher = trailingComma(cypher)
	r.multiCypher = append(r.multiCypher, cypher)
	return nil
}
