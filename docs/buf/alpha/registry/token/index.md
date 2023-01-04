---
id: index
title: token
sidebar_position: 1
---
Manage user tokens.

```
buf alpha registry token [flags]
```

### Flags

```
  -h, --help   help for token
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Subcommands

* [buf alpha registry token create](create.md)	 - Create a new token for a user.
* [buf alpha registry token delete](delete.md)	 - Delete a token by ID.
* [buf alpha registry token get](get.md)	 - Get a token by ID.
* [buf alpha registry token list](list.md)	 - List my tokens.

### Parent Command

* [buf alpha registry](index.md)	 - Manage assets on the Buf Schema Registry.
