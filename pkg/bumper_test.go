package bumper

import (
	"testing"
)

func TestNewPackageList(t *testing.T) {
	pl := NewPackageList()
	if pl.Head != nil {
		t.Errorf("Expected Head to not be nil, got %v", pl.Head)
	}
	if pl.Tail != nil {
		t.Errorf("Expected Tail to not be nil, got %v", pl.Tail)
	}
}

func TestAdd(t *testing.T) {
	pl := NewPackageList()
	p := NewPackage("foo", "1.0.0")
	pl.Add(p)
	if pl.Head != p {
		t.Errorf("Expected Head to be %v, got %v", p, pl.Head)
	}
	if pl.Tail != p {
		t.Errorf("Expected Tail to be %v, got %v", p, pl.Tail)
	}
	if p.Prev != nil {
		t.Errorf("Expected Prev to be nil, got %v", p.Prev)
	}
	if p.Next != nil {
		t.Errorf("Expected Next to be nil, got %v", p.Next)
	}
}

func TestAddMultiple(t *testing.T) {
	pl := NewPackageList()
	head := NewPackage("foo", "2.0.0")
	tail := NewPackage("foo", "1.0.0")
	pl.Add(head)
	pl.Add(tail)

	if pl.Head != head {
		t.Errorf("Expected Head to be %v, got %v", head, pl.Head)
	}
	if pl.Tail != tail {
		t.Errorf("Expected Tail to be %v, got %v", tail, pl.Tail)
	}
	if head.Prev != nil {
		t.Errorf("Expected Prev to be nil, got %v", head.Prev)
	}
	if head.Next != tail {
		t.Errorf("Expected Next to be %v, got %v", tail, head.Next)
	}
	if tail.Prev != head {
		t.Errorf("Expected Prev to be %v, got %v", head, tail.Prev)
	}
	if tail.Next != nil {
		t.Errorf("Expected Next to be nil, got %v", tail.Next)
	}
}

func TestAddMultipleReverse(t *testing.T) {
	pl := NewPackageList()
	head := NewPackage("foo", "2.0.0")
	tail := NewPackage("foo", "1.0.0")
	pl.Add(tail)
	pl.Add(head)

	if pl.Head != head {
		t.Errorf("Expected Head to be %v, got %v", head, pl.Head)
	}
	if pl.Tail != tail {
		t.Errorf("Expected Tail to be %v, got %v", tail, pl.Tail)
	}
	if head.Prev != nil {
		t.Errorf("Expected Prev to be nil, got %v", head.Prev)
	}
	if head.Next != tail {
		t.Errorf("Expected Next to be %v, got %v", tail, head.Next)
	}
	if tail.Prev != head {
		t.Errorf("Expected Prev to be %v, got %v", head, tail.Prev)
	}
	if tail.Next != nil {
		t.Errorf("Expected Next to be nil, got %v", tail.Next)
	}
}

func TestAddMultipleMiddle(t *testing.T) {
	pl := NewPackageList()
	head := NewPackage("foo", "3.0.0")
	middle := NewPackage("foo", "2.0.0")
	tail := NewPackage("foo", "1.0.0")
	pl.Add(middle)
	pl.Add(head)
	pl.Add(tail)

	if pl.Head != head {
		t.Errorf("Expected Head to be %v, got %v", head, pl.Head)
	}
	if pl.Tail != tail {
		t.Errorf("Expected Tail to be %v, got %v", tail, pl.Tail)
	}
	if head.Prev != nil {
		t.Errorf("Expected Prev to be nil, got %v", head.Prev)
	}
	if head.Next != middle {
		t.Errorf("Expected Next to be %v, got %v", middle, head.Next)
	}
	if middle.Prev != head {
		t.Errorf("Expected Prev to be %v, got %v", head, middle.Prev)
	}
	if middle.Next != tail {
		t.Errorf("Expected Next to be %v, got %v", tail, middle.Next)
	}
	if tail.Prev != middle {
		t.Errorf("Expected Prev to be %v, got %v", middle, tail.Prev)
	}
	if tail.Next != nil {
		t.Errorf("Expected Next to be nil, got %v", tail.Next)
	}
}

