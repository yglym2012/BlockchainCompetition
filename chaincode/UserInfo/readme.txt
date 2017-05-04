UserInfo智能合约：每一个用户都在这一个合约中创建，修改，删除...

数据结构：
	{
		"Balance": 0, 
		"CreditScore": {
		  "CurrentCreditScore": 0, 
		  "RateTimes": 0, 
		  "TotalCreditScore": 0
		}, 
		"Jobs": [
		  "123", 
		  "123"
		], 
		"UserInfo": {
		  "AgencyName": "", 
		  "BCID": "123", 
		  "Gender": 0, 
		  "Password": "123", 
		  "RealName": "", 
		  "Role": 0, 
		  "School": "", 
		  "Status": 1, 
		  "StuID": "", 
		  "Tele": "", 
		  "UserID": "", 
		  "Username": "abc"
		}
	}

world status：
	key		  value
	UserID	  用户的ID即身份证号
	UserInfo  用户信息（即上面数据结构中的信息）
	注释：key这一列下面，带引号就是key的实际值，不带引号的是变量，变量是什么，key的值是什么

deploy：
	init的参数有0个

	init的POST请求
	Post Addr:Port/chaincode
	{
	  "jsonrpc": "2.0",
	  "method": "deploy",
	  "params": {
	    "type": 1,
	    "chaincodeID":{
	        "path":"ChaincodePath"
	    },
	    "ctorMsg": {
	        "args":["init"]
	    },
	    "secureContext": VPName
	  },
	  "id": 1
	}

invoke：
	Add添加一个新用户
	参数有2个："UserId","UserInfo"

	Delete删除一个老用户
	参数有1个："UserId"

	Edit修改用户信息
	参数有2个："UserID","NewUserInfo"

	CreditScoreEdit修改用户信用积分
	参数有2个："UserID","NewScoreFromOthersNow"

	AddTX添加和该用户相关的兼职ID（对于中介）或TXID（对于学生）
	参数有2个："UserID","TXID"

	AutoSettle结算即修改用户账户余额
	参数有3个："StuID","AgencyID","Salary"

	invoke的POST请求
	Post Addr:Port/chaincode
	{
	  "jsonrpc": "2.0",
	  "method": "invoke",
	  "params": {
	    "type": 1,
	    "chaincodeID":{
	        "name":"ChaincodeID"
	    },
	    "ctorMsg": {
	        "args":["add","UserID","UserInfo"]
	    },
	    "secureContext": VPName
	  },
	  "id": 1
	}

query：
	QueryCurrentCreditScore查询当前用户的信用积分
	参数有1个："UserID"

	QueryUserInfo查询当前用户的信息
	参数有1个："UserID"

	query的POST请求
	Post Addr:Port/chaincode
	{
	  "jsonrpc": "2.0",
	  "method": "query",
	  "params": {
	    "type": 1,
	    "chaincodeID":{
	        "name":"ChaincodeID"
	    },
	    "ctorMsg": {
	        "args":["query",key]
	    },
	    "secureContext": VPName
	  },
	  "id": 3
	}