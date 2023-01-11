---
id: init
title: init
sidebar_position: 2
---
Initializes and writes a new buf.yaml configuration file.

```
buf mod init [flags]
```

### Flags

```
      --doc             Write inline documentation in the form of comments in the resulting configuration file.
  -h, --help            help for init
  -o, --output string   The directory to write the configuration file to. (default ".")
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Parent Command

* [buf mod](index)	 - Manage Buf modules.
