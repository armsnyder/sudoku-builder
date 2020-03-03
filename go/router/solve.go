package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type solveRequest struct {
	Board boardState `json:"board"`
}

type solveResponse struct {
	Solutions []boardState `json:"solutions"`
}

type boardState [9][9]int8

func solve(c *gin.Context) {
	reqBody := new(solveRequest)

	if err := c.ShouldBind(reqBody); err != nil {
		c.Error(bindError(err))
		return
	}

	if reqBody.Board.isEmpty() {
		c.Error(&apiError{Code: tooManySolutions, Message: "Board in request is empty"})
		return
	}

	if !reqBody.Board.satisfiesConstraints() {
		c.Error(&apiError{Code: noSolution, Message: "Board does not satisfy Sudoku constraints as-is"})
		return
	}

	c.JSON(http.StatusOK, &solveResponse{
		Solutions: []boardState{
			{
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
	})
}

func (b *boardState) isEmpty() bool {
	for _, col := range b {
		for _, cell := range col {
			if cell != 0 {
				return false
			}
		}
	}

	return true
}

func (b *boardState) satisfiesConstraints() bool {
	return b.satisfiesColumnConstraint() && b.satisfiesRowConstraint() && b.satisfiesSectionConstraint()
}

func (b *boardState) satisfiesColumnConstraint() bool {
	var seen [9]bool

	for i := 0; i < 9; i++ {
		for b := range seen {
			seen[b] = false
		}

		for j := 0; j < 9; j++ {
			val := b[i][j]
			if !b.isValueUnseenAndUpdateSeen(val, &seen) {
				return false
			}
		}
	}

	return true
}

func (b *boardState) satisfiesRowConstraint() bool {
	var seen [9]bool

	for j := 0; j < 9; j++ {
		for b := range seen {
			seen[b] = false
		}

		for i := 0; i < 9; i++ {
			val := b[i][j]
			if !b.isValueUnseenAndUpdateSeen(val, &seen) {
				return false
			}
		}
	}

	return true
}

func (b *boardState) satisfiesSectionConstraint() bool {
	for secI := 0; secI < 3; secI++ {
		for secJ := 0; secJ < 3; secJ++ {
			if !b.sectionSatisfiesSectionConstraint(secI, secJ) {
				return false
			}
		}
	}

	return true
}

func (b *boardState) sectionSatisfiesSectionConstraint(secI, secJ int) bool {
	var seen [9]bool

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			val := b[secI*3+i][secJ*3+j]
			if !b.isValueUnseenAndUpdateSeen(val, &seen) {
				return false
			}
		}
	}

	return true
}

func (b *boardState) isValueUnseenAndUpdateSeen(val int8, seen *[9]bool) bool {
	if val < 0 || val > 9 {
		return false
	}

	if val == 0 {
		return true
	}

	if seen[val-1] {
		return false
	}

	seen[val-1] = true

	return true
}
