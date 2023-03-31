package tpl

import (
	"bytes"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/samber/lo"
	"github.com/xuender/go-cli/utils"
)

type Env struct {
	Year    string
	User    string
	Package string
	Name    string
	Path    string
	Test    string
	License string
}

func NewEnv() *Env {
	currentUser := lo.Must1(user.Current())

	return &Env{
		Year: time.Now().Format("2006"),
		User: currentUser.Username,
	}
}

func NewEnvByGo(arg string) *Env {
	return NewEnvByFile(arg, ".go")
}

func NewEnvByFile(arg, ext string) *Env {
	env := NewEnv()
	pkg := filepath.Base(lo.Must1(os.Getwd()))
	dir := filepath.Dir(arg)
	name := filepath.Base(arg)

	if dir != "" && dir != "." {
		pkg = filepath.Base(dir)
	}

	path := filepath.Join(dir, utils.SnakeCase(name))

	if !strings.HasSuffix(path, ext) {
		path += ext
	}

	env.Package = pkg
	env.Name = name
	env.Path = path

	return env
}

func NewEnvByDir(arg string) *Env {
	env := NewEnv()
	pkg := filepath.Base(lo.Must1(os.Getwd()))
	dir := arg
	name := filepath.Base(arg)

	if dir != "" && dir != "." {
		pkg = dir
	}

	if data, err := os.ReadFile(filepath.Join(dir, ".git", "config")); err == nil {
		pkg = GitURL(data)
		name = pkg[strings.LastIndex(pkg, "/")+1:]
	}

	if name == "." {
		name = filepath.Base(pkg)
	}

	env.Package = pkg
	env.Name = name
	env.Path = arg

	return env
}

func (p *Env) Bytes(files fs.FS, path string) []byte {
	buf := &bytes.Buffer{}
	funcs := template.FuncMap{"dir": filepath.Dir}
	tmpl := lo.Must1(template.New("text").Funcs(funcs).ParseFS(files, path))

	lo.Must0(tmpl.ExecuteTemplate(buf, filepath.Base(path), p))

	return buf.Bytes()
}

func GitURL(data []byte) string {
	reg := regexp.MustCompile(`[A-Za-z0-9_\-./:]+\.git`)
	ret := string(reg.Find(data))
	ret = strings.TrimSuffix(ret, ".git")
	ret = strings.TrimPrefix(ret, "https://")
	ret = strings.TrimPrefix(ret, "http://")

	return strings.Replace(ret, ":", "/", 1)
}
