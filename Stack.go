/**
 * @fileoverview
 * @author Torres Sacristán, Jesús <ec-tsj@gmail.org>
 * @version 1.0
 * @copyright XCL System
 * @category System
 * @package System
 * @subpackage System.Base
 * @filesource
 */

package ds

import (
	"ec-tsj/container/list"
	"ec-tsj/core"
)

// Stack Definición
type (
	IStack interface {
		Push(...core.T)
		Pop() core.T
		Peek() core.T
	}

	Stack struct {
		*list.List
	}
)

// NewStack es el constructor
// @constructor
// @params {...core.T}
// @return {*Stack}
func NewStack(parms ...core.T) *Stack {
	stack := &Stack{list.Sx()} //&Stack{base{head: &node{content: "@Init@"}, size: 1}}
	if len(parms) > 0 {
		stack.Push(parms...)
	}
	return stack
}

//!+
// Push un elemento en la pila.
// @param {any} element
func (stack *Stack) Push(valueOf ...core.T) {
	stack.List.Push(valueOf...)
}

// Pop un elemento desde la pila
// Retorna el elemento en la top position de la pila y decrementa la variable length
func (stack *Stack) Pop() core.T {
	node, _ := stack.List.Pop()
	return node
}

// Peek obtiene un elemento de la pila. El que está próximo a salir.
func (stack *Stack) Peek() core.T {
	return stack.Locate(stack.Length() - 1).Value()
}

//!-
