Task Description:

Imagine an island archipelago that consists of several disconnected islands. Within the islands, the pair of cities are connected through a unique highway. You can assume that you can reach every city from a city within the island. There can be multiple unique roads between cities. You are given two text files -

– city_population.txtDownload city_population.txt: contains a list of all cities and their population. The file data format is in <city name>: <population>. For example, City A: 1000 means City A has a population of 1000. You can assume each city has a unique name.  
–road_network.txtDownload road_network.txt contains the list of all highways in the island archipelago. The file data format is in <city name 1>: <city name 2>. For example, City A: City B. Means City A and City B are connected by a unique highway. You can consider the highways as bi-directional. 
Complete the following tasks - 

Task-1: Implement a class City that will include the following fields - [4 pt]

Name of the city 
The population of the city
List of cities that are connected to this particular city. 

Task-2: Read the text file and construct a graph of cities. Use the objects of a class City to model a city and the graph. Note that conceptually, this is similar to the adjacency list representation of the graph. 

Task-3:  Given the list of City objects, write a function to return the number of islands in the archipelago. Note that this function would require finding the number of connected components in the graph. 

Task-4:  Given the list of City objects, write a function that would return the population of each island in the island archipelago. Note that this function would require you to find the population of each connected component in the graph. 

Task-5: Given two City objects, write a function that would return the minimum number of unique highways you can take to reach from one city to another. Note that this function requires you to find the distance, i.e., the number of unique highways between two cities.

