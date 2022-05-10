package main

import (
    "fmt"
    "flag"
    "os"
    "strings"
    "io/ioutil"
)

var dir = flag.String("d", "./template/init.cpp", "директория шаблона.")
var ext = flag.String("e", "cpp", "расширение файлов.")

type Env struct {
    DirTo, DirFrom, Ext, Content string
    Names []string
}

const (
    usage = "Usage: init <dir> [A..Z]"
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

    env.DirFrom, env.Ext = *dir, *ext
    env.DirTo, env.Names = flag.Args()[0], flag.Args()[1:]
    for strings.HasSuffix(env.DirTo, "/") && len(env.DirTo) > 0 {
        env.DirTo = env.DirTo[:len(env.DirTo)-1]
    }

    cont, err := ioutil.ReadFile(env.DirFrom)
    if err != nil {
        fmt.Fprintf(os.Stderr, "parseEnv: failed to read a template file.\n")
        fmt.Fprintf(os.Stderr, "err: %v\n", err)
        os.Exit(2)
    }

    env.Content = string(cont)
    return env
}

func mkenv(env *Env) {
    err := os.MkdirAll(env.DirTo, 0750)
    if err != nil {
        fmt.Fprintf(os.Stderr, "mkenv: failed to create directory.\n")
        fmt.Fprintf(os.Stderr, "err: %v\n", err)
        os.Exit(2)
    }

    for _, file := range env.Names {
        path := env.DirTo + "/" + file + "." + *ext
        desc, err := os.Create(path)
        if err != nil {
            fmt.Fprintf(os.Stderr, "mkenv: failed to create file \"%s\"\n", path)
            fmt.Fprintf(os.Stderr, "err: %v\n", err)
            continue
        }

        _, err = desc.Write([]byte(env.Content))
        if err != nil {
            fmt.Fprintf(os.Stderr, "mkenv: failed to write to destination file.\n")
            fmt.Fprintf(os.Stderr, "err: %v\n", err)
        }
        desc.Close()
    }
}

