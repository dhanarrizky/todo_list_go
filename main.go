package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/dhanarrizky/go_todoList/models"
	"github.com/dhanarrizky/go_todoList/validation"
)

func main() {
	inp := mainMenu()

	switch inp {
	case "1":
		AllTodo()
	case "2":
		AddNewTask()
	case "3":
		break
	default:
		fmt.Println("Enter numbers from 1-3, please repeat again")
		main()
	}

}

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func mainMenu() string {

	fmt.Println("=========================================================================================")
	fmt.Println("Main Menu")
	fmt.Print("1. All TodoList\n",
		"2. Add New Task\n",
		"3. Exit Application\n\n")
	fmt.Print("Choice One : ")
	inp := ""
	fmt.Scan(&inp)
	fmt.Println("=========================================================================================")
	return inp

}

func AllTodo() {
	data := models.GetAllData()
	fmt.Println("=========================================================================================")
	fmt.Println("|No\t| Name\t\t| StartTime\t\t| EndTime\t\t| Status\t|")
	fmt.Println("=========================================================================================")
	if len(data) != 0 {
		for index, value := range data {
			if value.Status == false {
				fmt.Println("|", index, "\t|", value.Name, "\t|", value.StartTime.Format("2006-01-02"), "\t\t|", value.EndTime.Format("2006-01-02"), "\t\t|",
					value.Status, "\t|")
				fmt.Println("-----------------------------------------------------------------------------------------")
			}
		}
		fmt.Println("=========================================================================================")
		fmt.Println("Do you want a checklist of completed work?")
		fmt.Println("1. Yes\n" +
			"2. No")
		var inp string
		fmt.Scan(&inp)
		switch inp {
		case "1":
			checklist()
		case "2":
			mainMenu()
		default:
			fmt.Println("Enter numbers from 1-3, please repeat again")
			AllTodo()
		}
	} else {
		fmt.Println("data is empty!!")
	}
	main()
}

func checklist() {
	var inp string
	fmt.Println("Which task number do you want to checklist?")
	fmt.Scan(&inp)

	data := models.GetAllData()

	result := false

	for index, val := range data {
		if inp == index {
			result = true
			val.Status = true
			models.UpdateData(val)
			fmt.Println("data update has been successfully!!")
			mainMenu()
		}
	}

	if result == false {
		fmt.Println("No not found, please try again")
		checklist()
	}
}

func AddNewTask() {
	var todo models.TodoList
	lenTodo := models.GetAllData()

	size := len(lenTodo)
	todo.Id = uint(size)

	var name string
	fmt.Println("=========================================================================================")
	fmt.Println("|\t\t\tLet's to do\t\t\t")
	fmt.Println("=========================================================================================")
	fmt.Print("Entery Task : ")
	fmt.Scan(&name)
	fmt.Println("Fotmating Time like this : 02/01/2006")
	fmt.Println("Entery Start Time : ")
	sy, sm, sd := validation.YearValidate()
	fmt.Println("Entery End Time : ")
	ey, em, ed := validation.YearValidate()
	fmt.Println("=========================================================================================")
	sTime := time.Date(sy, time.Month(sm), sd, 0, 0, 0, 0, time.UTC)
	eTime := time.Date(ey, time.Month(em), ed, 0, 0, 0, 0, time.UTC)

	todo.Name = name
	todo.StartTime = sTime
	todo.EndTime = eTime
	models.CreateNewData(todo)
	main()
}

func History() {
	data := models.GetAllData()
	fmt.Println("=========================================================================================")
	fmt.Println("|No\t| Name\t\t| StartTime\t\t| EndTime\t\t| Status\t|")
	fmt.Println("=========================================================================================")
	for index, value := range data {
		fmt.Println("|", index, "\t|", value.Name, "\t|", value.StartTime.Format("2006-01-02"), "\t\t|", value.EndTime.Format("2006-01-02"), "\t\t|",
			value.Status, "\t|")
		fmt.Println("-----------------------------------------------------------------------------------------")
	}

	fmt.Println("=========================================================================================")
	main()
}
