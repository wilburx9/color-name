# color-name

color-name is a tiny CLI tool that spits out a colour name when passed a colour hex.

## :fire:How to use
Clone the repo, cd to project root and run `go build`. Afterwards, you can run `./color-name -h HEX` where "HEX" is the hex value of the colour whose name you want.

## :rocket:Supported formats 
Supported hex formats include 3, 4, 6, and 8 character length (ignoring the leading "#", if any) hex values. The leading "#" sign can be omitted or added. 

Examples includes `FFF`, `#FFF`, `FFFF`, `#FFFF`, `FFFFFF`, `#FFFFFF`, `FFFFFFFF`, and `#FFFFFFFF`

## :muscle:Motivation
For some reason, I feel better when my projects have properly named colours but then I don't know all the names of the
 thousands of colours out there. 
 However, I have been using [this online tool](http://chir.ag/projects/name-that-color) by Chirag Mehta but I decided a
  CLI tool fits my need more closely. The online tool itself is open source, so I ported it to go.
  
## :zap:How it works
The colour hex is parsed, normalized and matched against a list of 1500+ predefined names. An exact match is returned if found, otherwise, the colour name of the closet matching RGB and HSL is returned. The predefined colour names were compiled by Chirag from [Wikipedia](http://en.wikipedia.org/wiki/List_of_colors), [Crayola](http://en.wikipedia.org/wiki/List_of_Crayola_crayon_colors) and [Color-Name Dictionaries](http://www-swiss.ai.mit.edu/~jaffer/Color/Dictionaries.html) like [Resene](http://www-swiss.ai.mit.edu/~jaffer/Color/resenecolours.txt).
  
  Please, note that there's no guarantee that names are 100% correct. 
  
 
