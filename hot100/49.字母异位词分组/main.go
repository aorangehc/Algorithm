package main

import (
	"fmt"
	"reflect"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}

	for _, str := range strs {
		s := []byte(str)
		slices.Sort(s)
		mp[string(s)] = append(mp[string(s)], str)
	}

	ans := make([][]string, 0, len(mp))

	for i := range mp {
		ans = append(ans, mp[i])
	}

	return ans
}

func main() {
	testCases := []struct {
		strs []string
		ans  [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			[]string{""},
			[][]string{
				{""},
			},
		},
		{
			[]string{"a"},
			[][]string{
				{"a"},
			},
		},
	}

	for _, t := range testCases {
		fmt.Println(reflect.DeepEqual(groupAnagrams(t.strs), t.ans))
	}
}
