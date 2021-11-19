package cmd

import (
	"fmt"
	"github.com/abhinav812/gopher-greeter/greeters"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "This command will ask a gopher to greet you.",
	Long:  `This command will greet you with greetings. A random asii-gopher will greet you`,
	Run: func(cmd *cobra.Command, args []string) {
		message := "Hello..."
		if len(args) == 0 {
			fmt.Printf("You have to give the greeting message... Using default greeting: %s\n", message)
		} else {
			message = strings.Join(args, " ")
		}
		printGreeting(message)
		// Generate a random integer depending on get the number of ascii files
		rand.Seed(time.Now().UnixNano())
		randInt := rand.Intn(getNbOfGopherFiles() - 1)

		// Display random gopher ASCII embed files
		fileData, err := greeters.EmbedGopherFiles().ReadFile("gopher" + strconv.Itoa(randInt) + ".txt")
		if err != nil {
			log.Fatal("Error during read gopher ascii file", err)
		}
		fmt.Println(string(fileData))
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// greetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// greetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	usage := "GopherGreet is inspired by Cowsay program.\n" +
		"GopherGreet allow you to display a message said by a cute random Gopher.\n\n" +
		"Usage:\n" +
		"   gopher-greeter greet MESSAGE\n\n" +
		"Example:\n" +
		"   gopher-greeter greet hello Gopher lovers!!"
	greetCmd.Flags().BoolP("help", "h", false, usage)
}

func getNbOfGopherFiles() int {
	files, err := greeters.EmbedGopherFiles().ReadDir(".")
	if err != nil {
		log.Fatal("Error during reading greeters folder", err)
	}

	nbOfFiles := 0
	for range files {
		nbOfFiles++
	}

	return nbOfFiles
}

func printGreeting(message string) {
	msgLen := len(message)
	line := " "
	for i := 0; i <= msgLen; i++ {
		line += "-"
	}
	fmt.Println(line)
	fmt.Println("< " + message + " >")
	fmt.Println(line)
	fmt.Println("        \\")
	fmt.Println("         \\")
}
