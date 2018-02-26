/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type AssetTrackingSmartContract struct {
}

type Asset struct {
	UUID         string  `json:"uuid"`
	SerialNumber string  `json:"serialNumber"`
	AssetType    string  `json:"assetType"`
	OwnerName    string  `json:"ownerName"`
	Description  string  `json:"description"`
	Events       []Event `json:"events"`
}

type Event struct {
	Summary            string `json:"summary"`
	Description        string `json:"description"`
	Date               string `json:"date"`
	BusinessProviderID string `json:"businessProviderId"`
	EncodedImage       string `json:"encodedImage"`
	EncodedFiles       string `json:"attachment"`
}

// Initializes the chaincode state
func (t *AssetTrackingSmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### Init ###########")
	asset := Asset{
		UUID:         "14a12ef0-9409-4872-9341-9ab003059ce9",
		SerialNumber: "3SZ1W278EJ8",
		AssetType:    "VEHICLE",
		OwnerName:    "THOMAS ALVA EDISON",
		Description:  "1908 FORD MODEL T\nCOLOR: ANYONE, AS LONG AS IT IS BLACK\nINTERIORS COLOR: BROWN LEATHER.\nENGINE: 2.9 L FOUR CYLINDERS\nENGINE NUMBER: 000015\nMADE IN: UNITED STATES\nBUYER NAME: THOMAS EDISON\nADDRESS: 345 Llewellyn Park, New Jersey, United States, USA.\nSELLER NAME: HENRY FORD",
	}
	jsonBlob, _ := json.Marshal(asset)
	err := stub.PutState(asset.UUID, jsonBlob)

	if err != nil {
		return shim.Error(err.Error())
	}

	asset = Asset{
		UUID:         "9d40ee4e-bf1e-4f74-8237-c5e9b6e8f6d3",
		SerialNumber: "3VW1W21KIBM312176",
		AssetType:    "VEHICLE",
		OwnerName:    "Jhonn Doe",
		Description:  "2011 VW JETTA STYLE ACTIVE MANUAL TRANSMISION. SIDE AIRBAGS PACKAGE, COLOR: WHITE CANDY     INTERIOR COLOR: BLACK FABRIC .    ENGINE: 2.5L FIVE CYLINDERS     ENGINE NUMBER: CCC094323     MADE IN: MEXICO      BUYER NAME: JHONN DOE     ADDRESS: 123 ABBY ROAD, THE DOMAIN. AUTIN TEXAS, USA.     SELLER NAME: RAY REDDINGTON",
	}
	jsonBlob, _ = json.Marshal(asset)
	err = stub.PutState(asset.UUID, jsonBlob)

	if err != nil {
		return shim.Error(err.Error())
	}

	asset = Asset{
		UUID:         "ab3af1a9-6d81-4be8-94f8-cd1667a894cb",
		SerialNumber: "157590103000100120006906040003",
		AssetType:    "REAL_ESTATE",
		OwnerName:    "Donald Trump",
		Description:  "Address: 1600 Pennsylvania Ave NW, Washington, DC 20500, USA Floor space: 5,110 m2       Construction started:    October 13, 1792,       Completed:   November 1, 1800",
	}
	jsonBlob, _ = json.Marshal(asset)
	err = stub.PutState(asset.UUID, jsonBlob)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

// Invoke makes queries... currently
func (t *AssetTrackingSmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### Invoke ###########")
	function, args := stub.GetFunctionAndParameters()

	if function != "invoke" {
		return shim.Error("Unknown function call")
	}

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting at least 2")
	}

	fn := args[0]
	switch fn {
	case "create":
		return t.create(stub, args)
	case "query":
		return t.query(stub, args)
	case "update":
		return shim.Error("Not yet implemented")
	case "delete":
		return shim.Error("Not yet implemented")
	default:
		return shim.Error("Unknown action, check the first argument, must be one of 'create', 'query', 'update' or 'delete'")
	}
}

// Finds an asset by id
func (t *AssetTrackingSmartContract) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### Query ###########")
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	id := args[1]

	assetBlob, err := stub.GetState(id)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if assetBlob == nil {
		return shim.Error("Not found!")
	}

	jsonResp := "{\"id\":\"" + id + "\",\"val\":\"" + string(assetBlob) + "\"}"

	a1 := Asset{}
	json.Unmarshal([]byte(assetBlob), &a1)
	fmt.Printf("-> %+v\n", a1)
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(assetBlob)
}

func (t *AssetTrackingSmartContract) create(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### Create ###########")
	err1 := t.validateAssetParams(args)
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	asset := t.createAsset(args)
	jsonBlob, err2 := json.Marshal(asset)
	if err2 != nil {
		return shim.Error(err2.Error())
	}

	fmt.Printf("Puting: " + string(jsonBlob))

	err3 := stub.PutState(asset.UUID, jsonBlob)
	if err3 != nil {
		return shim.Error(err3.Error())
	}

	return shim.Success(nil)
}

func (t *AssetTrackingSmartContract) createAsset(args []string) Asset {
	event := Event{
		Summary:            "CREATED",
		Description:        "CREATED",
		Date:               time.Now().Format(time.RFC3339),
		BusinessProviderID: args[6],
	}
	if len(args) >= 8 {
		if t.isBlank(args[7]) == false {
			event.EncodedImage = args[7]
		}
		if len(args) >= 9 && t.isBlank(args[8]) == false {
			event.EncodedFiles = args[8]
		}
	}

	asset := Asset{
		UUID:         args[1],
		SerialNumber: args[2],
		AssetType:    args[3],
		OwnerName:    args[4],
		Description:  args[5],
		Events:       []Event{event},
	}
	return asset
}

func (t *AssetTrackingSmartContract) validateAssetParams(args []string) error {
	if len(args) < 7 {
		return errors.New("Expecting 7 arguments, actual: " + strconv.Itoa(len(args)))
	}
	errorText := " "
	for i := 1; i <= 6; i++ {
		if t.isBlank(args[i]) {
			errorText += "[element at index " + strconv.Itoa(i) + " must not be blank] "
		}
	}
	errorText = strings.TrimSpace(errorText)
	if len(errorText) != 0 {
		return errors.New(errorText)
	}
	return nil
}

func (t *AssetTrackingSmartContract) isBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func main() {
	err := shim.Start(new(AssetTrackingSmartContract))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
