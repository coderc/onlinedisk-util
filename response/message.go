package response

const (
	SuccessCode                      = 0
	FileUploadFailedCode             = 1001
	FailedSelectFileListCode         = 1002
	FileDownloadFailedCode           = 1003
	FileDeleteFailedCode             = 1004
	FileUploadMultipleInitFailedCode = 1005
	FileUploadMultipleFailedCode     = 1006
	FileUploadMultipleNilChunkCode   = 1007
	SignupFailedCode                 = 2001
	SigninFailedCode                 = 2002
)

const (
	ErrorGetUserInfo          = "get user info failed"
	ErrorUsernameInvalid      = "username invalid"
	ErrorPasswordInvalid      = "password invalid"
	ErrorCreateUUIDFailed     = "create uuid failed"
	ErrorInsertUserInfoFailed = "insert user info failed"
	ErrorCreateUserFailed     = "create user failed"
)
