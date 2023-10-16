package main

const (
	todo status = iota
	inprogress
	done
)

type Task struct {
	status status
	title  string
	desc   string
}

func NewTask() *Task {
	return &Task{
		status: todo,
	}
}

func (t Task) FilterValue() string {
	return t.title
}

func main() {

}
