package handlers

const (
	INVALID_JSON_REQUEST       = "Invalid JSON Request!"
	INVALID_USER_INPUT         = "Invalid User Input!"
	INVALID_USER_EXIST         = "Username Already Exists!"
	INVALID_USER_NOT_FOUND     = "User Not Found!"
	FAILURE_HASHING_PASSWORD   = "An Error Occurred while Hashing Password"
	FAILURE_CREATE_USER        = "Failed to Create User!"
	FAILURE_USER_ALREADY_LOGIN = "User Already Logged In"
	FAILURE_USER_ROLE_HIGHEST  = "User Role is Already Highest"
	FAILURE_USER_ROLE_LOWEST   = "User Role is Already Lowest"
	SUCCESS_USER_FOUND         = "User Found!"
	SUCCESS_USER_CREATED       = "User Created Successfully!"
	SUCCESS_USER_DELETED       = "User Deleted Successfully!"
	SUCCESS_LOGIN              = "Login Successfully"
	SUCCESS_LOGOUT             = "Logout Successfully"
	SUCCESS_ROLE_UPGRADED      = "User Role Upgraded Successfully!"
	SUCCESS_ROLE_DOWNGRADED    = "User Role Downgraded Successfully!"
	SUCCESS_TOKEN_REFRESHED    = "Token Refreshed Successfully!"
)

const (
	JWT_COOKIE_NAME             = "jwt"
	TOKEN_CLAIMS_CONTEXT_KEY    = "token-claims"
	USER_CONTEXT_KEY            = "user"
	SUCCESS_MESSAGE_CONTEXT_KEY = "success-message"
	EXPIRATION_TIME_CONTEXT_KEY = "expiration-time"
	GITHUB_DATA_CONTEXT_KEY     = "github-data"
)

const (
	USER  = "user"
	ADMIN = "admin"
)
