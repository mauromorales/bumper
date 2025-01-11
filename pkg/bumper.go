package bumper

import (
	"sort"

	semver "github.com/Masterminds/semver/v3"
)

// Package respresents a software package that can have semantic version.
// It is a node in a linked list of packages.
type Package struct {
	Name            string
	Version         string
	PreviousVersion string
	Next            *Package
	Prev            *Package
}

// NewPackage creates a new Package.
func NewPackage(name, version string) *Package {
	return &Package{
		Name:    name,
		Version: version,
	}
}

func (p *Package) IsGreaterThan(other *Package) bool {
	v1, _ := semver.NewVersion(p.Version)
	v2, _ := semver.NewVersion(other.Version)
	return v1.GreaterThan(v2)
}

// PackageList represents a linked list of packages.
type PackageList struct {
	Head *Package
	Tail *Package
	Size int
}

// NewPackageList creates a new PackageList.
func NewPackageList() *PackageList {
	return &PackageList{}
}

// Add adds a package to the list in order of version. The newest package is the head of the list.
func (pl *PackageList) Add(p *Package) {
	if pl.Head == nil {
		pl.Head = p
		pl.Tail = p
		pl.Size++
		return
	}

	current := pl.Head
	for current != nil {
		if p.IsGreaterThan(current) {
			if current.Prev != nil {
				current.Prev.Next = p
				p.Prev = current.Prev
			} else {
				pl.Head = p
			}
			p.Next = current
			current.Prev = p
			pl.Size++
			return
		}
		if current.Next == nil {
			current.Next = p
			p.Prev = current
			pl.Tail = p
			pl.Size++
			return
		}
		current = current.Next
	}
}

// RemoveByVersion removes a package from the list by its semantic version.
func (pl *PackageList) RemoveByVersion(version string) {
	current := pl.Head
	for current != nil {
		if current.Version == version {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				pl.Head = current.Next
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				pl.Tail = current.Prev
			}
			pl.Size--
			return
		}
		current = current.Next
	}
}

// Bump takes a list of versions and bumps each package with a newer version.
// The list can have more versions than the list of packages.
// We start by evaluating the head. If there's a newer version we bump it, else we move to the next.
// When the greatest patch version of a minor version is found, we bump it and move to the next minor version in the list.
func (pl *PackageList) Bump(versions []string) {
	vs := make([]*semver.Version, len(versions))
	for i, r := range versions {
		v, _ := semver.NewVersion(r)

		vs[i] = v
	}
	sort.Sort(semver.Collection(vs))

	// reverse the versions
	for i, j := 0, len(vs)-1; i < j; i, j = i+1, j-1 {
		vs[i], vs[j] = vs[j], vs[i]
	}

	// start by evaluating the head
	current := pl.Head
	for _, v := range vs {
		for current != nil {
			if current.Prev != nil {
				previousVersion, _ := semver.NewVersion(current.Prev.Version)
				if previousVersion.Major() == v.Major() && previousVersion.Minor() == v.Minor() {
					break
				}
			}
			currentVersion, _ := semver.NewVersion(current.Version)
			if v.GreaterThan(currentVersion) {
				current.PreviousVersion = current.Version
				current.Version = v.String()
			}
			current = current.Next
		}
	}
}

// Diff returns a list of how each version has changed.
func (pl *PackageList) Diff() []string {
	var diffs []string
	current := pl.Head
	for current != nil {
		if current.PreviousVersion == "" {
			diffs = append(diffs, current.Name+" "+current.Version)
		} else {
			diffs = append(diffs, current.Name+" "+current.PreviousVersion+" -> "+current.Version)
		}
		current = current.Next
	}
	return diffs
}
