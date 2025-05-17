package tracker

import (
	"fmt"
	"sort"
)

type BruteForceTracker struct {
	domains     map[string]*Domain
	totalVisits int
	uniqueCount int
}

func NewBruteForceTracker() *BruteForceTracker {
	return &BruteForceTracker{
		domains:     make(map[string]*Domain),
		totalVisits: 0,
		uniqueCount: 0,
	}
}

func (t *BruteForceTracker) Visit(domainName string) {
	t.totalVisits++

	if domain, exists := t.domains[domainName]; exists {
		domain.Count++
	} else {
		t.uniqueCount++
		t.domains[domainName] = &Domain{Name: domainName, Count: 1}
	}
}

func (t *BruteForceTracker) GetTopN(n int) []Domain {
	allDomains := make([]Domain, 0, len(t.domains))
	for _, domain := range t.domains {
		allDomains = append(allDomains, *domain)
	}

	sort.Slice(allDomains, func(i, j int) bool {
		return allDomains[i].Count > allDomains[j].Count
	})

	if len(allDomains) > n {
		allDomains = allDomains[:n]
	}

	return allDomains
}

func (t *BruteForceTracker) Stats() string {
	return fmt.Sprintf("Total visits: %d, Unique domains: %d", t.totalVisits, t.uniqueCount)
}
