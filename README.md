<section align="center">

# 🏰 <br> SandCastle

**SandCastle** is a high performance, lightweight build tool for any language, OS or shell.

</section>

- [🏰  SandCastle](#--sandcastle)
- [Installation](#installation)
  - [Binaries](#binaries)
    - [Windows](#windows)
    - [Linux](#linux)
    - [MacOS](#macos)
  - [From Source](#from-source)
- [Usage](#usage)
  - [Arguments](#arguments)
  - [_castle.yaml_](#castleyaml)
  - [Recognised Shells](#recognised-shells)

# Installation

## Binaries

### Windows

- Download the `castle.exe` binary from [releases page](https://github.com/sam-the-programmer/sandcastle/releases/latest).
- Add it to your `$PATH` and you're good to go!

### Linux

- Download the `castle` binary from [releases page](https://github.com/sam-the-programmer/sandcastle/releases/latest).
- Add it to your `$PATH` and you're good to go!

### MacOS

- Download the `castle-macos` binary from [releases page](https://github.com/sam-the-programmer/sandcastle/releases/latest).
- Rename it to `castle`
- Add it to your `$PATH` and you're good to go!

## From Source

# Usage

## Arguments

- `init` initialise a **SandCastle** project.
- `<task-name>` run the specified task.

## _castle.yaml_

This is a configuration file that contains all of the tasks and things that you want **SandCastle** to do.

Create a _castle.yaml_ file, using `castle init` or any other method of your choice.

The contents are as follows.

```yaml
config: # the configuration options
  log_shell_cmds: true # should it give a title for shell steps in a task?

  # log_level is for the internal logging of SandCastle
  log_level: debug # choose one of "info", "debug", "warn", "error" or "none" - these are

tasks: # the tasks, use the name of the task as a CLI argument
  run: # the name of the task
    - echo "Hello, World!" # the steps of the task
    - echo "Second Step!"
  deploy: # another task
    - echo "Deployed!" # you can use shell commands, or any other terminal command
```

Now, if we wanted to run the `deploy` task, we would run `castle deploy`. Just like that!

## Recognised Shells

In order to run shell commands, you need to have a shell installed. **SandCastle** supports the following shells:

- bash
- zsh
- sh
- csh
- ksh
- tcsh
- dash
- fish
