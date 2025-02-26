package main

import (
	"fmt"
	"gator/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	// Читаем текущее состояние из конфига
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Записываем данные из конфига в структуру
	programState := &state{
		cfg: &cfg,
	}

	// Создаём экземпляр структуры с командами и добавляем их туда
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	// Получаем аргументы
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	// Передаём в run текущее состояние и структуру с аргументами
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		fmt.Println(err)
	}
}
