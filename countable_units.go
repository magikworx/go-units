package units

var (
	Count = UnitOptionQuantity("count")

	Whole     = NewUnit("whole", "whl", Count)
	One       = NewUnit("one", "1", Count)
	Half      = NewUnit("half", "1/2", Count)
	Third     = NewUnit("third", "1/3", Count)
	Quarter   = NewUnit("quarter", "1/4", Count)
	Fifth     = NewUnit("fifth", "1/5", Count)
	Sixth     = NewUnit("sixth", "1/6", Count)
	Seventh   = NewUnit("seventh", "1/7", Count)
	Eighth    = NewUnit("eighth", "1/8", Count)
	Tenth     = NewUnit("tenth", "1/10", Count)
	Sixteenth = NewUnit("sixteenth", "1/16", Count)
)

func init() {
	NewRatioConversion(One, Whole, 1.0)
	NewRatioConversion(Half, Whole, 0.5)
	NewRatioConversion(Third, Whole, 1.0/3.0)
	NewRatioConversion(Quarter, Whole, 0.25)
	NewRatioConversion(Fifth, Whole, 0.2)
	NewRatioConversion(Sixth, Whole, 1.0/6.0)
	NewRatioConversion(Seventh, Whole, 1.0/7.0)
	NewRatioConversion(Eighth, Whole, 0.125)
	NewRatioConversion(Tenth, Whole, 0.1)
	NewRatioConversion(Sixteenth, Whole, 1.0/16.0)
}
