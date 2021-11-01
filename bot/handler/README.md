## `github.com/artnoi43/fngobot/bot/handler`
Package `handler` provides FnGoBot with chat message handlers. The main program initialized a handler with `NewHandler`, which takes in `telebot.Bot` and `telebot.Message`. The initialized handler may be used to perform some methods, like sending quotes, price alerts, or user command error.

Because all handlers have access to all methods, we can reuse these methods anywhere. This is why `/alert` and `/track` commands still have access to `/quote` behaviors.