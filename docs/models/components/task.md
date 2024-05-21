# Task


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `ID`                                      | *int64*                                   | :heavy_check_mark:                        | The id of the task                        |
| `AttackID`                                | *int64*                                   | :heavy_check_mark:                        | The id of the attack                      |
| `StartDate`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | The time the task was started             |
| `Status`                                  | *string*                                  | :heavy_check_mark:                        | The status of the task                    |
| `Skip`                                    | **int64*                                  | :heavy_minus_sign:                        | The offset of the keyspace                |
| `Limit`                                   | **int64*                                  | :heavy_minus_sign:                        | The limit of the keyspace                 |