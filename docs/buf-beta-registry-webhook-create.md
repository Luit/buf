---
id: buf-beta-registry-webhook-create
title: buf beta registry webhook create
---
## buf beta registry webhook create

Create a repository webhook.

```
buf beta registry webhook create [flags]
```

### Options

```
      --callback-url string   The url for the webhook to callback to on a given event.
      --event string          The event type to create a webhook for. The proto enum string value is used for this input (e.g. 'WEBHOOK_EVENT_REPOSITORY_PUSH').
  -h, --help                  help for create
      --owner string          The owner name of the repository to create a webhook for.
      --remote string         The remote of the repository the created webhook will belong to.
      --repository string     The repository name to create a webhook for.
```

### Options inherited from parent commands

```
      --debug               Turn on debug logging.
      --log-format string   The log format [text,color,json]. (default "color")
      --timeout duration    The duration until timing out. (default 2m0s)
  -v, --verbose             Turn on verbose mode.
```

### Parent Command

* [buf beta registry webhook](buf-beta-registry-webhook.md)	 - Manage webhooks for a repository on the Buf Schema Registry.
