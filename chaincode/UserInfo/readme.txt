UserInfo智能合约：每一个用户都在这一个合约中创建，修改，删除...

数据结构：
	{
	    "Balance": "100",
	    "CreditScore": {
	        "CurrentCreditScore": "6",
	        "RateTimes": "1",
	        "TotalCreditScore": "6"
	    },
	    "Jobs": [
	        "123",
	        "123"
	    ],
	    "UserInfo": {
	        "AgencyName": "",
	        "BCID": "123",
	        "Gender": "0",
	        "Password": "123",
	        "RealName": "",
	        "Role": "0",
	        "School": "",
	        "Status": "1",
	        "StuID": "",
	        "Tele": "",
	        "UserID": "1",
	        "Username": "abc"
	    }
	}

world status：
	key		  value
	UserID	  UserInfo
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
	      "path": "https://github.com/yglym2012/BlockchainCompetition/chaincode/UserInfo"
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
	add添加一个新用户
	参数有2个："UserId","UserInfo"

	delete删除一个老用户
	参数有1个："UserId"

	edit修改用户信息
	参数有2个："UserID","NewUserInfo"

	creditScoreEdit修改用户信用积分
	参数有2个："UserID","NewScoreFromOthersNow"

	addTX添加和该用户相关的兼职ID（对于中介）或TXID（对于学生）
	参数有2个："UserID","TXID"

	autoSettle结算即修改用户账户余额
	参数有3个："StuID","AgencyID","Salary"

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
	queryCurrentCreditScore查询当前用户的信用积分
	参数有1个："UserID"

	queryUserInfo查询当前用户的信息
	参数有1个："UserID"

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

{\"Balance\": \"100\",\"CreditScore\": {\"CurrentCreditScore\": \"6\",\"RateTimes\": \"1\",\"TotalCreditScore\": \"6\"},\"Jobs\": [\"123\",\"123\"],\"UserInfo\": {\"AgencyName\": \"\",\"BCID\": \"123\",\"Gender\": \"0\",\"Password\": \"123\",\"RealName\": \"\",\"Role\": \"0\",\"School\": \"\",\"Status\": \"1\",\"StuID\": \"\",\"Tele\": \"\",\"UserID\": \"1\",\"Username\": \"abc\"}}

{\"Balance\": \"100\",\"CreditScore\": {\"CurrentCreditScore\": \"6\",\"RateTimes\": \"1\",\"TotalCreditScore\": \"6\"},\"Jobs\": [\"123\",\"123\"],\"UserInfo\": {\"AgencyName\": \"\",\"BCID\": \"123\",\"Gender\": \"0\",\"Password\": \"123\",\"RealName\": \"\",\"Role\": \"0\",\"School\": \"\",\"Status\": \"1\",\"StuID\": \"\",\"Tele\": \"\",\"UserID\": \"2\",\"Username\": \"abc\"}}

{\"Balance\": \"100\",\"CreditScore\": {\"CurrentCreditScore\": \"6\",\"RateTimes\": \"1\",\"TotalCreditScore\": \"6\"},\"Jobs\": [\"123\",\"123\"],\"UserInfo\": {\"AgencyName\": \"\",\"BCID\": \"123\",\"Gender\": \"0\",\"Password\": \"123\",\"RealName\": \"\",\"Role\": \"0\",\"School\": \"\",\"Status\": \"1\",\"StuID\": \"\",\"Tele\": \"\",\"UserID\": \"3\",\"Username\": \"abc\"}}

