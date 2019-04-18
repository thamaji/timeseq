package timeseq

import (
	"time"
)

type TimeSeq interface {
	Next() bool
	Value() time.Time
}

func From(start time.Time, step StepFunc) TimeSeq {
	return &InfiniteSeq{value: start, step: step}
}

type InfiniteSeq struct {
	value time.Time
	next  *time.Time
	step  StepFunc
}

func (seq *InfiniteSeq) Next() bool {
	if seq.next != nil {
		seq.value = *seq.next
	}
	next := seq.step(seq.value)
	seq.next = &next
	return true
}

func (seq *InfiniteSeq) Value() time.Time {
	return seq.value
}

func Range(start, end time.Time, step StepFunc) TimeSeq {
	return &RangeSeq{
		value: start,
		end:   end,
		step:  step,
	}
}

type RangeSeq struct {
	value time.Time
	next  *time.Time
	end   time.Time
	step  StepFunc
}

func (seq *RangeSeq) Next() bool {
	if seq.next != nil {
		seq.value = *seq.next
	}
	next := seq.step(seq.value)
	seq.next = &next
	return seq.value.Before(seq.end) // !(end < value)
}

func (seq *RangeSeq) Value() time.Time {
	return seq.value
}

type StepFunc func(time.Time) time.Time

func StepDuration(d time.Duration) StepFunc {
	return func(t time.Time) time.Time {
		return t.Add(d)
	}
}

func StepDays(days int) StepFunc {
	return func(t time.Time) time.Time {
		return t.AddDate(0, 0, days)
	}
}

func StepMonths(months int) StepFunc {
	return func(t time.Time) time.Time {
		return t.AddDate(0, months, 0)
	}
}

func StepYears(years int) StepFunc {
	return func(t time.Time) time.Time {
		return t.AddDate(years, 0, 0)
	}
}

func StepDate(years int, months int, days int) StepFunc {
	return func(t time.Time) time.Time {
		return t.AddDate(years, months, days)
	}
}
