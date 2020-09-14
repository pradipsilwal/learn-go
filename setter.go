package main

import (
	"errors"
	"fmt"
	"log"
)

type date struct {
	year  int
	month int
	day   int
}

func (d *date) setYear(year int) error {
	if year < 1 {
		return errors.New("Invalid Year")
	}
	d.year = year
	return nil
}

func (d *date) setMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid Month")
	}
	d.month = month
	return nil
}

func (d *date) setDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid Day")
	}
	d.day = day
	return nil
}

func setter() {
	dateObj := date{}
	err := dateObj.setYear(1994)
	if err != nil {
		log.Fatal(err)
	}
	err = dateObj.setMonth(14)
	if err != nil {
		log.Fatal(err)
	}
	err = dateObj.setDay(6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d-%d-%d\n", dateObj.day, dateObj.month, dateObj.year)
	fmt.Println(dateObj)
}
