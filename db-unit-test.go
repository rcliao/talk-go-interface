import "strconv"

func Run(
	dao DAO,
	statements, setupStatements, teardownStatements, solutions []Statement,
	selectedQuestions []string,
) ([]TestResult, error) {
	var testResult = []TestResult{}

	i := 0
	results, errs, err := dao.ExecuteStatements(setupStatements, teardownStatements, statements)
	solutionResults, _, err := dao.ExecuteStatements(setupStatements, teardownStatements, solutions)
	testcases := ConvertTablesToTestCases(solutionResults)

	for _, expected := range testcases {
		if i >= len(results) {
			testResult = append(
				testResult,
				TestResult{
					Expected: expected,
					Actual:   Result{},
					Pass:     false,
				},
			)
			i++
			continue
		}
	}

	return testResult, nil
}

// ConvertTablesToTestCases is a util method to add index to table
func ConvertTablesToTestCases(tables []Result) []TestCase {
	testcases := []TestCase{}
	for i, t := range tables {
		testcases = append(testcases, TestCase{
			Index:    strconv.Itoa(i + 1),
			Content:  t.Content,
			Question: "",
		})
	}

	return testcases
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
