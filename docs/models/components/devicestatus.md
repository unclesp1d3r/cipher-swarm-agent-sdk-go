# DeviceStatus


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `DeviceID`                                                     | *int64*                                                        | :heavy_check_mark:                                             | The id of the device                                           |
| `DeviceName`                                                   | *string*                                                       | :heavy_check_mark:                                             | The name of the device                                         |
| `DeviceType`                                                   | [components.DeviceType](../../models/components/devicetype.md) | :heavy_check_mark:                                             | The type of the device                                         |
| `Speed`                                                        | *int64*                                                        | :heavy_check_mark:                                             | The speed of the device                                        |
| `Utilization`                                                  | *int64*                                                        | :heavy_check_mark:                                             | The utilization of the device                                  |
| `Temperature`                                                  | *int64*                                                        | :heavy_check_mark:                                             | The temperature of the device, or -1 if unmonitored.           |