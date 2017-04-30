// ============================================================================================================================
// 本智能合约用于TX管理
// 功能包括：TX生成、查询，状态变更
// ============================================================================================================================

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

type SimpleChaincode struct {
}

// ============================================================================================================================
// TXInfo struct
// ============================================================================================================================
type TXInfoStruct struct {
	JobID       string
	UserID      string
	ApplyTime   string
	TxID        string
	Status      string
	StuScore    string
	AgencyScore string
}

func (t *SimpleChaincode) GetJobChaincodeToCall() string {
	chainCodeToCall := "59d5075e7146d8df37bdcb0c289c293c66ee2a56c0f8d833410ff7cd8d3dd6c65ff0c6966c1914109ed7e04423bc73967b7c693fbeb383bb1a057c75b37e2674"
	return chainCodeToCall
}

func (t *SimpleChaincode) GetUserChaincodeToCall() string {
	chainCodeToCall := "59d5075e7146d8df37bdcb0c289c293c66ee2a56c0f8d833410ff7cd8d3dd6c65ff0c6966c1914109ed7e04423bc73967b7c693fbeb383bb1a057c75b37e2674"
	return chainCodeToCall
}

// ============================================================================================================================
// Init function
// ============================================================================================================================

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

// ============================================================================================================================
// Invoke function is the entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "create" { //create a tx when a student applied a job and auto to check this application
		return t.Create(stub, args)
	} else if function == "artificialCheck" { //agency check this application when auto check not passed
		return t.ArtificialCheck(stub, args)
	} else if function == "evaluate" { //student and agancy evaluate each other
		return t.Evaluate(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

// ============================================================================================================================
// Create function is used to create a tx when a student applied a job and auto to check this application
// 1 input
// "TxID","TxInfo"
// ============================================================================================================================
func (t *SimpleChaincode) Create(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4. ")
	}
	TxID := args[0]
	TxInfo := args[1]
	TxTest, _ := stub.GetState(TxID)

	//test if the TX has been existed
	if TxTest != nil {
		return nil, errors.New("the Tx is existed")
	}

	// add the Tx
	err = stub.PutState(TxID, []byte(TxInfo))
	if err != nil {
		return nil, errors.New("Failed to add the user")
	}

	var TXInfoJsonType TXInfoStruct //json type to accept the TxInfo from state

	err = json.Unmarshal([]byte(TxInfo), &TXInfoJsonType)
	if err != nil {
		fmt.Println("error:", err)
	}

	//attach the TxID to related job
	//invoke JobInfo chaincode to add this TxID attach to the Job
	jobChainCodeToCall := t.GetJobChaincodeToCall()
	funcOfJobChaincode := "AddTX"
	invokeArgsOfJobChaincode := util.ToChaincodeArgs(funcOfJobChaincode, string(TXInfoJsonType.JobID), string(TXInfoJsonType.TxID))
	response1, err := stub.InvokeChaincode(jobChainCodeToCall, invokeArgsOfJobChaincode)
	if err != nil {
		errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	}
	fmt.Printf("Invoke chaincode successful. Got response %s", string(response))

	//attach the TxID to related student
	//invoke UserInfo chaincode to add this TxID attach to the student
	userChainCodeToCall := t.GetUserChaincodeToCall()
	funcOfUserChaincode := "AddTX"
	invokeArgsOfUserChaincode := util.ToChaincodeArgs(funcOfUserChaincode, string(TXInfoJsonType.UserID), string(TXInfoJsonType.TxID))
	response2, err := stub.InvokeChaincode(userChainCodeToCall, invokeArgsOfUserChaincode)
	if err != nil {
		errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	}
	fmt.Printf("Invoke chaincode successful. Got response %s", string(response))

	//auto check
	// Query User`s credit score
	f := "query"
	queryArgs := util.ToChaincodeArgs(f, string(TXInfoJsonType.UserID))
	response, err := stub.QueryChaincode(userChainCodeToCall, queryArgs)
	if err != nil {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	}
	Score, err := strconv.Atoi(string(response))
	if err != nil {
		errStr := fmt.Sprintf("Error retrieving state from ledger for queried chaincode: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	}
	if Score > 8 {
		TXInfoJsonType.Status = []byte("已通过审核待评价")
	} else {
		TXInfoJsonType.Status = []byte("未通过自动审核")
	}

	// put the new TxInfo into state
	a, err := json.Marshal(TXInfoJsonType)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// ============================================================================================================================
// ArtificialCheck function is used to check this application when auto check not passed by agency
// 2 input
// "TxID","Result(1:通过；2:未通过)"
// ============================================================================================================================
func (t *SimpleChaincode) ArtificialCheck(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. ")
	}
	TxID := args[0]
	Result := strconv.Atoi(args[1])
	TxInfo, err := stub.GetState(TxID)

	//test if the TX has been existed
	if err != nil {
		return nil, errors.New("The TX never been exited")
	}
	if TxInfo == nil {
		return nil, errors.New("The TX`s information is empty!")
	}

	var TXInfoJsonType TXInfoStruct //json type to accept the TxInfo from state

	err = json.Unmarshal(TxInfo, &TXInfoJsonType)
	if err != nil {
		fmt.Println("error:", err)
	}

	if strings.EqualFold(string(TXInfoJsonType.Status), "未通过自动审核") {
		if Result == 1 {
			TXInfoJsonType.Status = []byte("已通过审核待评价")
		} else {
			TXInfoJsonType.Status = []byte("未通过审核，已回绝")
		}
	} else {
		return nil, errors.New("Incorrect stage of status. Expecting 未通过自动审核. ")
	}

	// put the new TxInfo into state
	a, err := json.Marshal(TXInfoJsonType)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// ============================================================================================================================
// Evaluate function is used to evaluate each other by student and agancy
// 3 input
// "TxID","UserID","Score"
// ============================================================================================================================
func (t *SimpleChaincode) Evaluate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3. ")
	}
	TxID := args[0]
	UserID := args[1]
	Score := args[2]

	TxInfo, err := stub.GetState(TxID)

	//test if the TX has been existed
	if err != nil {
		return nil, errors.New("The TX never been exited")
	}
	if TxInfo == nil {
		return nil, errors.New("The TX`s information is empty!")
	}

	var TXInfoJsonType TXInfoStruct //json type to accept the TxInfo from state

	err = json.Unmarshal(TxInfo, &TXInfoJsonType)
	if err != nil {
		fmt.Println("error:", err)
	}

	if strings.EqualFold(string(TXInfoJsonType.UserID), UserID) {
		TXInfoJsonType.AgencyScore = []byte(Score)
	} else {
		TXInfoJsonType.StuScore = []byte(Score)
	}

	if TXInfoJsonType.StuScore != nil && TXInfoJsonType.AgencyScore != nil {
		if TXInfoJsonType.StuScore > 8 {
			// Query agency`s ID
			f := "queryAgencyIDandSalary"
			queryArgs := util.ToChaincodeArgs(f, string(TXInfoJsonType.JobID))
			AgencyID, Salary, err := stub.QueryChaincode(t.GetJobChaincodeToCall(), queryArgs)
			if err != nil {
				errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", err.Error())
				fmt.Printf(errStr)
				return nil, errors.New(errStr)
			}

			f2 := "autoSettle"
			invokeArgs2 := util.ToChaincodeArgs(f2, string(TXInfoJsonType.UserID), string(AgencyID), string(Salary))
			response2, err := stub.InvokeChaincode(t.GetUserChaincodeToCall(), invokeArgs2)
			if err != nil {
				errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", err.Error())
				fmt.Printf(errStr)
				return nil, errors.New(errStr)
			}

			fmt.Printf("Invoke chaincode successful. Got response %s", string(response))

			TXInfoJsonType.Status = []byte("已结算")
		} else {
			TXInfoJsonType.Status = []byte("已评价未通过自动结算")
		}
	} else {
		// put the new TxInfo into state
		a, err := json.Marshal(TXInfoJsonType)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	// put the new TxInfo into state
	a, err := json.Marshal(TXInfoJsonType)
	if err != nil {
		return nil, err
	}
	return nil, nil

}

// ============================================================================================================================
// Query function is the entry point for Queries
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "queryTxInfo" {
		return t.QueryTxInfo(stub, args)
	}

	return nil, errors.New("failed to query")

}

// ============================================================================================================================
// QueryTxInfo function is used to query the Tx`s information.
// 1 input
// "TxID"
// ============================================================================================================================
func (t *SimpleChaincode) QueryTxInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 ")
	}
	TxID := args[0]

	// Get the state from the ledger
	TxInfo, err := stub.GetState(TxID)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + TxID + "\"}"
		return nil, errors.New(jsonResp)
	}

	if TxInfo == nil {
		jsonResp := "{\"Error\":\"Nil content for " + TxID + "\"}"
		return nil, errors.New(jsonResp)
	}

	return TxInfo, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
