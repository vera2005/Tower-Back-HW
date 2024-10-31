package main

import (
	"fmt"
)

type bst struct {
	value    int
	leftVal  *bst
	rightVal *bst
}

func (elem *bst) Add(i int) {
	tekElem := elem
	tekVal := 0
	for {
		tekVal = tekElem.value
		if tekVal == i {
			return
		} else if i > tekVal {
			if tekElem.rightVal == nil {
				tekElem.rightVal = &bst{value: i}
				return
			}
			tekElem = tekElem.rightVal
		} else if i < tekVal {
			if tekElem.leftVal == nil {
				tekElem.leftVal = &bst{value: i}
				return
			}
			tekElem = tekElem.leftVal
		}
	}
}

func (elem *bst) Delete(i int) {
	var parent *bst
	tekElem := elem

	// Поиск узла для удаления
	for tekElem != nil && tekElem.value != i {
		parent = tekElem
		if i < tekElem.value {
			tekElem = tekElem.leftVal
		} else {
			tekElem = tekElem.rightVal
		}
	}

	// Если элемент не найден
	if tekElem == nil {
		fmt.Printf("Элемент %d НЕ существует в данном дереве\n", i)
		return
	}

	// Если узел имеет только одного ребенка или не имеет детей
	if tekElem.leftVal == nil {
		if parent == nil {
			*elem = *tekElem.rightVal // если это корень
		} else if parent.leftVal == tekElem {
			parent.leftVal = tekElem.rightVal
		} else {
			parent.rightVal = tekElem.rightVal
		}
	} else if tekElem.rightVal == nil {
		if parent == nil {
			*elem = *tekElem.leftVal // если это корень
		} else if parent.leftVal == tekElem {
			parent.leftVal = tekElem.leftVal
		} else {
			parent.rightVal = tekElem.leftVal
		}
	} else {
		//  Узел имеет двух детей
		minParent := tekElem
		minElem := tekElem.rightVal

		// Поиск минимального элемента в правом поддереве
		for minElem.leftVal != nil {
			minParent = minElem
			minElem = minElem.leftVal
		}
		// Копирование значение минимального узла в удаляемый узел
		tekElem.value = minElem.value
		// Удаление минимального узла из правого поддерева
		if minParent.leftVal == minElem {
			minParent.leftVal = minElem.rightVal
		} else {
			minParent.rightVal = minElem.rightVal
		}
	}
}

func (elem *bst) IsExit(i int) {
	tekElem := elem
	tekVal := 0
	for {
		if tekElem != nil {
			tekVal = tekElem.value
			if i == tekVal {
				fmt.Printf("Элемент %d существует в данном дереве\n", i)
				return
			} else if i > tekVal {
				tekElem = tekElem.rightVal
			} else {
				tekElem = tekElem.leftVal
			}
		} else {
			fmt.Printf("Элемент %d НЕ существует в данном дереве", i)
			return
		}
	}
}

func main() {
	myBst := bst{value: 5}
	values := []int{2, 3, 5, 1, 12, 56, 78, 34, 23, 12}
	for _, val := range values {
		myBst.Add(val)
	}
	myBst.IsExit(78)
	myBst.IsExit(100)
	myBst.Delete(56)
	for _, val := range values {
		myBst.IsExit(val)
	}
}
