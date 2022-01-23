## `github.com/artnoi43/fngobot/bot/tghandler`
Package `tghandler` provides FnGoBot with chat message handlers. The main program initialized a handler with `NewHandler`, which takes in `telebot.Bot` and `telebot.Message`. The initialized handler may be used to perform some methods, like sending quotes, price alerts, or user command error.

## Types and interfaces
The type for FnGoBot handlers is `handler`, which implements the `Handler` interface. There're also `Handlers`, which is a type alias for `[]*Handler`

## Flow

- User sends their chat message

- The main program captures the chat message, and wraps it into struct `UserCommand` with a command flag (field `Command`) corresponding to diiferent bot commands, i.e. sending helps, quotes, tracking, and price alerts

- Generated `UserCommand` will then call `Parse()`, which in turn return a `BotCommand` based on what the chat message look like

- `Parse()` parses a chat message into `BotCommand`. It collects all information from the chat message, and store such info in the return variable (type `BotCommand`)

- Generated `BotCommand` is passed to `NewHandler()`, where it is wrapped into a `handler`. `NewHandler` returns a `Handler` interface, which our `handler` implements

- The handler returned from `NewHandler()` can now call `Handle()`, which handles the session based on `BotCommand` embedded in the handler itself.