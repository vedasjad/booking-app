# Go Conference Booking Application

This is a simple Go Conference booking application designed to help users book tickets for a conference event. The application manages conference ticket availability, user input validation, booking, and email ticket delivery using goroutines.

## Overview

The Go Conference Booking Application provides the following features:

- Book tickets for the Go Conference event.
- Check ticket availability.
- Validate user input for name, email, and ticket quantity.
- Store user booking data.
- Send ticket confirmation emails asynchronously.

## Prerequisites

Before you can run this application, make sure you have Go installed on your system. You can download and install Go from the [official website](https://golang.org/dl/).

## Getting Started

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/vedasjad/booking-app.git
   ```

2. Change to the project directory:

   ```bash
   cd booking-app
   ```

3. Run the application:

   ```bash
   go run .
   ```

## Usage

1. Upon running the application, you will be greeted with information about the conference and available tickets.

2. Enter your first name, last name, email address, and the number of tickets you want to book.

3. The application will validate your input. It checks the length of your first and last name, the validity of the email address, and whether the number of tickets is within the available limit.

4. If your input is valid, it will subtract the booked tickets from the available tickets and store your booking information.

5. An asynchronous goroutine will simulate sending a confirmation email with ticket details to your provided email address after a 10-second delay.

6. The application will display the list of booked attendees' first names and the remaining available tickets.

7. If all tickets are booked, it will inform you that the conference is fully booked.

---

Thank you for using the Go Conference Booking Application!
