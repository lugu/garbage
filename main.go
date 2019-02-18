package main

import (
	"bufio"
	"fmt"
	"github.com/antzucaro/matchr"
	"os"
)

const (
	GreyList_Threshold = 3
	OSA_Threshold      = 20
)

type Filter map[string]int

func (f Filter) Match(line string) (bool, string) {
	for parent, _ := range f {
		distance := matchr.OSA(line, parent)
		if distance <= OSA_Threshold {
			return true, parent
		}
	}
	return false, line
}

func (f Filter) Contains(line string) bool {
	_, ok := f[line]
	return ok
}

func (f Filter) Increment(line string) {
	f[line]++
}

func (f Filter) Remove(line string) {
	delete(f, line)
}

func (f Filter) Count(line string) int {
	return f[line]
}

func main() {
	// initialize blacklist empty (growing buffer)
	// initialize greylist empty (sorted ring buffer)
	// initialize history empty (ring buffer)
	//
	// while read line
	// 	if line match black list:
	// 		skip
	// 	else if line match grey list:
	// 		if threshold meet:
	//			move to black list and skip
	// 		else :
	//			increment counter and print
	//  else if line match history:
	// 		move to grey list and print
	//  else add to history and print
	//
	// testing: create various test case depending on the
	// constants
	//
	// matching algo: share at least 80% characters
	//	option 1: hamming distance
	//	option 2: bag of words
	//	option 3: sum of same substring lenght
	//
	// data:
	//	blacklist:
	//		- contains(strings)
	//		- insert(strings)
	//	sorted ring buffer
	//	list
	// data structures:
	//	ring buffer:
	//	sorted ring buffer
	//	list
	blackList := Filter(make(map[string]int))
	grayList := Filter(make(map[string]int))
	history := Filter(make(map[string]int))

	input := bufio.NewReader(os.Stdin)
	for {
		line, err := input.ReadString('\n')
		if err != nil {
			break
		}

		matched, _ := blackList.Match(line)
		if matched {
			continue
		}

		matched, parent := grayList.Match(line)

		if matched {
			grayList.Increment(parent)
			if grayList.Count(parent) >= GreyList_Threshold {
				blackList.Increment(parent)
				grayList.Remove(parent)
				continue
			}
		} else {

			matched, parent = history.Match(line)
			if matched {
				grayList.Increment(parent)
				history.Remove(parent)
			} else {
				history.Increment(parent)
			}
			fmt.Print(line)
		}
	}
}
