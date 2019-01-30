# check.bin
This file describes navigation points and paths used by pedestrians and trains within the game.

Each point consists of various properties, such as area circle radius, point type or number of connecting links.

Links are stored as a block of data right after the block of points. where per each point, we take N links (based on EnterLinks) from the block that are associated to the point.

## Specification

