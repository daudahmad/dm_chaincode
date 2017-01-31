package documentmanager

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	copier "github.com/waqasburney/dm_chaincode/chaincode/document_manager_chaincode/copier"
	ecmdocument "github.com/waqasburney/dm_chaincode/chaincode/document_manager_chaincode/documenttemplate/ecmdocument"
)

func CreateDocument(docType string, docId string) ([]byte, error) {
	//var doc interface{}
	fmt.Printf("Inside CreateDocument for docType:  %s docId %s\n",docType, docId)
	switch docType {
	case "ECM_DOCUMENT":
		doc := ecmdocument.Document{DocIssuerID : "N/A"} // Need to Initialize with something else get null bytes when getting state from ledger
		dBytes, err := json.Marshal(doc)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Marshalled as : %s\n",string(dBytes))
		return dBytes,nil
	default:
		return nil, fmt.Errorf("DocumentType: No such type")
	}
}

func UpdateDocument(docType string, docId string, keyValuePairJson string, docBytes []byte) ([]byte, error) {
	var doc interface{}
	var f interface{}
	//fmt.Printf("UpdateDocument: docType: %s docId: %s keyValuePairJson: %s\n",docType, docId, keyValuePairJson)
	switch docType {
	case "ECM_DOCUMENT":
		fmt.Printf("UpdateDocument DOCTYPE: ECM_DOCUMENT\n")
		var t =ecmdocument.Document{}
		var t2 =ecmdocument.Document{}
		err := json.Unmarshal(docBytes, &t)
		if err != nil {
			return nil,err
		}
		fmt.Printf("Reached here\n")
		doc = &t
		f = &t2
	default:
		return nil, fmt.Errorf("DocumentType: No such type")
	}
	//err := json.Unmarshal([]byte(keyValuePairJson), &doc)
	err := json.Unmarshal([]byte(keyValuePairJson), &f)
	if err != nil {
		fmt.Println("Error in unmarshaling keyValuePairJson")
		return nil, err
	}
	fmt.Println("Reached Before Mapping\n")
	err = copier.Copy(doc,f)

	dBytes, err := json.Marshal(doc)
	if err != nil {
		return nil, err
	}
	return dBytes, nil
}

func QueryDocument(docType string, docId string, functionName string, argsJson string, docBytes []byte) ([]byte, error) {
	//var doc interface{}
	switch docType {
	case "ECM_DOCUMENT":
		{
			doc := ecmdocument.Document{}
			//TODO: Get document from ledger using docId key
			err := json.Unmarshal(docBytes, &doc)
			if err != nil {
				return nil, err
			}
			return ecmdocument.QueryDocument(functionName, argsJson, doc)
		}
	default:
		return nil, fmt.Errorf("DocumentType: Invalid ")
	}
}

func GetDocumentByFields(docType string, docId string, fieldNamesJson string, docBytes []byte) ([]byte, error) {
	var doc interface{}
	//var retStr string
	//var response interface{}
	//var fieldNames []string
	switch docType {
	case "ECM_DOCUMENT":
		var t ecmdocument.Document
		responseDoc := ecmdocument.Document{}
		err := json.Unmarshal(docBytes, &t) //stuff ledger bytes into t
		if err != nil {
			return nil, err
		}
		doc = &t
		retStr, err := getValuesByFields(fieldNamesJson, doc)
		if err != nil {
			fmt.Println("Get Values by Fields failed")
		}
		err = json.Unmarshal([]byte(retStr), &responseDoc) //This unmarshal is required to cast to the right type so that
		if err != nil {
			fmt.Println("Error in unmarshaling fieldNamesJson")
		}
		//marshalling can use the right json tags
		rBytes, err := json.Marshal(responseDoc)
		if err != nil {
			return nil, err
		}
		return rBytes, nil


	default:
		return nil, fmt.Errorf("DocumenType: Invalid type")
	}
}

