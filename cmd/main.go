package main

import (
	"fmt"

	"github.com/mms-gianni/GitGenie/pkg/genie"
)

var (
	// Version is the version of the application
	Version string
	// BuildDate is the date the application was built
	BuildDate string
)

func main() {
	genie.InitClient()
	diff := genie.Diff()
	//genie.Status()

	commitMessages := genie.SubmitToApi(diff)
	commitMessage := genie.SelectCommitMessage(commitMessages)
	commitMessage = genie.EditCommitMessage(commitMessage)

	//genie.Commit(commitMessage)
	fmt.Println(commitMessage)

}
