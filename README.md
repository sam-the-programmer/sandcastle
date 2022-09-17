<section align="center">

# 🏰 <br> Sandcastle

Sandcastle is a blazingly fast, lightweight build tool for any language or shell. With a simple call to the `castle` command, you can build and run your project in seconds.

</section>

## Installation
Download the latest release from our [GitHub Release](https://github.com/neuron-ai/sandcastle/releases/latest), and add it to your PATH environment variable. You may have to rename the file to **castle**, if you are on *MacOS*, as the file is called **castle-macos**

---

## Usage

To build and run, simply use...

```powershell
castle
```

<small>(Other CLI args are available - read on!)</small>

**castle.yaml**

```yaml
build: # castle build
  - build-script arg1 arg2
  - second-build-script arg1

run: # castle run
  - run-section arg1
  - run-more-sections arg1

test: # castle test
  - test-script

tasks: # castle task <task-name>
  - task-name:
    - task-script
```

## Arguments

The following arguments can be passed to the CLI, and they can be used as root level keys in the castle.yaml file.

```yaml
build
run
test
format
deploy
tasks
```

## Running Commands in Parallel

```yaml
config:
  batch-size: 2 # the number of steps to run in parallel
  parallel: # which default commands you want to enable parallelism for
    - build
    - test
  parallel-tasks: # which tasks you want to enable parallelism for
    - task-name
    - task-name-2
```

## Magic Commands

The following magic commands can be used for setting directories, and other command line builtins. They always end with `!`.

> **SETDIR! directory** - Sets the current directory to the specified directory. This is useful for setting the current directory to the project root, or any other directory.

> **GETDIR!** - Gets the current directory.

> **ECHO! string** - Prints the string to the console.

## CLI Arguments


```powershell
Usage of castle:
  -c, -config     Config YAML file to parse. (default "castle.yaml") 
  -v, -version    Show version.

  run              Run the project.
  build            Build the project.
  test             Test the project.
  deploy           Deploy the project.
  format           Format the project.
  task <taskName> Run a task.
```
