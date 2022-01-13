# ds

[![Home](https://godoc.org/github.com/gookit/event?status.svg)](file:///D:/EC-TSJ/Documents/CODE/SOURCE/Go/pkg/lib/cli)
[![Build Status](https://travis-ci.org/gookit/event.svg?branch=master)](https://travis-ci.org/)
[![Coverage Status](https://coveralls.io/repos/github/gookit/event/badge.svg?branch=master)](https://coveralls.io/github/)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/event)](https://goreportcard.com/report/github.com/)

> **[EN README](README.md)**

Ds es una librería para gestionar Stack y Queue.

## GoDoc

- [godoc for github](https://godoc.org/github.com/)

## Funciones Principales

> Objeto Base
- *tipo* `BaseError`  
- *tipo* `Base` 
- `Front() interface{}`
- `Back() interface{}` 
- `String() string `
- `Length() int` 
- `IsEmpty() bool` 
- `Clear()`

> Objeto Stack
- *tipo* `StackError` 
- *tipo* `Stack` *hereda al tipo* `Base`
- `NewStack() *Stack `
- `Push(valueOf ...interface{}) `
- `Pop() interface{} `
- `Peek() interface{}`

> Objeto Queue
- *tipo* `QueueError` 
- *tipo* `Queue` *hereda al tipo* `Base`
- `NewQueue() *Queue `
- `Enqueue(valueOf ...interface{})` 
- `Dequeue() interface{}` 
- `Peek() interface{} `



## Ejemplos
```go

	qt := ds.NewQueue()
	qt.Enqueue("uno", "dos", 5, "tres", "cuatro", "cinco", "seis")

	var ss interface{} = qt.Dequeue()
	if ss == nil {
		return
	}
	var ff interface{} = qt.Peek()
	if ff == nil {
		return
	}

	var tt int = qt.Length()
	var dd bool = qt.IsEmpty()
	if tt == 26 && dd {
		return
	}
	qt.Clear()

  
 	st := ds.NewStack()
	st.Push("uno", "dos", 5, "tres", "cuatro", "cinco")

	var s interface{} = st.Pop()
	if s == nil {
		return
	}
	var f interface{} = st.Peek()
	if f == nil {
		return
	}

	var t int = st.Length()
	var d bool = st.IsEmpty()
	if t == 26 && d {
		return
	}
	st.Clear()


```
## Notas




<!-- - [gookit/ini](https://github.com/gookit/ini) INI配置读取管理，支持多文件加载，数据覆盖合并, 解析ENV变量, 解析变量引用
-->
## LICENSE

**[MIT](LICENSE)**
