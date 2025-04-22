data <- c(10, 20, 30, 40, 50)
data
mean(data)




median(data)


get_mode <- function(v) {
  uniq_vals <- unique(v)
  uniq_vals[which.max(tabulate(match(v, uniq_vals)))]
}

get_mode(data)
