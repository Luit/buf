---
id: index
title: repository
sidebar_position: 4
---
Manage repositories.

```
buf beta registry repository [flags]
```

### Flags

```
  -h, --help   help for repository
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Subcommands

* [buf beta registry repository create](create.md)	 - Create a new repository on the BSR.
* [buf beta registry repository delete](delete.md)	 - Delete a BSR repository by name.
* [buf beta registry repository deprecate](deprecate.md)	 - Deprecate a repository on the BSR.
* [buf beta registry repository get](get.md)	 - Get a BSR repository by name.
* [buf beta registry repository list](list.md)	 - List BSR repositories.
* [buf beta registry repository undeprecate](undeprecate.md)	 - Undeprecate a BSR repository.
* [buf beta registry repository update](update.md)	 - Update a BSR repository settings.

### Parent Command

* [buf beta registry](index.md)	 - Manage assets on the Buf Schema Registry.
