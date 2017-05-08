TX智能合约：一笔交易

数据结构：
	{
	    "ApplyTime": "123",
	    "JobID": "123",
	    "Status": "123",
	    "UserID": "123",
	    "TxID": "tx1",
	    "StuScore": "",
	    "AgencyScore": ""
	}

world status：
	key		  value
	TxID	  TxInfo
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
	      "path": "https://github.com/yglym2012/BlockchainCompetition/chaincode/TX"
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
	create创建一个交易并且自动审核学生申请
	参数有2个："TxId","TxInfo"

	artificialCheck中介审核学生申请
	参数有2个："TxId","Result"

	evaluate中介、学生互评
	参数有3个："TxID","UserID","Score"

	invoke的POST请求
	Post https://a6377d73838047d39f8527f035520915-vp0.us.blockchain.ibm.com:5002/chaincode
	{
	  "jsonrpc": "2.0",
	  "method": "invoke",
	  "params": {
	    "type": 1,
	    "chaincodeID": {
	      "name": "96ef2509c45f6445eb511962014ec2385ab8cf3344fa0b4d6b73dcefe671c0dd1d5c25d927932535ade78d8c87d9c7f6e5fda4d49bff64c196f38f9ebe92a4ba"
	    },
	    "ctorMsg": {
	      "function": "create",
	      "args": [
	        "1","2"
	      ]
	    },
	    "secureContext": "user_type1_0"
	  },
	  "id": 0
	}

query：
	queryTxInfo查询交易信息
	参数有1个："TxID"

	query的POST请求
	Post https://a6377d73838047d39f8527f035520915-vp0.us.blockchain.ibm.com:5002/chaincode
	{
	  "jsonrpc": "2.0",
	  "method": "query",
	  "params": {
	    "type": 1,
	    "chaincodeID": {
	      "name": "96ef2509c45f6445eb511962014ec2385ab8cf3344fa0b4d6b73dcefe671c0dd1d5c25d927932535ade78d8c87d9c7f6e5fda4d49bff64c196f38f9ebe92a4ba"
	    },
	    "ctorMsg": {
	      "function": "queryTxInfo",
	      "args": [
	        "2"
	      ]
	    },
	    "secureContext": "user_type1_0"
	  },
	  "id": 0
	}

模拟数据

{\"ApplyTime\": \"123\",\"JobID\": \"j1\",\"Status\": \"\",\"UserID\": \"1\",\"TxID\": \"tx1\",\"StuScore\": \"\",\"AgencyScore\": \"\"}

{\"ApplyTime\": \"1494147008\",\"JobID\": \"590e9e4ee588d600d4fb0aaf\",\"Status\": \"\",\"UserID\": \"1234\",\"TxID\": \"590edfc0e588d601914ab18a\",\"StuScore\": \"\",\"AgencyScore\": \"6\"}

{\"AgencyScore\": \"6\", \"UserID\": \"1\", \"ApplyTime\": \"1494147008\", \"JobID\": \"590e9e4ee588d600d4fb0aaf\", \"State\": \"\", \"StuScore\": \"6\", \"TxID\": \"590edfc0e588d601914ab18a\"}


{\"ApplyTime\": \"123\",\"JobID\": \"590e9e4ee588d600d4fb0aaf\",\"Status\": \"\",\"UserID\": \"1\",\"TxID\": \"tx1\",\"StuScore\": \"\",\"AgencyScore\": \"\"}