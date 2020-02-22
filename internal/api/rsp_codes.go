package api

const (
	RspCodeSuccess                    				= 0
	RspCodeDbError														= 1001001
	RspCodeReqDataInvalid											= 1001002
	RspCodeSaveSessionError										= 1001003

	RspCodeNoAuth                 						= 2001001
	RspCodeUserErrorUsernameEmpty 						= 2001002
	RspCodeUserErrorPhoneEmpty    						= 2001003
	RspCodeUserErrorPasswordEmpty 						= 2001004
	RspCodeUserErrorPhoneExists   						= 2001005
	RspCodeUserErrorPhoneNoExists 						= 2001006
	RspCodeUserErrorPasswordInvalid 					= 2001007

)

