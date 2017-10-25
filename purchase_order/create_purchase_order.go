/*/*
-a-Licensed to the Apache Software Foundation (ASF) under one
or more contributor license Forms.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at
  http://www.apache.org/licenses/LICENSE-2.0 
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
"errors"
"fmt"
"strconv"
"encoding/json"

"github.com/hyperledger/fabric/core/chaincode/shim"
//"github.com/hyperledger/fabric/core/util"
)

// Po order example simple Chaincode implementation
type Manage_po_order struct {
}

var approved_po_order_entry = "approved_po_order_entry"	
type po_order struct{
								// Attributes of a Form 
	sap_po_order string `json:"sap_po_order"`	
	supplier string `json:"supplier"`
	venderso string `json:"venderso"`
	
	
}





// ============================================================================================================================
// Main - start the chaincode for Form management
// ============================================================================================================================
func main() {			
	err := shim.Start(new(Manage_po_order))
	if err != nil {
		fmt.Printf("Error starting Form management of po order chaincode: %s", err)
	}
}





// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *Manage_po_order) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var msg string
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	// Initialize the chaincode
	msg = args[0]
	fmt.Println("Manage Po Order chaincode is deployed successfully.");
	
	// Write the state to the ledger
	err = stub.PutState("abc", []byte(msg))	//making a test var "abc", I find it handy to read/write to it right away to test the network
	if err != nil {
		return nil, err
	}
	var po_order_form_empty []string
	po_order_form_empty_json_as_bytes, _ := json.Marshal(po_order_form_empty)								//marshal an emtpy array of strings to clear the index
	err = stub.PutState(approved_po_order_entry, po_order_form_empty_json_as_bytes)
	if err != nil {
		return nil, err
	}
	return nil, nil
}




// ============================================================================================================================
// Run - Our entry Formint for Invocations - [LEGACY] obc-peer 4/25/2016
// ============================================================================================================================
func (t *Manage_po_order) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)
	return t.Invoke(stub, function, args)
}




// ============================================================================================================================
// Invoke - Our entry Formint for Invocations
// ============================================================================================================================
func (t *Manage_po_order) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "create_po_order_id" {											//create a new Form
		return t.create_po_order_id(stub, args)
	} 
	fmt.Println("invoke did not find func: " + function)	
	jsonResp := "Error : Received unknown function invocation: "+ function 				//error
	return nil, errors.New(jsonResp)
}



// ============================================================================================================================
// Query - Our entry for Queries
// ============================================================================================================================
func (t *Manage_po_order) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query is running " + function)

	// Handle different functions
	if function == "get_all_po_order" {													//Read all Forms
		return t.get_all_po_order(stub, args)
	} 

	fmt.Println("query did not find func: " + function)				//error
	jsonResp := "Error : Received unknown function query: "+ function 
	return nil, errors.New(jsonResp)
}



func (t *Manage_po_order) create_po_order_id(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 9")
	}
	fmt.Println("Creating a new Form for po order id ")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	
	
	sap_po_order := args[0]
	supplier := args[1]
	venderso := args[2]
		
	
		
	//build the Form json string manually
	input := 	`{`+
		`"sap_po_order": "` + sap_po_order + `" , `+
		`"supplier": "` + supplier + `" , `+ 
		`"venderso": "` + venderso + `"`+
		`}`
		fmt.Println("input: " + input)
		fmt.Print("input in bytes array: ")
		fmt.Println([]byte(input))
	err = stub.PutState(sap_po_order, []byte(input))									//store Form with FAA_formNumber as key
	if err != nil {
		return nil, err
	}
	

	
	po_order_id_FormIndexAsBytes, err := stub.GetState(approved_po_order_entry)
	if err != nil {
		return nil, errors.New("Failed to get po order id  Form index")
	}
	var po_order_id_FormIndex []string
	fmt.Print("po_order_id_FormIndexAsBytes: ")
	fmt.Println(po_order_id_FormIndexAsBytes)
	
	json.Unmarshal(po_order_id_FormIndexAsBytes, &po_order_id_FormIndex)							//un stringify it aka JSON.parse()
	fmt.Print("po_order_id_FormIndex after unmarshal..before append: ")
	fmt.Println(po_order_id_FormIndex)
	//append
	po_order_id_FormIndex = append(po_order_id_FormIndex, sap_po_order)									//add Form transID to index list
	fmt.Println("! po order  Form index after appending po order id: ", po_order_id_FormIndex)
	jsonAsBytes, _ := json.Marshal(po_order_id_FormIndex)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(approved_po_order_entry, jsonAsBytes)						//store name of Form
	if err != nil {
		return nil, err
	}

	fmt.Println("Po order  created successfully.")
	return nil, nil
	
	
}



func (t *Manage_po_order) get_all_po_order(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	var jsonProposalResp,errResp string
	var po_order_id_FormIndex []string
	fmt.Println("Fetching All po order")
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting single space as an argument")
	}
	// fetching all po order
	po_order_id_FormIndexAsBytes, err := stub.GetState(approved_po_order_entry)
	if err != nil {
		return nil, errors.New("Failed to get all po order")
	}
	fmt.Print("po_order_id_FormIndexAsBytes : ")
	fmt.Println(po_order_id_FormIndexAsBytes)
	json.Unmarshal(po_order_id_FormIndexAsBytes, &po_order_id_FormIndex)								//un stringify it aka JSON.parse()
	fmt.Print("po_order_id_FormIndex : ")
	fmt.Println(po_order_id_FormIndex)
	// Proposal Data
	jsonProposalResp = "{"
	for i,val := range po_order_id_FormIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Proposal")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return nil, errors.New(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonProposalResp = jsonProposalResp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(proposal_id_FormIndex)-1 {
			jsonProposalResp = jsonProposalResp + ","
		}
	}
	fmt.Println("len(po_order_id_FormIndex) : ")
	fmt.Println(len(po_order_id_FormIndex))

	jsonProposalResp = jsonProposalResp + "}"
	fmt.Println([]byte(jsonProposalResp))
	fmt.Println("Fetched All proposal Forms successfully.")
	return []byte(jsonProposalResp), nil
}
