package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	InitDB()

	if len(os.Args) < 2 {
		fmt.Println("Uso: todo [add|list|done|delete] [args]")
		return
	}

	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		addCmd.Parse(os.Args[2:])
		desc := addCmd.Arg(0)
		if desc == "" {
			fmt.Println("Debes proporcionar una descripción.")
			return
		}
		if err := AddTask(desc); err != nil {
			fmt.Println("Error al agregar tarea:", err)
		} else {
			fmt.Println("✅ Tarea agregada.")
		}

	case "list":
		tasks, err := ListTasks()
		if err != nil {
			fmt.Println("Error al listar tareas:", err)
			return
		}
		for _, t := range tasks {
			status := " "
			if t.Done {
				status = "✔"
			}
			fmt.Printf("[%s] %d: %s (creado: %s)\n", status, t.ID, t.Description, t.CreatedAt)
		}

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Debes proporcionar el ID de la tarea.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if err := MarkDone(id); err != nil {
			fmt.Println("Error al marcar como hecha:", err)
		} else {
			fmt.Println("✅ Tarea completada.")
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Debes proporcionar el ID de la tarea.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if err := DeleteTask(id); err != nil {
			fmt.Println("Error al eliminar tarea:", err)
		} else {
			fmt.Println("🗑️ Tarea eliminada.")
		}

	default:
		fmt.Println("Comando no reconocido. Usa: add, list, done, delete.")
	}
}
