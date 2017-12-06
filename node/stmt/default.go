package stmt

import (
	"fmt"
	"io"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Default struct {
	node.SimpleNode
	token token.Token
	stmts node.Node
}

func NewDefault(token token.Token, stmts node.Node) node.Node {
	return Default{
		node.SimpleNode{Name: "Default", Attributes: make(map[string]string)},
		token,
		stmts,
	}
}

func (n Default) Print(out io.Writer, indent string) {
	fmt.Fprintf(out, "\n%v%v [%d %d] %q", indent, n.Name, n.token.StartLine, n.token.EndLine, n.token.Value)
	fmt.Fprintf(out, "\n%vstmts:", indent+"  ")
	n.stmts.Print(out, indent+"    ")
}
