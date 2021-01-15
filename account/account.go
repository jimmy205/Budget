package account

import (
	"budget/package/calendar"
	"budget/repository"

	"time"
)

// Accounting Accounting
type Accounting struct {
	budget map[string]float64
}

// NewAccounting NewAccounting
func NewAccounting() *Accounting {

	accounting := &Accounting{}

	return accounting
}

// TotalAmount 回傳可用預算
func (a *Accounting) TotalAmount(start, end time.Time) (total float64) {

	// 取得金額對照表
	a.setupBudgetMap()

	// 檢查非法時間
	if end.Before(start) {
		return
	}

	tmp := start
	for {

		total += a.monthBudget(tmp)

		// 每跑一次加一個月，直到相同的月份出現
		if a.sameMonth(tmp, end) {
			break
		}

		tmp = calendar.AddMonth(tmp)
	}

	// 減掉多餘的預算
	total -= a.subRedundant(start, end)

	return
}

// setupBudgetMap 初始化預算對照表
func (a *Accounting) setupBudgetMap() {
	budgets := repository.GetAll()

	a.budget = map[string]float64{}
	for _, budget := range budgets {
		a.budget[budget.YearMonth] = budget.Amount
	}

}

// sameMonth 是否為同一個月
func (a *Accounting) sameMonth(start, end time.Time) bool {
	return start.Format("200601") == end.Format("200601")
}

func (a *Accounting) subRedundant(start, end time.Time) (sub float64) {
	// 處理開始日前多餘的預算
	before := start.Day() - 1
	sub += float64(before) * a.dayBudget(start)

	// 處理結束日後多餘的預算
	after := calendar.GetDates(end) - end.Day()
	sub += float64(after) * a.dayBudget(end)

	return
}

// monthBudget 單個月的預算
func (a *Accounting) monthBudget(t time.Time) float64 {

	return a.budget[t.Format("200601")]
}

// dayBudget 一天的預算
func (a *Accounting) dayBudget(t time.Time) float64 {

	monthBudget := a.monthBudget(t)

	return monthBudget / float64(calendar.GetDates(t))
}
