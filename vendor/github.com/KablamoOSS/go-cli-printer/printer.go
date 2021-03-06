package printer

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/ttacon/chalk"
	"golang.org/x/crypto/ssh/terminal"
)

var instantiated *spinner.Spinner
var once sync.Once

// Set to true when testing, to never call os.Exit
var testing bool

var verbose bool
var color string
var spinnerStyle int

var previousProgressMessage string
var previousStepMessage string
var previousSubStepMessage string
var writer io.Writer

// Is the output a full terminal
var outputTTY bool

// Default way of terminating (from Fatal) is to os.Exit, but allow this to be
// changed for the purposes of testing.
var Terminate = ExitTerminate

func ExitTerminate(err error) {
	os.Exit(1)
}

func PanicTerminate(err error) {
	panic(err)
}

func init() {
	verbose = false
	color = "yellow"
	spinnerStyle = 14
	writer = os.Stdout
	testing = false

	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		outputTTY = true
	} else {
		outputTTY = false
	}
}

// Init the spinner with a verbose flag, and color.
// This is optional, and allows some customisation over the printer.
// You should do invoke this as early as posisble, before the first printer function
// is called.
func Init(initVerbose bool, initColor string, initSpinner int, writer io.Writer) {
	verbose = initVerbose
	color = initColor
	spinnerStyle = initSpinner
	getPrinter()
}

// Test to set testing to true, to prevent exiting on Fatal errors
func Test() {
	Terminate = PanicTerminate
}

// SetOutput to an io.Writer compatitble interface. Designed
// to capture error messages for use in testing
func SetOutput(out io.Writer) {
	s := getPrinter()
	s.Writer = out
}

// Create a singleton to the Spinner
// This ensures we only output to one line
func getPrinter() *spinner.Spinner {
	if instantiated == nil {
		once.Do(func() {
			instantiated = spinner.New(spinner.CharSets[spinnerStyle], 100*time.Millisecond)
			instantiated.Writer = writer
		})
	}
	return instantiated
}

// Progress message with a spinner
func Progress(message string) {
	if outputTTY {
		spinner := getPrinter()
		if message != previousProgressMessage {
			previousProgressMessage = message
			spinner.Suffix = fmt.Sprintf("  %s", message)
			spinner.Color(color)
		}
		spinner.Start()
	}
}

// Progress with formatted message
func Progressf(message string, args ...interface{}) {
	Progress(
		fmt.Sprintf(message, args),
	)
}

// Step prints a line console and stops the spinner
func Step(message string) {
	if outputTTY {
		if message != previousStepMessage {
			previousStepMessage = message
			spinner := getPrinter()

			spinner.Stop()
			fmt.Println(fmt.Sprintf("%s  %s", chalk.Yellow.Color("➜"), chalk.Bold.TextStyle(message)))
		}
	} else {
		log.Print(message)
	}
}

// Step with formatted message
func Stepf(message string, args ...interface{}) {
	Step(
		fmt.Sprintf(message, args),
	)
}

// SubStep prints a line console, at a given indent and stops the spinner
// ignoreVerboseRule true, ensures the SubStep always prints
func SubStep(message string, indent int, last, ignoreVerboseRule bool) {
	if outputTTY {
		if message != previousSubStepMessage {
			previousSubStepMessage = message
			// Substeps are only printed if the verbose flag is set at init
			// Unless it's the last substep
			if verbose || ignoreVerboseRule || last {
				var indentString string

				for i := 1; i <= indent; i++ {
					indentString = fmt.Sprintf("   %s", indentString)
				}

				icon := "├─"

				if last {
					icon = "└─"
				}
				spinner := getPrinter()

				spinner.Stop()
				fmt.Println(fmt.Sprintf("%s%s %s", chalk.Dim.TextStyle(indentString), chalk.Dim.TextStyle(icon), chalk.Dim.TextStyle(message)))
			}
		}
	} else {
		log.Print(message)
	}
}

// SubStep with formatted message
// FIXME: The extra arguments for SubStep being between the format string and
// it's formatting arguments is rather awkward
func SubStepf(message string, indent int, last, ignoreVerboseRule bool, args ...interface{}) {
	SubStep(
		fmt.Sprintf(message, args),
		indent,
		last,
		ignoreVerboseRule,
	)
}

