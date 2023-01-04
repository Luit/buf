---
id: index
title: index
sidebar_position: -20
---
Manage Protobuf plugins.

```
buf beta registry plugin [flags]
```

### Flags

```
  -h, --help   help for plugin
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log-format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Subcommands

* [buf beta registry plugin create](plugin/create.md)	 - Create a new Protobuf plugin.
* [buf beta registry plugin delete](plugin/delete.md)	 - Delete a Protobuf plugin by name.
* [buf beta registry plugin deprecate](plugin/deprecate.md)	 - Deprecate a plugin by name.
* [buf beta registry plugin list](plugin/list.md)	 - List plugins on the specified remote.
* [buf beta registry plugin undeprecate](plugin/undeprecate.md)	 - Undeprecate a plugin by name.
* [buf beta registry plugin version](plugin/version.md)	 - Manage Protobuf plugin versions.

### Parent Command

* [buf beta registry](../registry.md)	 - Manage assets on the Buf Schema Registry.
