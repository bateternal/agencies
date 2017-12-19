package DataStructure

// Order - a separate element. Can be changed. Only priority field is used for data structure
type Order struct {
	action   *Service
	priority int
	customer string
}

// MaxHeap - container for Order
type MaxHeap struct {
	Orders  []Order
}

// NewPQueue - return pointer to MaxHeap instance
func New() *MaxHeap {
	return &MaxHeap{
		Orders:  []Order{},
	}
}

func (p *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (p *MaxHeap) leftChild(i int) int {
	return 2*i + 1
}

func (p *MaxHeap) rightChild(i int) int {
	return 2*i + 2
}

func (p *MaxHeap) siftUp(i int) {
	for i > 0 && p.Orders[p.parent(i)].priority < p.Orders[i].priority {
		p.Orders[p.parent(i)], p.Orders[i] = p.Orders[i], p.Orders[p.parent(i)]
		i = p.parent(i)
	}
}

func (p *MaxHeap) siftDown(i int) {
	maxInd := i
	l := p.leftChild(i)
	if l < p.Size() && p.Orders[l].priority > p.Orders[maxInd].priority {
		maxInd = l
	}
	r := p.rightChild(i)
	if r < p.Size() && p.Orders[r].priority > p.Orders[maxInd].priority {
		maxInd = r
	}
	if i != maxInd {
		p.Orders[i], p.Orders[maxInd] = p.Orders[maxInd], p.Orders[i]
		p.siftDown(maxInd)
	}
}

// Insert - add new Order to the priority queue
func (p *MaxHeap) Insert(item Order) {
	p.Orders = append(p.Orders, item)
	p.siftUp(p.Size() - 1)
}

// ExctractMax - pop element with the biggest priority
func (p *MaxHeap) ExctractMax() Order {
	r := p.Orders[0]
	x := p.Size() - 1
	p.Orders[0] = p.Orders[x]

	p.Orders = p.Orders[:x]
	p.siftDown(0)
	return r
}

// Size - return the len of the slice
func (p *MaxHeap) Size() int {
	return len(p.Orders)
}
