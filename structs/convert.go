package structs

import (
	"fmt"
	"os"
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

func convertStruct(cmd *cobra.Command) *cobra.Command {
	cmd.Short = t.T("Convert struct")
	cmd.Long = t.T("Convert struct to other structs.")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) < base.Two {
			lo.Must0(cmd.Usage())

			return
		}

		sources := newSources(args)
		upper := cases.Title(language.English)

		for index := 1; index < len(sources); index++ {
			fromFunc := fmt.Sprintf("%s_From%s", sources[0].Structs[0].Name, sources[index].Structs[0].Name)
			if sources[0].Structs[0].Name == sources[index].Structs[0].Name {
				fromFunc = fmt.Sprintf("%s_From%s%s", sources[0].Structs[0].Name,
					upper.String(sources[index].Package), sources[index].Structs[0].Name)
			}

			toFunc := fmt.Sprintf("%s_To%s", sources[0].Structs[0].Name, sources[index].Structs[0].Name)
			if sources[0].Structs[0].Name == sources[index].Structs[0].Name {
				toFunc = fmt.Sprintf("%s_To%s%s", sources[0].Structs[0].Name,
					upper.String(sources[index].Package), sources[index].Structs[0].Name)
			}

			if lo.Contains(sources[0].Funcs, fromFunc) && lo.Contains(sources[0].Funcs, toFunc) {
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

			createFromFunc(sources, fromFunc, file, structName, fields)
			createToFunc(fields, sources, toFunc, file, structName)
		}
	}

	return cmd
}

func createFromFunc(sources []*utils.Source, fromFunc string, file *os.File, structName string, fields []string) {
	if lo.Contains(sources[0].Funcs, fromFunc) {
		return
	}

	fromFunc = fromFunc[len(sources[0].Structs[0].Name)+1:]
	_, _ = file.WriteString(fmt.Sprintf("\n// %s from %s.\n", fromFunc, structName))
	_, _ = file.WriteString(fmt.Sprintf("func (p *%s) %s(elem *%s) *%s {\n",
		sources[0].Structs[0].Name, fromFunc, structName, sources[0].Structs[0].Name))

	for _, field := range fields {
		_, _ = file.WriteString(fmt.Sprintf("  p.%s = elem.%s\n", field, field))
	}

	_, _ = file.WriteString("\n  return p\n}\n")

	logs.I.Println(t.T("create function %s", fromFunc))
}

func createToFunc(fields []string, sources []*utils.Source, toFunc string, file *os.File, structName string) {
	if lo.Contains(sources[0].Funcs, toFunc) {
		return
	}

	max := lo.Max(lo.Map(fields, func(field string, _ int) int { return len(field) }))
	toFunc = toFunc[len(sources[0].Structs[0].Name)+1:]
	_, _ = file.WriteString(fmt.Sprintf("\n// %s to %s.\n", toFunc, structName))
	_, _ = file.WriteString(fmt.Sprintf("func (p *%s) %s() *%s {\n", sources[0].Structs[0].Name, toFunc, structName))
	_, _ = file.WriteString(fmt.Sprintf("  return &%s{\n", structName))

	for _, field := range fields {
		_, _ = file.WriteString(fmt.Sprintf("    %s:%s p.%s,\n", field, strings.Repeat(" ", max-len(field)), field))
	}

	_, _ = file.WriteString("  }\n}\n")

	logs.I.Println(t.T("create function %s", toFunc))
}
