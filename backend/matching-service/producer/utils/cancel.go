package utils

import "fmt"

var cancelledUsers []string

func IsUserCancelled(user string) bool {
	fmt.Println("Checking if user is cancelled: " + user)
	for _, cancelledUser := range cancelledUsers {
		if cancelledUser == user {
			return true
		}
	}
	return false
}

func ResetUser(userToCancel string) {
	deletionIndex := -1
	for index, user := range cancelledUsers {
		if user == userToCancel {
			deletionIndex = index
		}
	}
	if deletionIndex != -1 {
		cancelledUsers[deletionIndex] = cancelledUsers[len(cancelledUsers)-1]
		cancelledUsers = cancelledUsers[:len(cancelledUsers)-1]
		return
	}
	return
}

func CancelUser(user string) {
	cancelledUsers = append(cancelledUsers, user)
}

func PrintCancelledUsers() {
	fmt.Println(cancelledUsers)
}
