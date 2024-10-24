package main

import "fmt"

func main() {
	// if - else Statement
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult.") // Jika umur 18 atau lebih, tampilkan ini
	} else {
		fmt.Println("You are a minor.") // Jika umur kurang dari 18, tampilkan ini
	}

	// if - else if - else Statement
	score := 85

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}

	// Variabel Temporary Pada if - else
	if age := 18; age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// switch Statement
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Today is Monday.")
	case "Tuesday":
		fmt.Println("Today is Tuesday.")
	case "Wednesday":
		fmt.Println("Today is Wednesday.")
	default:
		fmt.Println("Unknown day.")
	}

	// Pemanfaatan case Untuk Banyak Kondisi
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// Kurung Kurawal({}) Pada Keyword case & default
	switch day {
	case "Monday":
		{
			fmt.Println("Start of the week!")
			fmt.Println("Back to work.")
		}
	case "Friday":
		{
			fmt.Println("Almost weekend!")
		}
	default:
		fmt.Println("It's a regular day.")
	}

	// Penggunaan Style if-else Pada switch Statement
	switch {
	case age < 18:
		fmt.Println("You are a minor.")
	case age >= 18 && age < 65:
		fmt.Println("You are an adult.")
	default:
		fmt.Println("You are a senior.")
	}

	// Keyword fallthrough Dalam switch Statement
	grade := "B"

	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Good")
		fallthrough
	case "C":
		fmt.Println("You passed")
	default:
		fmt.Println("Invalid grade")
	}

	// Nesting Conditional Statement (Seleksi Kondisi Bersarang)
	country := "USA"

	if age >= 18 {
		fmt.Println("You are an adult.")

		// Bersarang dengan switch-case
		switch country {
		case "USA":
			fmt.Println("You are in the USA.")
		case "UK":
			fmt.Println("You are in the UK.")
		case "Canada":
			fmt.Println("You are in Canada.")
		default:
			fmt.Println("Country not recognized.")
		}

	} else {
		fmt.Println("You are a minor.")

		// Bersarang dengan if-else
		if score >= 80 {
			fmt.Println("You passed the exam with a good score.")
		} else if score >= 60 {
			fmt.Println("You passed the exam with an average score.")
		} else {
			fmt.Println("You failed the exam.")
		}
	}
}
