# Types and schemas
There's a few things we need to straighten out first:
1. Data types (should mirror Anki's types: Deck, Model, Media, Note, Card)
2. Schemas (should mirror Anki's schemas with sensible defaults that can be overridden)

Some initial observations:
* We must directly represent each _type_ in YAML
* We must represent those types using the correct _schema_
* We may omit some 'default' fields when doing so
* We may nest one type inside another in YAML (eg: notes defined under decks)

# Basic example JSON for API request/response

Request:
```
{
    "action": "${ACTION}",
    "version": 6,
    "params": {
        "${PARAM}": <data structure>
    }
}
```

Response:
```
{
    "result": ${RESULT},
    "error": ${ERROR}
}
```

# Example data schema (likely wrong, we should iterate)

## Models (separate from decks, notes, cards, media)
```
"modelName": "newModelName",
"inOrderFields": ["Field1", "Field2", "Field3"],
"css": ".card {\n font-family: arial;\n font-size: 20px;\n text-align: center;\n color: black;\n background-color: white;\n}\n",
"isCloze": false,
"cardTemplates": [
    {
        "Name": "My Card 1",
        "Front": "Front html {{Field1}}",
        "Back": "Back html  {{Field2}}"
    }
]
```

## Decks
```
"deck": "Golang::Stdlib::fmt"
```

## Notes
We should start with `canAddNotesWithErrorDetail`, which accepts an array of notes an tells you whether or not each one can be added...
```
"notes": [
    {
        "deckName": "Default",
        "modelName": "Basic",
        "fields": {
            "Front": "front content",
            "Back": "back content"
        },
        "tags": [
            "yomichan"
        ]
    },
    {
        "deckName": "Default",
        "modelName": "Basic",
        "fields": {
            "Front": "front content 2",
            "Back": "back content 2"
        },
        "tags": [
            "yomichan"
        ]
    }
]
```

This tells you whether or not you can add the notes. For those that can be added, use addNote. For those that cannot, parse the error and decide how to proceed. Eg: if it's a duplicate, use updateNote. (Now, the API is a bit stupid because canAddNotesWithErrorDetail does not return the note ID, and you need that for any of the note update calls (updateNoteFields, updateNoteTags, updateNote, updateNoteModel):

When updating a note, since we have the complete note config in YAML it is better to use updateNoteModel which takes the modelName, fields and tags (ie: the whole shebang)
```
"note": {
    "id": 1514547547030,
    "modelName": "NewModel",
    "fields": {
        "Front": "new front content",
        "Back": "new back content"
    },
    "tags": ["new", "tags"]
}
```

It's actually really stupid: to update a note, you need the note ID. But you don't get it in any of the useful outputs. For exampel, if you use `canAddNotesWithErrorDetail` it doesn't tell you the note ID. You're just told that these attributes clash with an existing note. You then have to get all notes (including ID) and find out which one has the same attributes in order to find the fucking ID. _Then_ you can use the ID to update the note attributes. 

