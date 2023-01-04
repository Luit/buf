---
id: index
title: index
sidebar_position: -40
---
Manage Protobuf templates on the Buf Schema Registry.

```
buf beta registry template [flags]
```

### Flags

```
  -h, --help   help for template
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log-format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Subcommands

* [buf beta registry template create](template/create.md)	 - Create a new Buf template.
* [buf beta registry template delete](template/delete.md)	 - Delete a template by name.
* [buf beta registry template deprecate](template/deprecate.md)	 - Deprecate a template by name.
* [buf beta registry template list](template/list.md)	 - List templates on the specified remote.
* [buf beta registry template undeprecate](template/undeprecate.md)	 - Undeprecate a template by name.
* [buf beta registry template version](template/version.md)	 - Manage Protobuf template versions.

### Parent Command

* [buf beta registry](../registry.md)	 - Manage assets on the Buf Schema Registry.
