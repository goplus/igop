/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package igop

const defaultStkSize = 64

// A Stack represents a FILO container.
type frameStack struct {
	data []*frame
	size int
}

// NewStack creates a Stack instance.
func NewStack() (p *frameStack) {
	return &frameStack{data: make([]*frame, defaultStkSize, defaultStkSize)}
}

// Init initializes this Stack object.
func (p *frameStack) Init() {
	p.data = make([]*frame, 0, defaultStkSize)
}

// Get returns the value at specified index.
func (p *frameStack) Get(idx int) *frame {
	return p.data[len(p.data)+idx]
}

// Set returns the value at specified index.
func (p *frameStack) Set(idx int, v *frame) {
	p.data[len(p.data)+idx] = v
}

// Push pushes a value into this stack.
func (p *frameStack) Push(v *frame) {
	//p.data = append(p.data, v)
	p.data[p.size] = v
	p.size++
}

// Pop pops a value from this stack.
func (p *frameStack) Pop() {
	//p.data = p.data[:len(p.data)-1]
	p.size--
}

// func (p *frameStack) Pop() *frame {
// 	n := len(p.data)
// 	v := p.data[n-1]
// 	p.data = p.data[:n-1]
// 	return v
// }

// Len returns count of stack elements.
func (p *frameStack) Len() int {
	return len(p.data)
}
