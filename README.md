# Debris

**Bridge the Gap Between Design Library Formats.**

[简体中文](docs/README/README_zh-Hans.md) [繁體中文](docs/README/README_zh-Hant.md)

Debris is a cross-platform command-line tool designed to simplify your design workflow by enabling **bi-directional conversion of design libraries** between different design library management applications.  Stop struggling with incompatible library formats and seamlessly move your valuable design assets across platforms.

Currently, Debris focuses on converting between **Billfish and Pixcall** libraries, leveraging Eagle format as an intermediary.  Future development will expand support to more EDA formats and direct conversion methods.

> **🚧  Work in Progress - Early Stage Development 🚧**
>
> Please note that Debris is currently under active development and is **not yet production-ready.**  While conversion between Billfish and Pixcall libraries (via Eagle as an intermediate format) is functional, expect potential issues and limitations. We appreciate your early interest and encourage you to contribute to its development!

## Features

**Currently Supported:**

* **Billfish ↔ Pixcall Conversion (via Eagle):** Convert Billfish libraries to Pixcall and vice-versa, using Eagle format as a stepping stone. This allows you to utilize libraries across these two popular EDA tools.
* **Command-Line Interface (CLI):**  Simple and intuitive command-line interface for easy integration into your scripts and workflows.
* **Cross-Platform Compatibility:** Thanks to the multi-platform compatibility of the Go programming language, debris can be run on major operating systems (macOS, Linux, Windows).

## Usage

Debris is used via the command line.

You need to specify the input and output formats, as well as the input and output paths.

Example: Converting a Billfish library to a Eagle library

```
debris --from billfish --to eagle --input ./Billfish --output ./Eagle
```

The basic command structure is:
```
debris --from <input_format> --to <output_format> --input <input_path> --output <output_path>
```

- `--from billfish` or `-f billfish`: Specifies the input library format as Billfish.
- `--to eagle` or `-t eagle`: Specifies the desired output library format as Pixcall.
- `--input ./Billfish` or `-i ./Billfish`: Indicates the path to the input Billfish library file or directory. Currently, it expects a Billfish library folder.
- `--output ./newFolder` or `-o ./newFolder`: Specifies the output directory where the converted library will be saved.
