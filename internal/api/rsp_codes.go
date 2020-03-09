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

	RspCodeAccountBookErrorAccountTimeEmpty		= 2002001
	RspCodeAccountBookErrorAccountTimeInvalid	= 2002002
	RspCodeAccountBookErrorAccountTypeInvalid = 2002003
	RspCodeAccountBookErrorMoneyNegative			= 2002004
	RspCodeAccountBookErrorDescriptionEmpty		= 2002005
	RspCodeAccountBookErrorAccountBooksEmpty	= 2002006
	RspCodeAccountBookErrorAccountIdEmpty			= 2002007

)

