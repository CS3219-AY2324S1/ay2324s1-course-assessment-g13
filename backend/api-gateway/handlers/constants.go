package handlers

const (
	INVALID_JSON_REQUEST       = "Invalid JSON Request!"
	INVALID_USER_INPUT         = "Invalid User Input!"
	INVALID_USER_EXIST         = "Username Already Exists!"
	INVALID_USER_NOT_FOUND     = "User Not Found!"
	FAILURE_HASHING_PASSWORD   = "An Error Occurred while Hashing Password"
	FAILURE_CREATE_USER        = "Failed to Create User!"
	FAILURE_USER_ALREADY_LOGIN = "User Already Logged In"
	SUCCESS_USER_FOUND         = "User Found!"
	SUCCESS_USER_CREATED       = "User Created Successfully!"
	SUCCESS_USER_DELETED       = "User Deleted Successfully!"
	SUCCESS_LOGIN              = "Login Successfully"
	SUCCESS_LOGOUT             = "Logout Successfully"
)

const (
	JWT_COOKIE_NAME  = "jwt"
	TOKEN_CLAIMS_KEY = "token-claims"
)
