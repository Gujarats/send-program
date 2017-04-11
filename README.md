#Send Program
This example program is used for generating dummy data using php, and creating API to save data using Goroutine so the user Don't need to wait for the result.

##Prerequisite
You need all of this tech to run this project.

    - Mysql
    - Go
    - PHP (for generating data) 

## The Case
Lets say in our application there are some module or API that run really slow. In this case the module is `generateToken` this is binary file for generating som random string.
For the sake of this case, lets say that we can't modify the `generateToken` module. So we must find a way to make our API runs smoothly and response the request in milliseconds.

the `generateToken` it self when you running it and give it som string it will give you a random string in this case is token in more than 2-3 seconds. The result will be saved in the table.
Imagine if one request have to wait for 2-3 seconds, what about several thousands of request in  seconds? the wait time will be enourmous and too long.

## Solutions 
We can create the API to use and save the data using goroutine when generating the token using `generateToken` and save the data to our table.


## result
This is the load testing result using Vegeta : 

![result_load_testing_from_local](https://github.com/Gujarats/send-program/blob/master/load-testing/result.png)

