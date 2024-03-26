# vert
Command line application
this tool is called vert, and it converts between different units of temeperature.

It defines a set of options (flags) that determine the target temperature. and takes a list of positional arguments in the form {value}{symbol}.
`vert -k 56c` can be read as convert 56 degrees Celsius to Kelvin.

*Note when using negative numbers*: prefix the argument(s) with `--` to indicate that what follows isn't a flag, if not the negative number will be interpreted as a flag by the program which raises an error. Read more [here](https://github.com/urfave/cli/issues/645#issuecomment-314882649) and [here](https://askubuntu.com/questions/278000/how-do-i-use-filenames-that-start-with-a-dash-as-command-arguments)
![ray-so-export(3)](https://github.com/themilar/vert/assets/53567551/5d0cefa9-3517-4d18-99b8-d87baa980852)
