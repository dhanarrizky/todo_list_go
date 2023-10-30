package models

import "strconv"

// function for inport or get data from database
var data = map[string]TodoList{}

func GetAllData() map[string]TodoList {
	return data
}

func CreateNewData(t TodoList) {
	idString := strconv.Itoa(len(data) + 1)
	data[idString] = t
}

func UpdateData(t TodoList) {
	idString := strconv.Itoa(int(t.Id))
	data[idString] = t
}
