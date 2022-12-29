# HotelManagement-Golang
This project contains basic API's to demonstrate the working of a HotelManagement System.

# 1) Hotel API's 
Hotel API's include the following features -

## 1.1) Adding a room to a Hotel.
A room in the hotel has following fields

RoomId      to give a unique Id to the rooms

RoomNo      Specifies the Room Number

FloorNo     Specifies the floor where the room is

RoomService Specifies whether room service is avilable

RoomPrice   The price of the room


To add a room use the following-

#### Method "POST"

/api/hotel/createRoom - 

It accepts a JSON
{
     "roomId":1,
    "roomNo":10,
    "floorNo":2,
    "roomService":"Yes"
    "roomPrice": 6000
}

## 1.2) Get all Rooms in a hotel 

#### Method "GET"

/api/hotel/getAllRooms

## 1.3) Get rooms based on a some price

#### Method "GET"

/api/hotel/queryRooms?Price=6000

This api gets the list of all rooms with price less than or equal to 6000

# 2) Customer API

A Customer has following fields

CustId      to give a unique Id to the customer

Name        Specifies the Name of the customer

Contact     Specifies the contact number of the customer

Email       Specifies the Email of the Customer


## 2.1) To add a cutomer use the following-

#### Method "POST"

/api/customer/addCustomer - 

It accepts a JSON
{
    "custId":1,
    "name":Aman,
    "contact":123456789,
    "mail":"aman@gmail.com"
}

## 2.2) Get list of all customer

#### Method "GET"

/api/customer/getInfo

This api wil give the list of all customers


## 2.3) Fetch Customer by Name

#### Method "GET"

/api/customer/fetchCustomer?name=Aman

This api will fetch the customer based on the Query param of "name"


# Booking API

The booking has the following fields

RoomId     to give a unique Id to the rooms

CustomerId to give a unique Id to the customer

This API allows the user to book a room in the hotel

#### Method "POST"

/api/booking/bookRoom?roomId=1&custId=1&name=Aman&email=aman@gmail.com&contact=123456789

the params will include-

RoomId of the room that the user wants to book
CustId of the customer
Name of the customer
Email of the customer
Contact of the customer

Work flow of booking a room

1) It first checks if the room is a valid room in the hotel by looking at the RoomId
2) Checks if the Room is not already Booked
3) If the above are true then it stores the customer in the customers table and also makes an entry in the bookings table

