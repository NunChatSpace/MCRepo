
# Overview
 - Backend
	- Controller 
		- Support API GET "/userinfo" to get transaction history
	 - MoonService
		 - Support API POST "/buy" to make transaction and MOON calculation
 - FrontEnd
	 - URL "localhost:3030/buy" is used to let you buy MOON 
	 - URL "localhost:3030/history" is used to show MOON coin transaction 

# Architecture
![Project Architecture](https://github.com/NunChatSpace/MCRepo/blob/main/ProjectArchitecture.png)

# API
 - GET "localhost:8079/userinfo"
	 - Input N/A
	 - Output
		 - When error Status is : 500
		 - When Success Status is 200
		```json
		  {
			"Status": 200,
			"Message": "Getting data is successfully",
			"DataLength": 1,
			"Data": [
			{
			"BuyDate": "2021-05-08 16:16:45",
			"Username": "User1620395604666",
			"THBT": 625.4275309208556,
			"MOON": 0.3040130799220151,
			"Rate": "1 MOON = 0.000486088419348 | 2057.238889462549196"
			}]
		}
 - POST "localhost:8080/buy"
	 - Input
		 - form-data
		```json
			 {
				 "Username": "text",
				 "BuyWith": 100,
				 "CurrentExchangeRate": 0.02,
				 "SlippageRateMin": 0.00,
				 "SlippageRateMax": 2
			 }
	- Output
		- When error Status is : 500
		- When success Status is : 200
		```json
		 {
			 "Status":200,
			 "Message":"Success",
			 "Data":
				 {
					 "BuyingStaus":"Success",
					 "MOON":1.8181818181818181,
					 "THBT":100
				}
		}

# Run Project
1. Run MongoDB
	- change directory to "MCRepo/Backend/MongoDB"
	- Execute command "docker-compose up"
2. Run ControllerService
	- change directory to "MCRepo/Backend/ControllerService"
	- Execute command "docker-compose up"
3. Run MoonService
	- change directory to "MCRepo/Backend/MoonService"
	- Execute command "docker-compose up"
4. Run Web service
	- change directory to "MCRepo/Frontend/moon-coin"
	- Execute command "docker-compose up"

# Testing
- Unit Test
	- ControllerService
		- Change directory to "MCRepo/Backend/ControllerService"
		- Execute "go test"
	- MoonService
		- Change directory to "MCRepo/Backend/MoonService"
		- Execute "go test"
- UI Test
	- Change directory to "MCRepo/API"
	- Execute command on Command line "Robot ." or "Robot {filename}.robot"
- API Test
	- Import file to Postman collection will got Collection name "MOONCoin_APITest"
	- Select API which you want to test
	- Press Send
	- At tab "Tests" on Response section you will get API test result