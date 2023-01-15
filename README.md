# HOW TO RUN AND TEST IT 
[work in progress]- next step: database connection

## Get all articles
Request: 

    curl --location --request GET 'localhost:10000/articles'

Response:

    [
        {
            "id": 1,
            "title": "Title one",
            "desc": "Desc one",
            "content": "Content one"
        },
        {
            "id": 2,
            "title": "Title two",
            "desc": "Desc two",
            "content": "Content Two"
        },
        {
            "id": 3,
            "title": "Title three",
            "desc": "Desc three",
            "content": "Content three"
        },
        {
            "id": 4,
            "title": "Title four",
            "desc": "Desc four",
            "content": "Content four"
        }
    ]

## Get article by Id

Request:

    curl --location --request GET 'localhost:10000/articles/1'

Response:

    {
        "id": 1,
        "title": "Title one",
        "desc": "Desc one",
        "content": "Content one"
    }

## Create article
Request:

    curl --location --request POST 'localhost:10000/articles' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "id": 5,
        "title": "Title five",
        "desc": "Desc five",
        "content": "Content five"
    }'

Response:

    {
        "id": 5,
        "title": "Title five",
        "desc": "Desc five",
        "content": "Content five"
    }

## Delete article by id
Request:

    curl --location --request DELETE 'localhost:10000/articles/5'

## Update article by id

Request:

    curl --location --request PUT 'localhost:10000/articles/1' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "title": "Lord of the Rings is the best movie ever",
        "desc": "Desc one",
        "content": "Content one"
    }'

Response:

    {
        "id": 1,
        "title": "Lord of the Rings is the best movie ever",
        "desc": "Desc one",
        "content": "Content one"
    }