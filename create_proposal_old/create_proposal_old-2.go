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

// Proposal example simple Chaincode implementation
type ManageProposal struct {
}

var approved_proposal_entry = "approved_proposal_entry"				//name for the key/value that will store a list of all known  Tier3 Form

type proposal struct{
								// Attributes of a Form 
	proposal_id string `json:"proposal_id"`	
	region string `json:"region"`
	country string `json:"country"`
	proposal_type string `json:"proposal_type"`
	
	proposal_date string `json:"proposal_date"`
	approval_date string `json:"approval_date"`
	shared_with_procurement_team_on string `json:"shared_with_procurement_team_on"`
	
	approver string `json:"approver"`
	number_of_tasks_covered string `json:"number_of_tasks_covered"`
	device_qty string `json:"device_qty"`
	accessary_periperal_qty string `json:"accessary_periperal_qty"`
	total_qty string `json:"total_qty"`
	status string `json:"status"`
	
	
}
// ============================================================================================================================
// Main - start the chaincode for Form management
// ============================================================================================================================
func main() {			
	err := shim.Start(new(ManageProposal))
	if err != nil {
		fmt.Printf("Error starting Form management chaincode: %s", err)
	}
}
// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *ManageProposal) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var msg string
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	// Initialize the chaincode
	msg = args[0]
	fmt.Println("ManageProposal chaincode is deployed successfully.");
	
	// Write the state to the ledger
	err = stub.PutState("abc", []byte(msg))	//making a test var "abc", I find it handy to read/write to it right away to test the network
	if err != nil {
		return nil, err
	}
	var proposal_form_empty []string
	proposal_form_empty_json_as_bytes, _ := json.Marshal(proposal_form_empty)								//marshal an emtpy array of strings to clear the index
	err = stub.PutState(approved_proposal_entry, proposal_form_empty_json_as_bytes)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
// ============================================================================================================================
// Run - Our entry Formint for Invocations - [LEGACY] obc-peer 4/25/2016
// ============================================================================================================================
func (t *ManageProposal) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)
	return t.Invoke(stub, function, args)
}
// ============================================================================================================================
// Invoke - Our entry Formint for Invocations
// ============================================================================================================================
func (t *ManageProposal) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "create_proposal_id" {											//create a new Form
		return t.create_proposal_id(stub, args)
	} 
	fmt.Println("invoke did not find func: " + function)	
	jsonResp := "Error : Received unknown function invocation: "+ function 				//error
	return nil, errors.New(jsonResp)
}

// ============================================================================================================================
// Query - Our entry for Queries
// ============================================================================================================================
func (t *ManageProposal) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query is running " + function)

	// Handle different functions
	if function == "get_all_proposal_data" {													//Read all Forms
		return t.get_all_proposal_data(stub, args)
	} else if function == "get_all_proposal_id" {													//Read all Forms
		return t.get_all_proposal_id(stub, args)
	} 

	fmt.Println("query did not find func: " + function)				//error
	jsonResp := "Error : Received unknown function query: "+ function 
	return nil, errors.New(jsonResp)
}


