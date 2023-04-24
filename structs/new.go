package structs

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/go-cli/utils"
	"github.com/xuender/kit/base"
	"github.com/xuender/kit/logs"
	"github.com/youthlin/t"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func newStruct(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("New struct")
	cmd.Long = t.T("Create a new struct function by other struct.")
	cmd.Example = "  go-cli s n target/struct from/struct"
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) < base.Two {
			lo.Must0(cmd.Usage())

			return
		}

		sources := newSources(args)
		upper := cases.Title(language.English)

		for index := 1; index < len(sources); index++ {
			funcName := fmt.Sprintf("New%sBy%s", sources[0].Structs[0].Name, sources[index].Structs[0].Name)
			if sources[0].Structs[0].Name == sources[index].Structs[0].Name {
				funcName = fmt.Sprintf("New%sBy%s%s",
					sources[0].Structs[0].Name,
					upper.String(sources[index].Package),
					sources[index].Structs[0].Name,
				)
			}

			if lo.Contains(sources[0].Funcs, funcName) {
				continue
			}

			file := utils.AppendFile(args[0])
			defer file.Close()

			structName := sources[index].Structs[0].Name
			if sources[0].Package != sources[index].Package {
				structName = sources[index].Package + "." + structName
			}

			fields := lo.Intersect(sources[0].Structs[0].FieldNames, sources[index].Structs[0].FieldNames)
			if len(fields) == 0 {
				panic(t.T("no duplicate name field"))
			}

			max := lo.Max(lo.Map(fields, func(field string, _ int) int { return len(field) }))

			_, _ = file.WriteString(fmt.Sprintf("\n// %s creates a new %s by %s.\n",
				funcName, sources[0].Structs[0].Name,
				structName,
			))
			_, _ = file.WriteString(fmt.Sprintf("func %s(elem *%s) *%s {\n", funcName, structName, sources[0].Structs[0].Name))
			_, _ = file.WriteString(fmt.Sprintf("  return &%s{\n", sources[0].Structs[0].Name))

			for _, field := range fields {
				_, _ = file.WriteString(fmt.Sprintf("    %s:%s elem.%s,\n", field, strings.Repeat(" ", max-len(field)), field))
			}

			_, _ = file.WriteString("  }\n}\n")

			logs.I.Println(t.T("create function %s", funcName))
		}
	}

	return cmd
}

func newSources(args []string) []*utils.Source {
	sources := make([]*utils.Source, len(args))
	for index, arg := range args {
		sources[index] = lo.Must1(utils.NewSource(arg))

		if len(sources[index].Structs) != 1 {
			panic(t.T("Golang source file %s struct num is not one.", arg))
		}
	}

	return sources
}