// Finish prints message to the console and stops the spinner with success.
// This is best used to indicated the end of a task
func Finish(message string) {
	if outputTTY {
		spinner := getPrinter()
		spinner.Stop()
		fmt.Println(fmt.Sprintf("%s  %s", chalk.Green.Color("✔"), chalk.Bold.TextStyle(message)))
	} else {
		log.Printf("✔ %s", message)
	}
}

// Finish with formatted message
func Finishf(message string, args ...interface{}) {
	Finish(
		fmt.Sprintf(message, args),
	)
}

// Warn prints a warning to the screen. It's formatted like other errors, and coloured yellow.
func Warn(err error, resolution string, link string) {
	if outputTTY {
		spinner := getPrinter()

		spinner.Stop()
		errMessage := fmt.Sprintf(
			"%s %s",
			chalk.Bold.TextStyle(chalk.Yellow.Color("!  Error:")),
			chalk.Yellow.Color(err.Error()),
		)
		if resolution != "" {
			errMessage = fmt.Sprintf(
				"%s\n%s%s",
				errMessage,
				chalk.Dim.TextStyle(chalk.Bold.TextStyle("☞  Resolution: ")),
				chalk.Dim.TextStyle(resolution),
			)
		}

		if link != "" {
			errMessage = fmt.Sprintf(
				"%s\n%s%s",
				errMessage,
				chalk.Dim.TextStyle(chalk.Bold.TextStyle("∞  More info: ")),
				chalk.Italic.TextStyle(link),
			)
		}

		fmt.Println(errMessage)

	} else {
		log.Printf("WARN: %s", err.Error())
	}
}

// Error prints an error to the screen. As it's intended reader is a user of your program,
// it expects both the error message, a way for the reader to resolve the error, and if
// possible a link to futher information.
// If the error doesn't have a link, pass a blank string ""
func Error(err error, resolution string, link string) {
	if outputTTY {
		spinner := getPrinter()

		spinner.Stop()
		errMessage := fmt.Sprintf(
			"%s %s",
			chalk.Bold.TextStyle(chalk.Red.Color("✖  Error:")),
			chalk.Red.Color(err.Error()),
		)
		if resolution != "" {
			errMessage = fmt.Sprintf(
				"%s\n%s%s",
				errMessage,
				chalk.Dim.TextStyle(chalk.Bold.TextStyle("☞  Resolution: ")),
				chalk.Dim.TextStyle(resolution),
			)
		}

		if link != "" {
			errMessage = fmt.Sprintf(
				"%s\n%s%s",
				errMessage,
				chalk.Dim.TextStyle(chalk.Bold.TextStyle("∞  More info: ")),
				chalk.Italic.TextStyle(link),
			)
		}

		fmt.Println(errMessage)

	} else {
		log.Printf("ERROR: %s", err.Error())
	}
}

// Fatal prints an error in the exact same way as Error, except it prefixes with "Fatal",
// and the function ends with a panic
func Fatal(err error, resolution string, link string) {
	if outputTTY {
		spinner := getPrinter()

		errMessage := fmt.Sprintf(
			"%s %s",
			chalk.Bold.TextStyle(chalk.Red.Color("✖  Fatal:")),
			chalk.Red.Color(err.Error()),
		)

		if resolution != "" {
			errMessage = fmt.Sprintf(
				"%s\n%s%s",
				errMessage,
				chalk.Dim.TextStyle(chalk.Bold.TextStyle("☞  Resolution: ")),
				chalk.Dim.TextStyle(resolution),
			)
		}
		// Add the link if a valid one was supplied
		if link != "" {
			errMessage = fmt.Sprintf(
				"%s\n%s%s",
				errMessage,
				chalk.Dim.TextStyle(chalk.Bold.TextStyle("∞  More info: ")),
				chalk.Italic.TextStyle(link),
			)
		}

		spinner.Stop()
		fmt.Println(errMessage)

	} else {
		log.Printf("FATAL: %s", err.Error())
	}

	Terminate(err)
}

// Stop the spinner
func Stop() {
	spinner := getPrinter()
	spinner.Stop()
}
