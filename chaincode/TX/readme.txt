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
	      "name": "160d9b88e83856238d689e329768e86e319047ad61aebf9e15a2c0d8636f4ad30621d60352f46012dfaf150f25d160cdb2f3cf148c611997777e1189cd218c7b"
	    },
	    "ctorMsg": {
	      "function": "autoSettle",
	      "args": [
	        "1","2","50"
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
	      "name": "160d9b88e83856238d689e329768e86e319047ad61aebf9e15a2c0d8636f4ad30621d60352f46012dfaf150f25d160cdb2f3cf148c611997777e1189cd218c7b"
	    },
	    "ctorMsg": {
	      "function": "queryUserInfo",
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

