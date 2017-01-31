package main

import (
	dm "github.com/waqasburney/dm_chaincode/chaincode/document_manager_chaincode/documentmanager"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sha "crypto/sha256"
)

const MYVERSION string = "1.0.0"

//This is the structure used to specify the version of document manager chaincode being invoked
type DMCodeInfo struct {
	Version string `json:"version"`
}

type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	var stateArg DMCodeInfo

	if len(args) != 1 {
		return nil, errors.New("init expects one argument, a JSON string with tagged version string")
	}
	err = json.Unmarshal([]byte(args[0]), &stateArg)
	if err != nil {
		return nil, errors.New("Version argument unmarshal failed: " + fmt.Sprint(err))
	}
	if stateArg.Version != MYVERSION {
		return nil, errors.New("Document Manager version " + MYVERSION + " must match version argument: " + stateArg.Version)
	}
	return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//args = [docId,docType]
	fmt.Printf("Invoke: Function called is %s\n", function)
	docType := args[0]
	docId := args[1]
	if function == "CreateDocument" {
		//NOTE:This only creates an empty document. To Initialize fields, call UpdateDocument
		rBytes, err := dm.CreateDocument(docType, docId)
		if err != nil {
			return nil, err
		}
		return t.putBytesToLedger(stub, docId, rBytes)
	} else if function == "UpdateDocument" {
		//fmt.Printf("UpdateDocument called\n")
		keyValuePairJson := args[2]
		//fmt.Printf("CC: keyValuePairJson %s\n",keyValuePairJson)
		docBytes, err := t.getBytesFromLedger(stub, docId)
		if err != nil {
			return nil, err
		}
		rBytes, err := dm.UpdateDocument(docType, docId, keyValuePairJson, docBytes)
		if err != nil {
			return nil, err
		}
		return t.putBytesToLedger(stub, docId, rBytes)
	}
	return nil, fmt.Errorf("Function Name: Invalid\n")
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query: Function called is %s\n", function)
	if function == "GetDocument" {
		docId := args[0]
		fmt.Printf("GetDocument for docId: %s\n",docId)
		return t.getBytesFromLedger(stub, docId)
	} else if function == "GetDocumentByFields" {
		docId := args[0]
		docType := args[1]
		fieldNamesJson := args[2]
		docBytes, err := t.getBytesFromLedger(stub, docId)
		if err != nil {
			return nil, err
		}
		return dm.GetDocumentByFields(docType, docId, fieldNamesJson, docBytes)
	} else if function == "QueryDocument" {
		docId := args[0]
		docType := args[1]
		functionName := args[2]
		argsJson := args[3]
		docBytes, err := t.getBytesFromLedger(stub, docId)
		if err != nil {
			return nil, err
		}
		return dm.QueryDocument(docType, docId, functionName, argsJson, docBytes)
	} else if function == "GenerateHash" {
			docStr := args[0]
			hashStr := sha.Sum256([]byte(docStr))
			return []byte(hex.EncodeToString(hashStr[:])), nil;
	}
	return nil, nil
}


func (t *SimpleChaincode) getBytesFromLedger(stub shim.ChaincodeStubInterface, key string) ([]byte, error) {

	lBytes, err := stub.GetState(key)
	if err != nil {
		return nil, errors.New("Error in getting State from Ledger")
	}
	//TODO: lBytes can be nil for empty document
	/*
		if lBytes == nil {
			return nil, fmt.Errorf("Key: Could not get state from ledger for %s", key)
		}
	*/
	fmt.Printf("Bytes From ledger: %s\n", string(lBytes))
	return lBytes, nil
}

//TODO: See if returning back the byte array makes sense??

func (t *SimpleChaincode) putBytesToLedger(stub shim.ChaincodeStubInterface, key string, value []byte) ([]byte, error) {
	fmt.Printf("putBytesToLedger: Putting key: %s Value: %s\n",key,string(value))
	err := stub.PutState(key, value)
	if err != nil {
		fmt.Printf("Failed PUT to ledger while using %s\n", key)
		return nil, fmt.Errorf("Failed PUT to ledger while using key:  %s\n", key)
	}
	return value, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
