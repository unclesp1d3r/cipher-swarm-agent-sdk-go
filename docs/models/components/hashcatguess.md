# HashcatGuess


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `GuessBase`                                                 | *string*                                                    | :heavy_check_mark:                                          | The base value used for the guess (for example, the mask)   |
| `GuessBaseCount`                                            | *int64*                                                     | :heavy_check_mark:                                          | The number of times the base value was used                 |
| `GuessBaseOffset`                                           | *int64*                                                     | :heavy_check_mark:                                          | The offset of the base value                                |
| `GuessBasePercentage`                                       | *float64*                                                   | :heavy_check_mark:                                          | The percentage completion of the base value                 |
| `GuessMod`                                                  | *string*                                                    | :heavy_check_mark:                                          | The modifier used for the guess (for example, the wordlist) |
| `GuessModCount`                                             | *int64*                                                     | :heavy_check_mark:                                          | The number of times the modifier was used                   |
| `GuessModOffset`                                            | *int64*                                                     | :heavy_check_mark:                                          | The offset of the modifier                                  |
| `GuessModPercentage`                                        | *float64*                                                   | :heavy_check_mark:                                          | The percentage completion of the modifier                   |
| `GuessMode`                                                 | *int64*                                                     | :heavy_check_mark:                                          | The mode used for the guess                                 |