//getValuesByFields is a utility function that returns the values of the given fields from the input document structure
func getValuesByFields(fieldNamesJson string, doc interface{})(string,error){
	var fieldNames []string
	var retStr string
	err := json.Unmarshal([]byte(fieldNamesJson), &fieldNames)
	if err != nil {
		fmt.Println("Error in unmarshaling fieldNamesJson")
		return "", err
	}
	//retStr = fmt.Sprintf("{") //use this to construct the response string
	for _, v := range fieldNames {
		val, err := getValue(v, doc)
		fmt.Printf("Field: %s Value: %v\n",v,val)
		if err != nil {
			return retStr, err
		}
		jVal,err := json.Marshal(val)
		if err != nil {
			return retStr,err
		}
		if len(retStr) == 0 {
    	retStr = fmt.Sprintf("\"%s\":%s", v, string(jVal))
		} else
		{
			retStr = fmt.Sprintf(" %s,\"%s\":%v", retStr,v, string(jVal))
		}
	}
	retStr = fmt.Sprintf("{%s}", retStr)
	fmt.Printf("RETSTR: %s\n",retStr)
	return retStr,nil
}



func getValue(fieldName string, t interface{}) (interface{}, error) {
	// t has to contain pointer type
	var value interface{}
	s := reflect.ValueOf(t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		// f := s.Field(i)
		structFieldName := strings.ToLower(typeOfT.Field(i).Name)
		inputFieldName := strings.ToLower(fieldName)
		if structFieldName == inputFieldName {
		//if typeOfT.Field(i).Name == fieldName {
			fmt.Printf("Type is %s\n", s.Field(i).Kind())
			value = s.Field(i).Interface()
			return value, nil
		}

		// fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	return value, fmt.Errorf("FieldName: Does not exist")
}

/*
func main() {
	fmt.Println("Hello, playground")
	nameJson := "{\"DocIssuerID\":\"burney1\", \"DocIssuerName\":\"Waqas Burney\", \"DocDescription\":\"My test document.\", \"DocURL\":\"www.google.com\", \"DocID\":\"1234\",  \"DocHash\":\"1231241asd23e\", \"DocUserID\":\"waqasb\" , \"DocUserIDList\":[\"Bill\",\"Tom\"] ,\"AccessOnceOnly\":\"false\", \"IsDeleted\":\"false\"}"
  //nameJson := "{\"EstimatedLoadingDate\" : \"01/12/2016\",\"PlaceOfLoading:\":\"Singapore\",\"ShipperReferenceNumber\":\"213124\",\"HouseBillOfLadingReference\":\"123\",\"ShipperName\":\"Masai\",\"ShipmentContainers\":[{\"GrossWeight\":123,\"NetWeight\":2,\"VerifiedGrossMass\":23},{\"GrossWeight\":6,\"NetWeight\":15,\"VerifiedGrossMass\":26}]}"
	fmt.Println("NAME JSON: %s", nameJson)

	doc := ecmdocument.Document{DocIssuerID : "burney1",DocIssuerName : "Muhammad Burney",DocDescription : "Its an empty Document"}
	fmt.Printf("Document Before is %+v\n",doc);
	dByte, _ := json.Marshal(doc)
	docByte , _ := UpdateDocument("ECM_DOCUMENT", "1234", nameJson , dByte)
	err := json.Unmarshal(docByte, &doc)
	fmt.Printf("Document After 1 is %+v\n",doc);

	//dByte, _ = json.Marshal(doc)
	//docByte , _ = UpdateDocument("ECM_DOCUMENT", "burney1", nameJson2 , dByte)
	//err = json.Unmarshal(docByte, &doc)
	//fmt.Printf("Document After 2 is %+v\n",doc);
	//mapping := f.(map[string]interface{})
	//for k, v := range mapping {
//		fmt.Printf("FieldName: %s\n", k)
		//setValue(k, v, &doc)
//	}

	if err != nil {
		fmt.Println("Error in unmarshaling docByte")
	}

	//fmt.Printf("ShipperReferenceNumber: %s HouseBillOfLadingReference: %s\n", doc.ShipperReferenceNumber, doc.HouseBillOfLadingReference)
	//fmt.Printf("ShipmentContainers: %v \n", doc.ShipmentContainers)

}*/
