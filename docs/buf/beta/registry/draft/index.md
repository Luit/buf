---
id: index
title: index
sidebar_position: -8
---
Manage a repository's drafts.

```
buf beta registry draft [flags]
```

### Flags

```
  -h, --help   help for draft
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log-format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Subcommands

* [buf beta registry draft delete](draft/delete.md)	 - Delete a draft of a BSR repository by name.
* [buf beta registry draft list](draft/list.md)	 - List drafts for the specified repository.

### Parent Command

* [buf beta registry](../registry.md)	 - Manage assets on the Buf Schema Registry.
