# Invoice-generator

Invoice-generator is a simple tool for generating invoices in automatic mode.
There is no need to fill out anything every month. You do it once and that's it.

It can be used for several projects at once, just create different configuration
files and put them in different folders where you want to save invoices.

## Installation

### Binaries

Pre-built [binaries](https://github.com/coffeewasmyidea/invoice-generator/releases/latest).

### Build from source

Clone the repo:

```shell
git clone https://github.com/coffeewasmyidea/invoice-generator.git
```

Run the command:

```shell
make install
```

## Configuration

 When you first run generating an invoice without an `invoice-generator.toml`
 file in the directory, the program will prompt a question:

```shell
The file invoice-generator.toml was not found in the current directory. Want to create it? (Y/n)
```

Type `Y`, and then you need to open the created `invoice-generator.toml` file
and fill in the parameters of your invoice there. This completes the
configuration setup, now you can generate invoices.

## General usage

To generate a new invoice, use the command:

```shell
invoice-generator g
```

In addition, you can pass the date (month.year) of the desired service period as
an argument as follows:

```shell
invoice-generator g 10.2023
```

You will get an invoice for the desirable period (previous month for example)
with the correct dates for all presented fields.

## Help information

```shell
$ invoice-generator
The file invoice-generator.toml was not found in the current directory. Want to create it? (Y/n) y
invoice-generator.toml created!
NAME:
   invoice-generator - cli invoice assistant

USAGE:
   invoice-generator [global options] command [command options] [arguments...]

VERSION:
   0.6.3.d1ca33e

DESCRIPTION:
   This is a simple tool for generating invoices in automatic mode. There is no need to fill out anything every month. You do it once and that's it.

COMMANDS:
   generate, g  Generate a new invoice based on invoice-generator.toml information and current date.
                Alternatively, you can pass the date (month.year) of the desired service period as anargument like this: invoice-generator g 10.2023
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Invoice example:

<div align="center">
<img src="examples/SE-010123.png" max-width="880px" style="margin:10px 0 15px 0">
</div>

## Requiremrnts

There are no requirements and additional dependencies it will work on (Windows/Linux/MacOS).
