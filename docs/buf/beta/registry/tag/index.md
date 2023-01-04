---
id: index
title: index
sidebar_position: -31
---
Manage a repository's tags.

```
buf beta registry tag [flags]
```

### Flags

```
  -h, --help   help for tag
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log-format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Subcommands

* [buf beta registry tag create](tag/create.md)	 - Create a tag for the specified commit.
* [buf beta registry tag list](tag/list.md)	 - List tags for the specified repository.

### Parent Command

* [buf beta registry](../registry.md)	 - Manage assets on the Buf Schema Registry.
