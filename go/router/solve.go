package router

import (
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

	c.JSON(200, &solveResponse{
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
	var seen [9]bool

	// Check column constraint
	for i := 0; i < 9; i++ {
		for b := range seen {
			seen[b] = false
		}
		for j := 0; j < 9; j++ {
			val := b[i][j]
			if val < 0 || val > 9 {
				return false
			}
			if val == 0 {
				continue
			}
			if seen[val-1] {
				return false
			}
			seen[val-1] = true
		}
	}

	// Check row constraint
	for j := 0; j < 9; j++ {
		for b := range seen {
			seen[b] = false
		}
		for i := 0; i < 9; i++ {
			val := b[i][j]
			if val < 0 || val > 9 {
				return false
			}
			if val == 0 {
				continue
			}
			if seen[val-1] {
				return false
			}
			seen[val-1] = true
		}
	}

	// Check section constraint
	for iOffset := 0; iOffset < 9; iOffset += 3 {
		for jOffset := 0; jOffset < 9; jOffset += 3 {
			for b := range seen {
				seen[b] = false
			}
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					val := b[iOffset+i][jOffset+j]
					if val < 0 || val > 9 {
						return false
					}
					if val == 0 {
						continue
					}
					if seen[val-1] {
						return false
					}
					seen[val-1] = true
				}
			}
		}
	}

	return true
}
