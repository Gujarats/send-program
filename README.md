# Send Program

This example program is used for generating dummy data using php, and creating API to save data using Goroutine so the user Don't need to wait for the result.

## Prerequisite
You need all of this tech to run this project.

    - Mysql
    - Go
    - PHP (for generating data) 

## The Case
Lets say in our application there are some module or API that run really slow. In this case the module is `registermo` this is binary file for generating som random string.
For the sake of this case, lets say that we can't modify the `registermo` module. So we must find a way to make our API runs smoothly and response the request in milliseconds.

the `registermo` it self when you running it and give it som string it will give you a random string in this case is token in more than 2-3 seconds. The result will be saved in the table.
Imagine if one request have to wait for 2-3 seconds, what about several thousands of request in  seconds? the wait time will be enourmous and too long.

## Solutions 
We can create the API to use and save the data using goroutine when generating the token using `registermo` and save the data to our table.


## result
This is the load testing result using Vegeta : 

![result_load_testing_from_local](https://github.com/Gujarats/send-program/blob/master/load-testing/result.png)

Above test is using this command : 
```shell
echo "GET http://localhost:8080/send/mo?msisdn=60123456789&operatorid=3&shortcodeid=8&text=ON+GAMES" | vegeta attack -duration=10s -rate=2000 | sudo tee results_find_driver.bin | vegeta report
```
It means that we're creating 2000 request persecond in 10 seconds!. I can create more request to seveal thousands more like 10.000 request but my machine will eat rams and cpu to create this processs.

This is my spec of my laptop : 

 - Ubuntu 16.04 64bit
 - 8 Gb RAM
 - Intel core i5-6200U CPU @2.30Fhz x 4
