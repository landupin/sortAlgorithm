# sortAlgorithm
this repo will show you some implementions of well known sorting algorithms
for license see the LICENSE file (GNU General Public License v3.0)

# Mergesort

### how it works:
**1. divide**
	divide the array into subarrays by cutting them in two halfes until the subarrays include only one object

**2. conquer**
	put the subarrays into the rootarray by sorting them
	
	*	compare the first left object with the first right object
	*	put the smaller object into the rootarray
	*	continue comparing with the next object, but the not changed side stays with its object
	*	if one side is empty the objets from the other side are all moved to the rootarray without comparision
	
# Selectionsort

### how it works:

**iterate through the array:**
	*	search the lowest object
	*	swap it with the object of iteration

# Insertionsort

### how it works:

**iterate through array starting with the second object**
	1.	save the object of iteration
	2.	go through the array backwards until the next object isn't smaller than the saved one or you are inspecting the first object of the array
			overwrite the object of iteration with index of iteration with the object of
	3.	overwrite this object with the saved one
		
# Binarytreesort (in work yet)

### how it works:
**1. inserting the numbers out of the array into a binary tree**

	*	number of the array as tree root
	*	numbers smaller than the node you are inspecting into the left subtree
	*	numbers higher than the node you are inspecting into the right subtree
**2. printing the binary tree inorder:**

	1.	lowest node of the left side from the root
	2.	upper node
	3.	right subtree
	4.	start again with root = current inspected node

# Swapsort (in work yet)

### how it works:
1. start with the first element of the array
2. count all numbers lower than the inspected element
3. swap the inspected element with the element of the array with index = number of lower numbers than the inspected number
4. if index of the inspected number = number of lower numbers than the inspected one: **go to the next number** 
5. **if index of the inspected number = size of the array (number of numbers):** END of sorting **else:** repeat from step 2

The problem with this algorithm is, that no number can appear twice, because in this case there is no determination.


