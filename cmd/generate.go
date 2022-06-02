/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/utils"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/logs"
	"github.com/xuender/oils/oss"
)

// nolint
func init() {
	root := getRoot()
	generateCmd := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"g"},
		Short:   Printer.Sprintf("generate short"),
		Long:    Printer.Sprintf("generate long"),
		Example: "  go-cli g\n  go-cli g cmd\n  go-cli g service pkg/service",
		Run: func(cmd *cobra.Command, args []string) {
			typeStr := ""
			name := ""
			mod := ""
			size := len(args)

			if size == 0 {
				typeStr = selectType()
			}

			if size > 0 {
				typeStr = strings.ToLower(args[0])
			}

			if size < 1 {
				name = promptName()
			}

			if size > 1 {
				name = args[1]
			}

			if size > base.Two {
				mod = args[base.Two]
			}

			logs.Debugw("g", "type", typeStr, "name", name)

			switch typeStr {
			case "cmd", "c":
				createCmd(name)
			case "service", "s":
				createService(name)
			case "enum", "e":
				createEnum(name, mod)
			case "proto", "protobuf", "p":
				createProto(name, mod)
			}
		},
	}

	root.AddCommand(generateCmd)
}

func createProto(name, mod string) {
	if mod == "" {
		prompt := promptui.Select{
			Label:     Printer.Sprintf("generate select proto"),
			Items:     []string{"message", "enum"},
			Templates: NewSelectTemplates(),
		}
		index, _ := base.Must2(prompt.Run())

		if index > 0 {
			mod = "enum"
		}
	}

	if mod == "enum" {
		createFile(name, "proto_enum.tpl", ".proto")

		return
	}

	createFile(name, "proto_message.tpl", ".proto")
}

func createEnum(name, mod string) {
	if mod == "" {
		prompt := promptui.Select{
			Label:     Printer.Sprintf("generate select enum"),
			Items:     []string{Printer.Sprintf("generate increment"), Printer.Sprintf("generate allergen")},
			Templates: NewSelectTemplates(),
		}
		index, _ := base.Must2(prompt.Run())

		if index > 0 {
			mod = "allergen"
		}
	}

	if mod == "allergen" {
		createFile(name, "enum_allergen.tpl", ".go")

		return
	}

	createFile(name, "enum_increment.tpl", ".go")
}

func createService(name string) {
	createFile(name, "service.tpl", ".go")
}

func createFile(name, tpl, ext string) {
	dir := filepath.Dir(name)
	baseName := filepath.Base(name)
	file := filepath.Join(dir, utils.FileName(baseName)+ext)

	if oss.Exist(file) {
		return
	}

	env := NewEnv(".")
	env.Path = dir

	if dir == "." {
		if index := strings.LastIndex(env.Package, "/"); index > 0 {
			env.Package = env.Package[index+1:]
			if index := strings.LastIndex(env.Package, "-"); index > 0 {
				env.Package = env.Package[index+1:]
			}
		}
	} else {
		env.Package = filepath.Base(dir)
	}

	env.Name = utils.TypeName(baseName)

	_ = os.MkdirAll(dir, oss.DefaultDirFileMod)
	tmpl := base.Must1(template.ParseFS(static, filepath.Join(staticName, "generate", tpl)))
	buffer := &bytes.Buffer{}

	base.Must(tmpl.Execute(buffer, env))
	base.Must(os.WriteFile(file, buffer.Bytes(), oss.DefaultFileMode))
}

func createCmd(name string) {
	file, err := utils.Parse("cmd", "root.go")
	useCobra := err == nil && utils.UseCobra(file)

	logs.Debugw("cmd", "name", name)

	if useCobra {
		if oss.Exist("cmd", name+".go") {
			panic(Printer.Sprintf("cmd %s exist", name))
		}

		base.Must(oss.Exec("cobra-cli", "add", name))

		return
	}

	if oss.Exist("cmd", name) {
		panic(Printer.Sprintf("cmd %s exist", name))
	}

	base.Must(os.MkdirAll(filepath.Join("cmd", name), oss.DefaultDirFileMod))

	tmpl := base.Must1(template.ParseFS(static, filepath.Join(staticName, "generate", "cmd.tpl")))
	buffer := &bytes.Buffer{}
	env := &Env{Name: name}

	base.Must(tmpl.Execute(buffer, env))
	base.Must(os.WriteFile(filepath.Join("cmd", name, "main.go"), buffer.Bytes(), oss.DefaultFileMode))
}

func selectType() string {
	prompt := promptui.Select{
		Label:     Printer.Sprintf("generate select type"),
		Items:     []string{"cmd", "service", "enum", "protobuf"},
		Templates: NewSelectTemplates(),
	}
	_, res := base.Must2(prompt.Run())
	res = strings.ToLower(res)

	return res
}

func promptName() string {
	prompt := promptui.Prompt{
		Label:    Printer.Sprintf("generate prompt name"),
		Validate: validate,
	}

	return base.Must1(prompt.Run())
}

func validate(input string) error {
	if input == "" {
		return ErrEmpty
	}

	return nil
}
