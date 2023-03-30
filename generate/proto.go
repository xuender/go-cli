package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/tpl"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
	"github.com/youthlin/t"
)

func protoCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringP(_type, "t", "message", t.T("message or enum"))
	cmd.Short = t.T("generate protobuf")
	cmd.Long = t.T("generate protobuf")
	// nolint: lll
	cmd.Example = t.T("  # create message\n  go-cli g p pb/Book\n  # create enum\n  go-cli g p BookType -t enum -o pb/book.proto")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			lo.Must0(cmd.Usage())

			return
		}

		output, _ := cmd.Flags().GetString(_output)

		for _, arg := range args {
			env := tpl.NewEnvByFile(arg, ".proto")

			if output != "" {
				env.Path = output
			}

			createProto(env, strings.ToLower(lo.Must1(cmd.Flags().GetString(_type))))
		}
	}

	return cmd
}

func createProto(env *tpl.Env, typeCode string) {
	logs.D.Println(t.T("create Proto: %s", env.Name))

	var file *os.File
	defer file.Close()

	if oss.Exist(env.Path) {
		file = AppendFile(env.Path)
	} else {
		file = CreateFile(env.Path)

		lo.Must1(file.Write(env.Bytes(_static, filepath.Join(_staticPath, "proto.tpl"))))
	}

	lo.Must1(file.Write(env.Bytes(_static, filepath.Join(_staticPath, fmt.Sprintf("proto_%s.tpl", typeCode)))))
}
