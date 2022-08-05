# eas slack bot

This is a service for webhooks of eas. When build or submit, relevant information is passed through the Slack bot.

## Useage

```bash
docker run -d -p ${CONTAINER_PORT}:${APP_PORT} --name ${CONTAINER_NAME} \
-e APP_PORT=${APP_PORT} \
-e SECRET_WEBHOOK_KEY=${SECRET_WEBHOOK_KEY} \
-e SLACK_TOKEN=${SLACK_TOKEN} \
-e SLACK_CHANNEL_NAME=${SLACK_CHANNEL_NAME} \
ghcr.io/himitery/eas-slack-bot:latest
```

### Environments

|    Environment     | Description                                                                        |
| :----------------: | :--------------------------------------------------------------------------------- |
|      APP_PORT      | internal service port                                                              |
| SECRET_WEBHOOK_KEY | SECRET_WEBHOOK_KEY has to match SECRET value set with `eas webhook:create` command |
|    SLACK_TOKEN     | bot user oauth token for your workspace                                            |
| SLACK_CHANNEL_NAME | channel name in your workspace to which messages will be sent                      |


## Reference

- [eas webhooks](https://docs.expo.dev/eas/webhooks)