package bonds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYieldPerPeriod(t *testing.T) {
	delta := 0.0001
	// no change for one period per year
	assert.InDelta(t, 1.12, yieldPerPeriod(12.0, 1), delta)
	// less than 6% due to compound interest
	assert.InDelta(t, 1.0583, yieldPerPeriod(12.0, 2), delta)
	// less than 3% due to compound interest
	assert.InDelta(t, 1.0287, yieldPerPeriod(12.0, 4), delta)
	// less than 1% due to compound interest
	assert.InDelta(t, 1.0095, yieldPerPeriod(12.0, 12), delta)
}

func TestBondPrice(t *testing.T) {
	delta := 0.001

	t.Run("on par", func(t *testing.T) {
		params := BondParams{
			FaceValue:      1000,
			YieldPerYear:   12,
			CouponRate:     12,
			CouponsPerYear: 1,
			CouponsTotal:   1,
		}
		assert.InDelta(t, params.FaceValue, BondPrice(params), delta, "1 year")
		params.CouponsTotal = 2 // 2 years
		assert.InDelta(t, params.FaceValue, BondPrice(params), delta, "2 years")
	})

	t.Run("more coupons per year", func(t *testing.T) {
		params := BondParams{
			FaceValue:      1000,
			YieldPerYear:   12,
			CouponRate:     12,
			CouponsPerYear: 2,
			CouponsTotal:   2,
		}
		assert.InDelta(t, 1003.123, BondPrice(params), delta, "2 coupons per year")
		params.CouponsPerYear = 4
		params.CouponsTotal = 4
		assert.InDelta(t, 1004.708, BondPrice(params), delta, "4 coupons per year")
	})

	t.Run("higher expected return", func(t *testing.T) {
		params := BondParams{
			FaceValue:      1000,
			YieldPerYear:   12,
			CouponRate:     10,
			CouponsPerYear: 1,
			CouponsTotal:   1,
		}
		assert.InDelta(t, 982.143, BondPrice(params), delta, "1 year")
		params.CouponsTotal = 2
		assert.InDelta(t, 966.199, BondPrice(params), delta, "2 years")
	})

	t.Run("lower expected return", func(t *testing.T) {
		params := BondParams{
			FaceValue:      1000,
			YieldPerYear:   10,
			CouponRate:     12,
			CouponsPerYear: 1,
			CouponsTotal:   1,
		}
		assert.InDelta(t, 1018.182, BondPrice(params), delta, "1 year")
		params.CouponsTotal = 2
		assert.InDelta(t, 1034.711, BondPrice(params), delta, "2 years")
	})

	t.Run("ОФЗ-26249", func(t *testing.T) {
		params := BondParams{
			FaceValue:      1000,
			YieldPerYear:   16,
			CouponRate:     11,
			CouponsPerYear: 2,
			CouponsTotal:   12, // 6 years (16.06.2032)
		}
		assert.InDelta(t, 831.375, BondPrice(params), delta, "1 year")
	})

	t.Run("ОФЗ-26251", func(t *testing.T) {
		params := BondParams{
			FaceValue:      1000,
			YieldPerYear:   16.8,
			CouponRate:     12,
			CouponsPerYear: 2,
			CouponsTotal:   22, // 11 years (10.06.2037)
		}
		assert.InDelta(t, 789.665, BondPrice(params), delta, "1 year")
	})

	t.Run("ФосА1П1 USD", func(t *testing.T) {
		params := BondParams{
			FaceValue:      100,
			YieldPerYear:   8.6,
			CouponRate:     6.25,
			CouponsPerYear: 4,
			CouponsTotal:   12, // 3 years (31.05.2029)
		}
		assert.InDelta(t, 94.514, BondPrice(params), delta, "1 year")
	})
}
