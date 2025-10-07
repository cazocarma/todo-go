package main

type Task struct {
	ID          int
	Description string
	Done        bool
	CreatedAt   string
}

func AddTask(description string) error {
	_, err := DB.Exec("INSERT INTO tasks (description) VALUES ($1)", description)
	return err
}

func ListTasks() ([]Task, error) {
	rows, err := DB.Query("SELECT id, description, done, created_at FROM tasks ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Description, &t.Done, &t.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func MarkDone(id int) error {
	_, err := DB.Exec("UPDATE tasks SET done = true WHERE id = $1", id)
	return err
}

func DeleteTask(id int) error {
	_, err := DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
