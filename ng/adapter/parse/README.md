## `github.com/artnoi43/fngobot/lib/parse`
Package `parse` parses FnGoBot chat commands into structs to be sent to handlers in `github.com/artnoi43/fngobot/lib/bot/handlers`. It parses `UserCommand`, a struct representing user input, into `BotCommand` - a struct representing used internally by the handlers.

### `type UserCommand struct`
This struct has 2 fields: `Command` and `Chat`. `UserCommand.Command` is an int (enumerated) representing the bot commands to be used according to the enumerated int. `UserCommand.Chat` is the whole chat message.
### `type BotCommand struct`
This struct has `quoteCommand`, `trackCommand`, and `bot.Alert` embedded.
### `type quoteCommand struct`
This struct holds information just enough for `/quote` handlers to perform its tasks. It has only one field, `Securities`, which is a slice of `bot.Securities`.
### `type trackCommand struct`
This struct has `quoteCommand` embedded, but it also has one extra field: `TrackTimes`, which is an int specifying for how many times should the `/track` handlers iterate.
### There's no `alertCommand`
Because type `bot.Alert` is already pretty good at holding information relevant to `/alert` handlers, we can use `bot.Alert` as our bot commands for `/alert`. The handler methood `PriceAlert()` accepts just `bot.Alert` to work, so there's no need to add type `alertCommand` for `/alert` handlers.
### `func (c UserCommand) Parse() (cmd BotCommand, parseError int)`
This `UserCommand` method returns `BotCommand` or a parsing error. The logic for parsing is all in this function. Its use case switch on `c.Command` to determine how to parse the chat message.