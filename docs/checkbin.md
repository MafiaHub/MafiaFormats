# check.bin
This file describes navigation points and paths used by pedestrians and trains within the game.

Each point consists of various properties, such as area circle radius, point type or number of connecting links.

Links are stored as a block of data right after the block of points. where per each point, we take N links (based on EnterLinks) from the block that are associated to the point.

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
| Trains_Forward = 4 |  |
| Trains_Reverse = 0x8400 |  |
| Other = 0x1000 |  |

### Spec: FormatCheckBIN

| Type | Name | Description |
| ---- | ---- | ----------- |
| uint32 | Magic | Should be 0x1ABCEDF |
| uint32 | NumPoints | Number of navigation nodes |
| Point | Points | N definitions of Point;  |
| Link | Links | N definitions of Link; Size depends on number of links used by all points |
### Spec: Point

| Type | Name | Description |
| ---- | ---- | ----------- |
| PointType | Type |  |
| uint16 | ID |  |
| uint16 | AreaSize |  |
| uint8 | Unk | plain array of 10 elements; Unknown values |
| uint8 | EnterLinks | How many consequent links belong to this node |
| uint8 | ExitLinks | Is the same as EnterLinks |
### Spec: Link

| Type | Name | Description |
| ---- | ---- | ----------- |
| uint16 | TargetPoint | Refers back to a point it connects to |
| LinkType | LinkType |  |
| float32 | Unk |  |

