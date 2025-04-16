## low-level design
### movement
for all pieces there could only be 1 or 2 movement(s) from these 4 possible movements:
1. diagonal
1. vertical
1. horizontal
1. L movement

for all the movements below:
- `[0]` means selected coordinate
- `[1]` means destination coordinate
- `Z` means the column of coordinate
- `9` means the row of coordinate
#### diagonal
```
if [0]Z > [1]Z && [0]9 < [1]9
	move top-left e.g. H1,G2
else if [0]Z < [1]Z && [0]9 < [1]9
	move top-right e.g. A1,B2
else if [0]Z < [1]Z && [0]9 > [1]9
	move bottom-right e.g. A8,B7
else if [0]Z > [1]Z && [0]9 > [1]9
	move bottom-left e.g. H8,G7
```
#### vertical
```
if [0]Z == [1]Z && [0]9 < [1]9
	move top e.g. A1,A2
else if [0]Z == [1]Z && [0]9 > [1]9
	move bottom e.g. A8,A7
```
#### horizontal
```
if [0]Z > [1]Z && [0]9 == [1]9
	move left e.g. H1,G1
else if [0]Z < [1]Z && [0]9 == [1]9
	move right e.g. A8,B8
```
#### L movement
```
if [0]Z == [1]Z-1 && [0]9 == [1]9+2
or [0]Z == [1]Z-2 && [0]9 == [1]9+1
  move L bottom-right e.g. A8,B6; A8,C7
if [0]Z == [1]Z+1 && [0]9 == [1]9+2
or [0]Z == [1]Z+2 && [0]9 == [1]9+1
  move L bottom-left e.g. H8,G6; H8,F7
if [0]Z == [1]Z+1 && [0]9 == [1]9-2
or [0]Z == [1]Z+2 && [0]9 == [1]9-1
  move L top-left e.g. H1,G3; H1,F2
if [0]Z == [1]Z-1 && [0]9 == [1]9-2
or [0]Z == [1]Z-2 && [0]9 == [1]9-1
  move L top-right e.g. A1,B3; A1,C2
```