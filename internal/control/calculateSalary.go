package control

const personalDeductionRate = 11000000.0
const dependentDeductionRate = 4400000.0

const bhxhRate = 0.08
const bhytRate = 0.015
const bhtnRate = 0.01

func CalculatePersonalIncomeTax(taxable float64) float64 {
	tax := 0.0
	steps := []struct {
		limit float64
		rate  float64
	}{
		{5000000, 0.05},
		{5000000, 0.10},
		{8000000, 0.15},
		{14000000, 0.20},
		{20000000, 0.25},
		{28000000, 0.30},
		{1e18, 0.35},
	}

	for _, step := range steps {
		if taxable <= 0 {
			break
		}
		apply := step.limit
		if taxable < step.limit {
			apply = taxable
		}
		tax += apply * step.rate
		taxable -= apply
	}
	return tax
}

func CalculateNetSalary(gross float64, dependents int) float64 {
	bhxh := gross * bhxhRate
	bhyt := gross * bhytRate
	bhtn := gross * bhtnRate
	insurance := bhxh + bhyt + bhtn

	beforeTax := gross - insurance

	personalDeduction := personalDeductionRate
	dependentDeduction := float64(dependents) * dependentDeductionRate

	taxable := beforeTax - personalDeduction - dependentDeduction
	if taxable < 0 {
		taxable = 0
	}

	personalIncomeTax := CalculatePersonalIncomeTax(taxable)

	net := beforeTax - personalIncomeTax
	return net
}
