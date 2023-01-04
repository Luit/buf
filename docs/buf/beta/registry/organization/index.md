---
id: index
title: index
sidebar_position: -12
---
Manage organizations.

```
buf beta registry organization [flags]
```

### Flags

```
  -h, --help   help for organization
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log-format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Subcommands

* [buf beta registry organization create](organization/create.md)	 - Create a new organization on the BSR.
* [buf beta registry organization delete](organization/delete.md)	 - Delete an organization by name on the BSR.
* [buf beta registry organization get](organization/get.md)	 - Get an organization on the BSR by name.

### Parent Command

* [buf beta registry](../registry.md)	 - Manage assets on the Buf Schema Registry.
