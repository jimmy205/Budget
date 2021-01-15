package repository

// Budget 預算
type Budget struct {
	Data map[string]int
}

// DefaultBudget DefaultBudget
var DefaultBudget *Budget

// NewBudget 新的預算
func NewBudget() {

	DefaultBudget = &Budget{
		Data: make(map[string]int),
	}

	return
}

// BudgetData BudgetData
type BudgetData struct {
	YearMonth string
	Amount    int
}

// GetAll 假的取得全部資料
func GetAll() map[string]int {

	// for yearMonth, amount := range b.Data {
	// 	result = append(result, BudgetData{
	// 		YearMonth: yearMonth,
	// 		Amount:    amount,
	// 	})
	// }

	return DefaultBudget.Data
}
