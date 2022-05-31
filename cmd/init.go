/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/utils"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/i18n"
	"github.com/xuender/oils/logs"
	"github.com/xuender/oils/oss"
)

//go:embed static
var static embed.FS

const staticName = "static"

type Env struct {
	Year    string
	User    string
	Package string
	Name    string
}

func NewEnv(dir string) *Env {
	currentUser := base.Must1(user.Current())
	git := ""
	name := ""

	if data, err := os.ReadFile(filepath.Join(dir, ".git", "config")); err == nil {
		git = utils.GitURL(data)
		name = git[strings.LastIndex(git, "/")+1:]
	}

	return &Env{
		Year:    time.Now().Format("2006"),
		User:    currentUser.Username,
		Package: git,
		Name:    name,
	}
}

// nolint
func init() {
	root := getRoot()
	initCmd := &cobra.Command{
		Use:     "init",
		Short:   Printer.Sprintf("init short"),
		Long:    Printer.Sprintf("init long"),
		Example: "  go-cli init\n  go-cli init path",
		Run: func(cmd *cobra.Command, args []string) {
			if debug, err := cmd.Flags().GetBool("debug"); err != nil || !debug {
				logs.SetInfoLevel()
			}

			if language, err := cmd.Flags().GetString("language"); err == nil && language != "" {
				Printer = i18n.GetPrinter(i18n.GetTag([]string{language}))
			}

			dir := "."
			if len(args) > 0 {
				dir = args[0]
			}

			logs.Debugw("init", "dir", dir)
			env := NewEnv(dir)

			for _, entry := range base.Must1(static.ReadDir(staticName)) {
				file, data := readStatic(dir, entry.Name(), env)

				if oss.Exist(file) {
					continue
				}

				Printer.Printf("init file %s", file)
				base.Must(os.WriteFile(file, data, oss.DefaultFileMode))
			}
		},
	}

	root.AddCommand(initCmd)
}

func readStatic(dir, name string, env *Env) (string, []byte) {
	out := name
	if strings.HasPrefix(name, "d_") {
		out = "." + name[2:]
	}

	tpl := false

	if strings.HasSuffix(out, ".tpl") {
		out = out[:len(out)-4]
		tpl = true
	}

	file := base.Must1(filepath.Abs(filepath.Join(dir, out)))

	if tpl {
		tmpl := base.Must1(template.ParseFS(static, staticName+"/"+name))
		buffer := &bytes.Buffer{}

		base.Must(tmpl.Execute(buffer, env))

		return file, buffer.Bytes()
	}

	reader := base.Must1(static.Open(filepath.Join(staticName, name)))
	defer reader.Close()

	return file, base.Must1(io.ReadAll(reader))
}
