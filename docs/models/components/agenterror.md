# AgentError


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Message`                                                   | *string*                                                    | :heavy_check_mark:                                          | The error message                                           |
| `Metadata`                                                  | [*components.Metadata](../../models/components/metadata.md) | :heavy_minus_sign:                                          | Additional metadata about the error                         |
| `Severity`                                                  | [components.Severity](../../models/components/severity.md)  | :heavy_check_mark:                                          | The severity of the error                                   |
| `AgentID`                                                   | *int64*                                                     | :heavy_check_mark:                                          | The agent that caused the error                             |
| `TaskID`                                                    | **int64*                                                    | :heavy_minus_sign:                                          | The task that caused the error, if any                      |