# Trimmed Mean Calculator

This repository contains code that data scientists may leverage for calculating the trimmed mean of a slice of floats or of a slice of integers in Golang.  This trimmed mean calculation code can be found in this repository's main.go program. As you can see, the program's trimmedMean function accepts three inputs:
1. numbers - a slice of floats or integers for which you'd like to calculate a trimmed mean
2. lowerTrim - the proportion of numbers that should be removed from the lower end of the slice after sorting but before calculating the trimmed mean
3. upperTrim - the proportion of numbers that should be removed from the upper end of the slice after sorting but before calculating the trimmed mean

Data scientists may use this code to calculate the mean after trimming from either one end or from bothe ends of the sorted slice.  Furthermore, as we can see from the two example calculations provided within the main.go program, the data scientist has the option to provide one or two trimming parameters (lowerTrim and upperTrim).  If the data scientist provides two trimming parameters, then the program will trim the requested proportion of numbers from either end of the sorted slice. If the data scientist provides one trimming parameter, then the program will assume that the user wants the requested proportion of numbers trimmed from both the upper and lower ends of the sorted slice in a symmetric fashion.

To incorporate this Golang code into your own code, simply import the code into your own program and leverage its trimmedMean function. For an exmaple of how to do this, feel free to refer to this repository's trimmed_mean_test.go program, which leverages the this function for 7 unit test cases, which all pass. These passing test cases cover a range of possible scenarios, including: 
1. A scenario where we trim a slice of floats symmetrically
2. A scenario where we trim just the lower end of a slice of integers
3. A scenario where we trim just the upper end of a slice of floats
4. A scenario where we don't trim 0% of a slice of floats
5. A scenario where the slice is empty
6. A scenario where we trim 5% of the values from either end of a sorted slice of 100 integers
7. A scenario where we trim 5% of the values from either end of a sorted slice of 100 floats

Each of these unit tests passed, and the test outputs were written to the test_results_Golang.txt file.

To provide ourselves extra certainty that our program was functioning properly, we validated some of our unit tests by calculating the trimmed means for identical vectors of integers and numerics in R.  Specifically, we repeated unit tests 6 and 7 from above in R and printed the results to this repository's trimmedMeanOutputsR.txt file. As we see, the trimmed mean calculations in R produced identical answers to those produced by the trimmed mean calculaitons in Golang, so we can feel confident leveraging this Golang code in future programs.

I hope that this program is helpful for your future projects involving trimmed mean calculations!
