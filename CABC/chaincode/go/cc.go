package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type Atc struct {
	//ObjectType string `json:"docType"`
	ID        string `json:"ID"`
	Time      string `json:"Time"`
	Publisher string `json:"Publisher"`
	Company   string `json:"Company"`
	Type      string `json:"Type"`
	Address   string `json:"Address"`
	Signature string `json:"Signature"`
	IsIPFS    bool   `json:"IsIPFS"`
	Flight    string `json:"Flight"`
	Content   string `json:"Content"`

	Historys []HistoryItem // 当前edu的历史记录
}

type HistoryItem struct {
	TxId string
	Atc  Atc
}

type AtcChaincode struct {
}

func (t *AtcChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Print("===========  Init  ===========")

	return shim.Success(nil)
}

func (t *AtcChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fun, args := stub.GetFunctionAndParameters()

	if fun == "addAtc" {
		return t.addAtc(stub, args) // 添加信息
	} else if fun == "queryAtcInfoByID" {
		return t.queryAtcInfoByID(stub, args) // 根据证书编号及姓名查询信息
	} else if fun == "queryAtcByQueryString" {
		return t.queryAtcByQueryString(stub, args) // 根据证书编号及姓名查询信息
	} else if fun == "updateAtc" {
		return t.updateAtc(stub, args) // 根据证书编号及姓名查询信息
	} else if fun == "delAtc" {
		return t.delAtc(stub, args) // 根据证书编号及姓名查询信息
	}

	return shim.Error("Wrong")
}

func (t *AtcChaincode) addAtc(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	var atc Atc
	err := json.Unmarshal([]byte(args[0]), &atc)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}

	//// 查重: 身份证号码必须唯一
	//_, exist := GetEduInfo(stub, edu.EntityID)
	//if exist {
	//	return shim.Error("要添加的身份证号码已存在")
	//}

	_, bl := PutAtc(stub, atc)
	if !bl {
		return shim.Error("保存信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}

func (t *AtcChaincode) queryAtcInfoByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 根据身份证号码查询edu状态
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据身份证号码查询信息失败")
	}

	if b == nil {
		return shim.Error("根据身份证号码没有查询到相关的信息")
	}

	// 对查询到的状态进行反序列化
	var atc Atc
	err = json.Unmarshal(b, &atc)
	if err != nil {
		return shim.Error("反序列化edu信息失败")
	}

	// 迭代处理
	//var historys []HistoryItem
	//var hisAtc Atc
	//for iterator.HasNext() {
	//	hisData, err := iterator.Next()
	//	if err != nil {
	//		return shim.Error("获取edu的历史变更数据失败")
	//	}

	//	var historyItem HistoryItem
	//	historyItem.TxId = hisData.TxId
	//	json.Unmarshal(hisData.Value, &hisAtc)

	//	if hisData.Value == nil {
	//		var empty Atc
	//		historyItem.Atc = empty
	//	} else {
	//		historyItem.Atc = hisAtc
	//	}

	//	historys = append(historys, historyItem)

	//}
	items, err := getHistory(stub, atc.ID)

	if err != nil {
		return shim.Error("获取历史记录失败")
	}

	atc.Historys = *items

	// 返回
	result, err := json.Marshal(atc)
	if err != nil {
		return shim.Error("序列化edu信息时发生错误")
	}
	return shim.Success(result)
}

func getHistory(stub shim.ChaincodeStubInterface, atc_id string) (*[]HistoryItem, error) {

	// 获取历史变更数据
	iterator, err := stub.GetHistoryForKey(atc_id)
	if err != nil {
		return nil, err
	}
	defer iterator.Close()

	var historys []HistoryItem
	var hisAtc Atc
	for iterator.HasNext() {
		hisData, err := iterator.Next()
		if err != nil {
			return nil, err
		}

		var historyItem HistoryItem
		historyItem.TxId = hisData.TxId
		json.Unmarshal(hisData.Value, &hisAtc)

		if hisData.Value == nil {
			var empty Atc
			historyItem.Atc = empty
		} else {
			historyItem.Atc = hisAtc
		}

		historys = append(historys, historyItem)
	}

	return &historys, nil
}

