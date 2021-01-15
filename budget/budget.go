package budget

import (
	"budget/repository"
	"time"
)

// Accounting Accounting
type Accounting struct {
}

// NewAccounting NewAccounting
func NewAccounting() *Accounting {
	return &Accounting{}
}

// TotalAmount 回傳可用預算
func (a *Accounting) TotalAmount(start, end time.Time) (total float64) {

	// 檢查非法時間
	if end.Sub(start).Hours() < 0 {
		return 0
	}

	// 檢查是否有跨月
	if !a.isSameMonth(start, end) {

		return a.crossMonth(start, end)
	}

	// get budget data
	budgetData := repository.GetAll()

	startStr := start.Format("200601")

	amount := budgetData[startStr]

	amountOfDay := amount / a.getDayOfMonth(start)

	diffDay := (end.Sub(start).Hours() / 24) + 1

	return float64(amountOfDay * int(diffDay))
}

func (a *Accounting) crossMonth(start, end time.Time) float64 {

	// 取得資料
	budgetData := repository.GetAll()

	// 取得開始的日期
	startDay := start.Day()
	endOfStartMonth := a.getDayOfMonth(start)
	fistDay := endOfStartMonth - startDay + 1
	firstAmount := budgetData[start.Format("200601")]
	first := (firstAmount / endOfStartMonth) * fistDay

	// 取得結束的日期
	endDay := start.Day()
	endAmount := budgetData[end.Format("200601")] / a.getDayOfMonth(end)

	return float64(first + endAmount*endDay)
}

func (a *Accounting) getDayOfMonth(date time.Time) (
	dayOfMonth int,
) {

	year, month, _ := date.Date()

	dayOfMonth = time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	return
}

func (a *Accounting) isSameMonth(start, end time.Time) bool {

	if start.Month() == end.Month() {
		return true
	}

	return false
}
