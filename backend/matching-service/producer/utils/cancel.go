package utils

import "time"

var cancelledUsers map[string][]string
var userInQueueMap map[string]bool

func Init() {
	cancelledUsers = make(map[string][]string)
	for _, criteria := range MatchCriterias {
		cancelledUsers[string(criteria)] = []string{}
	}

	userInQueueMap = make(map[string]bool)
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

func SetUserInQueue(user string) {
	userInQueueMap[user] = true
}

func RemoveUserFromQueueImmediate(user string) {
	delete(userInQueueMap, user)
}

func RemoveUserFromQueueDelay(user string) {
	time.Sleep(time.Second * 1) // Ensure match occurs first before removing from queue
	delete(userInQueueMap, user)
}

func IsUserInQueue(user string) bool {
	if v, ok := userInQueueMap[user]; ok {
		return v
	}
	return false
}
