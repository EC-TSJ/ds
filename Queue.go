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
	IQueue interface {
		Enqueue(...core.T)
		Dequeue() core.T
		Peek() core.T
	}

	Queue struct {
		*list.List
	}
)

// NewQueue ...
// @constructor
// @params {...core.T}
// @return {*Queue}
func NewQueue(parms ...core.T) *Queue {
	queue := &Queue{list.Sx()} //&Queue{base{head: &node{content: "@Init@"}, size: 0}}
	if len(parms) > 0 {
		queue.Enqueue(parms...)
	}
	return queue
}

// Enqueue añade un elemento al final de una cola
// @param {any} element
func (queue *Queue) Enqueue(valueOf ...core.T) { //push
	queue.List.Push(valueOf...)
}

//!+
// Dequeue remueve un elemento desde el frente de una cola
func (queue *Queue) Dequeue() core.T { //pop
	node, _ := queue.List.Pop(0)
	return node
}

// Peek Obtiene un elemento de la pila. El que está próximo a salir.
func (queue *Queue) Peek() core.T {
	return queue.Locate(0).Value()
}

//!-
