package main

import "fmt"

func main() {
	ranges()
	shadowingVariables()
	typeEmbedding()
	valueReceivers()
}

// ranges
func ranges() {
	valuesStr := []string{"a", "b", "c"}
	for index, value := range valuesStr {
		fmt.Println(index, value)
	}

	valuesInt := []int{4, 8, 15, 16, 23, 42}
	for value := range valuesInt { // получаем только индекс, а не элемент слайса
		fmt.Println(value)
	}
}

// shadowing variables
func shadowingVariables() {
	node := Node{
		value:  "original",
		weight: 1,
	}

	t1 := func(n Node) (Node, error) {
		return Node{
			value:  n.value + " t1",
			weight: n.weight + 1,
		}, nil
	}

	t2 := func(n Node) (Node, error) {
		return Node{
			value:  n.value + " t2",
			weight: n.weight + 2,
		}, nil
	}

	transformed, err := applyTansformations(node, []Transformation{t1, t2})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", transformed) // [{original t1 2} {original t2 3}]
	}

	transformed2, err := applyTansformationsValid(node, []Transformation{t1, t2})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", transformed2) // [{original t1 2} {original t1 t2 4}]
	}
}

type Node struct {
	value  string
	weight int
}

type Transformation func(Node) (Node, error)

func applyTansformations(n Node, ts []Transformation) ([]Node, error) {
	var steps []Node
	for _, t := range ts {
		n, err := t(n)
		if err != nil {
			return nil, err
		}
		steps = append(steps, n)
	}
	return steps, nil
}

func applyTansformationsValid(n Node, ts []Transformation) ([]Node, error) {
	var steps []Node
	for _, t := range ts {
		newNode, err := t(n)
		n = newNode
		if err != nil {
			return nil, err
		}
		steps = append(steps, n)
	}
	return steps, nil
}

// Встраиваемый тип.
// Структура Wrapped и Node имеют одинаковые методы Name,
// тут нужно вызывать явно указывая полный путь к методы через вызовы встриваемых полей,
// иначе можно получить неожидаемый вызов метода.
func typeEmbedding() {
	v1 := Wrapper{Node{"123", 0}}
	fmt.Println(v1.Name(), v1.Node.Name())
	fmt.Println(v1.value, v1.Node.value)
	fmt.Println(v1.Name())
	fmt.Println(v1.Node.Name())
}

func (n Node) Name() string {
	return n.value
}

type Wrapper struct { // когда структура встраивает другую структуру, она наследует все методы и поля встраиваемого типа
	Node
}

func (w Wrapper) Name() string {
	return "wrapped value: " + w.Node.value
}

// value receivers
func valueReceivers() {
	// func (n Node) Name() string {...} // Value receiver
	// func (n *Node) Name() string {...} // Pointer receiver

	n := Node{value: "my value"}
	n.SetValue("new value")
	fmt.Println(n)
}

func (n *Node) SetValue(value string) {
	n.value = value
}
