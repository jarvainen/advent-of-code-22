package calendar

import (
	__day_one "github.com/jarvainen/advent-of-code-22/calendar/1_day_one"
	__day_two "github.com/jarvainen/advent-of-code-22/calendar/2_day_two"
	__day_three "github.com/jarvainen/advent-of-code-22/calendar/3_day_three"
	__day_four "github.com/jarvainen/advent-of-code-22/calendar/4_day_four"
	__day_five "github.com/jarvainen/advent-of-code-22/calendar/5_day_five"
	__day_six "github.com/jarvainen/advent-of-code-22/calendar/6_day_six"
	__day_seven "github.com/jarvainen/advent-of-code-22/calendar/7_day_seven"
	__day_eight "github.com/jarvainen/advent-of-code-22/calendar/8_day_eight"
)

type Solution func()

var Solutions = [24]Solution{
	__day_one.Solution,
	__day_two.Solution,
	__day_three.Solution,
	__day_four.Solution,
	__day_five.Solution,
	__day_six.Solution,
	__day_seven.Solution,
	__day_eight.Solution,
}
