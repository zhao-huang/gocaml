package gcil

import (
	"fmt"
	"io"
	"strings"
)

type Val interface {
	Print(io.Writer)
}

type OperatorKind int

// Operators
const (
	NOT OperatorKind = iota
	NEG
	FNEG
	ADD
	SUB
	FADD
	FSUB
	FMUL
	FDIV
	LESS
	EQ
)

var opTable = [...]string{
	NOT:  "not",
	NEG:  "-",
	FNEG: "-.",
	ADD:  "+",
	SUB:  "-",
	FADD: "+.",
	FSUB: "-.",
	FMUL: "*.",
	FDIV: "/.",
	LESS: "<",
	EQ:   "=",
}

// Kind of function call.
type AppKind int

const (
	// Means to call a function without closure
	DIRECT_CALL AppKind = iota
	CLOSURE_CALL
	EXTERNAL_CALL
)

var appTable = [...]string{
	DIRECT_CALL:   "",
	CLOSURE_CALL:  "cls",
	EXTERNAL_CALL: "x",
}

type (
	Unit struct{}
	Bool struct {
		Const bool
	}
	Int struct {
		Const int
	}
	Float struct {
		Const float64
	}
	Unary struct {
		Op    OperatorKind
		Child string
	}
	Binary struct {
		Op  OperatorKind
		Lhs string
		Rhs string
	}
	Ref struct {
		Ident string
	}
	If struct {
		Cond string
		Then *Block
		Else *Block
	}
	Fun struct {
		Params []string
		Body   *Block
	}
	App struct {
		Callee string
		Args   []string
		Kind   AppKind
	}
	Tuple struct {
		Elems []string
	}
	Array struct {
		Size string
		Elem string
	}
	TplLoad struct { // Used for each element of LetTuple
		From  string
		Index int
	}
	ArrLoad struct {
		From  string
		Index string
	}
	ArrStore struct {
		To    string
		Index string
		Rhs   string
	}
	XRef struct {
		Ident string
	}
	NOP struct {
	}
	// Introduced at closure-transform.
	MakeCls struct {
		Vars []string
		Fun  string
	}
)

var (
	UnitVal = &Unit{}
	NOPVal  = &NOP{}
)

func (v *Unit) Print(out io.Writer) {
	fmt.Fprintf(out, "unit")
}
func (v *Bool) Print(out io.Writer) {
	fmt.Fprintf(out, "bool %v", v.Const)
}
func (v *Int) Print(out io.Writer) {
	fmt.Fprintf(out, "int %d", v.Const)
}
func (v *Float) Print(out io.Writer) {
	fmt.Fprintf(out, "float %f", v.Const)
}
func (v *Unary) Print(out io.Writer) {
	fmt.Fprintf(out, "unary %s %s", opTable[v.Op], v.Child)
}
func (v *Binary) Print(out io.Writer) {
	fmt.Fprintf(out, "binary %s %s %s", opTable[v.Op], v.Lhs, v.Rhs)
}
func (v *Ref) Print(out io.Writer) {
	fmt.Fprintf(out, "ref %s", v.Ident)
}
func (v *If) Print(out io.Writer) {
	fmt.Fprintf(out, "if %s", v.Cond)
}
func (v *Fun) Print(out io.Writer) {
	fmt.Fprintf(out, "fun %s", strings.Join(v.Params, ","))
}
func (v *App) Print(out io.Writer) {
	fmt.Fprintf(out, "app%s %s %s", appTable[v.Kind], v.Callee, strings.Join(v.Args, ","))
}
func (v *Tuple) Print(out io.Writer) {
	fmt.Fprintf(out, "tuple %s", strings.Join(v.Elems, ","))
}
func (v *Array) Print(out io.Writer) {
	fmt.Fprintf(out, "array %s %s", v.Size, v.Elem)
}
func (v *TplLoad) Print(out io.Writer) {
	fmt.Fprintf(out, "tplload %d %s", v.Index, v.From)
}
func (v *ArrLoad) Print(out io.Writer) {
	fmt.Fprintf(out, "arrload %s %s", v.Index, v.From)
}
func (v *ArrStore) Print(out io.Writer) {
	fmt.Fprintf(out, "arrstore %s %s %s", v.Index, v.To, v.Rhs)
}
func (v *XRef) Print(out io.Writer) {
	fmt.Fprintf(out, "xref %s", v.Ident)
}
func (v *NOP) Print(out io.Writer) {
	fmt.Fprintf(out, "nop")
}
func (v *MakeCls) Print(out io.Writer) {
	fmt.Fprintf(out, "makecls (%s) %s", strings.Join(v.Vars, ","), v.Fun)
}