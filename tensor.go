package tensor2go

import (
	"fmt"

	"github.com/joernweissenborn/propertygraph2go"
)

const (
	ROOT     = "ROOT"
	DIM_EDGE = "D"
)

type Tensor struct {
	graph propertygraph2go.PropertyGraph

	tensorRoot propertygraph2go.Vertex

	indexPool *propertygraph2go.IndexPool
}

func NewTensor(dims ...int) (t *Tensor) {
	t = &Tensor{
		graph:     propertygraph2go.NewInMemoryGraph(),
		indexPool: propertygraph2go.NewIndexPool(),
	}

	for _, dim := range dims {
		fmt.Println("Adding", dim)
		t.addDim(dim)
	}
	return
}

func (t *Tensor) createTensorRoot(p interface{}) {
	fmt.Println("creating root")
	t.tensorRoot = t.graph.CreateVertex(ROOT, p)
}

func (t *Tensor) addDim(size int) {
	if t.tensorRoot == nil {
		t.createTensorRoot(&tensorProperties{0})
	}
	tp, ok := t.tensorRoot.Properties().(*tensorProperties)

	if !ok {
		panic("Adding dim on scalar")
	}

	newDimIndex := tp.dims
	tp.raiseDims()
	indices := t.Indices()
	dimLabel := fmt.Sprintf("%d", newDimIndex)
	fmt.Println("Adding dinmd", dimLabel)
	fmt.Println("Adding dinmd", indices)

	if len(indices) == 0 {
		for i := 0; i < size; i++ {

			t.addItem(Zero{}, []int{i})
		}
	} else {

		for _, index := range indices {
			item := t.getItemVertex(index)
			t.graph.CreateEdge(t.indexPool.GetEdgeIndex(), dimLabel, t.tensorRoot, item, 0)
		}

		for i := 1; i < size; i++ {
			for _, index := range indices {
				t.addItem(Zero{}, append(index, i))
			}
		}

	}
}

func (t *Tensor) Indices() (indices [][]int) {
	indices = [][]int{}
	for dim := 0; dim < t.Dims(); dim++ {
		old := indices
		indices = [][]int{}
		for i := 0; i < t.DimSize(dim); i++ {
			for _, j := range old {
				indices = append(indices, append(j, i))
			}
		}
	}
	return
}

func (t *Tensor) DimSize(dim int) (size int) {
	fmt.Println(t.tensorRoot)
	return len(t.graph.GetOutgoingEdgesByLabel(ROOT, fmt.Sprintf("%d", dim)))
}

func (t *Tensor) Dims() (dims int) {
	tp, ok := t.tensorRoot.Properties().(*tensorProperties)
	if ok {
		dims = tp.dims
	}
	return
}

func (t *Tensor) getItemVertex(indices []int) (v propertygraph2go.Vertex) {
	if len(indices) != 0 {
		panic("Tried to get item vertex without suppling enough indices")
	}
	candidateEdges := t.graph.GetOutgoingEdgesByLabel(ROOT, fmt.Sprintf("%d", 0))
	var nextCandidateEdges []propertygraph2go.Edge

	for i, index := range indices {
		for _, c := range candidateEdges {
			if c.Properties().(int) == index {
				if i < len(indices)-1 {
					for _, e := range t.graph.GetIncomingEdgesByLabel(c.Head().Id(), fmt.Sprintf("%d", i+1)) {
						nextCandidateEdges = append(nextCandidateEdges, e)
					}
				} else {
					v = c.Head()
				}
			}
		}
	}
	return
}

func (t *Tensor) addItem(val interface{}, index []int) {
	item := t.graph.CreateVertex(t.indexPool.GetVertexIndex(), val)
	for dim := range index {
		if dim > t.Dims() {
			panic(fmt.Sprintf("Tried to add item on non existend dim. Tensor rank is %d, requested dim is %d", t.Dims(), dim))
		}
		label := fmt.Sprintf("%d", dim)
		t.graph.CreateEdge(
			t.indexPool.GetEdgeIndex(),
			label,
			t.tensorRoot,
			item,
			len(t.graph.GetOutgoingEdgesByLabel(ROOT, label)),
		)
	}
}

//func (Tensor) Contract(Tensor) Tensor {

//}
