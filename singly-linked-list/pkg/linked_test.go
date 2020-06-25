package linked_test

import (
	linked "go-data-structures/singly-linked-list/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_InsertLast_AddsNodeAtTheEnd(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	list.InsertLast("4")
	list.InsertLast("3")
	list.InsertLast("2")
	list.InsertLast("1")

	expected := []linked.Data{
		"4",
		"3",
		"2",
		"1",
	}

	test.EqualValues(expected, dump(list))
}

func TestList_InsertFront_AddsNodeAtTheBeginning(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	list.InsertFront("4")
	list.InsertFront("3")
	list.InsertFront("2")
	list.InsertFront("1")

	expected := []linked.Data{
		"1",
		"2",
		"3",
		"4",
	}

	test.EqualValues(expected, dump(list))
}

func TestList_InsertBefore_InsertNilMark(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	_ = list.InsertLast("1")
	_ = list.InsertLast("2")
	_ = list.InsertBefore("3", nil)

	expected := []linked.Data{
		"1",
		"2",
	}

	test.EqualValues(expected, dump(list))
}

func TestList_InsertBefore_AddsNodeBeforeGivenNode(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	n1 := list.InsertLast("1")
	n2 := list.InsertLast("2")
	n3 := list.InsertLast("3")
	n4 := list.InsertLast("4")

	list.InsertBefore("x", n3)

	_ = n1
	_ = n2
	_ = n3
	_ = n4

	expected := []linked.Data{
		"1",
		"2",
		"x",
		"3",
		"4",
	}

	test.EqualValues(expected, dump(list))
}

func TestList_InsertBefore_UpdatesHead(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	n1 := list.InsertLast("1")
	x := list.InsertBefore("x", n1)

	expected := []linked.Data{
		"x",
		"1",
	}

	test.Equal(expected, dump(list))
	test.Equal(x, list.Head())
}

func TestList_InsertBefore_DoesNotAddNodeIfUnknownNodePassed(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	n1 := list.InsertLast("1")
	list.Delete(n1)

	list.InsertBefore("1", n1)

	expected := []linked.Data{}

	test.EqualValues(expected, dump(list))
}

func TestList_InsertAfter_InsertNilMark(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	_ = list.InsertLast("1")
	_ = list.InsertLast("2")
	_ = list.InsertAfter("3", nil)

	expected := []linked.Data{
		"1",
		"2",
	}

	test.EqualValues(expected, dump(list))
}

func TestList_InsertAfter_AddsNodeAfterGivenNode(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	n1 := list.InsertLast("1")
	n2 := list.InsertLast("2")
	n3 := list.InsertLast("3")
	n4 := list.InsertLast("4")

	list.InsertAfter("x", n2)

	_ = n1
	_ = n3
	_ = n4

	expected := []linked.Data{
		"1",
		"2",
		"x",
		"3",
		"4",
	}

	test.EqualValues(expected, dump(list))
	test.Equal(n4, list.Tail())
}

func TestList_InsertAfter_UpdatesTail(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	n1 := list.InsertLast("1")
	x := list.InsertAfter("x", n1)

	expected := []linked.Data{
		"1",
		"x",
	}

	test.EqualValues(expected, dump(list))
	test.Equal(x, list.Tail())
}

func TestList_InsertAfter_DoesNotAddNodeIfUnknownNodePassed(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	n1 := list.InsertFront("1")
	n2 := list.InsertLast("3")
	n3 := list.InsertAfter("2", n1)
	list.Delete(n2)
	list.Delete(n3)
	list.Delete(n1)

	list.InsertAfter("1", n1)

	expected := []linked.Data{}

	test.EqualValues(expected, dump(list))
}

func TestList_Delete_RemovesItemAndConnectsNodes(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	_ = list.InsertLast("1")
	n2 := list.InsertLast("2")
	_ = list.InsertLast("3")

	list.Delete(n2)

	expected := []linked.Data{
		"1",
		"3",
	}

	test.EqualValues(expected, dump(list))
}

func TestList_Delete_DoesNotPanicWithNil(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	_ = list.InsertLast("1")
	_ = list.InsertLast("2")
	_ = list.InsertLast("3")

	list.Delete(nil)

	expected := []linked.Data{
		"1",
		"2",
		"3",
	}

	test.EqualValues(expected, dump(list))
}

func TestList_Delete_EmptiesHeadAndTail(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	n1 := list.InsertLast("1")

	list.Delete(n1)

	expected := []linked.Data{}

	test.EqualValues(expected, dump(list))
	test.Nil(list.Head(), "list head must be nil")
	test.Nil(list.Tail(), "list tail must be nil")
}

func TestList_Size_ReturnsZeroForEmptyList(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	test.EqualValues(0, list.Size())
}

func TestList_Size_ReturnsCorrectNumber(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	_ = list.InsertLast("1")
	_ = list.InsertLast("2")
	_ = list.InsertLast("3")

	test.EqualValues(3, list.Size())
}

func TestList_Size_ReturnsCorrectNumberAfterDelete(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	_ = list.InsertLast("1")
	_ = list.InsertLast("2")

	n3 := list.InsertLast("3")

	list.Delete(n3)

	test.EqualValues(2, list.Size())
}

func TestList_Prev_ReturnsPreviousNodeToGivenNode(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	test.Nil(list.Prev(nil), "prev must return nil for empty list")

	n1 := list.InsertFront("1")

	test.Nil(list.Prev(n1), "prev must return nil for head")

	n3 := list.InsertLast("3")

	test.EqualValues(n1, list.Prev(n3), "prev must return previous node for n3")

	n2 := list.InsertAfter("2", n1)

	test.EqualValues(n2, list.Prev(n3), "prev must return previous node for n3")
}

func TestList_NodeNext_ReturnsNextNodeFromGivenNode(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	n1 := list.InsertLast("1")
	_ = list.InsertLast("2")

	test.Equal("2", n1.Next().Data())
}

func TestList_NextNode_ReturnsNodeAfterTail(t *testing.T) {
	test := assert.New(t)

	var list linked.List

	_ = list.InsertLast("1")
	n2 := list.InsertLast("2")

	test.Nil(n2.Next(), "node must be nil")
}

func dump(list linked.List) []linked.Data {
	result := []linked.Data{}

	list.Range(func(node *linked.Node) {
		result = append(result, node.Data())
	})

	return result
}
