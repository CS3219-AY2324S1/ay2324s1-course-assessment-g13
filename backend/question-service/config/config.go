package config

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"question-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var ctx = context.TODO()

const minDocuments int64 = 5

func ConnectDb() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load environment variables, with error: ", err)
		os.Exit(2)
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database("questions-service").Collection("questions")
}

func PopulateDb() {

	count, err := Collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if count >= minDocuments {
		return
	}

	questions := []models.Question{
		{
			Title: "Reverse a String",
			Description: `Write a function that reverses a string. The input string is given as an array of characters s.
					You must do this by modifying the input array in-place with O(1) extra memory.
					Example 1:
					Input: s = ["h","e","l","l","o"]
					Output: ["o","l","l","e","h"]

					Example 2:
					Input: s = ["H","a","n","n","a","h"]
					Output: ["h","a","n","n","a","H"]

					Constraints:
					• 1 <= s.length <= 105
					• s[i] is a printable ascii character.`,
			Categories: []string{"Strings", "Algorithms"},
			Complexity: "Easy",
		},
		{
			Title: "Course Schedule",
			Description: `There are a total of numCourses courses you have to take, labeled from 0 to
			numCourses - 1. You are given an array prerequisites where prerequisites[i]
			= [ai, bi] indicates that you must take course bi first if you want to take
			course ai.
			For example, the pair [0, 1], indicates that to take course 0 you have to first
			take course 1.
			Return true if you can finish all courses. Otherwise, return false.
			
			Example 1:
			Input: numCourses = 2, prerequisites = [[1,0]]
			Output: true
			Explanation: There are a total of 2 courses to take.
			To take course 1 you should have finished course 0. So it is possible.
			Example 2:
			Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
			Output: false
			Explanation: There are a total of 2 courses to take.
			To take course 1 you should have finished course 0, and to take course 0
			you should also have finished course 1. So it is impossible.
			
			Constraints:
			• 1 <= numCourses <= 2000
			• 0 <= prerequisites.length <= 5000
			• prerequisites[i].length == 2
			• 0 <= ai, bi < numCourses
			• All the pairs prerequisites[i] are unique.`,
			Categories: []string{"Data Structures", "Algorithms"},
			Complexity: "Medium",
		},
		{
			Title: "Add Binary",
			Description: `Given two binary strings a and b, return their sum as a binary string.
			Example 1:
			Input: a = "11", b = "1"
			Output: "100"
			Example 2:
			Input: a = "1010", b = "1011"
			Output: "10101"
			
			Constraints:
			•  1 <= a.length, b.length <= 104
			•  a and b consist only of '0' or '1' characters.
			•  Each string does not contain leading zeros except for the zero itself.`,
			Categories: []string{"Bit Manipulation", "Algorithms"},
			Complexity: "Easy",
		},
		{
			Title: "Longest Common Subsequence",
			Description: `Given two strings text1 and text2, return the length of their longest
			common subsequence. If there is no common subsequence, return 0.
			A subsequence of a string is a new string generated from the original string
			with some characters (can be none) deleted without changing the relative
			order of the remaining characters.
			For example, "ace" is a subsequence of "abcde".
			A common subsequence of two strings is a subsequence that is common to
			both strings.
			Example 1:
			Input: text1 = "abcde", text2 = "ace"
			Output: 3
			Explanation: The longest common subsequence is "ace" and its length is 3.
			Example 2:
			Input: text1 = "abc", text2 = "abc"
			Output: 3
			Explanation: The longest common subsequence is "abc" and its length is 3.
			Example 3:
			Input: text1 = "abc", text2 = "def"
			Output: 0
			Explanation: There is no such common subsequence, so the result is 0.
			
			Constraints:
			• 1 <= text1.length, text2.length <= 1000
			• text1 and text2 consist of only lowercase English characters.`,
			Categories: []string{"Strings", "Algorithms"},
			Complexity: "Medium",
		},
		{
			Title: "Chalkboard XOR Game",
			Description: `You are given an array of integers nums represents the numbers written on
			a chalkboard.
			Alice and Bob take turns erasing exactly one number from the chalkboard,
			with Alice starting first. If erasing a number causes the bitwise XOR of all the
			elements of the chalkboard to become 0, then that player loses. The bitwise
			XOR of one element is that element itself, and the bitwise XOR of no
			elements is 0.
			Also, if any player starts their turn with the bitwise XOR of all the elements
			of the chalkboard equal to 0, then that player wins.
			Return true if and only if Alice wins the game, assuming both players play
			optimally.
			Example 1:
			Input: nums = [1,1,2]
			Output: false
			Explanation:
			Alice has two choices: erase 1 or erase 2.
			If she erases 1, the nums array becomes [1, 2]. The bitwise XOR of all the
			elements of the chalkboard is 1 XOR 2 = 3. Now Bob can remove any
			element he wants, because Alice will be the one to erase the last element
			and she will lose.
			If Alice erases 2 first, now nums become [1, 1]. The bitwise XOR of all the
			elements of the chalkboard is 1 XOR 1 = 0. Alice will lose.
			Example 2:
			Input: nums = [0,1]
			Output: true
			Example 3:
			Input: nums = [1,2,3]
			Output: true
			Constraints:
			• 1 <= nums.length <= 1000
			• 0 <= nums[i] < 2^16`,
			Categories: []string{"Brainteaser"},
			Complexity: "Hard",
		},
	}

	for _, question := range questions {
		if err := Collection.FindOne(context.TODO(), bson.M{"title": question.Title}).Err(); err == nil {
			continue
		}
		_, err := Collection.InsertOne(context.TODO(), question)
		if err != nil {
			log.Fatal(err)
		}
	}
}
