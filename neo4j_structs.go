package neo4j_extended

//NeoRequest helper struct for building neo4j Requests
type NeoRequest struct {
	multiCypher []string
	params      map[string]interface{}
}

//NewNeoRequest default constructor
func NewNeoRequest() NeoRequest {
	neoReq := NeoRequest{
		multiCypher: []string{},
		params:      make(map[string]interface{}),
	}
	return neoReq
}

//NeoField name value parameter tuple
type NeoField struct {
	Name string
	Val  interface{}
}

//NeoNode name label variable tuple
type NeoNode struct {
	Name     string
	Label    string
	Fields   []NeoField
	Relation *NeoRelation
}

//NeoRelation relation
type NeoRelation struct {
	Name   string
	Label  string
	Fields []NeoField
	Node   *NeoNode
}
