package main

import (
	"fmt"
	"log"
	"time"
)

func getWeekSelection() time.Time {
	var selection int
	currentWeek := getCurrentWeek()
	previousWeek := getPreviousWeek()

	fmt.Println("(1) Current Week -", currentWeek.Format(time.DateOnly))
	fmt.Println("(2) Previous Week -", previousWeek.Format(time.DateOnly))
	fmt.Println("(3) Custom Week")

	fmt.Scanln(&selection)

	if selection == 1 {
		return currentWeek
	} else if selection == 2 {
		return previousWeek
	} else if selection == 3 {
		var date string
		fmt.Println("Enter date in the following format: yyyy-MM-dd")
		fmt.Scanln(&date)
		parsedDate, err := time.Parse(time.DateOnly, date)
		if err != nil {
			log.Fatal("Entered invalid date format")
		} else {
			return getWeekTimeStartingSunday(parsedDate)
		}
	} else {
		log.Fatal("Invalid selection. Select 1, 2, or 3.")
	}

	log.Fatal("An invalid selection was made.")
	return time.Now()
}

func getCurrentWeek() time.Time {
	return getWeekTimeStartingSunday(time.Now())
}

func getPreviousWeek() time.Time {
	previousWeek := time.Now().AddDate(0, 0, -7)
	return getWeekTimeStartingSunday(previousWeek)
}

func getWeekTimeStartingSunday(time time.Time) time.Time {
	currentWeekDay := time.Weekday()
	switch currentWeekDay {
	case 0: // Sunday
		return time
	case 1: // Monday
		return time.AddDate(0, 0, -1)
	case 2: // Tuesday
		return time.AddDate(0, 0, -2)
	case 3: // Wednesday
		return time.AddDate(0, 0, -3)
	case 4: // Thursday
		return time.AddDate(0, 0, -4)
	case 5: // Friday
		return time.AddDate(0, 0, -5)
	case 6: // Saturday
		return time.AddDate(0, 0, -6)
	default:
		log.Fatal("Invalid date sent to getWeekTimeStartingSunday")
	}

	return time
}
