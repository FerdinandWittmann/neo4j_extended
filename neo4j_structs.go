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
	name string
	val  interface{}
}

//NeoNode name label variable tuple
type NeoNode struct {
	name     string
	label    string
	fields   []NeoField
	relation *NeoRelation
}

//NeoRelation relation
type NeoRelation struct {
	name   string
	label  string
	fields []NeoField
	node   *NeoNode
}
