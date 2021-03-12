package neo4j_extended

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
