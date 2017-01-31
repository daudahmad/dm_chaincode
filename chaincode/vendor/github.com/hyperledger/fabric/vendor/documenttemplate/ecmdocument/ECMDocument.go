package ecmdocument
import (
    "encoding/json"
    "encoding/hex"
    "errors"
    "fmt"
    sha "crypto/sha256"
)


type Document struct {
  DocIssuerID   string `json:"docIssuerID,omitempty"`
  DocIssuerName string `json:"docIssuerName,omitempty"`
  DocDescription string `json:"docDescription,omitempty"`
  DocURL string  `json:"docURL,omitempty"`
  DocID  string `json:"docID,omitempty"`
  DocHash string  `json:"docHash,omitempty"`
  DocUserID string `json:"docUserID,omitempty"`
  DocUserIDList []string `json:"docUserIDList,omitempty"`
  AccessOnceOnly string `json:"accessOnceOnly,omitempty"`
  IsDeleted string `json:"isDeleted,omitempty"`
}

func QueryDocument(functionName string, argsJson string, doc Document)([]byte,error){
  switch functionName {
      case "IsValidDocument" :
        status,err := IsValidDocument(doc,argsJson)
        return status,err
      case  "GetDocument" :
        ecmDoc, err := GetDocument(argsJson)
        return ecmDoc,err
      default:
          return nil, fmt.Errorf("Function Name: Invalid")
  }
}

func GetDocument(argsJson string)([]byte,error){
  var err error
  var ecmDoc Document
  err = json.Unmarshal([]byte(argsJson), &ecmDoc)
  if err != nil{
    returnJSON, _ := json.Marshal("false")
    return returnJSON, errors.New("Unmarshal failed: " + fmt.Sprint(err))
  }

  returnJSON,_ := json.Marshal(ecmDoc)
  return returnJSON,nil
}


func GenerateHash(docStr string) ([]byte,error){
  //docBytes, err := json.Marshal(docStr)
  hashStr := sha.Sum256([]byte(docStr))
  return []byte(hex.EncodeToString(hashStr[:])), nil;
}

func IsValidDocument(ecmDoc Document, argsJson string)([]byte,error){
  var err error

  if len(argsJson) == 0 {
      returnJSON, _ := json.Marshal("false")
      return returnJSON, errors.New("Empty argument found")
  }

  err = json.Unmarshal([]byte(argsJson), &ecmDoc)
  if err != nil {
      returnJSON, _ := json.Marshal("false")
      return returnJSON, errors.New("Unmarshal failed: " + fmt.Sprint(err))
  }

  if ecmDoc.DocIssuerID == "" || ecmDoc.DocID == "" {
      returnJSON, _ := json.Marshal("false")
      return returnJSON, errors.New("Doc Issuer ID and Doc ID are mandatory")
  }
  returnJSON := []byte("true")
	return returnJSON,nil
}

/*
func main() {
    fmt.Printf("Starting...\n")
    //var ecmDoc Document
    //result, _ := QueryDocument("GetDocument","{\"DocIssuerID\":\"burney1\", \"DocIssuerName\":\"Waqas Burney\", \"DocDescription\":\"My test document.\", \"DocURL\":\"www.google.com\", \"DocID\":\"1234\",  \"DocHash\":\"www.google.com\", \"DocUserID\":\"waqasb\" , \"DocUserIDList\":[\"Bill\",\"Tom\"] ,\"AccessOnceOnly\":\"false\", \"IsDeleted\":\"false\"}",ecmDoc)
    result1, _  := GenerateHash("WAQAS IS THE BOSS 1")
    result2, _  := GenerateHash("WAQAS IS THE BOSS 2")
    //stringResult, _ := json.Marshal(ecmDoc)
    //fmt.Println(string(stringResult))
    fmt.Println(string(result1))
    fmt.Println(string(result2))
}*/
