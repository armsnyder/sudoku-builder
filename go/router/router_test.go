package router_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"

	"sudoku-builder/router"
)

func TestSolveAPI(t *testing.T) {
	tests := []struct {
		name                           string
		body                           interface{}
		expectedHTTPStatus             int
		expectedErrorCode              string
		expectedErrorMessageContains   string
		expectedErrorMessageNoContains string
		expectedSolutions              []interface{}
	}{
		{
			name:                           "Solve a malformed board",
			body:                           gin.H{"board": 1},
			expectedHTTPStatus:             400,
			expectedErrorCode:              "MALFORMED_REQUEST_BODY",
			expectedErrorMessageContains:   "field 'board'",
			expectedErrorMessageNoContains: "go",
		},

		{
			name:               "Solve an empty board",
			body:               gin.H{"board": []int{}},
			expectedHTTPStatus: 400,
			expectedErrorCode:  "TOO_MANY_SOLUTIONS",
		},

		{
			name: "Solve a board with a column contradiction",
			body: gin.H{"board": [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 3},
				{0, 0, 5, 0, 0, 0, 0, 8, 0},
				{6, 0, 9, 0, 0, 8, 0, 0, 5},
				{1, 0, 0, 0, 5, 0, 8, 0, 6},
				{4, 0, 6, 0, 0, 0, 2, 0, 7},
				{7, 9, 0, 0, 6, 3, 0, 0, 1},
				{0, 1, 2, 6, 0, 0, 9, 0, 8},
				{0, 0, 7, 9, 5, 2, 0, 0, 0},
				{9, 6, 4, 7, 8, 0, 3, 1, 0},
			}},
			expectedHTTPStatus: 400,
			expectedErrorCode:  "NO_SOLUTION",
		},

		{
			name: "Solve a board with a row contradiction",
			body: gin.H{"board": [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 3},
				{0, 0, 5, 0, 0, 0, 0, 8, 0},
				{6, 0, 9, 0, 0, 8, 0, 0, 5},
				{1, 0, 0, 0, 5, 0, 8, 0, 6},
				{4, 0, 6, 0, 0, 0, 2, 0, 7},
				{7, 9, 0, 0, 6, 9, 0, 0, 1},
				{0, 1, 2, 6, 0, 0, 9, 0, 8},
				{0, 0, 7, 9, 1, 2, 0, 0, 0},
				{9, 6, 4, 7, 8, 0, 3, 1, 0},
			}},
			expectedHTTPStatus: 400,
			expectedErrorCode:  "NO_SOLUTION",
		},

		{
			name: "Solve a board with a section contradiction",
			body: gin.H{"board": [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 3},
				{0, 0, 5, 0, 0, 0, 0, 8, 0},
				{6, 0, 9, 0, 0, 8, 0, 0, 5},
				{1, 0, 0, 0, 5, 0, 8, 0, 6},
				{4, 0, 6, 0, 0, 5, 2, 0, 7},
				{7, 9, 0, 0, 6, 3, 0, 0, 1},
				{0, 1, 2, 6, 0, 0, 9, 0, 8},
				{0, 0, 7, 9, 1, 2, 0, 0, 0},
				{9, 6, 4, 7, 8, 0, 3, 1, 0},
			}},
			expectedHTTPStatus: 400,
			expectedErrorCode:  "NO_SOLUTION",
		},

		{
			name: "Solve a board with an illegal digit",
			body: gin.H{"board": [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 3},
				{0, 0, 5, 0, 0, 0, 0, 8, 0},
				{6, 0, 9, 0, 0, 8, 0, 0, 5},
				{1, 0, 0, 0, 5, 0, 8, 0, 6},
				{4, 0, 10, 0, 0, 0, 2, 0, 7},
				{7, 9, 0, 0, 6, 3, 0, 0, 1},
				{0, 1, 2, 6, 0, 0, 9, 0, 8},
				{0, 0, 7, 9, 1, 2, 0, 0, 0},
				{9, 6, 4, 7, 8, 0, 3, 1, 0},
			}},
			expectedHTTPStatus: 400,
			expectedErrorCode:  "NO_SOLUTION",
		},

		{
			name: "Solve a real board with only one solution",
			body: gin.H{"board": [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 3},
				{0, 0, 5, 0, 0, 0, 0, 8, 0},
				{6, 0, 9, 0, 0, 8, 0, 0, 5},
				{1, 0, 0, 0, 5, 0, 8, 0, 6},
				{4, 0, 6, 0, 0, 0, 2, 0, 7},
				{7, 9, 0, 0, 6, 3, 0, 0, 1},
				{0, 1, 2, 6, 0, 0, 9, 0, 8},
				{0, 0, 7, 9, 1, 2, 0, 0, 0},
				{9, 6, 4, 7, 8, 0, 3, 1, 0},
			}},
			expectedHTTPStatus: 200,
			expectedSolutions: []interface{}{
				[][]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{4, 5, 6, 7, 8, 9, 1, 2, 3},
					{7, 8, 9, 1, 2, 3, 4, 5, 6},
					{2, 1, 4, 3, 6, 5, 8, 9, 7},
					{3, 6, 5, 8, 9, 7, 2, 1, 4},
					{8, 9, 7, 2, 1, 4, 3, 6, 5},
					{5, 3, 1, 6, 4, 2, 9, 7, 8},
					{6, 4, 2, 9, 7, 8, 5, 3, 1},
					{9, 7, 8, 5, 3, 1, 6, 4, 2},
				},
			},
		},
	}

	gin.SetMode(gin.TestMode)

	srv := httptest.NewServer(router.New(true, ""))
	defer srv.Close()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp := httpexpect.New(t, srv.URL).
				POST("/solve").
				WithJSON(test.body).
				Expect().
				ContentType("application/json", "utf-8").
				Status(test.expectedHTTPStatus).
				JSON().
				Object()
			if test.expectedErrorCode != "" {
				resp.Value("code").String().Equal(test.expectedErrorCode)
			}
			if test.expectedErrorMessageContains != "" {
				resp.Value("message").String().ContainsFold(test.expectedErrorMessageContains)
			}
			if test.expectedErrorMessageNoContains != "" {
				resp.Value("message").String().NotContainsFold(test.expectedErrorMessageNoContains)
			}
			if test.expectedSolutions != nil {
				resp.Value("solutions").Array().ContainsOnly(test.expectedSolutions...)
			}
		})
	}
}
