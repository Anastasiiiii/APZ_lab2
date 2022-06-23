package lab2

import (
	"errors"
	"strconv"
	"strings"
)

type Node struct {
	el   string
	next *Node
}

func newNode(el string, top *Node) *Node {
	return &Node{
		el,
		top,
	}
}

type Stack struct {
	top  *Node
	size int
}

func createStack() *Stack {
	return &Stack{
		nil,
		0,
	}
}

func (this *Stack) pushEl(el string) {
	this.top = newNode(el, this.top)
	this.size++
}

func (this Stack) isEmptyStack() bool {
	if this.size > 0 {
		return false
	} else {
		return true
	}
}

func (this *Stack) popEl() string {
	var tmpStr string = ""
	if !this.isEmptyStack() {
		tmpStr = this.top.el
		this.top = this.top.next
		this.size--
	}
	return tmpStr
}

func isOperator(symbol byte) bool {
	switch symbol {
	case
		'+',
		'-',
		'*',
		'/',
		'^':
		return true
	}
	return false
}

func isANumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func PrefixToPostfix(prefix string) (string, error) {
	if prefix == "" || prefix == " " {
		return "", errors.New("Empty string is passed to function")
	}
	var arr []string = strings.Split(prefix, " ")
	var stack *Stack = createStack()
	var el string = ""
	var num1 string = ""
	var num2 string = ""
	var res string
	var err error

	for i := len(arr) - 1; i >= 0; i-- {
		if isOperator(arr[i][0]) {
			if stack.size > 1 {
				num1 = stack.popEl()
				num2 = stack.popEl()
				el = num1 + " " + num2 + " " + string(arr[i][0])
				stack.pushEl(el)
			} else {
				res = ""
				err = errors.New("Invalid syntax")
			}
		} else if isANumber(arr[i]) {
			el = string(arr[i])
			stack.pushEl(el)
		} else {
			res = ""
			err = errors.New("Invalid syntax")
			break
		}
	}

	res = stack.popEl()
	err = nil

	return res, err
}
