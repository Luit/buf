---
id: get
title: get
sidebar_position: 1
---
Get details about a commit.

```
buf beta registry commit get <buf.build/owner/repo:ref> [flags]
```

### Flags

```
      --format string   The output format to use. Must be one of [text,json] (default "text")
  -h, --help            help for get
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Parent Command

* [buf beta registry commit](index.md)	 - Manage a repository's commits.
