---
id: update
title: update
sidebar_position: 6
---
Update a module's dependencies by updating the buf.lock file.

### Synopsis

Update a module&#39;s dependencies by updating the buf.lock file.

Fetch the latest digests for the specified references in the config file, and write them and their transitive dependencies to the buf.lock file. The first argument is the directory of the local module to update. Defaults to &#34;.&#34; if no argument is specified.

```
buf mod update <directory> [flags]
```

### Flags

```
  -h, --help           help for update
      --only strings   The name of the dependency to update. When set, only this dependency is updated (along with any of its sub-dependencies). May be passed multiple times.
```

### Flags inherited from parent commands

```
      --debug               Turn on debug logging.
      --log_format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Parent Command

* [buf mod](index.md)	 - Manage Buf modules.
