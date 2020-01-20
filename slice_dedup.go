// Deduplicate - De-duplication of elements in a []string{}
// Returns a separate []string with only unique elements *IN NATURAL SORTED ORDER*
// The original input slice is NOT modified (no side-effects)
// Inspired by https://github.com/golang/go/wiki/SliceTricks
func Deduplicate(original []string) []string {
	var inputSlice = []string{}

	// handle empty slices
	if len(original) == 0 {
		return inputSlice
	}

	inputSlice = append(inputSlice, original...)
	sort.Strings(inputSlice)
	j := 0
	for i := 1; i < len(inputSlice); i++ {
		if inputSlice[j] == inputSlice[i] {
			continue
		}
		j++
		// preserve the original data
		// inputSlice[i], inputSlice[j] = inputSlice[j], inputSlice[i]
		// only set what is required
		inputSlice[j] = inputSlice[i]
	}
	return inputSlice[:j+1]
}


////////
// TEST
///////
func TestDeduplicate(t *testing.T) {
	inputs := [][]string{
		{"zebra", "aardvark", "Antarctica", "member", "zebra", "America", "member"},
		{"12Ccbu12", "C1sco123", "cisco@123", "C1sco123", "c1sco123=", "cisco@123"},
		{"testword", "testword", "TESTWORD"},
		{"abcd", "abcd"},
		{},
	}

	ExpectedOutputs := [][]string{
		{"America", "Antarctica", "aardvark", "member", "zebra"},
		{"12Ccbu12", "C1sco123", "c1sco123=", "cisco@123"},
		{"TESTWORD", "testword"},
		{"abcd"},
		{},
	}

	for count, inputSlice := range inputs {
		outputSlice := Deduplicate(inputSlice)
		if !reflect.DeepEqual(ExpectedOutputs[count], outputSlice) {
			t.Fatalf("For input slice %v, expected %v, but got %v", inputSlice, ExpectedOutputs[count], outputSlice)
		}
	}
}
