package path

const (
	REGISTER                      = "/auth/register"
	SIGNUP                        = "/auth/signup"
	LOGIN                         = "/auth/login"
	LOGOUT                        = "/auth/logout"
	REFRESH                       = "/auth/refresh"
	AUTH_USER                     = "/auth/user"
	AUTH_USERS                    = "/auth/users"
	AUTH_USER_UPGRADE             = "/auth/user/upgrade"
	AUTH_USER_UPGRADE_SUPER_ADMIN = "/auth/user/upgrade-super-admin"
	AUTH_USER_DOWNGRADE           = "/auth/user/downgrade"
	HISTORY                       = "/history"
	HISTORIES                     = "/histories*"
	ALL_USER_SERVICE              = "/users*"
	ALL_QUESTION_SERVICE          = "/questions*"
	ALL_MATCHING_SERVICE          = "/match*"
	ALL_COLLAB_SERVICE            = "/ws*"
)
