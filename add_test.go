package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestAdd_Modifies_ConstantExpressionByAddingValue(t *testing.T) {
	// Arrange
	eq := &Equation{
		Left: &ConstExpression{Value: 1},
		Right: &ConstExpression{Value: 2},
	}

	// Act
	Add(eq, func() int { return 1})

	// Assert
	assert.Equal(t, 2, eq.Left.(*ConstExpression).Value)
	assert.Equal(t, 3, eq.Right.(*ConstExpression).Value)
}

func TestAdd_Modifies_AddConstantExpressionByAddingValue(t *testing.T) {
	// Arrange
	eq := &Equation{
		Left: &AddConstantExpression{Value: 1},
		Right: &AddConstantExpression{Value: 2},
	}

	// Act
	Add(eq, func() int { return 1})

	// Assert
	assert.Equal(t, 2, eq.Left.(*AddConstantExpression).Value)
	assert.Equal(t, 3, eq.Right.(*AddConstantExpression).Value)
}

func TestAdd_Wraps_OtherExpressionsByAddConstantExpression(t *testing.T) {
	// Arrange
	eq := &Equation{
		Left: &VariableExpression{},
		Right: &VariableExpression{},
	}

	// Act
	Add(eq, func() int { return 1})

	// Assert
	assert.Equal(t, 1, eq.Left.(*AddConstantExpression).Value)
	assert.IsType(t, &VariableExpression{}, eq.Left.(*AddConstantExpression).Parent)

	assert.Equal(t, 1, eq.Right.(*AddConstantExpression).Value)
	assert.IsType(t, &VariableExpression{}, eq.Right.(*AddConstantExpression).Parent)
}

func TestAddConstantExpression_String(t *testing.T) {
	expr := AddConstantExpression{ Value: 1, Parent: &VariableExpression{}}

	assert.Equal(t, "x + 1", expr.String())
}
