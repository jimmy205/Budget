package repository

// Budget 預算
type Budget struct {
	Data []BudgetDetail
}

// BudgetDetail 預算詳細資料
type BudgetDetail struct {
	YearMonth string
	Amount    float64
}

// DefaultBudget DefaultBudget
var DefaultBudget *Budget

// NewBudget 新的預算
func NewBudget() {

	DefaultBudget = &Budget{}

	return
}

// BudgetData BudgetData
type BudgetData struct {
	YearMonth string
	Amount    int
}

// GetAll 取得全部資料
func GetAll() []BudgetDetail {
	return DefaultBudget.Data
}
