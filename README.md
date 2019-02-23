# min-cut

// coursera Stanford Algorithms part1-Divide and Conquer, Sorting and Searching, and Randomized Algorithms Assignment 4

Question:
The file contains the adjacency list representation of a simple undirected graph. There are 200 vertices labeled 1 to 200. The first column in the file represents the vertex label, and the particular row (other entries except the first column) tells all the vertices that the vertex is adjacent to. So for example, the  6th row looks like : "6   155 56 52 12 This just means that the vertex with label 6 is adjacent to (i.e., shares an edge with) the vertices with labels 155,56,52,120,......,etc  Your task is to code up and run the randomized contraction algorithm for the min cut problem and use it on the above graph to compute the min cut. (HINT: Note that you'll have to figure out an implementation of edge contractions. Initially, you might want to do this naively, creating a new graph from the old every time there's an edge contraction. But you should also think about more efficient implementations.) (WARNING: As per the video lectures, please make sure to run the algorithm many times with different random seeds, and remember the smallest cut that you ever find.) Write your numeric answer in the space provided. So e.g., if your answer is 5, just type 5 in the space provided.



Solution:
The solution contains 4 parts.
1. Read file, then save vertices and there adjacent vertices as two-dimensional array (I transfer the file from .txt to .xlsx, because it's easier to read every vertice from the excel file.).
2. Select two vertices (u,v) randomly.
3. Connect u with vertices which were connecting with v, then delete v.
4. Delete all self loops.
5. Count minimun cuts many times to get the minimum value.