// 根据证书编号及姓名查询信息
// args: CertNo, name
func (t *AtcChaincode) queryAtcByQueryString(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 拼装CouchDB所需要的查询字符串(是标准的一个JSON串)
	// queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"eduObj\", \"CertNo\":\"%s\"}}", CertNo)
	//queryString := fmt.Sprintf("{\"selector\":{\"Publisher\":\"%s\", \"CertNo\":\"%s\", \"Name\":\"%s\"}}", DOC_TYPE, CertNo, name)
	//queryString := fmt.Sprintf("{\"selector\":{\"Publisher\":\"%s\"}}", args[0])
	queryString := fmt.Sprintf("{\"selector\": %s}", args[0])

	// 查询数据
	//result, err := getAtcByQueryString(stub, queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error("do not find: " + err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	//var buffer bytes.Buffer
	var atcs = make([]Atc, 0)

	//bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("data get failed: " + err.Error())
		}
		//// Add a comma before array members, suppress it for the first array member
		//if bArrayMemberAlreadyWritten == true {
		//	buffer.WriteString(",")
		//}

		//// Record is a JSON object, so we write as-is
		//buffer.WriteString(string(queryResponse.Value))
		//bArrayMemberAlreadyWritten = true

		var atc Atc
		if err := json.Unmarshal(queryResponse.Value, &atc); err != nil {
			fmt.Println(err.Error())
			return shim.Error("data unmarshal failed: " + err.Error())
		}

		items, err := getHistory(stub, atc.ID)

		if err != nil {
			return shim.Error("获取历史记录失败")
		}

		atc.Historys = *items
		atcs = append(atcs, atc)
	}

	data, err := json.Marshal(atcs)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("data marshal failed: " + err.Error())
	}
	//if err != nil {
	//	return shim.Error("根据证书编号及姓名查询信息时发生错误")
	//}
	//if result == nil {
	//	return shim.Error("根据指定的证书编号及姓名没有查询到相关的信息")
	//}
	return shim.Success(data)
}

//// 根据指定的查询字符串实现富查询
//func getAtcByQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
//	resultsIterator, err := stub.GetQueryResult(queryString)
//	if err != nil {
//		return nil, err
//	}
//	defer resultsIterator.Close()
//
//	// buffer is a JSON array containing QueryRecords
//	//var buffer bytes.Buffer
//	var atcs = make([]Atc, 0)
//
//	//bArrayMemberAlreadyWritten := false
//	for resultsIterator.HasNext() {
//		queryResponse, err := resultsIterator.Next()
//		if err != nil {
//			return nil, err
//		}
//		//// Add a comma before array members, suppress it for the first array member
//		//if bArrayMemberAlreadyWritten == true {
//		//	buffer.WriteString(",")
//		//}
//
//		//// Record is a JSON object, so we write as-is
//		//buffer.WriteString(string(queryResponse.Value))
//		//bArrayMemberAlreadyWritten = true
//
//		var atc Atc
//		if err := json.Unmarshal(queryResponse.Value, &atc); err != nil {
//			fmt.Println(err.Error())
//			return shim.Error("data unmarshal failed: " + err.Error())
//		}
//		users = append(users, user)
//	}
//
//	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
//
//	return buffer.Bytes(), nil
//}

func (t *AtcChaincode) updateAtc(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	var atcinfo Atc
	err := json.Unmarshal([]byte(args[0]), &atcinfo)
	if err != nil {
		return shim.Error("反序列化edu信息失败")
	}

	// 根据身份证号码查询信息
	result, bl := GetAtcInfo(stub, atcinfo.ID)
	if !bl {
		return shim.Error("根据身份证号码查询信息时发生错误")
	}

	result.ID = atcinfo.ID
	result.Time = atcinfo.Time
	result.Address = atcinfo.Address
	result.Signature = atcinfo.Signature
	result.Content = atcinfo.Content
	result.IsIPFS = atcinfo.IsIPFS

	_, bl = PutAtc(stub, result)
	if !bl {
		return shim.Error("保存信息信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息更新成功"))
}

func (t *AtcChaincode) delAtc(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	/*var edu Education
	result, bl := GetEduInfo(stub, info.EntityID)
	err := json.Unmarshal(result, &edu)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}*/

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("删除信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息删除成功"))
}

func PutAtc(stub shim.ChaincodeStubInterface, atc Atc) ([]byte, bool) {

	//edu.ObjectType = DOC_TYPE

	b, err := json.Marshal(atc)
	if err != nil {
		return nil, false
	}

	// 保存edu状态
	err = stub.PutState(atc.ID, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

func valOf(i ...interface{}) []reflect.Value {
	var rt []reflect.Value
	for _, i2 := range i {
		rt = append(rt, reflect.ValueOf(i2))
	}
	return rt
}

func GetAtcInfo(stub shim.ChaincodeStubInterface, ID string) (Atc, bool) {
	var atc Atc
	// 根据身份证号码查询信息状态
	b, err := stub.GetState(ID)
	if err != nil {
		return atc, false
	}

	if b == nil {
		return atc, false
	}

	// 对查询到的状态进行反序列化
	err = json.Unmarshal(b, &atc)
	if err != nil {
		return atc, false
	}

	// 返回结果
	return atc, true
}

func main() {
	err := shim.Start(new(AtcChaincode))
	if err != nil {
		fmt.Printf("启动EducationChaincode时发生错误: %s", err)
	}
}
