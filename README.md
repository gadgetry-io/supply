# supply
supply

## Introduction

**SUPPLY** is a command line (cli) tool developed by [Gadgetry.io](https://gadgetry.io)

It is a utility that our practitioners have designed to support our day-to-day work with the patterns 
we use to solve commonly occurring problems with Cloud Architecture, Infrastructure as Code, and DevOps Automation.


*SUPPLY*

	supply (noun)
	- the provisions and equipment necessary for people engaged in a particular project
    - the quantity or amount of a commodity that is needed or available

	supply (verb)
	- to make available for use
	- to provide what is needed or wanted
    - to furnish; to equip

SUPPLY uses "packages" as a way to programmatically interact with patterns in variety of contexts:

### USAGE:

    supply <package> <command> [flags]

## Getting Started

- Download the CLI
- Configure the CLI
- List Supply Packages
- Create a new Project from a "Supply Package"
- Interact with your Project using Supply

    supply project list-packages
    supply project get-package --package <package_name> --name <project_name> 
    supply project create-package
    supply project tar-package


    # PACKAGES COMING SOON!
    supply ansible
    supply aws
    supply chef
    supply cloudformation
    supply cobra
    supply docker
    supply gin 
    supply jekyll
    supply packer
    supply terraform

### Project Structure


    /supply

        /cmd
            /common         Supply Common Utility Functions
            /configure      Supply Configure Sub-Command
            /guide          Supply Guide Sub-Command
            /project        Supply Project Sub-Command
            /version        Supply Version Sub-Command
            root.go         Supply Command (e.g. the base command when called without any subcommands)
        
        /version
            version.go      Supply Version
        
        main.go             Main Entry Point to Application

        LICENSE             Mozilla Public License v2.0

        README.md           The File You Are Reading Now



### Developing Supply


#### Common Package
The Common Package (/cmd/common) is a package of exported Utility Functions which keeps the Project DRY, as well as delivers consistency in user experience.

#### Formatting User Output Messages
To deliver a more consistent user experience, there are several common functions which are used to format output:

    # common.FmtMessage(format string, message string)
    # This function is used to create formatted Print Messages for the User

    COLORS
    - white
    - blue
    - green
    - yellow
    - red

    FORMATS
    - h1-<color>    (Header 1)
    - h2-<color>    (Header 2)
    - h3-<color>    (Header 3)
    - text-<color>  (Normal, Mixed-Case Text)
    - upper-<color> (All Upper-Case Text)
    - lower-<color> (All Upper-Case Text)
    - pass          (Fmt Prefix: PASS    | Color: Green)
    - warn          (Fmt Prefix: WARN    | Color: Yellow)
    - fail          (Fmt Prefix: FAIL    | Color: Red)
    - error         (Fmt Prefix: ERROR   | Color: Red)
    - fatal         (Fmt Prefix: FATAL   | Color: Red  | os.Exit(1) )
    - default       (Fmt Prefix: DEFAULT | Color: Blue)
    - success       (Fmt Prefix: SUCCESS | Color: Green)
    - warning       (Fmt Prefix: WARNING | Color: Yello)
    - failure       (Fmt Prefix: FAILURE | Color: Red)


    # common.LogMessage(format string, message string)
    This utility function is used to create prefixed and color formatted Log Messages for the User

    FORMATS
    - info      (Log Prefix: INFO    | Color: Blue)
    - pass      (Log Prefix: PASS    | Color: Green)
    - warn      (Log Prefix: WARN    | Color: Yellow)
    - fail      (Log Prefix: FAIL    | Color: Red)
    - error     (Log Prefix: ERROR   | Color: Red)
    - fatal     (Log Prefix: FATAL   | Color: Red  | os.Exit(1) )
    - default   (Log Prefix: DEFAULT | Color: Blue)
    - success   (Log Prefix: SUCCESS | Color: Green)
    - warning   (Log Prefix: WARNING | Color: Yello)
    - failure   (Log Prefix: FAILURE | Color: Red)

    USAGE:
    common.LogMessage("info", "Example Log Message)

    YYYY/MM/DD HH:mm:ss : INFO: Example Log Message 

    EXAMPLE
    2019/10/10 08:00:00 DEBUG: [package: guide][file: getTopic.go][Cmd Run: func()]
    2019/10/10 08:00:00 DEBUG: [package: project][file: getPackage.go][function parsePackageTemplate()]
    2019/10/10 08:00:00 DEBUG: [package: project][file: getPackage.go][context 'Parsing Package Template File']


    # common.DebugMessage(message string)
    # This function is used to created formatted Debug Messages for User Debugging
    # Debug Messages are only created if debug.trace == true in the supply config file (~/.supply/config.yaml).
    # The purpose of Debug Messages is to add context that helps to quickly identify problems in the code base.

    USAGE:
    YYYY/MM/DD HH:mm:ss DEBUG: [package][file][function or context]

    EXAMPLE
    2019/10/10 08:00:00 DEBUG: [package: guide][file: getTopic.go][Cmd Run: func()]
    2019/10/10 08:00:00 DEBUG: [package: project][file: getPackage.go][function parsePackageTemplate()]
    2019/10/10 08:00:00 DEBUG: [package: project][file: getPackage.go][context 'Parsing Package Template File']
