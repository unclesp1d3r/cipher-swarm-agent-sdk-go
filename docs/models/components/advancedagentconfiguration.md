# AdvancedAgentConfiguration


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `AgentUpdateInterval`                                               | *int64*                                                             | :heavy_check_mark:                                                  | The interval in seconds to check for agent updates                  |
| `UseNativeHashcat`                                                  | *bool*                                                              | :heavy_check_mark:                                                  | Use the hashcat binary already installed on the client system       |
| `BackendDevice`                                                     | *string*                                                            | :heavy_check_mark:                                                  | The device to use for hashcat, separated by commas                  |
| `OpenclDevices`                                                     | **string*                                                           | :heavy_minus_sign:                                                  | The OpenCL device types to use for hashcat, separated by commas     |
| `EnableAdditionalHashTypes`                                         | *bool*                                                              | :heavy_check_mark:                                                  | Causes hashcat to perform benchmark-all, rather than just benchmark |