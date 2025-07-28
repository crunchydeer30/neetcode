package main

func countStudents(students []int, sandwiches []int) int {
	studentsPreferences := frequencyAnalysis(students)
	studentsInQueue := len(students)

	for _, sandwich := range sandwiches {
		if studentsPreferences[sandwich] == 0 {
			break
		}
		studentsPreferences[sandwich]--
		studentsInQueue--
	}

	return studentsInQueue
}

func frequencyAnalysis(arr []int) map[int]int {
	hashmap := make(map[int]int)

	for _, val := range arr {
		hashmap[val]++
	}

	return hashmap
}
