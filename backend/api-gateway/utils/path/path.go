package path

const (
	REGISTER             = "/auth/register"
	REGISTER_GITHUB      = "/auth/login/github"
	LOGIN                = "/auth/login"
	LOGIN_GITHUB         = "/auth/login/github"
	GITHUB_CALLBACK      = "/auth/login/github/callback"
	LOGOUT               = "/auth/logout"
	REFRESH              = "/auth/refresh"
	AUTH_USER            = "/auth/user"
	AUTH_USER_UPGRADE    = "/auth/user/upgrade"
	AUTH_USER_DOWNGRADE  = "/auth/user/downgrade"
	ALL_USER_SERVICE     = "/users*"
	ALL_QUESTION_SERVICE = "/questions*"
)
