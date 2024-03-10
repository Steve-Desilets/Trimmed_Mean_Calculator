
# Function to log messages
logMessage <- function(message) {
  cat(message, "\n")
  # Write the log message to a log file
  log_file <- file("C:/Users/steve/OneDrive/Desktop/Github/Trimmed_Mean_Golang_Package/trimmedMeanOutputsR.txt", open="a")
  writeLines(paste(Sys.time(), message), log_file)
  close(log_file)
}

# Let's calculate the trimmed mean of a vector of integers
double_vector_1 <- rep(-1000, 5)
double_vector_2 <- 1:90
double_vector_3 <- rep(1000000, 5)

joint_double_vector <- c(double_vector_1, double_vector_2, double_vector_3)

joint_integer_vector <- as.integer(joint_double_vector) 

trimmed_mean_of_integers <- mean(x = joint_integer_vector, trim = 0.05)
logMessage(paste("Test Case One: One Hundred Integer Inputs"))
logMessage(paste("The program-calculated trimmed mean after trimming 5% of the values from either end of the sorted vector of numerics is: ", trimmed_mean_of_integers))
logMessage(paste("The expected correct trimmed mean value is 45.5."))



# Let's calculate the trimmed mean of a vector of numerics
double_vector_4 <- rep(-1000.54321, 5)
double_vector_5 <- c(0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0,
                     1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9, 2.0,
                     2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7, 2.8, 2.9, 3.0,
                     3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 3.9, 4.0,
                     4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 4.7, 4.8, 4.9, 5.0,
                     5.1, 5.2, 5.3, 5.4, 5.5, 5.6, 5.7, 5.8, 5.9, 6.0,
                     6.1, 6.2, 6.3, 6.4, 6.5, 6.6, 6.7, 6.8, 6.9, 7.0,
                     7.1, 7.2, 7.3, 7.4, 7.5, 7.6, 7.7, 7.8, 7.9, 8.0,
                     8.1, 8.2, 8.3, 8.4, 8.5, 8.6, 8.7, 8.8, 8.9, 9.0)
double_vector_6 <- rep(3141592.65358979, 5)

joint_double_vector_2 <- c(double_vector_4, double_vector_5, double_vector_6)

joint_numeric_vector <- as.numeric(joint_double_vector_2) 

trimmed_mean_of_numerics <- mean(x = joint_numeric_vector, trim = 0.05)


logMessage(paste("Test Case Two: One Hundred Float Inputs"))
logMessage(paste("The program-calculated trimmed mean after trimming 5% of the values from either end of the sorted vector of numerics is: ", trimmed_mean_of_numerics))
logMessage(paste("The expected correct trimmed mean value is 4.55."))

