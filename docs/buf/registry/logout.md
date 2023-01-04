---
id: logout
title: logout
sidebar_position: 2
---
Log out of the Buf Schema Registry.

### Synopsis

Log out of the Buf Schema Registry.

This command removes any BSR credentials from your .netrc file. 

```
buf registry logout [flags]
```

### Flags

```
  -h, --help   help for logout
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Parent Command

* [buf registry](index.md)	 - Manage assets on the Buf Schema Registry.
