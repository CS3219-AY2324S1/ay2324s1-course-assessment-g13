package utils

var cancelledUsers map[string][]string

func Init() {
	cancelledUsers = make(map[string][]string)
	for _, criteria := range MatchCriterias {
		cancelledUsers[string(criteria)] = []string{}
	}
}

func IsUserCancelled(user string, criteria string) bool {
	for _, cancelledUser := range cancelledUsers[criteria] {
		if cancelledUser == user {
			return true
		}
	}
	return false
}

func ResetUser(userToCancel string, criteria string) {
	deletionIndex := -1
	for index, user := range cancelledUsers[criteria] {
		if user == userToCancel {
			deletionIndex = index
		}
	}
	if deletionIndex != -1 {
		cancelledUsers[criteria][deletionIndex] = cancelledUsers[criteria][len(cancelledUsers[criteria])-1]
		cancelledUsers[criteria] = cancelledUsers[criteria][:len(cancelledUsers[criteria])-1]
	}
}

func CancelUser(user string, criteria string) {
	ResetUser(user, criteria)                                         // Resets user first by removing user if they exist
	cancelledUsers[criteria] = append(cancelledUsers[criteria], user) // Then cancel to prevent duplicates
}
