# Days 11-12: Giving the Owner of the Book Store Authority
The owner of the bookstore you are creating a backend for has requested that his account have an admin property and that a route, /earnings should only be accessible to him. Can you do this?

## Setup
- Navigate to the root of this repo.
- Run the command ```go run ./main.go``` to start the server.
- Visit the following url endpoints:
    |METHOD|DESCRIPTION|ENDPOINT|SAMPLE BODY|
    |------|-----------|--------|----|
    |POST  |Register user|http://127.0.0.1:3000/register|{"username":"username","password":"password"}|
    |POST  |Login user   |http://127.0.0.1:3000/login|{"username":"username","password":"password"}|
    |POST  |Refresh token   |http://127.0.0.1:3000/login|{"refresh_token": "refresh_token"}|
    |GET   |Show earnings for admin|http://127.0.0.1:3000/earnings|-|
    |GET   |Get all books added by logged-in user|http://127.0.0.1:3000/books|-|
    |GET   |Get a book added by logged-in user|http://127.0.0.1:3000/books/{id}|-|
    |POST  |Create a book for logged-in user|http://127.0.0.1:3000/books|{"author":"book_author","title":"book_title"}|
