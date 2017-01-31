##Registering

##1.Init


peer chaincode deploy -p github.com/hyperledger/fabric/examples/chaincode/go/document_manager_chaincode -c '{"function":"Init", "Args": ["{\"version\":\"1.0.0\"}"]}'


peer chaincode invoke -l golang -n 062fd3207f8b0f10964c8e0225107ef6a8c8dedc712fda6fdf1453351bdc6d8ec6e97d95f6ef62ba70422af83c4f8584f8d5db800d0b98e48f0e6372a85ca04f -c '{"Function":"CreateDocument", "Args": ["ECM_DOCUMENT","ABC1234"]}'

peer chaincode invoke -l golang -n 062fd3207f8b0f10964c8e0225107ef6a8c8dedc712fda6fdf1453351bdc6d8ec6e97d95f6ef62ba70422af83c4f8584f8d5db800d0b98e48f0e6372a85ca04f -c '{"Function":"UpdateDocument", "Args": ["ECM_DOCUMENT","ABC1234","{\"DocIssuerID\":\"Masai\"}"]}'

peer chaincode query -l golang -n 062fd3207f8b0f10964c8e0225107ef6a8c8dedc712fda6fdf1453351bdc6d8ec6e97d95f6ef62ba70422af83c4f8584f8d5db800d0b98e48f0e6372a85ca04f -c '{"Function":"GetDocument", "Args": ["ABC123"]}'


peer chaincode invoke -l golang -n 062fd3207f8b0f10964c8e0225107ef6a8c8dedc712fda6fdf1453351bdc6d8ec6e97d95f6ef62ba70422af83c4f8584f8d5db800d0b98e48f0e6372a85ca04f -c '{"Function":"UpdateDocument", "Args": ["ECM_DOCUMENT","ABC123","{\"DocDescription\":\"This is a document\"}"]}'


peer chaincode query -l golang -n 062fd3207f8b0f10964c8e0225107ef6a8c8dedc712fda6fdf1453351bdc6d8ec6e97d95f6ef62ba70422af83c4f8584f8d5db800d0b98e48f0e6372a85ca04f -c '{"Function":"GetDocument", "Args": ["ABC123"]}'


peer chaincode query -l golang -n 062fd3207f8b0f10964c8e0225107ef6a8c8dedc712fda6fdf1453351bdc6d8ec6e97d95f6ef62ba70422af83c4f8584f8d5db800d0b98e48f0e6372a85ca04f -c '{"Function":"GetDocumentFields", "Args": ["ABC123","SAMPLE","[\"ShipperName\"]"]}'
peer chaincode query -l golang -n 70ddae626a2bc318ef19e1d32d904643913a6db642daa560e7916274d504d12c6cac2f5fd21bdede98c1f963d37860a3e9b6726108175186bac9251c45b7156b -c '{"Function":"QueryDocument", "Args": ["ABC123","SAMPLE","IsValidDocument","ARGS"]}'
peer chaincode invoke -l golang -n 0059edb858e57956ebd37762aaf05f0c141d97a49db1f9c254ee93dfdab21ee469f713928fd53adfee249d0b2fbfc3c28ce52c0294a7abe8d15ac7b92aa4b560 -c '{"Function":"ProcessDocument", "Args": ["SAMPLE","ABC123","SpecialProcessDocument","ARGS"]}'
