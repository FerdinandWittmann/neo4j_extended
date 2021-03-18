package neo4j_extended

import (
	"errors"
	"strconv"
)

//nodeOp any operation that does work with nodes-rel system
func nodeOp(key string, node *NeoNode, req *NeoRequest) (err error) {
	cypherkey := key + " "
	cypher, err := node.toCypher(req)
	if err != nil {
		return err
	}
	cypher = cypherkey + cypher
	req.multiCypher = append(req.multiCypher, cypher)
	return nil
}

//addFields adds Fields to cypher string
func getFieldsCypher(r *NeoRequest, fields *[]NeoField) (out string, err error) {
	if fields == nil {
		return "", nil
	}
	out = "{"
	for _, field := range *fields {
		//Check if parameter with same Name already exists
		fieldName := field.Name
		_, ok := r.params[field.Name]
		i := 1
		for ok {
			field.Name = field.Name + strconv.Itoa(i)
			i += 1
			_, ok = r.params[field.Name]
		}
		//Default supported types
		switch val := field.Val.(type) {
		case int:
			field.Val = val
			if val != 0 {
				out += fieldName + ": $" + field.Name + ", "
				r.addParam(field)
			}
		case string:
			field.Val = val
			if val != "" {
				out += fieldName + ": $" + field.Name + ", "
				r.addParam(field)
			}
		case float64, float32:
			field.Val = val
			if val != 0 {
				out += fieldName + ": $" + field.Name + ", "
				r.addParam(field)
			}
		case []string:
			field.Val = val
			if len(val) > 0 {
				out += fieldName + ": $" + field.Name + ", "
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

//MultiCypherToCypher takes a cyper reduces it to one
func MultiCypherToCypher(multiCypher []string) (cypher string) {
	for _, singleCypher := range multiCypher {
		cypher = cypher + singleCypher + " "
	}
	return cypher
}

//trailingComma removes ", "
func trailingComma(s string) string {
	//TODO Edge Case > 2 Analyse
	if len(s) > 2 {
		if s[len(s)-2:] == ", " {
			s = s[:len(s)-2]
		}
		return s
	}
	return s

}
