# supply
supply

## Introduction

**SUPPLY** is a command line (cli) tool developed by [Gadgetry.io](https://gadgetry.io)

It is a utility that our practitioners have designed to support our day-to-day work with the patterns 
we use to solve commonly occurring problems with Cloud Architecture, Infrastructure as Code, and DevOps Automation.

*SUPPLY*

	supply (verb)
	- to make available for use
	- to provide what is needed or wanted
    - to furnish; to equip

	supply (noun)
	- the provisions necessary for a particular project

## Architecture

*API*

*CLI*

*WEB*

## Terminology

*BLUEPRINT*

    blueprint (noun)
    - a guide for making something
    - a pattern or design that can be followed

*PATTERN*

    pattern (noun)
    - something designed or used as a model for making things

*CONSTRUCT*

    construct (verb)
    - to make or form by combining or arranging parts or elements
    - to build something according to a set of instructions or rules

    construct (noun)
    - a thing which is deliberately built or formed

    NOTES:
    A construct is the basic building block
    A construct encapsulates everything needed to create a component of a solution

*COMPONENT*

    component (noun)
    - a part or element of a larger whole

    component (adjective)
    - constituting part of a larger whole
    - a constituent.
    
    NOTES:
    A component is an identifiable part of a larger program or construction
    A component provides a particular resource, service, stack, function or group of related functions

*STACK*

    stack (noun)
    - a collection of resources that you can manage as a single unit (e.g. CloudFormation Stack)

*SOLUTION*

    solution (noun)
    - a means of solving a problem
    - an action or process of solving a problem
    - a product or service designed to meet a particular need
    - a correct answer to a problem, puzzle

*LIBRARY*

    library (noun)
    - a collection of blueprints, solutions, and patterns 

*PROJECT*

    project (verb)
    - to devise in the mind (design)
    - to put or set forth

    project (noun)
    - a specific plan or design

*COMPOSITION*

    composition (noun)
    - the nature of something's constituents
    - the way in which a whole or solution is made up

    NOTES:
    - A high-level construct can be composed from any number of lower-level construct
    - A composition of constructs means that you can define reusable components and share them like any other code.

*COMPOSE*

    compose (verb)
    - to form by putting together

## USAGE

* SUPPLY uses *"blueprints"* as a way to programmatically interact with solution patterns in variety of contexts:
* SUPPLY uses *"constructs"* as a way to programmatically interact with components
* SUPPLY composes *"constructs"* as components of a larger solution

*supply init*
This command initializes a new project via blueprint

*supply create*
This command creates a new blueprint, construct, 

*supply build*
This command runs the build command specified

*supply compose*
This command is used to compose a new blueprint / construct

*supply deploy*

*supply delete*
This command deletes a construct from 




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
