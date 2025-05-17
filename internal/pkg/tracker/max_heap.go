package tracker

import "fmt"

type MaxHeap struct {
	domains     []Domain
	domainIdx   map[string]int
	totalVisits int
	uniqueCount int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		domains:   make([]Domain, 0),
		domainIdx: make(map[string]int),
	}
}

func (h *MaxHeap) size() int {
	return len(h.domains)
}

func (h *MaxHeap) isEmpty() bool {
	return len(h.domains) == 0
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h *MaxHeap) rightChild(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) getMax() (Domain, error) {
	if h.isEmpty() {
		return Domain{}, fmt.Errorf("heap is empty")
	}

	return h.domains[0], nil
}

func (h *MaxHeap) findIndexByName(name string) int {
	idx, exists := h.domainIdx[name]
	if !exists {
		return -1
	}

	return idx
}

func (h *MaxHeap) insertOrUpdate(domainName string) {
	index := h.findIndexByName(domainName)

	if index == -1 {
		h.uniqueCount++

		h.domains = append(
			h.domains, Domain{
				Name:  domainName,
				Count: 1,
			})

		h.domainIdx[domainName] = len(h.domains) - 1

		h.siftUp(len(h.domains) - 1)

		return
	}

	h.domains[index].Count++

	h.siftUp(index)
	h.siftDown(index)
}

func (h *MaxHeap) updateNameToIDMap(i, j int) {
	h.domainIdx[h.domains[i].Name] = i
	h.domainIdx[h.domains[j].Name] = j
}

func (h *MaxHeap) siftUp(index int) {
	parent := h.parent(index)
	for index > 0 && h.domains[parent].Count < h.domains[index].Count {
		h.domains[parent], h.domains[index] = h.domains[index], h.domains[parent]

		h.updateNameToIDMap(parent, index)

		index = parent
		parent = h.parent(index)
	}
}

func (h *MaxHeap) siftDown(index int) {
	largest := index
	left := h.leftChild(index)
	right := h.rightChild(index)
	size := len(h.domains)

	if left < size && h.domains[left].Count > h.domains[largest].Count {
		largest = left
	}

	if right < size && h.domains[right].Count > h.domains[largest].Count {
		largest = right
	}

	if largest != index {
		h.domains[index], h.domains[largest] = h.domains[largest], h.domains[index]

		h.updateNameToIDMap(index, largest)

		h.siftDown(largest)
	}
}

func (h *MaxHeap) extractMax() (Domain, error) {
	if h.isEmpty() {
		return Domain{}, fmt.Errorf("heap is empty")
	}

	max := h.domains[0]
	delete(h.domainIdx, max.Name)

	lastIndex := len(h.domains) - 1
	h.domains[0] = h.domains[lastIndex]

	h.domainIdx[h.domains[0].Name] = 0

	h.domains = h.domains[:lastIndex]

	if !h.isEmpty() {
		h.siftDown(0)
	}

	return max, nil
}

func (h *MaxHeap) display() {
	fmt.Println("Heap domains:")

	for i, domain := range h.domains {
		fmt.Printf("[%d] %s: %d\n", i, domain.Name, domain.Count)
	}

	fmt.Println()
}

func (h *MaxHeap) Visit(domainName string) {
	h.totalVisits++

	h.insertOrUpdate(domainName)
}

func (h *MaxHeap) GetTopN(n int) []Domain {
	if n <= 0 {
		return []Domain{}
	}

	heapCopy := &MaxHeap{
		domains:   make([]Domain, len(h.domains)),
		domainIdx: map[string]int{},
	}

	copy(heapCopy.domains, h.domains)
	heapCopy.domainIdx = h.domainIdx

	result := make([]Domain, 0, n)
	for i := 0; i < n && !heapCopy.isEmpty(); i++ {
		max, _ := heapCopy.extractMax()
		result = append(result, max)
	}

	return result
}
