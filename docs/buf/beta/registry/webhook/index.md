---
id: index
title: webhook
sidebar_position: 7
---
Manage webhooks for a repository on the Buf Schema Registry.

```
buf beta registry webhook [flags]
```

### Flags

```
  -h, --help   help for webhook
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Subcommands

* [buf beta registry webhook create](create.md)	 - Create a repository webhook.
* [buf beta registry webhook delete](delete.md)	 - Delete a repository webhook.
* [buf beta registry webhook list](list.md)	 - List repository webhooks.

### Parent Command

* [buf beta registry](index.md)	 - Manage assets on the Buf Schema Registry.
