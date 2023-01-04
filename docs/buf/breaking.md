---
id: breaking
title: breaking
sidebar_position: 3
---
Verify that the input location has no breaking changes compared to the against location.

### Synopsis

Verify that the input location has no breaking changes compared to the against location.

The first argument is the source, module, or image to check for breaking changes.
The first argument must be one of format [bin,dir,git,json,mod,protofile,tar,zip].
If no argument is specified, defaults to &#34;.&#34;. 

```
buf breaking <input> --against <against-input> [flags]
```

### Flags

```
      --against string          Required. The source, module, or image to check against. Must be one of format [bin,dir,git,json,mod,protofile,tar,zip].
      --against-config string   The file or data to use to configure the against source, module, or image.
      --config string           The file or data to use for configuration.
      --disable-symlinks        Do not follow symlinks when reading sources or configuration from the local filesystem.
                                By default, symlinks are followed in this CLI, but never followed on the Buf Schema Registry.
      --error-format string     The format for build errors or check violations printed to stdout. Must be one of [text,json,msvs,junit]. (default "text")
      --exclude-imports         Exclude imports from breaking change detection.
      --exclude-path strings    Exclude specific files or directories, for example "proto/a/a.proto" or "proto/a".
                                If specified multiple times, the union is taken.
  -h, --help                    help for breaking
      --limit-to-input-files    Only run breaking checks against the files in the input.
                                When set, the against input contains only the files in the input.
                                Overrides --path.
      --path strings            Limit to specific files or directories, for example "proto/a/a.proto" or "proto/a".
                                If specified multiple times, the union is taken.
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Parent Command

* [buf](index.md)	 - The Buf CLI
