package account

import (
	"budget/repository"

	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type AccountingSuite struct {
	suite.Suite
	*Accounting
}

// TestAccountingInit TestAccountingInit
func TestAccountingInit(t *testing.T) {
	suite.Run(t, new(AccountingSuite))
}

func (t *AccountingSuite) SetupTest() {
	t.Accounting = NewAccounting()
}

// 測試非法日期
func (t *AccountingSuite) Test_InvalidTime() {

	repository.NewBudget()

	mockData := []repository.BudgetDetail{
		{
			YearMonth: "202101",
			Amount:    3100,
		},
	}

	repository.DefaultBudget.Data = mockData

	startDate, _ := time.Parse("2006-01-02", "2021-01-02")
	endDate, _ := time.Parse("2006-01-02", "2021-01-01")

	t.Equal(float64(0), t.Accounting.TotalAmount(startDate, endDate))

}

// 測試取得單日
func (t *AccountingSuite) Test_GetSameDay() {

	// 初始化一個新的預算表
	repository.NewBudget()

	mockData := []repository.BudgetDetail{
		{
			YearMonth: "202101",
			Amount:    3100,
		},
	}

	repository.DefaultBudget.Data = mockData

	startDate, _ := time.Parse("2006-01-02", "2021-01-01")
	endDate, _ := time.Parse("2006-01-02", "2021-01-01")

	t.Equal(float64(100), t.Accounting.TotalAmount(startDate, endDate))
}

// 測試取得當月多日
func (t *AccountingSuite) Test_GetPartOfMonth() {

	repository.NewBudget()

	mockData := []repository.BudgetDetail{
		{
			YearMonth: "202101",
			Amount:    3100,
		},
	}

	repository.DefaultBudget.Data = mockData

	startDate, _ := time.Parse("2006-01-02", "2021-01-01")
	endDate, _ := time.Parse("2006-01-02", "2021-01-02")

	t.Equal(float64(200), t.Accounting.TotalAmount(startDate, endDate))

}

// 測試取得無預算
func (t *AccountingSuite) Test_ZeroAmount() {

	repository.NewBudget()

	mockData := []repository.BudgetDetail{
		{
			YearMonth: "202101",
			Amount:    0,
		},
	}

	repository.DefaultBudget.Data = mockData

	startDate, _ := time.Parse("2006-01-02", "2021-01-01")
	endDate, _ := time.Parse("2006-01-02", "2021-01-02")

	t.Equal(float64(0), t.Accounting.TotalAmount(startDate, endDate))
}

// 測試取得跨月
func (t *AccountingSuite) Test_CrossMonth() {

	repository.NewBudget()

	mockData := []repository.BudgetDetail{
		{
			YearMonth: "202101",
			Amount:    3100,
		}, {
			YearMonth: "202102",
			Amount:    280,
		},
	}

	repository.DefaultBudget.Data = mockData

	startDate, _ := time.Parse("2006-01-02", "2021-01-30")
	endDate, _ := time.Parse("2006-01-02", "2021-02-02")

	t.Equal(float64(220), t.Accounting.TotalAmount(startDate, endDate))
}

// 測試取得跨年
func (t *AccountingSuite) Test_CrossYear() {

	repository.NewBudget()

	mockData := []repository.BudgetDetail{
		{
			YearMonth: "202112",
			Amount:    3100,
		}, {
			YearMonth: "202201",
			Amount:    31,
		},
	}

	repository.DefaultBudget.Data = mockData

	startDate, _ := time.Parse("2006-01-02", "2021-12-30")
	endDate, _ := time.Parse("2006-01-02", "2022-01-02")

	t.Equal(float64(202), t.Accounting.TotalAmount(startDate, endDate))
}
