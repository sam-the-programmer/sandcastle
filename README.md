<section align="center">

# 🏰 <br> Sandcastle

Sandcastle is a blazingly fast, lightweight build tool for any language or shell. With a simple call to the `castle` command, you can build and run your project in seconds.

</section>

## Installation
Download the latest release from our [GitHub Release](https://github.com/neuron-ai/sandcastle/releases/latest), and add it to your PATH environment variable. You may have to rename the file to **castle**, if you are on *MacOS*, as the file is called **castle-macos**

## Usage

To build and run, simply use...

```bash
castle
```

Where castle.yaml is similar to...

```yaml
build:
  - build_script arg1 arg2
  - second_build_script arg1

run:
  - run_section arg1
  - run_more_sections arg1
```

## CLI Arguments


```bash
Usage of castle.exe:
  -b, -build      Build the project.
  -c, -config     Config YAML file to parse. (default "castle.yml") 
  -r, -run        Run the project.
  -v, -version    Show version.
```