// ============================================================================================================================
// create Form - create a new Form for proposal id, store into chaincode state
// ============================================================================================================================
func (t *ManageProposal) create_proposal_id(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	if len(args) != 13 {
		return nil, errors.New("Incorrect number of arguments. Expecting 9")
	}
	fmt.Println("Creating a new Form for proposal id ")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	
	
	proposal_id := args[0]
	region := args[1]
	country := args[2]
	proposal_type := args[3]
	
	proposal_date := args[4]
	approval_date := args[5]
	shared_with_procurement_team_on := args[6]
	
	
	
	approver := args[7]
	number_of_tasks_covered := args[8]
	device_qty := args[9]
	accessary_periperal_qty := args[10]
	total_qty := args[11]
	status := args[12]
	
		
	//build the Form json string manually
	input := 	`{`+
		`"proposal_id": "` + proposal_id + `" , `+
		`"region": "` + region + `" , `+ 
		`"country": "` + country + `"`+
	        `"proposal_type": "` + proposal_type + `" , `+ 
	
	
	        `"proposal_date": "` + proposal_date + `" , `+ 
		`"approval_date": "` + approval_date + `" , `+ 
		`"shared_with_procurement_team_on": "` + shared_with_procurement_team_on + `" , `+ 
	
		`"approver": "` + approver + `" , `+ 
		`"number_of_tasks_covered": "` + number_of_tasks_covered + `" , `+ 
		`"device_qty": "` + device_qty + `" , `+ 
		`"accessary_periperal_qty": "` + accessary_periperal_qty + `" , `+ 
		`"total_qty": "` + total_qty + `" , `+ 
		`"status": "` + status + `"` +	
	
		`}`
	
		fmt.Println("input: " + input)
		fmt.Print("input in bytes array: ")
		fmt.Println([]byte(input))
	err = stub.PutState(proposal_id, []byte(input))									//store Form with FAA_formNumber as key
	if err != nil {
		return nil, err
	}
	

	
	proposal_id_FormIndexAsBytes, err := stub.GetState(approved_proposal_entry)
	if err != nil {
		return nil, errors.New("Failed to get proposal id  Form index")
	}
	var proposal_id_FormIndex []string
	fmt.Print("proposal_id_FormIndexAsBytes: ")
	fmt.Println(proposal_id_FormIndexAsBytes)
	
	json.Unmarshal(proposal_id_FormIndexAsBytes, &proposal_id_FormIndex)							//un stringify it aka JSON.parse()
	fmt.Print("proposal_id_FormIndex after unmarshal..before append: ")
	fmt.Println(proposal_id_FormIndex)
	//append
	proposal_id_FormIndex = append(proposal_id_FormIndex, proposal_id)									//add Form transID to index list
	fmt.Println("! Proposal  Form index after appending proposal id: ", proposal_id_FormIndex)
	jsonAsBytes, _ := json.Marshal(proposal_id_FormIndex)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(approved_proposal_entry, jsonAsBytes)						//store name of Form
	if err != nil {
		return nil, err
	}

	fmt.Println("Proposal Form created successfully.")
	return nil, nil
	
	
}



func (t *ManageProposal) get_all_proposal_data(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	var jsonProposalResp,errResp string
	var proposal_id_FormIndex []string
	fmt.Println("Fetching All Proposals")
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting single space as an argument")
	}
	// fetching all Proposal
	proposal_id_FormIndexAsBytes, err := stub.GetState(approved_proposal_entry)
	if err != nil {
		return nil, errors.New("Failed to get all Proposals")
	}
	fmt.Print("proposal_id_FormIndexAsBytes : ")
	fmt.Println(proposal_id_FormIndexAsBytes)
	json.Unmarshal(proposal_id_FormIndexAsBytes, &proposal_id_FormIndex)								//un stringify it aka JSON.parse()
	fmt.Print("proposal_id_FormIndex : ")
	fmt.Println(proposal_id_FormIndex)
	// Proposal Data
	jsonProposalResp = "{"
	for i,val := range proposal_id_FormIndex{
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
	fmt.Println("len(proposal_id_FormIndex) : ")
	fmt.Println(len(proposal_id_FormIndex))

	jsonProposalResp = jsonProposalResp + "}"
	fmt.Println([]byte(jsonProposalResp))
	fmt.Println("Fetched All Proposal Data successfully.")
	return []byte(jsonProposalResp), nil
}



func (t *ManageProposal) get_all_proposal_id(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	var jsonProposalResp,errResp string
	var proposal_id_FormIndex []string
	fmt.Println("Fetching All Proposals")
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting single space as an argument")
	}
	// fetching all Proposal
	proposal_id_FormIndexAsBytes, err := stub.GetState(approved_proposal_entry)
	if err != nil {
		return nil, errors.New("Failed to get all Proposals")
	}
	fmt.Print("proposal_id_FormIndexAsBytes : ")
	fmt.Println(proposal_id_FormIndexAsBytes)
	json.Unmarshal(proposal_id_FormIndexAsBytes, &proposal_id_FormIndex)								//un stringify it aka JSON.parse()
	fmt.Print("proposal_id_FormIndex : ")
	fmt.Println(proposal_id_FormIndex)
	// Proposal Data
	jsonProposalResp = "{ "
	for i,val := range proposal_id_FormIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Proposal")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return nil, errors.New(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonProposalResp = jsonProposalResp + "\""+ val + "\""
		if i < len(proposal_id_FormIndex)-1 {
			jsonProposalResp = jsonProposalResp + ","
		}
	}
	fmt.Println("len(proposal_id_FormIndex) : ")
	fmt.Println(len(proposal_id_FormIndex))

	jsonProposalResp = jsonProposalResp + "}"
	fmt.Println([]byte(jsonProposalResp))
	fmt.Println("Fetched All Proposal ID successfully.")
	return []byte(jsonProposalResp), nil
}
