package habit

type Habit struct {
	ID         uint64 `json:"id"`
	UserId     uint64 `json:"userId"`
	DaysOfWeek [7]int `json:"daysOfWeek"`
}
