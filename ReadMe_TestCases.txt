===========================================================
Below are the test cases executed in GoLang
Input given by user in console
===========================================================

#### Test Case 1: (Which is given for discussion)
	Enter building array of tuples - example: ((0,4),(1,4)) ::
	((0,4),(1,4)),((1,1),(3,1)),((2,2),(4,2)),((3,3),(5,3)),((6,1),(8,1))

	Tuple array:  [0 4 1 4 1 1 3 1 2 2 4 2 3 3 5 3 6 1 8 1] 
	Tuple length:  20

	Total number of building are 5 and its coordinates are:
	{0 4}{1 4}
	{1 1}{3 1}
	{2 2}{4 2}
	{3 3}{5 3}
	{6 1}{8 1}

	Output:
	[{0 4} {1 4} {1 1} {2 1} {2 2} {3 2} {3 3} {5 3} {5 0} {6 0} {6 1} {8 1} {8 0}]

	Process finished with exit code 0

===========================================================
### Test cases executed other than example given
#### Test Case 2:
	Enter building array of tuples - example: ((0,4),(1,4)) ::
	((2,10),(9,10)),((3,15),(6,15)),((5,12),(12,12)),((13,10),(16,10)),((15,5),(17,5))

	Tuple array:  [2 10 9 10 3 15 6 15 5 12 12 12 13 10 16 10 15 5 17 5] 
	Tuple length:  20

	Total number of building are 5 and its coordinates are:
	{2 10}{9 10}
	{3 15}{6 15}
	{5 12}{12 12}
	{13 10}{16 10}
	{15 5}{17 5}

	Output:
	[{2 0} {2 10} {3 10} {3 15} {6 15} {6 12} {12 12} {12 0} {13 0} {13 10} {16 10} {16 5} {17 5} {17 0}]

	Process finished with exit code 0

===========================================================
#### Test Case 3 : (Corner Case - 2 building start from same coordinate)
	Enter building array of tuples - example: ((0,4),(1,4)) ::
	((3,8),(8,8)),((3,10),(10,10))

	Tuple array:  [3 8 8 8 3 10 10 10] 
	Tuple length:  8

	Total number of building are 2 and its coordinates are:
	{3 8}{8 8}
	{3 10}{10 10}

	Output:
	[{3 0} {3 10} {10 10} {10 0}]

	Process finished with exit code 0

===========================================================
#### Test Case 4 : (Corner Case - 2 building end at from same coordinate)
	Enter building array of tuples - example: ((0,4),(1,4)) ::
	((2,10),(10,10)),((4,8),(10,8))

	Tuple array:  [2 10 10 10 4 8 10 8] 
	Tuple length:  8

	Total number of building are 2 and its coordinates are:
	{2 10}{10 10}
	{4 8}{10 8}

	Output:
	[{2 0} {2 10} {10 10} {10 0}]

	Process finished with exit code 0

===========================================================
#### Test Case 5 : (If each building Y position is not matching or X lying under same position)
Enter building array of tuples - example: ((0,4),(1,4)) ::
((1,4),(2,6))

Tuple array:  [1 4 2 6] 
Tuple length:  4
Invalid rectangle coordinates.
Builing count 0
Enter coordinates again

===========================================================
#### Test Case 6 : (If inputs less than 4)
	Enter building array of tuples - example: ((0,4),(1,4)) ::
	((1,4))

	Tuple array:  [1 4] 
	Tuple length:  2
	Less than 4 inputs. Enter building coordinates again
	
===========================================================
#### Test Case 7 : (If inputs are not balanced then only valid building coordinates will be taken)
	Enter building array of tuples - example: ((0,4),(1,4)) ::
	((0,4),(1,4)),((1,1))

	Tuple array:  [0 4 1 4 1 1] 
	Tuple length:  6
	Input points are not balanced and buildings considered are 1

	Total number of building are 1 and its coordinates are:
	{0 4}{1 4}

	Output:
	[{0 4} {1 4} {1 0}]

	Process finished with exit code 0
===========================================================
