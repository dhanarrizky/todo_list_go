package validation

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/dhanarrizky/go_todoList/models"
)

func YearValidate() (year, month, day int) {
	var Year, Month, Day string
	fmt.Print("Year :")
	fmt.Scan(&Year)
	fmt.Print("Month :")
	fmt.Scan(&Month)
	fmt.Print("Day :")
	fmt.Scan(&Day)
	if len(Year) < 4 {
		err := errors.New("Maybe you entered the wrong year!!!, please Try again!!")
		fmt.Println(err)
		return YearValidate()
	}

	y, err := strconv.Atoi(Year)
	if err != nil {
		fmt.Println(err)
		return YearValidate()
	}

	m, err := strconv.Atoi(Month)
	if err != nil {
		fmt.Println(err)
		return YearValidate()
	}

	if m > 12 || m < 1 {
		err := errors.New("Maybe you entered the wrong month!!!, please try again!!")
		fmt.Println(err)
		return YearValidate()
	}

	d, err := strconv.Atoi(Day)
	if err != nil {
		fmt.Println(err)
		return YearValidate()
	}

	if d > 12 || d < 1 {
		err := errors.New("Maybe you entered the wrong month!!!, please try again!!")
		fmt.Println(err)
		return YearValidate()
	}

	value := checkTime(y, m, d)
	if value == false {
		fmt.Println("time not available!!, please try again !!!")
		return YearValidate()
	}

	return y, m, d
}

func checkTime(y, m, d int) bool {
	checktimes := models.GetAllData()
	Times := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)

	result := true

	for _, value := range checktimes {
		if !(Times.Before(value.StartTime) || Times.After(value.EndTime)) {
			result = false
			return result
		}
	}

	return result
}
