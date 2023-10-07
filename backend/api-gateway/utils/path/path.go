package path

const (
	REGISTER             = "/auth/register"
	LOGIN                = "/auth/login"
	GITHUB_LOGIN         = "/auth/login/github"
	LOGOUT               = "/auth/logout"
	REFRESH              = "/auth/refresh"
	AUTH_USER            = "/auth/user"
	AUTH_USER_UPGRADE    = "/auth/user/upgrade"
	AUTH_USER_DOWNGRADE  = "/auth/user/downgrade"
	ALL_USER_SERVICE     = "/users*"
	ALL_QUESTION_SERVICE = "/questions*"
	ALL_MATCHING_SERVICE = "/match*"
)
