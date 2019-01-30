# check.bin
This file describes navigation points and paths used by pedestrians and trains within the game.

## Specification
### Enum: PointType

| Value | Description |
| ----- | ----------- |
| Pedestrian = 0x1 |  |
| AI = 0x2 |  |
| Vehicle = 0x4 |  |
| Tram_Station = 0x8 |  |
| Special = 0x10 |  |
### Enum: LinkType

| Value | Description |
| ----- | ----------- |
| Pedestrian = 1 |  |
| AI = 2 |  |
| TrainsAndSalinas_Forward = 4 |  |
| TrainsAndSalinas_Reverse = 0x8400 |  |
| Other = 0x1000 |  |

### Spec: Header

| Type | Name | Description |
| ---- | ---- | ----------- |
| uint32 | Magic |  |
| uint32 | NumPoints |  |
### Spec: Point

| Type | Name | Description |
| ---- | ---- | ----------- |
| uint16 | Type |  |
| uint16 | ID |  |
| uint16 | AreaSize |  |
| uint8 | Unk | plain array of 10 elements; Unknown values |
| uint8 | EnterLinks |  |
| uint8 | ExitLinks | Is the same as EnterLinks |
### Spec: Link

| Type | Name | Description |
| ---- | ---- | ----------- |
| uint16 | TargetPoint |  |
| uint16 | LinkType |  |
| float32 | Unk |  |

