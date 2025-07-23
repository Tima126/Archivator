package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"io"
	"mycobraapp/lib/vlc"
	"os"
	"path/filepath"
	"strings"
)

const packedExtension = "vlc"

var ErrorEmptyPath = errors.New("path to file is not specified")

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

func pack(_ *cobra.Command, args []string) {

	if len(args) == 0 || args[0] == "" {
		Handler(ErrorEmptyPath)

	}

	filepath := args[0]

	r, err := os.Open(filepath)

	if err != nil {
		Handler(err)
	}

	defer r.Close()

	data, err := io.ReadAll(r)

	if err != nil {
		Handler(err)
	}

	packed := vlc.Encode(string(data))

	err = os.WriteFile(packedFileName(filepath), []byte(packed), 0644)

	if err != nil {
		Handler(err)
	}
}

func packedFileName(path string) string {

	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension

}
func init() {
	packCmd.AddCommand(vlcCmd)

}
