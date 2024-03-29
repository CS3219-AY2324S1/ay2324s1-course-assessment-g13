package handlers

const (
	ERROR_OCCURRED                     = "An Error Occurred"
	INVALID_JSON_REQUEST               = "Invalid JSON Request!"
	INVALID_USER_INPUT                 = "Invalid User Input!"
	INVALID_USER_EXIST                 = "User Already Exists!"
	INVALID_USER_NOT_FOUND             = "User Not Found!"
	INVALID_DB_ERROR                   = "Error Occured when Accessing Database"
	INVALID_SUPERADMIN_KEY             = "Super Admin Key Invalid"
	FAILURE_HASHING_PASSWORD           = "Double Check Your Credentials"
	FAILURE_CREATE_USER                = "Failed to Create User!"
	FAILURE_USER_ALREADY_LOGIN         = "User Already Logged In"
	FAILURE_USER_ROLE_HIGHEST          = "User Role is Already Highest"
	FAILURE_USER_ROLE_LOWEST           = "User Role is Already Lowest"
	FAILURE_UPGRADE_SUPERADMIN         = "Failed to Upgrade to Super Admin"
	FAILURE_ALREADY_SUPERADMIN         = "User is Currently Super Admin"
	FAILURE_DELETE_SUPERADMIN          = "Super Admin User Cannot be Deleted!"
	FAILURE_NOT_SUPERADMIN_CHANGE_ROLE = "User not Super Admin, Cannot Change Role!"
	FAILURE_NOT_SUPERADMIN_GET_USERS   = "User not Super Admin, Cannot Get Users!"
	SUCCESS_USER_FOUND                 = "User Found!"
	SUCCESS_USER_CREATED               = "User Created Successfully!"
	SUCCESS_USER_DELETED               = "User Deleted Successfully!"
	SUCCESS_LOGIN                      = "Login Successfully!"
	SUCCESS_LOGOUT                     = "Logout Successfully!"
	SUCCESS_ROLE_UPGRADED              = "User Role Upgraded Successfully!"
	SUCCESS_ROLE_UPGRADED_SUPER_ADMIN  = "User Role Upgraded to Super Admin Successfully!"
	SUCCESS_ROLE_DOWNGRADED            = "User Role Downgraded Successfully!"
	SUCCESS_TOKEN_REFRESHED            = "Token Refreshed Successfully!"
)

const (
	TOKEN_CLAIMS_CONTEXT_KEY    = "token-claims"
	USER_CONTEXT_KEY            = "user"
	SUCCESS_MESSAGE_CONTEXT_KEY = "success-message"
)

const (
	ACCESS_TOKEN_COOKIE_NAME  = "access-token"
	REFRESH_TOKEN_COOKIE_NAME = "refresh-token"
)

const (
	USER        = "user"
	ADMIN       = "admin"
	SUPER_ADMIN = "super admin"
)

const USER_ROLE_KEY_REQUEST_HEADER = "X-User-Role"
