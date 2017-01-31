##Registering

##1.Init


peer chaincode deploy -p github.com/hyperledger/fabric/examples/chaincode/go/document_manager_chaincode -c '{"function":"Init", "Args": ["{\"version\":\"1.0.0\"}"]}'


peer chaincode invoke -l golang -n 70ddae626a2bc318ef19e1d32d904643913a6db642daa560e7916274d504d12c6cac2f5fd21bdede98c1f963d37860a3e9b6726108175186bac9251c45b7156b -c '{"Function":"CreateDocument", "Args": ["ECM_DOCUMENT","ABC123"]}'

peer chaincode invoke -l golang -n 70ddae626a2bc318ef19e1d32d904643913a6db642daa560e7916274d504d12c6cac2f5fd21bdede98c1f963d37860a3e9b6726108175186bac9251c45b7156b -c '{"Function":"UpdateDocument", "Args": ["ECM_DOCUMENT","ABC123","{\"docIssuerID\":\"Masai\"}"]}'

peer chaincode query -l golang -n 70ddae626a2bc318ef19e1d32d904643913a6db642daa560e7916274d504d12c6cac2f5fd21bdede98c1f963d37860a3e9b6726108175186bac9251c45b7156b -c '{"Function":"GetDocument", "Args": ["ABC123"]}'


peer chaincode invoke -l golang -n 70ddae626a2bc318ef19e1d32d904643913a6db642daa560e7916274d504d12c6cac2f5fd21bdede98c1f963d37860a3e9b6726108175186bac9251c45b7156b -c '{"Function":"UpdateDocument", "Args": ["ECM_DOCUMENT","ABC123","{\"docDescription\":\"This is a document\"}"]}'


peer chaincode query -l golang -n 70ddae626a2bc318ef19e1d32d904643913a6db642daa560e7916274d504d12c6cac2f5fd21bdede98c1f963d37860a3e9b6726108175186bac9251c45b7156b -c '{"Function":"GetDocument", "Args": ["ABC123"]}'


peer chaincode query -l golang -n 0059edb858e57956ebd37762aaf05f0c141d97a49db1f9c254ee93dfdab21ee469f713928fd53adfee249d0b2fbfc3c28ce52c0294a7abe8d15ac7b92aa4b560 -c '{"Function":"GetDocumentFields", "Args": ["ABC123","SAMPLE","[\"ShipperName\"]"]}'
peer chaincode query -l golang -n 0059edb858e57956ebd37762aaf05f0c141d97a49db1f9c254ee93dfdab21ee469f713928fd53adfee249d0b2fbfc3c28ce52c0294a7abe8d15ac7b92aa4b560 -c '{"Function":"QueryDocument", "Args": ["ABC123","SAMPLE","IsValidDocument","ARGS"]}'
peer chaincode invoke -l golang -n 0059edb858e57956ebd37762aaf05f0c141d97a49db1f9c254ee93dfdab21ee469f713928fd53adfee249d0b2fbfc3c28ce52c0294a7abe8d15ac7b92aa4b560 -c '{"Function":"ProcessDocument", "Args": ["SAMPLE","ABC123","SpecialProcessDocument","ARGS"]}'
