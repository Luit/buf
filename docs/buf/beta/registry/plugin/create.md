---
id: create
title: create
sidebar_position: 0
---
Create a new Protobuf plugin.

```
buf beta registry plugin create <buf.build/owner/plugins/plugin> [flags]
```

### Flags

```
      --format string       The output format to use. Must be one of [text,json] (default "text")
  -h, --help                help for create
      --visibility string   The plugin's visibility setting. Must be one of [public,private].
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Parent Command

* [buf beta registry plugin](index.md)	 - Manage Protobuf plugins.
