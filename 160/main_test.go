package main

import "testing"

type input struct {
	grade       int
	presentList []string
}
type output struct {
	present string
}
type expect struct {
	i input
	o output
	e bool
}

func TestChoosePresentByGrade(t *testing.T) {
	var expects = []expect{
		{input{1, []string{"a", "b", "c"}}, output{"a"}, true},
		{input{2, []string{"a", "b", "c"}}, output{"b"}, true},
		{input{3, []string{"a", "b", "c"}}, output{"c"}, true},
		{input{0, []string{"a", "b", "c"}}, output{"c"}, false},
		{input{4, []string{"a", "b", "c"}}, output{"c"}, false},
		{input{3, []string{"cartoy", "plamodel", "gameconsole"}}, output{"gameconsole"}, true},
		{input{2, []string{"gloves", "muffler", "sweater"}}, output{"muffler"}, true},
		{input{1, []string{"pizza", "steak", "sushi"}}, output{"pizza"}, true},
		{input{1, []string{"cartoy", "plamodel", "gameconsole"}}, output{"gameconsole"}, false},
		{input{2, []string{"gloves", "muffler", "sweater"}}, output{"sushi"}, false},
		{input{3, []string{"pizza", "steak", "sushi"}}, output{"pizza"}, false},
	}

	for i, v := range expects {
		if v.e != (v.o.present == choosePresentByGrade(v.i.grade, v.i.presentList)) {
			t.Errorf("test case [%d] failed. v = %v", i, v)
		}
	}
}
