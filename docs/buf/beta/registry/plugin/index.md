---
id: index
title: plugin
sidebar_position: 4
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
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Subcommands

* [buf beta registry plugin create](create.md)	 - Create a new Protobuf plugin.
* [buf beta registry plugin delete](delete.md)	 - Delete a Protobuf plugin by name.
* [buf beta registry plugin deprecate](deprecate.md)	 - Deprecate a plugin by name.
* [buf beta registry plugin list](list.md)	 - List plugins on the specified remote.
* [buf beta registry plugin undeprecate](undeprecate.md)	 - Undeprecate a plugin by name.
* [buf beta registry plugin version](version/index.md)	 - Manage Protobuf plugin versions.

### Parent Command

* [buf beta registry](index.md)	 - Manage assets on the Buf Schema Registry.
