package bonds

// BondParams structure defines a bond's parameters for price calculation.
// Notice that accrued interest is not (currently) supported -
// the bond price is calculated for the specified number of whole coupon periods.
// If the bond pays more than one coupon per year - incomplete years are supoorted.
type BondParams struct {
	FaceValue      float64 // bond's principal amount (face value)
	YieldPerYear   float64 // expected return rate (percentage, per year). E.g. 12%
	CouponRate     float64 // coupon rate (percentage, per year). E.g. 10%
	CouponsPerYear int     // number of coupon payment periods per year
	CouponsTotal   int     // total number of coupon payment periods, i.e. couponsPerYear * numberOfYears
}
