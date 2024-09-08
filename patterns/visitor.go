package main

import "fmt"

// Паттерн "Посетитель" позволяет отслеживать состояние сложных структур данных

type Node interface {
	Accept(visitor Visitor)
}

type Visitor interface {
	VisitLeaf(leaf *Leaf)
	VisitBranch(branch *Branch)
}

type Leaf struct {
	Value int
}

func (l *Leaf) Accept(visitor Visitor) {
	visitor.VisitLeaf(l)
}

type Branch struct {
	Left  Node
	Right Node
}

func (b *Branch) Accept(visitor Visitor) {
	visitor.VisitBranch(b)
}

type PrintVisitor struct{}

func (p *PrintVisitor) VisitLeaf(leaf *Leaf) {
	fmt.Println("Leaf value:", leaf.Value)
}

func (p *PrintVisitor) VisitBranch(branch *Branch) {
	fmt.Println("Branch")
	branch.Left.Accept(p)
	branch.Right.Accept(p)
}

func main() {
	tree := &Branch{
		Left: &Leaf{Value: 1},
		Right: &Branch{
			Left:  &Leaf{Value: 2},
			Right: &Leaf{Value: 3},
		},
	}

	printVisitor := &PrintVisitor{}
	tree.Accept(printVisitor)
}
