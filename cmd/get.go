package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/abhinav812/gopher-greeter/internal/asciiconvertor"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get the desired Gopher png.",
	Long:  `This get command will call GitHub repository in order to return the desired Gopher...`,
	Run: func(cmd *cobra.Command, args []string) {
		var gopherName = "dr-who"
		var pngFolder = "gophers"

		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
		}

		url := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"
		fmt.Println("Try to get '" + gopherName + "' Gopher...")

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic(err)
			}
		}(resp.Body)

		if resp.StatusCode == http.StatusOK {
			// Create gophers dir
			err := os.MkdirAll(filepath.Join(".", pngFolder), os.ModePerm)
			if err != nil {
				fmt.Println(err)
			}

			// Create the file
			out, err := os.Create(filepath.Join(pngFolder, filepath.Base(gopherName+".png")))
			if err != nil {
				fmt.Println(err)
			}
			defer func(out *os.File) {
				err := out.Close()
				if err != nil {
					panic(err)
				}
			}(out)

			// Write the body to the file
			_, err = io.Copy(out, resp.Body)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Perfect! Just saved in " + out.Name() + "!")

			// convert the png into ascii text
			asciiText := asciiconvertor.ConvertToAsii(out.Name())
			fmt.Println(asciiText)

			writeToFile(asciiText, out.Name())

		} else {
			fmt.Println("Error: " + gopherName + " does not exists!!")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func writeToFile(asciiText string, fileName string) {
	// Create the file
	out, err := os.Create(filepath.Join("gophers", filepath.Base(stripExtension(fileName)+".text")))
	if err != nil {
		fmt.Println(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			panic(err)
		}
	}(out)

	// Write the body to the file
	//_, err = io.Copy(out, asciiText)
	_, err = out.WriteString(asciiText)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Perfect! Just saved in " + out.Name() + "!")
}

func stripExtension(fileName string) string {
	extension := filepath.Ext(fileName)
	return fileName[0 : len(fileName)-len(extension)]
}
