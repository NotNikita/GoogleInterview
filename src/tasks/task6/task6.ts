interface Todo {
  title: string;
  description: string;
  completed: boolean;
}

type TodoPicked = MyPick<Todo, 'title' | 'completed'>;

const todo: TodoPicked = {
  title: 'Clean room',
  completed: true,
};

type MyPick<T, Keys extends keyof T> = {[k in Keys]: T[k]};

type MyOmit<T, Keys extends keyof T> = {[k in Keys as k extends Keys ? never : k]: T[k]};

type TodoOmitted = MyOmit<Todo, 'completed' | 'description'>;
const todo2: TodoOmitted = {
  title: 'Do homework',
};
