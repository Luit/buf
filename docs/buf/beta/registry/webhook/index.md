---
id: index
title: index
sidebar_position: -44
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
      --log-format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Subcommands

* [buf beta registry webhook create](webhook/create.md)	 - Create a repository webhook.
* [buf beta registry webhook delete](webhook/delete.md)	 - Delete a repository webhook.
* [buf beta registry webhook list](webhook/list.md)	 - List repository webhooks.

### Parent Command

* [buf beta registry](../registry.md)	 - Manage assets on the Buf Schema Registry.
