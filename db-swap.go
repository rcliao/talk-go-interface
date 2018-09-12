import (
	"net/http"

	"github.com/rcliao/sql-unit-test/db"
	"github.com/rcliao/sql-unit-test/parser"
)

func RunTest(factory *db.Factory) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		submission := r.FormValue("statements")
		subject := r.FormValue("subject")
		dao := factory.CreateDAO(getSubjectType(subject))
		selectedQuestions := r.Form["question_numbers"]
		statements := parser.ParseSQL(submission, "#")
		testResult, err := tester.Run(
			dao,
			statements,
			setupStatements,
			teardownStatements,
			solutions,
			selectedQuestions,
		)
	})
}

