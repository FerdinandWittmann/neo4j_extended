package neo4j_extended

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var NDriver *neo4j.Driver

//SendSimple takes a singler cypher instruction and returns a single value
func SendSimple(session neo4j.Session, cypher string) (res interface{}, err error) {
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
func (n *NeoRequest) Send(session neo4j.Session) (result neo4j.Result, err error) {
	cypherList := append(n.multiCypher)
	cypher := MultiCypherToCypher(cypherList)
	fmt.Println(cypher)
	result, err = session.Run(cypher, n.params)
	return result, err
}

//MultiCypherToCypher takes a cyper reduces it to one
func MultiCypherToCypher(multiCypher []string) (cypher string) {
	for _, singleCypher := range multiCypher {
		cypher = cypher + singleCypher + " "
	}
	return cypher
}

//CreateSession creates neo3j Session from driver
func CreateSession(accessMode neo4j.AccessMode) (session neo4j.Session, err error) {
	driver := (*NDriver)
	session, err = driver.NewSession(neo4j.SessionConfig{
		AccessMode:   accessMode,
		DatabaseName: "coli",
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}

//SetDriver sets the driver variable should be locally initialized
func SetDriver(_driver *neo4j.Driver) {
	NDriver = _driver
}
