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

This tells you whether or not you can add the notes. For those that can be added, use addNote. For those that cannot, parse the error and decide how to proceed. Eg: if it's OK, add the note but if it's a duplicate, update the note using one of the update calls. There are several calls that can do specific things, and these are stupidly named, because of what they actually do:
- `updateNoteFields` only updates fields
- `updateNoteTags` only updates tags
- `updateNoteModel` updates the entire thing, including fields, tags and model
- `updateNote` updates the note fields and tags but not the model

That's not very intuitive. We will _only_ use updateNoteModel and fling everything we have at it. _Hopefully_ it won't fail if the tags are empty:
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

This is where it gets extremely stupid: to update a note, you need the note ID. But you don't get the note ID from `canAddNotesWithErrorDetail`, so for a particular set of parameters in a given note you'll be told "Nuh uh, that's a duplicate", but won't tell you the ID of the note that it duplicates. So you need to ask for ALL of the notes and check which ID matches _all_ the parameters you're trying to create a note for. Fucking _daft_. 

It's kinda moot anyway because if you fundamentally change a note, your new YAML no longer matches up with whatever you previously wrote. Right? So we have two choices:
1. Delete any note that isn't currently represented in YAML and add _all_ notes in YAML
2. Update the YAML with the note IDs after pushing (is that a thing? Can we? Should we then just ignore the ID with requests that can't accept it?)
3. ??? I don't know ???
