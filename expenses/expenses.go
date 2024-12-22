package expenses

import (
	"errors"
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

var cat = map[string]bool{}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	output := make([]Record, 0)
	for _, r := range in {
		cat[r.Category] = true
		if predicate(r) {
			output = append(output, r)
		}
	}
	return output
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	temp := Filter(in, ByDaysPeriod(p))
	total := 0.0
	for _, r := range temp {
		total += r.Amount
	}
	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	temp1 := Filter(in, ByCategory(c))
	_, ok := cat[c]
	if !ok {
		errmsg := fmt.Sprintf("unknown category %s", c)
		return 0.0, errors.New(errmsg)
	}

	//But it would be better to just check that the length of temp1 is not 0
	// if len(temp1) == 0 {
	//		errmsg := fmt.Sprintf("unknown category %s", c)
	//		return 0.0, errors.New(errmsg)
	//}

	total := TotalByPeriod(temp1, p)
	return total, nil

}
