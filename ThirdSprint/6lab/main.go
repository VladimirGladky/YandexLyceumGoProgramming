package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

// Node представляет узел в связанном списке.
type Node struct {
	value int            // Значение узла
	next  unsafe.Pointer // Указатель на следующий узел в списке
}

// List представляет связанный список.
type List struct {
	head unsafe.Pointer // Указатель на голову списка
}

// Add добавляет новый узел с заданным значением в начало списка.
func (l *List) Add(value int) {
	node := &Node{value: value}

	for {
		// Загружаем текущую голову списка
		oldHead := atomic.LoadPointer(&l.head)
		// Устанавливаем новый узел как следующий для добавляемого узла
		node.next = oldHead

		// Если голову можно атомарно заменить на новый узел, то прерываем цикл
		if atomic.CompareAndSwapPointer(&l.head, oldHead, unsafe.Pointer(node)) {
			break
		}
	}
}

// Print выводит значения узлов списка.
func (l *List) Print() {
	// Загружаем голову списка
	curr := atomic.LoadPointer(&l.head)

	// Проходим по всем узлам в списке
	for curr != nil {
		// Преобразуем указатель в структуру Node
		node := (*Node)(curr)
		// Выводим значение узла
		fmt.Println(node.value)
		// Загружаем указатель на следующий узел
		curr = atomic.LoadPointer(&node.next)
	}
}

func main() {
	// Создаем новый связанный список
	list := &List{}
	// Добавляем узлы со значениями 1, 2, 3 в начало списка
	list.Add(1)
	list.Add(2)
	list.Add(3)
	// Выводим значения узлов списка
	list.Print()

	var value1 int32 = 42
	var value2 int32 = 24

	AtomicSwap(&value1, &value2)

	if value1 != 24 || value2 != 42 {
		fmt.Errorf("Expected values after swap: 24, 42; got: %v, %v", value1, value2)
	} else {
		fmt.Println("Values after swap: 24, 42")
	}
}
