---
id: create
title: create
sidebar_position: 0
---
Create a new token for a user.

```
buf alpha registry token create <buf.build> [flags]
```

### Flags

```
  -h, --help           help for create
      --note string    A human-readable note that identifies the token.
      --ttl duration   How long the token should live. Set to 0 for no expiry. (default 720h0m0s)
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Parent Command

* [buf alpha registry token](index.md)	 - Manage user tokens.
