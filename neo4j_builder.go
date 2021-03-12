package neo4j_extended

import "errors"

//addFields adds Fields to cypher string
func getFieldsCypher(r *NeoRequest, fields *[]NeoField) (out string, err error) {
	if fields == nil {
		return "", nil
	}
	out = "{"
	for _, field := range *fields {
		//Default supported types
		switch val := field.Val.(type) {
		case int:
			field.Val = val
			if val != 0 {
				out += field.Name + ": $" + field.Name + ", "
				r.addParam(field)
			}
		case string:
			field.Val = val
			if val != "" {
				out += field.Name + ": $" + field.Name + ", "
				r.addParam(field)
			}
		case float64:
			field.Val = val
			if val != 0 {
				out += field.Name + ": $" + field.Name + ", "
				r.addParam(field)
			}
		case []string:
			field.Val = val
			if len(val) > 0 {
				out += field.Name + ": $" + field.Name + ", "
				r.addParam(field)
			}
		default:
			return "", errors.New("@AddCreate: wrong type")
		}
	}
	out = trailingComma(out)
	out += "}"
	return out, nil
}

//add adds a cypher line
func (n *NeoRequest) add(cypher string, param NeoField) {
	n.addCypher(cypher)
	n.addParam(param)
}

//addCypher adds a cypher
func (n *NeoRequest) addCypher(cypher string) {
	n.multiCypher = append(n.multiCypher, cypher)
}

//addCypherAll adds multiple cypher
func (n *NeoRequest) addCypherAll(multicypher []string) {
	n.multiCypher = append(n.multiCypher, multicypher...)
}

//addParam adds a prameter
func (n *NeoRequest) addParam(param NeoField) {
	n.params[param.Name] = param.Val
}

//addParamAll adds multiple params
func (n *NeoRequest) addParamAll(params []NeoField) {
	for _, param := range params {
		n.addParam(param)
	}
}

//trailingComma removes ", "
func trailingComma(s string) string {
	if s[len(s)-2:] == ", " {
		s = s[:len(s)-2]
	}
	return s

}
