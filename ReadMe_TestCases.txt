===========================================================
Below are the test cases executed in GoLang
Input given by user in console
===========================================================

#### Test Case 1: (Which is given for discussion)
	Enter building count of the city: 5
	Building coordinates has to be given by user
	Enter building 1 points as below
	 X1:0
	 Y1:4
	 X2:1
	 Y2:4
	Enter building 2 points as below
	 X1:1
	 Y1:1
	 X2:3
	 Y2:1
	Enter building 3 points as below
	 X1:2
	 Y1:2
	 X2:4
	 Y2:2
	Enter building 4 points as below
	 X1:3
	 Y1:3
	 X2:5
	 Y2:3
	Enter building 5 points as below
	 X1:6
	 Y1:1
	 X2:8
	 Y2:1

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
	Enter building count of the city: 5
	Building coordinates has to be given by user
	Enter building 1 points as below
	 X1:2
	 Y1:10
	 X2:9
	 Y2:10
	Enter building 2 points as below
	 X1:3
	 Y1:15
	 X2:6
	 Y2:15
	Enter building 3 points as below
	 X1:5
	 Y1:12
	 X2:12
	 Y2:12
	Enter building 4 points as below
	 X1:13
	 Y1:10
	 X2:16
	 Y2:10
	Enter building 5 points as below
	 X1:15
	 Y1:5
	 X2:17
	 Y2:5

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
	Enter building count of the city: 2
	Building coordinates has to be given by user
	Enter building 1 points as below
	 X1:3
	 Y1:8
	 X2:8
	 Y2:8
	Enter building 2 points as below
	 X1:3
	 Y1:10
	 X2:10
	 Y2:10

	Total number of building are 2 and its coordinates are:
	{3 8}{8 8}
	{3 10}{10 10}

	Output:
	[{3 0} {3 10} {10 10} {10 0}]

	Process finished with exit code 0
===========================================================

#### Test Case 4 : (Corner Case - 2 building end at from same coordinate)
	Enter building count of the city: 2
	Building coordinates has to be given by user
	Enter building 1 points as below
	 X1:2
	 Y1:10
	 X2:10
	 Y2:10
	Enter building 2 points as below
	 X1:4
	 Y1:8
	 X2:10
	 Y2:8

	Total number of building are 2 and its coordinates are:
	{2 10}{10 10}
	{4 8}{10 8}

	Output:
	[{2 0} {2 10} {10 10} {10 0}]

	Process finished with exit code 0
===========================================================

#### Test Case 5 : (If building count is not positive integer)
	Enter building count of the city: 0
	Building count is not valid. Enter again
===========================================================

#### Test Case 6 : (If each building Y position is not matching or X lying under same position)
	Enter building count of the city: 1
	Building coordinates has to be given by user
	Enter building 1 points as below
	 X1:1
	 Y1:4
	 X2:2
	 Y2:6
	Invalid rectangle coordinates. Enter again
	Enter building 1 points as below
	 X1:
===========================================================
