package main

import "testing"

func TestPushFrontSize(t *testing.T) {
	l := newLinkedList[int]()
	l.pushFront(1)
	l.pushFront(2)
	l.pushFront(3)

	if l.size != 3 {
		t.Fatalf("expected:%v, got:%v", 3, l.size)
	}

	if l.string() != "[ 3 -> 2 -> 1 ]" {
		t.Errorf("expected:%v, got:%v", "[ 3 -> 2 -> 1 ]", l.string())
	}
}

func TestPushBackSize(t *testing.T) {
	l := newLinkedList[int]()
	l.pushBack(1)
	l.pushBack(2)

	if l.size != 2 {
		t.Fatalf("expected:%v, got:%v", 2, l.size)
	}

	if l.string() != "[ 1 -> 2 ]" {
		t.Errorf("expected:%v, got:%v", "[ 1 -> 2 ]", l.string())
	}
}

// get bounds-checks against size, so a wrong size makes
// front-pushed values unreachable
func TestGetAfterPushFront(t *testing.T) {
	l := newLinkedList[int]()
	l.pushFront(1)
	l.pushBack(10)
	l.pushFront(20)

	want := []int{20, 1, 10}
	for i, expected := range want {
		got, ok := l.get(i)
		if !ok {
			t.Fatalf("get(%v) not ok, size is %v", i, l.size)
		}
		if got != expected {
			t.Errorf("get(%v) expected:%v, got:%v", i, expected, got)
		}
	}

	if _, ok := l.get(3); ok {
		t.Errorf("get(3) expected out of range on a list of %v", l.size)
	}

	if _, ok := l.get(-1); ok {
		t.Errorf("get(-1) expected out of range")
	}
}

func TestPopFront(t *testing.T) {
	l := newLinkedList[int]()
	l.pushBack(1)
	l.pushBack(2)

	value, ok := l.popFront()
	if !ok || value != 1 {
		t.Fatalf("expected:%v true, got:%v %v", 1, value, ok)
	}

	if l.size != 1 {
		t.Errorf("expected:%v, got:%v", 1, l.size)
	}

	if _, ok := l.popFront(); !ok {
		t.Fatal("expected the last pop to succeed")
	}

	if _, ok := l.popFront(); ok {
		t.Error("expected pop on an empty list to fail")
	}

	if l.head != nil || l.tail != nil {
		t.Errorf("expected head and tail nil, got:%v %v", l.head, l.tail)
	}
}

func TestPopBack(t *testing.T) {
	l := newLinkedList[int]()
	l.pushBack(1)
	l.pushBack(2)

	value, ok := l.popBack()
	if !ok || value != 2 {
		t.Fatalf("expected:%v true, got:%v %v", 2, value, ok)
	}

	if l.size != 1 {
		t.Errorf("expected:%v, got:%v", 1, l.size)
	}

	if _, ok := l.popBack(); !ok {
		t.Fatal("expected the last pop to succeed")
	}

	if _, ok := l.popBack(); ok {
		t.Error("expected pop on an empty list to fail")
	}

	if l.head != nil || l.tail != nil {
		t.Errorf("expected head and tail nil, got:%v %v", l.head, l.tail)
	}
}

func TestRemove(t *testing.T) {
	l := newLinkedList[int]()
	l.pushBack(1)
	l.pushBack(2)
	l.pushBack(3)

	// middle
	if ok := l.remove(2); !ok {
		t.Fatal("expected remove(2) to succeed")
	}

	if l.size != 2 {
		t.Errorf("expected:%v, got:%v", 2, l.size)
	}

	if l.string() != "[ 1 -> 3 ]" {
		t.Errorf("expected:%v, got:%v", "[ 1 -> 3 ]", l.string())
	}

	// head, so tail links get rewired
	if ok := l.remove(1); !ok {
		t.Fatal("expected remove(1) to succeed")
	}

	if l.head.value != 3 || l.tail.value != 3 {
		t.Errorf("expected head and tail to be 3, got:%v %v", l.head.value, l.tail.value)
	}

	if ok := l.remove(99); ok {
		t.Error("expected remove of a missing value to fail")
	}
}

func TestInsertAt(t *testing.T) {
	l := newLinkedList[int]()
	l.pushBack(1)
	l.pushBack(3)

	// middle
	if ok := l.insertAt(1, 2); !ok {
		t.Fatal("expected insertAt(1, 2) to succeed")
	}

	// front
	if ok := l.insertAt(0, 0); !ok {
		t.Fatal("expected insertAt(0, 0) to succeed")
	}

	// back, idx == size
	if ok := l.insertAt(l.size, 4); !ok {
		t.Fatal("expected insertAt at the end to succeed")
	}

	if l.size != 5 {
		t.Errorf("expected:%v, got:%v", 5, l.size)
	}

	if l.string() != "[ 0 -> 1 -> 2 -> 3 -> 4 ]" {
		t.Errorf("expected:%v, got:%v", "[ 0 -> 1 -> 2 -> 3 -> 4 ]", l.string())
	}

	if ok := l.insertAt(99, 1); ok {
		t.Error("expected insertAt past the end to fail")
	}

	if ok := l.insertAt(-1, 1); ok {
		t.Error("expected insertAt with a negative index to fail")
	}
}

func TestEmptyList(t *testing.T) {
	l := newLinkedList[string]()

	if l.size != 0 {
		t.Errorf("expected:%v, got:%v", 0, l.size)
	}

	if _, ok := l.get(0); ok {
		t.Error("expected get on an empty list to fail")
	}

	if l.string() != "[  ]" {
		t.Errorf("expected:%v, got:%v", "[  ]", l.string())
	}
}
