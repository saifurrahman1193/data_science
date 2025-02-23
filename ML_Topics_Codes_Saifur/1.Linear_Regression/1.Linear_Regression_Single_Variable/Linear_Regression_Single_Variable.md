![Alt text](assets/1.png)
![Alt text](assets/2.png)
![Alt text](assets/3.png)
![Alt text](assets/4.png)
![Alt text](assets/5.png)
![Alt text](assets/6.png)
![Alt text](assets/6_2.png)



The **least squares method** is widely used for finding the **line of best fit** in linear regression. The idea is to minimize the sum of the squared differences (errors) between the observed data points and the predicted values on the line. Here's a breakdown of the process:

1. **The Goal**: The goal of the line of best fit is to find the straight line that best represents the data in a way that minimizes the vertical distance between the points and the line. These distances are also referred to as "errors" or "residuals."

2. **Why Least Squares**: Instead of just minimizing the raw distance between the data points and the line, **least squares** focuses on minimizing the square of these distances. Squaring the distances ensures that negative and positive errors don't cancel each other out, and it penalizes larger errors more than smaller ones.

3. **Visualizing It**: 
   - **Diagram 1**: A line that is too steep or too flat could have large vertical distances from many points, making it less accurate.
   - **Diagram 2**: A line that looks close but doesn't minimize the sum of the squared distances could also be suboptimal.
   - **Diagram 3**: The line that minimizes the sum of squared distances from all the points is considered the best fit.

### Why Diagram 3 is Best:
- The line in **Diagram 3** has been **calculated using the least squares method**, ensuring it has the smallest possible sum of squared differences (errors).
- This method guarantees the line is the most accurate, even if other lines might visually look close.

### Key Takeaways:
- The line of best fit is the one that minimizes the **sum of squared residuals**.
- **Diagram 3** is the best because it follows this principle, not just a subjective visual assessment of "how close" it looks.


![Alt text](assets/7.png)
![Alt text](assets/8.png)
![Alt text](assets/9.png)
![Alt text](assets/10.png)
![Alt text](assets/11.png)
![Alt text](assets/12.png)
![Alt text](assets/13.png)
![Alt text](assets/14.png)
![Alt text](assets/15.png)
![Alt text](assets/16.png)
![Alt text](assets/17.png)
![Alt text](assets/18.png)
![Alt text](assets/19.png)
![Alt text](assets/20.png)
![Alt text](assets/21.png)
![Alt text](assets/22.png)
![Alt text](assets/23.png)
![Alt text](assets/24.png)
![Alt text](assets/25.png)
![Alt text](assets/26.png)
![Alt text](assets/27.png)
![Alt text](assets/28.png)
![Alt text](assets/29.png)
![Alt text](assets/30.png)
![Alt text](assets/31.png)
![Alt text](assets/32.png)
![Alt text](assets/34.png)
![Alt text](assets/35.png)
![Alt text](assets/36.png)
![Alt text](assets/37.png)










## Intercept/Constant: 
- The baseline sales when advertising spend is zero.
- Intercept: The value of the dependent variable when all independent variables are zero.

## Slope/Gradient/rate of change/Coefficient: 
- The increase in sales for each unit increase in advertising spend.





## 📌 Business Interpretation of Intercept in Regression

In business terms, an **intercept** represents a **baseline or starting point** for the dependent variable when all independent variables are set to zero. It helps quantify the impact of changes in independent variables from this baseline.

## 📊 **Understanding the Intercept in Business**
- The **intercept** reflects the portion of the dependent variable that is **not influenced** by independent variables.
- It serves as the **starting point** for evaluating the effects of independent variables on the dependent variable.

### **Examples**
- **Sales Prediction Model:** The intercept might represent **expected sales when all marketing efforts (predictors) are at zero**.
- **Finance:** The intercept can represent **fixed or overhead costs** that are incurred regardless of activity levels.




---

## **📈 Types of Intercepts in Regression**
### **1️⃣ Positive Intercept**
- If the intercept is **positive**, the predicted dependent variable (**Y**) when the independent variable (**X**) is zero is **positive**.
- 🔹 The regression line crosses the **y-axis above zero**.
- Positive Intercept: Indicates a positive baseline value.

### **2️⃣ Negative Intercept**
- If the intercept is **negative**, the predicted value of **Y** when **X = 0** is **negative**.
- 🔹 The regression line crosses the **y-axis below zero**.
- Negative Intercept: Indicates a negative baseline value.

### **3️⃣ Zero Intercept**
- If the intercept is **zero**, the regression line **passes through the origin (0,0)**.
- 🔹 This means there is **no additional constant term** in the equation.
- 🔹 This situation is **rare and highly theoretical**.
- Zero Intercept: Indicates no baseline value, often theoretical.








---

## **💡 Summary**
- The **intercept** shows the expected value of **Y** when **X = 0**.
- It helps understand the **baseline value** in business models.
- **Positive, negative, and zero intercepts** give different interpretations based on the context.




# Resources
- https://www.youtube.com/watch?v=8jazNUpO3lQ&list=PLeo1K3hjS3uvCeTYTeyfe0-rN5r8zn9rw&index=2
- https://github.com/codebasics/py/tree/master/ML