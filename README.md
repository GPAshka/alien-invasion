# Alien invasion task

General implementation flow:
1. main function opens input file
2. New game object created. At this moment aliens are created and world map is being read 
from the input file line by line, then each line is parsed and added to the Map object.
3. All aliens are randomly set to the starting cities.
4. Main game loop starts. Move aliens and then fight them until​ ​all​ ​the​ ​aliens​ ​have​ ​been destroyed,​ ​
or​ ​each​ ​alien​ ​has​ ​moved​ ​at​ ​least​ ​10,000​ ​times.
5. After game finished, print remaining world map.  

Implementation notes:
1. If alien is trapped (there are now roads anymore from the current city),
he just stays there, he is not destroyed and it's not possible to move it to other city.
2. It's said that if there are two aliens in one city, they destroy each other and the city.
But it could happen that there are more than 2 aliens in the same city, and in this case 
I destroy all aliens and the city.
3. If city in the input file, all occurrences will be merged in one record (input data example:  
Foo north=Bar west=Baz   
Foo south=Qu-ux  
Result: Map will contain one city Foo with all three directions). However, it;s not checked 
whether directions for the particular city are not duplicated (for example, if there are two 
North directions for city Foo).  

Sample input data:  
Foo​ ​north=Bar​ ​west=Baz​ ​south=Qu-ux  
Bar​ ​south=Foo​ ​west=Bee

Sample output data:  
1. Some cities are left on the Map  
Bar has​ ​been​ ​destroyed​ ​by​ ​alien​s [0 1]!  
Foo west=Baz south=Qu-ux
2. No cities are left on the Map  
Bee has​ ​been​ ​destroyed​ ​by​ ​alien​s [2 3 4 6 0]!  
Foo has​ ​been​ ​destroyed​ ​by​ ​alien​s [5 14 13 15]!  
Qu-ux has​ ​been​ ​destroyed​ ​by​ ​alien​s [7 1 9 10 17]!  
Bar has​ ​been​ ​destroyed​ ​by​ ​alien​s [8 11 18 16 19 12]!  
