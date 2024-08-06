# State

The state of the agent:
                       * `pending` - The agent needs to perform the setup process again.
                       * `active` - The agent is ready to accept tasks, all is good.
                       * `error` - The agent has encountered an error and needs to be checked.
                       * `stopped` - The agent has been stopped by the user.


## Values

| Name           | Value          |
| -------------- | -------------- |
| `StatePending` | pending        |
| `StateStopped` | stopped        |
| `StateError`   | error          |