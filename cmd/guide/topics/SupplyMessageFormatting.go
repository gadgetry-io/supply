package topics

import (
	"fmt"

	"github.com/gadgetry-io/supply/cmd/common"
)

// SupplyMessageFormatting - Topic on How to Format Messages During Supply CLI Development
func SupplyMessageFormatting() {

	common.FmtMessage("h3-blue", "SUPPLY GUIDE")
	common.FmtMessage("h3-blue", "Topic: Formatting Messages in SUPPLY CLI")
	common.FmtMessage("text-blue", "This guide covers how to format messages in your supply packages")
	common.FmtMessage("text-blue", "We have 3 types of messages\n -> Fmt Messages\n -> Log Messages\n -> Debug Messages\n")

	common.FmtMessage("h3", "common.FmtMessage")

	common.FmtMessage("h3", "common.LogMessage")

	common.FmtMessage("h3", "common.DebugMessage")

	fmt.Printf("\n\ncommon.DebugMessage\n")
	common.DebugMessage("[package: schematic] func Run()")

	fmt.Printf("\n\ncommon.LogMessage\n")
	common.LogMessage("info", "Example info log message")
	common.LogMessage("pass", "Example pass log message")
	common.LogMessage("warn", "Example warn log message")
	common.LogMessage("fail", "Example fail log message")
	common.LogMessage("error", "Example error log message\n")
	common.LogMessage("default", "Example default log message")
	common.LogMessage("success", "Example success log message")
	common.LogMessage("warning", "Example warning log message")
	common.LogMessage("failure", "Example failure log message\n")

	//fmt.Printf("Use 'supply schematic --help' for more information\n\n")
	fmt.Printf("\n\ncommon.FmtMessage Formatted Text\n")
	common.FmtMessage("h1-white", "Example h1 Header")
	common.FmtMessage("h2-white", "Example h2 Header")
	common.FmtMessage("h3-white", "Example h3 Header")
	common.FmtMessage("text-white", "Example Text")

	fmt.Printf("\n\ncommon.FmtMessage Colors\n")
	common.FmtMessage("text-white", "Example White Text")
	common.FmtMessage("text-blue", "Example Blue Text")
	common.FmtMessage("text-yellow", "Example Yellow Text")
	common.FmtMessage("text-green", "Example Green Text")
	common.FmtMessage("text-red", "Example Red Text")

	fmt.Printf("\n\ncommon.FmtMessage Upper/Lower Case\n")
	common.FmtMessage("upper-white", "Example upper-white")
	common.FmtMessage("lower-white", "Example lower-white")

	fmt.Printf("\n\n")
}
