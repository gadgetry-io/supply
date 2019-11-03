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
