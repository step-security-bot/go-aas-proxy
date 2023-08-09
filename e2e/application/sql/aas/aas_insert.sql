INSERT INTO aas VALUES ('smart.festo.com/demo/aas/1/1/454576463545648365874', '{"modelType": {"name": "AssetAdministrationShell"},"idShort": "Festo_3S7PM0CP4BD","identification": {"idType": "IRI","id": "smart.festo.com/demo/aas/1/1/454576463545648365874"},"dataSpecification": [],"embeddedDataSpecifications": [],"submodels": [{"keys": [{"type": "AssetAdministrationShell","local": false,"value": "smart.festo.com/demo/aas/1/1/454576463545648365874","idType": "IRI"},{"type": "Submodel","local": true,"value": "www.company.com/ids/sm/4343_5072_7091_3242","idType": "IRI"}]},{"keys": [{"type": "AssetAdministrationShell","local": false,"value": "smart.festo.com/demo/aas/1/1/454576463545648365874","idType": "IRI"},{"type": "Submodel","local": true,"value": "www.company.com/ids/sm/6053_5072_7091_5102","idType": "IRI"}]},{"keys": [{"type": "AssetAdministrationShell","local": false,"value": "smart.festo.com/demo/aas/1/1/454576463545648365874","idType": "IRI"},{"type": "Submodel","local": true,"value": "smart.festo.com/demo/sm/instance/1/1/13B7CCD9BF7A3F24","idType": "IRI"}]},{"keys": [{"type": "AssetAdministrationShell","local": false,"value": "smart.festo.com/demo/aas/1/1/454576463545648365874","idType": "IRI"},{"type": "Submodel","local": true,"value": "www.company.com/ids/sm/2543_5072_7091_2660","idType": "IRI"}]},{"keys": [{"type": "AssetAdministrationShell","local": false,"value": "smart.festo.com/demo/aas/1/1/454576463545648365874","idType": "IRI"},{"type": "Submodel","local": true,"value": "www.company.com/ids/sm/6563_5072_7091_4267","idType": "IRI"}]}],"asset": {"keys": [{"type": "Asset","local": true,"value": "","idType": "IRDI"}],"modelType": {"name": "Asset"},"dataSpecification": [],"embeddedDataSpecifications": [],"idShort": "","identification": {"idType": "IRDI","id": ""},"kind": "Instance"},"views": [],"conceptDictionary": [],"category": "CONSTANT","assetRef": {"keys": [{"type": "Asset","local": true,"value": "HTTP://PK.FESTO.COM/3S7PM0CP4BD","idType": "IRI"}]}}', current_timestamp, current_timestamp);


INSERT INTO submodel_proxy VALUES 
('www.company.com/ids/sm/4343_5072_7091_3242', 'Nameplate', 'https://www.hsu-hh.de/aut/aas/nameplate', 'smart.festo.com/demo/aas/1/1/454576463545648365874', current_timestamp,current_timestamp),
-- ('smart.festo.com\demo\sm\instance\1\1\13B7CCD9BF7A3F24','DeviceDescriptionFiles', 'http://admin-shell/sample/submodel/type/device-description-files', 'smart.festo.com/demo/aas/1/1/454576463545648365874', current_timestamp,current_timestamp),
-- ('www.company.com/ids/sm/6053_5072_7091_5102','Service', 'https://www.hsu-hh.de/aut/aas/service', 'smart.festo.com/demo/aas/1/1/454576463545648365874', current_timestamp,current_timestamp),
('www.company.com/ids/sm/6563_5072_7091_4267','Identification', 'https://www.hsu-hh.de/aut/aas/identification', 'smart.festo.com/demo/aas/1/1/454576463545648365874', current_timestamp,current_timestamp),
-- ('www.company.com/ids/sm/2543_5072_7091_2660','Document', 'https://www.hsu-hh.de/aut/aas/document', 'smart.festo.com/demo/aas/1/1/454576463545648365874', current_timestamp,current_timestamp)
;


