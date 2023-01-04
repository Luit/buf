---
id: index
title: buf
sidebar_position: 0
---
version 1.11.1-dev

The Buf CLI

### Synopsis

The Buf CLI

A tool for working with Protocol Buffers and managing resources on the Buf Schema Registry (BSR). 

```
buf [flags]
```

### Flags

```
      --debug               Turn on debug logging.
  -h, --help                help for buf
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
      --version             Print the version.
```

### Subcommands

* [buf beta](beta/index.md)	 - Beta commands. Unstable and likely to change.
* [buf breaking](breaking.md)	 - Verify that the input location has no breaking changes compared to the against location.
* [buf build](build.md)	 - Build all Protobuf files from the specified input and output a Buf image.
* [buf convert](convert.md)	 - Convert a message from binary to JSON or vice versa
* [buf export](export.md)	 - Export the files from the input location to an output location.
* [buf format](format.md)	 - Format all Protobuf files from the specified input and output the result.
* [buf generate](generate.md)	 - Generate stubs for protoc plugins using a template.
* [buf lint](lint.md)	 - Verify that the input location passes lint checks.
* [buf ls-files](ls-files.md)	 - List all Protobuf files for the input.
* [buf mod](mod/index.md)	 - Manage Buf modules.
* [buf push](push.md)	 - Push a module to a registry.
* [buf registry](registry/index.md)	 - Manage assets on the Buf Schema Registry.

