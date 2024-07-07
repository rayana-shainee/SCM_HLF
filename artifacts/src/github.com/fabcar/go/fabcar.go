package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/flogging"
)

// SmartContract Define the Smart Contract structure
type SmartContract struct {
}

// Product : Define the product structure
type Product struct {
	ProductID   string `json:"productID"`
	ProductName string `json:"productName"`
	Quantity    string `json:"quantity"`
	Cost        string `json:"cost"`
	Role        string `json:"role"`
	Organic     string `json:"isOrganic"`
	Desc        string `json:"desc"`
	Status      string `json:"status"`
	Date        string `json:"date"`
}

// Init ; Method for initializing smart contract
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

var logger = flogging.MustGetLogger("supplychain_cc")

// Invoke : Method for invoking smart contract
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()

	logger.Infof("Function name is: %s", function)
	logger.Infof("Args length is : %d", len(args))

	if function == "queryProduct" {
		return s.queryProduct(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createProduct" {
		return s.createProduct(APIstub, args)
	} else if function == "setOrganic" {
		return s.setOrganic(APIstub, args)
	} else if function == "updateStatus" {
		return s.updateStatus(APIstub, args)
	} else if function == "buyProduct" {
		return s.buyProduct(APIstub, args)
	} else if function == "queryAllProducts" {
		return s.queryAllProducts(APIstub)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	productAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(productAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	products := []Product{
		Product{ProductID: "PROD1", ProductName: "Wheat", Quantity: "100", Cost: "50", Role: "farmer", Organic: "false", Desc: "Grown using traditional methods", Status: "harvested", Date: "2024-06-01"},
		Product{ProductID: "PROD2", ProductName: "Rice", Quantity: "200", Cost: "70", Role: "farmer", Organic: "false", Desc: "Grown in flooded fields", Status: "harvested", Date: "2024-06-01"},
	}

	i := 0
	for i < len(products) {
		productAsBytes, _ := json.Marshal(products[i])
		APIstub.PutState("PROD"+strconv.Itoa(i), productAsBytes)
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 9 {
		return shim.Error("Incorrect number of arguments. Expecting 9")
	}

    if args[4] != "farmer" {
		return shim.Error("Only farmers can add products")
	}
	var product = Product{ProductID: args[0], ProductName: args[1], Quantity: args[2], Cost: args[3], Role: args[4], Organic: "false", Desc: args[6], Status: args[7], Date: args[8]}

	productAsBytes, _ := json.Marshal(product)
	APIstub.PutState(args[0], productAsBytes)

	return shim.Success(productAsBytes)
}

func (s *SmartContract) setOrganic(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	productAsBytes, _ := APIstub.GetState(args[0])
	product := Product{}

	json.Unmarshal(productAsBytes, &product)
	if args[1] != "govt" {
		return shim.Error("Only government can set the product as organic")
	}

	product.Organic = "true"

	productAsBytes, _ = json.Marshal(product)
	APIstub.PutState(args[0], productAsBytes)

	return shim.Success(productAsBytes)
}

func (s *SmartContract) updateStatus(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	productAsBytes, _ := APIstub.GetState(args[0])
	product := Product{}

	json.Unmarshal(productAsBytes, &product)
	if args[2] != "distributor" && (args[1] != "Delivered" && args[1] != "out for delivery") {
		return shim.Error("Only distributors can update the status to 'Delivered' or 'out for delivery'")
	}

	product.Status = args[1]

	productAsBytes, _ = json.Marshal(product)
	APIstub.PutState(args[0], productAsBytes)

	return shim.Success(productAsBytes)
}

func (s *SmartContract) buyProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	productAsBytes, _ := APIstub.GetState(args[0])
	product := Product{}

	json.Unmarshal(productAsBytes, &product)
	if product.Organic != "true" {
		return shim.Error("Product is not organic, cannot be bought by customer")
	}

	if args[2] != "customer" {
		return shim.Error("Only customers can buy products")
	}

	quantityToBuy, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Quantity to buy must be a numeric value")
	}

	currentQuantity, err := strconv.Atoi(product.Quantity)
	if err != nil {
		return shim.Error("Current quantity must be a numeric value")
	}

	if quantityToBuy > currentQuantity {
		return shim.Error("Not enough quantity available")
	}

	updatedQuantity := currentQuantity - quantityToBuy
	product.Quantity = strconv.Itoa(updatedQuantity)

	productAsBytes, _ = json.Marshal(product)
	APIstub.PutState(args[0], productAsBytes)

	return shim.Success(productAsBytes)
}

func (s *SmartContract) queryAllProducts(APIstub shim.ChaincodeStubInterface) sc.Response {
	startKey := "PROD0"
	endKey := "PROD999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer strings.Builder
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if bArrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return shim.Success([]byte(buffer.String()))
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}