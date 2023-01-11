---
id: deprecate
title: deprecate
sidebar_position: 3
---
Deprecate a plugin by name.

```
buf beta registry plugin deprecate <buf.build/owner/plugins/plugin> [flags]
```

### Flags

```
  -h, --help             help for deprecate
      --message string   The message to display with deprecation warnings.
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Parent Command

* [buf beta registry plugin](index)	 - Manage Protobuf plugins.
