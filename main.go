package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	dir = flag.String("d", "./template/init.cpp", "директория шаблона.")
	ext = flag.String("e", "", "расширение файлов.")
)

type Env struct {
	DirTo, DirFrom, Ext, Content string
	Names                        []string
}

const (
	usage = "Usage: ginit <dir> [A..Z]"
)

func main() {
	env := parseEnv()
	mkenv(&env)
}

func parseEnv() Env {
	var env Env
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println(usage)
		flag.PrintDefaults()
		os.Exit(2)
	}
	if *ext == "" {
		env.Ext = parseExt(*dir)
	} else {
		env.Ext = *ext
	}
	env.DirFrom = *dir
	env.DirTo, env.Names = flag.Args()[0], flag.Args()[1:]
	for strings.HasSuffix(env.DirTo, "/") && len(env.DirTo) > 0 {
		env.DirTo = env.DirTo[:len(env.DirTo)-1]
	}
	cont, err := ioutil.ReadFile(env.DirFrom)
	if err != nil {
		log.Println("failed to read a template file.")
		log.Fatal(err)
	}
	env.Content = string(cont)
	return env
}

func parseExt(name string) string {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			return name[i+1:]
		}
	}
	return "cpp"
}

func mkenv(env *Env) {
	err := os.MkdirAll(env.DirTo, 0750)
	if err != nil {
		log.Println("failed to create directory.")
		log.Fatal(err)
	}
	for _, file := range env.Names {
		path := env.DirTo + "/" + file + "." + env.Ext
		desc, err := os.Create(path)
		if err != nil {
			log.Println("failed to create file", fmt.Sprintf("\"%s\"", path))
			log.Println(err)
			continue
		}
		_, err = desc.Write([]byte(env.Content))
		if err != nil {
			log.Println("failed to write to destination file.")
			log.Println(err)
		}
		desc.Close()
	}
}
