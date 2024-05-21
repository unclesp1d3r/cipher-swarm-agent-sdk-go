# HashcatBenchmark


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `HashType`                                       | *int64*                                          | :heavy_check_mark:                               | The hashcat hash type                            |
| `Runtime`                                        | *int64*                                          | :heavy_check_mark:                               | The runtime of the benchmark in milliseconds.    |
| `HashSpeed`                                      | *float64*                                        | :heavy_check_mark:                               | The speed of the benchmark in hashes per second. |
| `Device`                                         | *int64*                                          | :heavy_check_mark:                               | The device used for the benchmark                |