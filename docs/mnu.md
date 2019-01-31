# MNU format
This format specifies game menus used within 1.3 version of the game.

## Specification

### Spec: FormatMNU

| Type | Name | Description |
| ---- | ---- | ----------- |
| uint8 | Magic | 4 definitions of Magic; should be 'Menu' |
| uint32 | Unknown |  |
| uint32 | NumControls |  |
| Control | Controls | N definitions of Control;  |
### Spec: Control

| Type | Name | Description |
| ---- | ---- | ----------- |
| uint32 | Unknown |  |
| uint8 | Type | String consisting of 4 characters;  |
| float32 | Position | plain array of 2 elements;  |
| float32 | Scale | plain array of 2 elements;  |
| uint32 | TextID | taken from Textdb_xx.def |
| uint16 | TextColor |  |
| uint16 | BackgroundColor |  |

