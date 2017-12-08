package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Catch struct {
	node.SimpleNode
	token    token.Token
	types    []node.Node
	variable node.Node
	stmts    []node.Node
}

func NewCatch(token token.Token, types []node.Node, variable node.Node, stmts []node.Node) node.Node {
	return Catch{
		node.SimpleNode{Name: "Catch", Attributes: make(map[string]string)},
		token,
		types,
		variable,
		stmts,
	}
}

func (n Catch) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)

	fmt.Fprintf(out, "\n%vtypes:", indent+"  ")
	for _, nn := range n.types {
		nn.Print(out, indent+"    ")
	}

	fmt.Fprintf(out, "\n%vvariable:", indent+"  ")
	n.variable.Print(out, indent+"    ")

	if n.stmts != nil {
		fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
		for _, nn := range n.stmts {
			nn.Print(out, indent+"    ")
		}
	}
}