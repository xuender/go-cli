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

	"github.com/manifoldco/promptui"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/utils"
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
	Path    string
}

func NewEnv(dir string) *Env {
	currentUser := lo.Must1(user.Current())
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
	license := ""
	root := getRoot()
	initCmd := &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
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

			if licenseStr, err := cmd.Flags().GetString("license"); err == nil && licenseStr != "" {
				license = licenseStr
			}

			dir := "."
			if len(args) > 0 {
				dir = args[0]
			}

			logs.Debugw("init", "dir", dir)
			env := NewEnv(dir)

			for _, entry := range lo.Must1(static.ReadDir(filepath.Join(staticName, "init"))) {
				file, data := readStatic(dir, entry.Name(), env)

				if oss.Exist(file) {
					continue
				}

				Printer.Printf("init file %s", file)
				lo.Must0(os.WriteFile(file, data, oss.DefaultFileMode))
			}

			selectLicense(dir, license, env)
		},
	}

	initCmd.Flags().StringP("license", "s", "", Printer.Sprintf("init license"))
	root.AddCommand(initCmd)
}

func selectLicense(dir, license string, env *Env) {
	file := filepath.Join(dir, "LICENSE")
	if oss.Exist(file) {
		return
	}

	if license == "" {
		prompt := promptui.Select{
			Label:     Printer.Sprintf("init select license"),
			Items:     []string{"MIT", "APACHE2", "BSD3"},
			Templates: NewSelectTemplates(),
		}
		_, license = lo.Must2(prompt.Run())
		license = strings.ToLower(license)
	}

	tmpl := lo.Must1(template.ParseFS(static, filepath.Join(staticName, "license", license)+".tpl"))
	buffer := &bytes.Buffer{}

	lo.Must0(tmpl.Execute(buffer, env))
	lo.Must0(os.WriteFile(file, buffer.Bytes(), oss.DefaultFileMode))
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

	file := lo.Must1(filepath.Abs(filepath.Join(dir, out)))

	if tpl {
		tmpl := lo.Must1(template.ParseFS(static, filepath.Join(staticName, "init", name)))
		buffer := &bytes.Buffer{}

		lo.Must0(tmpl.Execute(buffer, env))

		return file, buffer.Bytes()
	}

	reader := lo.Must1(static.Open(filepath.Join(staticName, "init", name)))
	defer reader.Close()

	return file, lo.Must1(io.ReadAll(reader))
}
