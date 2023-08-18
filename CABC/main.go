package main

import (
	"fmt"
	"os"
	"sdkInit"
)

//var App sdkInit.Application

func main() {
	fmt.Println("nice")

	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser: "Admin",    //admin
			OrgName:      "Baiyi",    //name
			OrgMspId:     "BaiyiMSP", //msp
			OrgUser:      "User1",    //user
			//OrgMspClient          *mspclient.Client //running
			//OrgAdminClientContext *contextAPI.ClientProvider
			//OrgResMgmt            *resmgmt.Client
			OrgPeerNum: 1, //num
			//Peers                 []fab.Peer //peers
			OrgAnchorFile: "/root/go/src/CABC/fixtures/channel-artifacts/BaiyiMSPanchors.tx",
		},
		{
			OrgAdminUser: "Admin",       //admin
			OrgName:      "AirChina",    //name
			OrgMspId:     "AirChinaMSP", //msp
			OrgUser:      "User1",       //user
			//OrgMspClient          *mspclient.Client //running
			//OrgAdminClientContext *contextAPI.ClientProvider
			//OrgResMgmt            *resmgmt.Client
			OrgPeerNum: 1, //num
			//Peers                 []fab.Peer //peers
			OrgAnchorFile: "/root/go/src/CABC/fixtures/channel-artifacts/AirChinaMSPanchors.tx",
		},
		{
			OrgAdminUser: "Admin",           //admin
			OrgName:      "OtherAirline",    //name
			OrgMspId:     "OtherAirlineMSP", //msp
			OrgUser:      "User1",           //user
			//OrgMspClient          *mspclient.Client //running
			//OrgAdminClientContext *contextAPI.ClientProvider
			//OrgResMgmt            *resmgmt.Client
			OrgPeerNum: 1, //num
			//Peers                 []fab.Peer //peers
			OrgAnchorFile: "/root/go/src/CABC/fixtures/channel-artifacts/OtherAirlineMSPanchors.tx",
		}}

	info := sdkInit.SdkEnvInfo{
		//channel  info
		ChannelID:     "mychannel",
		ChannelConfig: "/root/go/src/CABC/fixtures/channel-artifacts/channel.tx",
		//orgs
		Orgs: orgs,
		//orderer info
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "Orderer",
		OrdererEndpoint:  "orderer.cabc.com",
		//OrdererClientContext *contextAPI.ClientProvider

		//chaincode info
		ChaincodeID: "cabc_chaincode",
		//ChaincodeGoPath      string
		ChaincodePath:    "/root/go/src/CABC/chaincode/go",
		ChaincodeVersion: "1.0.0",
		//Client           *channel.Client
	}

	sdkentity, err := sdkInit.Setup("config.yaml", &info)

	if err != nil {
		fmt.Println(">> Sdk set error", err)
		os.Exit(-1)
	}

	if err := sdkInit.CreateChannel(&info); err != nil {
		fmt.Println(">> Create Channel error", err)
		os.Exit(-1)
	}

	if err := sdkInit.JoinChannel(&info); err != nil {
		fmt.Println(">> Join Channel error", err)
		os.Exit(-1)
	}
	packageID, err := sdkInit.InstallCC(&info)
	fmt.Println(packageID)

	if err != nil {
		fmt.Println(">> install chaincode error", err)
		os.Exit(-1)
	}
	if err := sdkInit.ApproveLifecycle(&info, 1, packageID); err != nil {
		fmt.Println(">> approve chaincode error", err)
		os.Exit(-1)
	}

	if err := sdkInit.InitCC(&info, false, sdkentity); err != nil {
		fmt.Println(">> init chaincode error", err)
		os.Exit(-1)
	}
	fmt.Println(">> 链码状态设置完成")
}
