package neo4j_extended

import (
	"encoding/json"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var NDriver *neo4j.Driver

//SetDriver sets the driver variable should be locally initialized
func SetDriver(_driver *neo4j.Driver) {
	NDriver = _driver
}

//CreateSession creates neo3j Session from driver
func CreateSession(accessMode neo4j.AccessMode) (session *neo4j.Session, err error) {
	driver := (*NDriver)
	_session, err := driver.NewSession(neo4j.SessionConfig{
		AccessMode:   accessMode,
		DatabaseName: "coli",
	})
	if err != nil {
		return nil, err
	}
	session = (&_session)
	return session, nil
}

//SendSimple takes a singler cypher instruction and returns a single value
func SendSimple(session neo4j.Session, cypher string) (res interface{}, err error) {
	fmt.Println(cypher)
	result, err := session.Run(cypher, nil)
	if err != nil {
		return nil, err
	}
	if result.Next() {
		return result.Record().Values()[0], nil
	}
	return nil, result.Err()
}

//Send Sends a new Cypher
func (n *NeoRequest) SendNew(accessMode neo4j.AccessMode) (result *neo4j.Result, err error) {
	session, err := CreateSession(accessMode)
	defer (*session).Close()
	if err != nil {
		return nil, err
	}
	cypherList := append(n.multiCypher)
	cypher := MultiCypherToCypher(cypherList)
	//fmt.Println(cypher)
	_result, err := (*session).Run(cypher, n.params)
	if err != nil {
		return nil, err
	}
	if err = _result.Err(); err != nil {
		return nil, err
	}
	result = &_result
	return result, err
}

//Send Sends a new Cypher
func (n *NeoRequest) Send(session *neo4j.Session) (result *neo4j.Result, err error) {
	cypherList := append(n.multiCypher)
	cypher := MultiCypherToCypher(cypherList)
	//fmt.Println(cypher)
	_result, err := (*session).Run(cypher, n.params)
	if err != nil {
		return nil, err
	}
	result = (&_result)
	return result, err
}

//PrettyPrintValues prints the map of a neo4j Node result takes as input rec.values()
func PrettyPrintValues(res *neo4j.Result) {
	for (*res).Next() {
		rec := (*res).Record()
		for _, value := range rec.Keys() {
			_values, ok := rec.Get(value)
			if !ok {
				fmt.Println("what")
			}
			switch o := _values.(type) {
			case neo4j.Node:
				values_map := o.Props()
				b, err := json.MarshalIndent(values_map, "", " ")
				if err != nil {
					fmt.Println("@PrettyPrintValues: ", err)
				}
				fmt.Print(string(b))
			case neo4j.Relationship:
				values_map := o.Props()
				b, err := json.MarshalIndent(values_map, "", " ")
				if err != nil {
					fmt.Println("@PrettyPrintValues: ", err)
				}
				fmt.Print(string(b))
			}
		}
	}

}
