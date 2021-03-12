package neo4j_extended

//AddCreate adds a create Cypher
func (r *NeoRequest) AddCreate(n NeoNode) (err error) {
	cypher := "CREATE (" + n.Name + ":" + n.Label + " "
	fieldsString, err := getFieldsCypher(r, n.Fields)
	if err != nil {
		return err
	}
	cypher += fieldsString
	cypher += ")"
	r.multiCypher = append(r.multiCypher, cypher)
	return nil
}
