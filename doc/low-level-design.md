## low-level design
### movement
for all pieces there could only be 1 or 2 movement(s) from these 4 possible movements:
1. diagonal
1. vertical
1. horizontal
1. L movement
#### diagonal
```
if origin.col > dest.col && origin.row < dest.row
	move top-left e.g. H1,G2
else if origin.col < dest.col && origin.row < dest.row
	move top-right e.g. A1,B2
else if origin.col < dest.col && origin.row > dest.row
	move bottom-right e.g. A8,B7
else if origin.col > dest.col && origin.row > dest.row
	move bottom-left e.g. H8,G7
```
#### vertical
```
if origin.col == dest.col && origin.row < dest.row
	move top e.g. A1,A2
else if origin.col == dest.col && origin.row > dest.row
	move bottom e.g. A8,A7
```
#### horizontal
```
if origin.col > dest.col && origin.row == dest.row
	move left e.g. H1,G1
else if origin.col < dest.col && origin.row == dest.row
	move right e.g. A8,B8
```
#### L movement
```
if origin.col == dest.col+1 && origin.row == dest.row-2
or origin.col == dest.col+2 && origin.row == dest.row-1
  move L top-left e.g. H1,G3; H1,F2
if origin.col == dest.col-1 && origin.row == dest.row-2
or origin.col == dest.col-2 && origin.row == dest.row-1
  move L top-right e.g. A1,B3; A1,C2
if origin.col == dest.col-1 && origin.row == dest.row+2
or origin.col == dest.col-2 && origin.row == dest.row+1
  move L bottom-right e.g. A8,B6; A8,C7
if origin.col == dest.col+1 && origin.row == dest.row+2
or origin.col == dest.col+2 && origin.row == dest.row+1
  move L bottom-left e.g. H8,G6; H8,F7
```
### validate pieces movement
there are several approaches:
1. pre-generate all possible squares from given origin square and check whether the destination square is in those squares or not
1. trace the origin square to destination square and determine if the movement is one of horizontal, vertical, diagonal, or L based one how rows and cols changes every step

ultimately the pre-generate approach seems to be the simplest one, now to make searching destination square in possible squares fast we can use:
1. `map` treated as a set, the key would be the square
1. 2D index like K-D Tree

`map` treated as a set is far simpler than having to implement K-D Tree

### dependency
main->usecase->state->entity