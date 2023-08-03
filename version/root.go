package version

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
	"github.com/xuender/kit/types"
	"github.com/youthlin/t"
)

func NewCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Update version")
	cmd.Long = t.T("Update version and git tag, push.")
	cmd.Example = t.T("  # Veersion update\n  go-cli version")

	cmd.Flags().StringP("ver", "v", "", t.T("Specify the current version."))
	cmd.Run = run

	return cmd
}

func run(cmd *cobra.Command, _ []string) {
	ver, _ := cmd.Flags().GetString("ver")

	name := getHistory()
	if name == "" {
		name = "History.md"
	}

	if ver == "" {
		data := lo.Must1(os.ReadFile(name))
		ver = GetVer(data)
		logs.D.Println("read", name, ver)
	}

	ver = IncVer(ver)
	command := exec.Command("git", "changelog", "-t", ver, "-x")
	data := lo.Must1(command.Output())

	logs.I.Println("git", "changelog", "-t", ver, "-x")
	lo.Must0(os.WriteFile(name, data, oss.DefaultFileMode))

	ver = "v" + ver

	logs.D.Println("update version", ver)
	msg := "chore: " + ver

	lo.Must0(exec.Command("git", "commit", "-am", msg).Run())
	logs.D.Println("git commit")

	lo.Must0(exec.Command("git", "tag", ver).Run())
	logs.D.Println("git tag")

	lo.Must0(exec.Command("git", "push").Run())
	logs.D.Println("git push")

	lo.Must0(exec.Command("git", "push", "--tag").Run())
	logs.D.Println("git push --tag")
}

// nolint: gochecknoglobals
var _mds = [...]string{
	"CHANGELOG.md",
	"History.md",
}

func IncVer(ver string) string {
	last := strings.LastIndex(ver, ".")
	num := lo.Must1(types.ParseInteger[int](ver[last+1:]))
	num++

	return fmt.Sprintf("%s.%d", ver[:last], num)
}

func SetVer(data []byte, ver string) []byte {
	return bytes.Replace(data, []byte("x.x.x / "), []byte(ver+" / "), 1)
}

func GetVer(data []byte) string {
	reg := regexp.MustCompile(`([vV])?(\d+\.\d+\.\d+) / `)
	ver := string(reg.Find(data))

	if ver == "" {
		return "0.0.1"
	}

	ver = strings.ToLower(ver)

	start := 0
	end := strings.Index(ver, " ")

	index := strings.Index(ver, "v")
	if index > -1 {
		start++
	}

	return ver[start:end]
}

func getHistory() string {
	for _, entry := range lo.Must1(os.ReadDir(lo.Must1(oss.Abs(".")))) {
		for _, name := range _mds {
			if name == entry.Name() {
				return name
			}
		}
	}

	return ""
}
