# UpdateAgentRequestBody


## Fields

| Field                             | Type                              | Required                          | Description                       |
| --------------------------------- | --------------------------------- | --------------------------------- | --------------------------------- |
| `ID`                              | *int64*                           | :heavy_check_mark:                | The id of the agent               |
| `Name`                            | *string*                          | :heavy_check_mark:                | The hostname of the agent         |
| `ClientSignature`                 | *string*                          | :heavy_check_mark:                | The signature of the client       |
| `OperatingSystem`                 | *string*                          | :heavy_check_mark:                | The operating system of the agent |
| `Devices`                         | []*string*                        | :heavy_check_mark:                | N/A                               |