func TestRemoveByVersion(t *testing.T) {
	pl := NewPackageList()
	p := NewPackage("foo", "1.0.0")
	pl.Add(p)
	pl.RemoveByVersion("1.0.0")
	if pl.Head != nil {
		t.Errorf("Expected Head to be nil, got %v", pl.Head)
	}
	if pl.Tail != nil {
		t.Errorf("Expected Tail to be nil, got %v", pl.Tail)
	}
	if pl.Size != 0 {
		t.Errorf("Expected Size to be 0, got %v", pl.Size)
	}
}

func TestBump(t *testing.T) {
	pl := NewPackageList()
	p := NewPackage("foo", "1.0.0")
	pl.Add(p)
	pl.Bump([]string{"1.0.1"})
	if p.Version != "1.0.1" {
		t.Errorf("Expected Version to be 1.0.1, got %v", p.Version)
	}
}

func TestBumpMultiple(t *testing.T) {
	pl := NewPackageList()
	p1 := NewPackage("foo", "1.0.0")
	p2 := NewPackage("foo", "1.1.0")
	pl.Add(p1)
	pl.Add(p2)
	pl.Bump([]string{"1.0.0", "1.1.0", "1.0.1", "1.1.1"})
	if pl.Tail.Version != "1.0.1" {
		t.Errorf("Expected Version to be 1.0.1, got %v", pl.Tail.Version)
	}
	if pl.Head.Version != "1.1.1" {
		t.Errorf("Expected Version to be 1.1.1, got %v", pl.Head.Version)
	}
}

func TestBumpMultiple2(t *testing.T) {
	pl := NewPackageList()
	p1 := NewPackage("foo", "1.0.0")
	p2 := NewPackage("foo", "1.1.0")
	p3 := NewPackage("foo", "1.2.0")
	pl.Add(p1)
	pl.Add(p2)
	pl.Add(p3)
	pl.Bump([]string{"1.0.0", "1.1.0", "1.0.1", "1.1.1", "2.0.0", "1.2.1"})
	if pl.Tail.Version != "1.1.1" {
		t.Errorf("Expected Version to be 1.1.1, got %v", pl.Tail.Version)
	}
	if pl.Head.Version != "2.0.0" {
		t.Errorf("Expected Version to be 2.0.0, got %v", pl.Head.Version)
	}
	if pl.Head.Next.Version != "1.2.1" {
		t.Errorf("Expected Version to be 1.2.1, got %v", pl.Head.Next.Version)
	}
}

func TestDiff(t *testing.T) {
	pl := NewPackageList()
	p1 := NewPackage("foo", "1.0.0")
	p2 := NewPackage("foo", "1.1.0")
	p3 := NewPackage("foo", "1.2.0")
	pl.Add(p1)
	pl.Add(p2)
	pl.Add(p3)
	pl.Bump([]string{"1.0.0", "1.1.0", "1.0.1", "1.1.1", "2.0.0", "1.2.1"})
	diff := pl.Diff()
	if diff[0] != "foo 1.2.0 -> 2.0.0" {
		t.Errorf("Expected Diff to be foo 1.2.0 -> 2.0.0, got %v", diff[0])
	}
	if diff[1] != "foo 1.1.0 -> 1.2.1" {
		t.Errorf("Expected Diff to be foo 1.1.0 -> 1.2.1, got %v", diff[1])
	}
	if diff[2] != "foo 1.0.0 -> 1.1.1" {
		t.Errorf("Expected Diff to be foo 1.0.0 -> 1.1.1, got %v", diff[2])
	}
}

func TestDiff2(t *testing.T) {
	pl := NewPackageList()
	p1 := NewPackage("foo", "1.0.0")
	p2 := NewPackage("foo", "1.1.0")
	p3 := NewPackage("foo", "1.2.0")
	pl.Add(p1)
	pl.Add(p2)
	pl.Add(p3)
	pl.Bump([]string{"1.0.0", "1.1.0", "1.0.1", "1.2.1"})
	diff := pl.Diff()
	if diff[0] != "foo 1.2.0 -> 1.2.1" {
		t.Errorf("Expected Diff to be foo 1.2.0 -> 1.2.1, got %v", diff[0])
	}
	if diff[1] != "foo 1.1.0" {
		t.Errorf("Expected Diff to be foo 1.1.0, got %v", diff[1])
	}
	if diff[2] != "foo 1.0.0 -> 1.0.1" {
		t.Errorf("Expected Diff to be foo 1.0.0 -> 1.0.1, got %v", diff[2])
	}
}
