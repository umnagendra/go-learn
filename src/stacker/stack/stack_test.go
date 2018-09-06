// Unit tests for stacker/stack
package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func pushElements(testStack *Stack, count int) {
	for i := 0; i < count; i++ {
		testStack.Push(i * i)
	}
}

func TestLen(t *testing.T) {
	testStack := Stack{}
	// length should be 0 to begin with
	assert.Equal(t, 0, testStack.Len(), "Stack length incorrect")
	pushElements(&testStack, 100)
	assert.Equal(t, 100, testStack.Len(), "Stack length incorrect")
}

func TestCap(t *testing.T) {
	testStack := Stack{}
	// capacity should be 0 to begin with
	assert.Equal(t, 0, testStack.Cap(), "Stack capacity incorrect")
	pushElements(&testStack, 100)
	// capacity should be at least 100
	assert.True(t, testStack.Cap() >= 100, "Stack capacity incorrect")
}

func TestIsEmpty(t *testing.T) {
	testStack := Stack{}
	assert.True(t, testStack.IsEmpty(), "Stack supposed to be empty")
	pushElements(&testStack, 100)
	assert.False(t, testStack.IsEmpty(), "Stack supposed to be NOT empty")
}

func TestTop(t *testing.T) {
	testStack := Stack{}
	item, err := testStack.Top()
	assert.Nil(t, item, "Expected error since stack is empty")
	assert.NotNil(t, err, "Expected error since stack is empty")

	pushElements(&testStack, 100)
	item, err = testStack.Top()

	assert.Nil(t, err, "Expected value since stack is not empty")
	assert.NotNil(t, item, "Expected value since stack is not empty")
}

func TestPop(t *testing.T) {
	testStack := Stack{}
	pushElements(&testStack, 100)

	for i := 0; i < 100; i++ {
		item, err := testStack.Pop()
		assert.Nil(t, err, "Expected value since stack is not empty")
		assert.NotNil(t, item, "Expected value since stack is not empty")
	}

	// pop one last time on an empty stack, should fail
	item, err := testStack.Pop()
	assert.Nil(t, item, "Expected error since stack is empty")
	assert.NotNil(t, err, "Expected error since stack is empty")
}
