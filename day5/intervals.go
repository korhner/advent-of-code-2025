package main

import "sort"
import "strconv"

type Interval struct {
	Start int
	End   int
}

func (iv Interval) Overlaps(other Interval) bool {
	return (other.Start >= iv.Start && other.Start <= iv.End) || (iv.Start >= other.Start && iv.Start <= other.End)
}

func (iv Interval) Contains(point int) bool {
	return point >= iv.Start && point <= iv.End
}

func (iv Interval) Range() int {
	return iv.End - iv.Start + 1
}

type IntervalSet []Interval

func NewIntervalSet(intervals ...Interval) IntervalSet {
	set := IntervalSet(intervals)
	sort.Slice(set, func(i, j int) bool {
		return set[i].Start < set[j].Start
	})

	merged := make(IntervalSet, 0, len(set))
	for _, iv := range set {
		if len(merged) == 0 {
			merged = append(merged, iv)
			continue
		}

		last := &merged[len(merged)-1]
		if last.Overlaps(iv) {
			last.Start = min(last.Start, iv.Start)
			last.End = max(last.End, iv.End)
		} else {
			merged = append(merged, iv)
		}
	}
	return merged
}

func (set IntervalSet) Contains(point int) bool {
	for _, iv := range set {
		if point < iv.Start {
			break
		}

		if iv.Contains(point) {
			return true
		}
	}
	return false
}

func (set IntervalSet) Range() int {
	total := 0
	for _, iv := range set {
		total += iv.Range()
	}
	return total
}

func (set IntervalSet) toString() string {
	result := ""
	for _, iv := range set {
		result += "[" + strconv.Itoa(iv.Start) + "-" + strconv.Itoa(iv.End) + "] "
	}
	return result
}
