package font

import (
	"os"
	"os/exec"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/xuender/kit/logs"
	"github.com/youthlin/t"
)

func subsetCmd(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringP("text", "t", "", t.T("Specify characters to include in the subset, as UTF-8 string."))
	cmd.Flags().StringP(
		"text-file",
		"", "",
		t.T("Like --text but reads from a file. Newline character are not added to the subset."),
	)
	cmd.Flags().StringP(
		"output-file",
		"", "",
		t.T("The output font file. If not specified, the subsetted font will be saved in as font-file.subset."),
	)

	cmd.Short = t.T("Subset font")
	cmd.Long = t.T("OpenType font subsetter and optimizer.")
	cmd.Example = t.T("  # Subset font\n  go-cli font subset font.ttf --text-file=zh_CN.po --output-file=new.ttf")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			lo.Must0(cmd.Usage())

			return
		}

		var (
			text, _       = cmd.Flags().GetString("text")
			textFile, _   = cmd.Flags().GetString("text-file")
			outputFile, _ = cmd.Flags().GetString("output-file")
			all, _        = cmd.Flags().GetBool("all")
		)

		if outputFile == "" || (text == "" && textFile == "") {
			lo.Must0(cmd.Usage())

			return
		}

		if text != "" {
			text = strings.ReplaceAll(text, "'", "")
			logs.I.Printf("fonttools subset %s --text='%s' --output-file=%s", args[0], text, outputFile)

			Exec("fonttools", "subset", args[0], "--text='"+text+"'", "--output-file="+outputFile)

			return
		}

		file := lo.Must1(os.CreateTemp(os.TempDir(), "go-cli."))
		defer func() {
			file.Close()
			os.Remove(file.Name())
		}()

		lo.Must1(file.WriteString(load(all, textFile)))

		logs.I.Printf("fonttools subset %s --text-file=%s --output-file=%s", args[0], file.Name(), outputFile)
		Exec("fonttools", "subset", args[0], "--text-file="+file.Name(), "--output-file="+outputFile)
	}

	return cmd
}

func Exec(name string, args ...string) {
	command := exec.Command(name, args...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	lo.Must0(command.Run())
}
