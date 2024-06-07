# AgentHeartbeatResponseState

The state of the agent:
                       * `pending` - The agent needs to perform the setup process again.
                       * `active` - The agent is ready to accept tasks, all is good.
                       * `stopped` - The agent has been stopped by the user.


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `AgentHeartbeatResponseStatePending` | pending                              |
| `AgentHeartbeatResponseStateStopped` | stopped                              |
| `AgentHeartbeatResponseStateError`   | error                                |