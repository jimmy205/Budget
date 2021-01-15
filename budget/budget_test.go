package budget

import (
	"budget/repository"
	"time"

	"testing"

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

func (t *AccountingSuite) Test_GetSameDay() {

	repository.NewBudget()

	repository.DefaultBudget.Data["202101"] = 3100

	startDate, _ := time.Parse("2006-01-02", "2021-01-01")
	endDate, _ := time.Parse("2006-01-02", "2021-01-01")

	t.Equal(float64(100), t.Accounting.TotalAmount(startDate, endDate))
}

func (t *AccountingSuite) Test_GetPartOfMonth() {

	repository.NewBudget()

	repository.DefaultBudget.Data["202101"] = 3100

	startDate, _ := time.Parse("2006-01-02", "2021-01-01")
	endDate, _ := time.Parse("2006-01-02", "2021-01-02")

	t.Equal(float64(200), t.Accounting.TotalAmount(startDate, endDate))

}

func (t *AccountingSuite) Test_InvalidTime() {

	repository.NewBudget()

	repository.DefaultBudget.Data["202101"] = 3100

	startDate, _ := time.Parse("2006-01-02", "2021-01-02")
	endDate, _ := time.Parse("2006-01-02", "2021-01-01")

	t.Equal(float64(0), t.Accounting.TotalAmount(startDate, endDate))

}

func (t *AccountingSuite) Test_ZeroAmount() {

	repository.NewBudget()

	repository.DefaultBudget.Data["202101"] = 0

	startDate, _ := time.Parse("2006-01-02", "2021-01-02")
	endDate, _ := time.Parse("2006-01-02", "2021-01-01")

	t.Equal(float64(0), t.Accounting.TotalAmount(startDate, endDate))
}

func (t *AccountingSuite) Test_DataNotFound() {

	repository.NewBudget()

	repository.DefaultBudget.Data["202102"] = 0

	startDate, _ := time.Parse("2006-01-02", "2021-01-02")
	endDate, _ := time.Parse("2006-01-02", "2021-01-01")

	t.Equal(float64(0), t.Accounting.TotalAmount(startDate, endDate))
}

func (t *AccountingSuite) Test_CrossMonth() {

	repository.NewBudget()

	repository.DefaultBudget.Data["202101"] = 3100
	repository.DefaultBudget.Data["202102"] = 280

	startDate, _ := time.Parse("2006-01-02", "2021-01-30")
	endDate, _ := time.Parse("2006-01-02", "2021-02-02")

	t.Equal(float64(220), t.Accounting.TotalAmount(startDate, endDate))
}
