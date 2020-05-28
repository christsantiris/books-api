To Test API (For now)
to get all books: 
curl http://localhost:3000/api/v1/books

to get book by id
curl http://localhost:3000/api/v1/book/1

to create book
curl localhost:3000/api/v1/book -X POST -d '{"title": "LOR", "author": "Tolkein", "hrating": 10}' -H "Content-Type: application/json"

to delete book by id
curl -X DELETE http://localhost:3000/api/v1/book/1
