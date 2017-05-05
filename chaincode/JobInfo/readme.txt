JobInfo智能合约：每一个兼职信息都在这一个合约中创建，修改，删除...

数据结构：
	{
	    "JobDetail": {
	        "Demand": "",
	        "Day": "",
	        "Place": "",
	        "Salary": "50",
	        "JobTime": ""
	    },
	    "JobID": "j1",
	    "UserID": "1",
	    "UserName": "",
	    "TotalApplied": "123",
	    "TotalHired": "123",
	    "TotalSettled": "123",
	    "TotalWaitCheck": "123",
	    "Txs": [
	        "123",
	        "123"
	    ]
	}

world status：
	key		  value
	JobID	  JobInfo
	注释：key这一列下面，带引号就是key的实际值，不带引号的是变量，变量是什么，key的值是什么

deploy：
	init的参数有0个

	init的POST请求
	Post https://a6377d73838047d39f8527f035520915-vp0.us.blockchain.ibm.com:5002/chaincode
	{
	  "jsonrpc": "2.0",
	  "method": "deploy",
	  "params": {
	    "type": 1,
	    "chaincodeID": {
	      "path": "https://github.com/yglym2012/BlockchainCompetition/chaincode/JobInfo"
	    },
	    "ctorMsg": {
	      "function": "deploy",
	      "args": [
	      ]
	    },
	    "secureContext": "user_type1_0"
	  },
	  "id": 0
	}

invoke：
	add添加一条新兼职信息，并将该兼职ID捆绑到发布信息中介个人信息中
	参数有2个："JobId","JobInfo"

	delete删除一条已发布的兼职信息
	参数有1个："JobId"

	edit修改兼职信息
	参数有2个："JobId","NewJobInfo"

	addTotalApplied该兼职总申请数加一
	参数有1个："JobId"

	addTotalWaitCheck该兼职总待审核数加一
	参数有1个："JobId"

	addTotalHired该兼职总录用数加一
	参数有1个："JobId"

	addTotalSettled该兼职总结算数加一
	参数有1个："JobId"

	addTX添加和该兼职信息相关的TXID
	参数有2个："JobId","TXID"

	invoke的POST请求
	Post https://a6377d73838047d39f8527f035520915-vp0.us.blockchain.ibm.com:5002/chaincode
	{
	  "jsonrpc": "2.0",
	  "method": "invoke",
	  "params": {
	    "type": 1,
	    "chaincodeID": {
	      "name": "8790b59c32feacc3c5eedf10ef8e8362d99da1c4753aa60b476e6f9ea9cea02143c482abce760a9e38e191f5fcb501ca1e7668143c296b68d68e95fb3965eadc"
	    },
	    "ctorMsg": {
	      "function": "addTX",
	      "args": [
	        "1","2"
	      ]
	    },
	    "secureContext": "user_type1_0"
	  },
	  "id": 0
	}

query：
	queryAgencyID查询当前兼职对应的中介ID
	参数有1个："JobId"

	querySalary查询当前兼职对应的薪水
	参数有1个："JobId"

	queryJobInfo查询当前兼职的信息
	参数有1个："JobId"

	query的POST请求
	Post https://a6377d73838047d39f8527f035520915-vp0.us.blockchain.ibm.com:5002/chaincode
	{
	  "jsonrpc": "2.0",
	  "method": "query",
	  "params": {
	    "type": 1,
	    "chaincodeID": {
	      "name": "8790b59c32feacc3c5eedf10ef8e8362d99da1c4753aa60b476e6f9ea9cea02143c482abce760a9e38e191f5fcb501ca1e7668143c296b68d68e95fb3965eadc"
	    },
	    "ctorMsg": {
	      "function": "queryJobInfo",
	      "args": [
	        "2"
	      ]
	    },
	    "secureContext": "user_type1_0"
	  },
	  "id": 0
	}

模拟数据

{\"JobDetail\": {\"Demand\": \"\",\"Day\": \"\",\"Place\": \"\",\"Salary\": \"50\",\"JobTime\": \"\"},\"JobID\": \"j1\",\"UserID\": \"1\",\"UserName\": \"\",\"TotalApplied\": \"123\",\"TotalHired\": \"123\",\"TotalSettled\": \"123\",\"TotalWaitCheck\": \"123\",\"Txs\": [\"123\",\"123\"]}

{\"JobDetail\": {\"Demand\": \"\",\"Day\": \"\",\"Place\": \"\",\"Salary\": \"50\",\"JobTime\": \"\"},\"JobID\": \"j2\",\"UserID\": \"2\",\"UserName\": \"\",\"TotalApplied\": \"123\",\"TotalHired\": \"123\",\"TotalSettled\": \"123\",\"TotalWaitCheck\": \"123\",\"Txs\": [\"123\",\"123\"]}
