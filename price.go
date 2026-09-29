package bonds

import "math"

// yieldPerPeriod calculates a yield rate per period so that, after a year,
// the resulting compound interest is equal to the given yield per year.
// yieldPerYear should be specified as a percentage, e.g. 12% per year.
// The result is returned as 1 + percentage / 100, e.g. 1.12.
// Notice that due to compound interest
// the resulting value is less than just yieldPerYear/periodsPerYear.
func yieldPerPeriod(yieldPerYear float64, periodsPerYear int) float64 {
	return math.Pow(1+yieldPerYear/100, 1/float64(periodsPerYear))
}

// BondPrice calculates a price of a bond with the given parameters
// by repeatedly discounting coupons and face value with the yield rate per period.
func BondPrice(input BondParams) float64 {
	yieldPerPeriod := yieldPerPeriod(input.YieldPerYear, input.CouponsPerYear)
	couponRatePerPeriod := input.CouponRate / float64(input.CouponsPerYear)

	coupon := input.FaceValue * couponRatePerPeriod / 100
	value := input.FaceValue

	var sum float64
	for i := 0; i < input.CouponsTotal; i++ {
		coupon /= yieldPerPeriod
		value /= yieldPerPeriod
		sum += coupon
	}

	return sum + value
}
