package Listas

import "fmt"

type edge struct {
	node  string
	label string
}
type graph struct {
	nodes map[string][]edge
}

func newGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) addEdge(from, to, label string) {
	g.nodes[from] = append(g.nodes[from], edge{node: to, label: label})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (e *edge) String() string {
	return fmt.Sprintf("%v", e.node)
}

func (g *graph) String() string {
	out := `digraph finite_state_machine {
		rankdir=LR;
		size="8,5"
		node [shape = circle];`
	for k := range g.nodes {
		for _, v := range g.getEdges(k) {
			out += fmt.Sprintf("\t%s -> %s\t[ label = \"%s\" ];\n", k, v.node, v.label)
		}
	}
	out += "}"
	return out
}

func Print1() {
	g := newGraph()
	// https://graphviz.gitlab.io/_pages/Gallery/directed/fsm.html

	g.addEdge("LR_0", "LR_2", "SS(B)")
	g.addEdge("LR_0", "LR_1", "SS(S)")
	g.addEdge("LR_1", "LR_3", "S($end)")
	g.addEdge("LR_2", "LR_6", "SS(b)")
	g.addEdge("LR_2", "LR_5", "SS(a)")
	g.addEdge("LR_2", "LR_4", "S(A)")
	g.addEdge("LR_5", "LR_7", "S(b)")
	g.addEdge("LR_5", "LR_5", "S(a)")
	g.addEdge("LR_6", "LR_6", "S(b)")
	g.addEdge("LR_6", "LR_5", "S(a)")
	g.addEdge("LR_7", "LR_8", "S(b)")
	g.addEdge("LR_7", "LR_5", "S(a)")
	g.addEdge("LR_8", "LR_6", "S(b)")
	g.addEdge("LR_8", "LR_5", "S(a)")

	fmt.Println(g)
}
