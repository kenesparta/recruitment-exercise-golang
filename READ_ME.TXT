Before Starting:
Copy this proyect into your personal git account and commit before making any changes. We want to see the history of all commits you do.

Proyect Description:
This appliction emulates a factory that assembles cars. This is a small factory with 5 spots only, meaning it could assemble up to 5 cars at the same time. 
If all spots are busy at a certain moment, no new cars can be done until one spot idles. Every single car part takes one second to be assembled so the 
total assemble time could take up to 7s. if you run this app now you will see a console message "Assembling vehicle..." every 7 seconds. Meaning every 7
seconds a new car is assembled, we should assemble 5 cars every 7 seconds instead.

Tasks:
This application is syncronous, make it work CONCURRENTLY! You can use goroutines, channels/buffered channels, waitgroups, mutexes, etc at your 
discression.
This project has several go files, you should focus on these 3 functions:
1. assemblyspot.AssembleVehicle
2. factory.StartAssemblingProcess
3. main.StartAssemblingProcess
4. Implement Unit Test (please use factory_test.go as base file)
Each function has a hint comment
IMPORTANT: have in mind racing conditions.

At the end of the assemble process of each vehicle, main.go should receive the vehicle and display its TestingLog and AssembleLog. Please have in mind do not wait for all
vehicles to be done to return them all to main, once each single vehicle is assembled send it over to main for log display right away
