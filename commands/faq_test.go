package commands

import (
	"testing"
)

var testMenu = [5]faqButton{
	{
		display: "1",
		command: "1",
	},
	{
		display: "2",
		command: "2",
	},
	{
		display: "3",
		command: "3",
	},
	{
		display: "4",
		command: "4",
	},
	{
		display: "5",
		command: "5",
	},
}

type testSplitData struct {
	split      int
	expectRows []int
}

var testSplits = [5]testSplitData{
	{
		split: 1,
		expectRows: []int{
			1,
			1,
			1,
			1,
			1,
		},
	},
	{
		split: 5,
		expectRows: []int{
			5,
		},
	},
	{
		split: 2,
		expectRows: []int{
			2,
			2,
			1,
		},
	},
}

func TestCreateFaqMenu(t *testing.T) {
	_, errBelowZero := createFAQMenu(-1, []faqButton{})
	_, errZero := createFAQMenu(0, []faqButton{})
	if errBelowZero == nil || errZero == nil {
		t.FailNow()
	}
	if errBelowZero.Error() != "split is below a zero" {
		t.Fatal(errBelowZero)
	}

	if errZero.Error() != "split is a zero" {
		t.Fatal(errZero)
	}

	for _, splitData := range testSplits {
		inlineKeyboard, err := createFAQMenu(splitData.split, testMenu[:])
		if err != nil && len(splitData.expectRows) != 0 && splitData.split != -1 {
			t.Log("Unexpected fail", inlineKeyboard.InlineKeyboard, err, splitData)
			t.FailNow()
		}
		if len(inlineKeyboard.InlineKeyboard) != len(splitData.expectRows) {
			t.Log("Failed on rows count", inlineKeyboard.InlineKeyboard, err, splitData)
			t.FailNow()
		}

		for i, data := range inlineKeyboard.InlineKeyboard {
			expect := splitData.expectRows[i]
			if len(data) != expect {
				t.Log("Failed on row count", inlineKeyboard.InlineKeyboard, err, splitData)
				t.FailNow()
			}
		}

	}
}
