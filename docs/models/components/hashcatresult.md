# HashcatResult


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `Timestamp`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | The time the hash was cracked             |
| `Hash`                                    | *string*                                  | :heavy_check_mark:                        | The hash value                            |
| `PlainText`                               | *string*                                  | :heavy_check_mark:                        | The plain text value                      |