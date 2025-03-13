# All Supported Actions
This yaml2ankiconnect project will support the following types of actions:
- Card Actions
- Deck Actions
- Graphical Actions
- Media Actions
- Miscellaneous Actions
- Model Actions
- Note Actions

# Card Actions

## findCards

Returns an array of card IDs for a given query. Functionally identical to guiBrowse but doesn't use the GUI for better performance.
Sample request:

```
{
    "action": "findCards",
    "version": 6,
    "params": {
        "query": "deck:current"
    }
}
```

Sample result:

```
{
    "result": [1494723142483, 1494703460437, 1494703479525],
    "error": null
}
```

## cardsToNotes

Returns an unordered array of note IDs for the given card IDs. For cards with the same note, the ID is only given once in the array.
Sample request:

```
{
    "action": "cardsToNotes",
    "version": 6,
    "params": {
        "cards": [1502098034045, 1502098034048, 1502298033753]
    }
}
```

Sample result:

```
{
    "result": [1502098029797, 1502298025183],
    "error": null
}
```

## cardsModTime

Returns a list of objects containings for each card ID the modification time. This function is about 15 times faster than executing cardsInfo.
Sample request:

```
{
    "action": "cardsModTime",
    "version": 6,
    "params": {
        "cards": [1498938915662, 1502098034048]
    }
}
```

Sample result:

```
{
    "result": [
        {
            "cardId": 1498938915662,
            "mod": 1629454092
        }
    ],
    "error": null
}
```

## cardsInfo

Returns a list of objects containing for each card ID the card fields, front and back sides including CSS, note type, the note that the card belongs to, and deck name, last modification timestamp as well as ease and interval.
Sample request:

```
{
    "action": "cardsInfo",
    "version": 6,
    "params": {
        "cards": [1498938915662, 1502098034048]
    }
}
```

Sample result:

```
{
    "result": [
        {
            "answer": "back content",
            "question": "front content",
            "deckName": "Default",
            "modelName": "Basic",
            "fieldOrder": 1,
            "fields": {
                "Front": {"value": "front content", "order": 0},
                "Back": {"value": "back content", "order": 1}
            },
            "css":"p {font-family:Arial;}",
            "cardId": 1498938915662,
            "interval": 16,
            "note":1502298033753,
            "ord": 1,
            "type": 0,
            "queue": 0,
            "due": 1,
            "reps": 1,
            "lapses": 0,
            "left": 6,
            "mod": 1629454092
        },
        {
            "answer": "back content",
            "question": "front content",
            "deckName": "Default",
            "modelName": "Basic",
            "fieldOrder": 0,
            "fields": {
                "Front": {"value": "front content", "order": 0},
                "Back": {"value": "back content", "order": 1}
            },
            "css":"p {font-family:Arial;}",
            "cardId": 1502098034048,
            "interval": 23,
            "note":1502298033753,
            "ord": 1,
            "type": 0,
            "queue": 0,
            "due": 1,
            "reps": 1,
            "lapses": 0,
            "left": 6
        }
    ],
    "error": null
}
```

## changeDeck

Moves cards with the given IDs to a different deck, creating the deck if it doesn't exist yet.
Sample request:

```
{
    "action": "changeDeck",
    "version": 6,
    "params": {
        "cards": [1502098034045, 1502098034048, 1502298033753],
        "deck": "Japanese::JLPT N3"
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

# Deck Actions
## createDeck

Create a new empty deck. Will not overwrite a deck that exists with the same name.
Sample request:

```
{
    "action": "createDeck",
    "version": 6,
    "params": {
        "deck": "Japanese::Tokyo"
    }
}
```

Sample result:

```
{
    "result": 1519323742721,
    "error": null
}
```

## deckNames

Gets the complete list of deck names for the current user.
Sample request:

```
{
    "action": "deckNames",
    "version": 6
}
```

Sample result:

```
{
    "result": ["Default"],
    "error": null
}
```

## deckNamesAndIds

Gets the complete list of deck names and their respective IDs for the current user.
Sample request:

```
{
    "action": "deckNamesAndIds",
    "version": 6
}
```

Sample result:

```
{
    "result": {"Default": 1},
    "error": null
}
```

## getDecks

Accepts an array of card IDs and returns an object with each deck name as a key, and its value an array of the given cards which belong to it.
Sample request:

```
{
    "action": "getDecks",
    "version": 6,
    "params": {
        "cards": [1502298036657, 1502298033753, 1502032366472]
    }
}
```

Sample result:

```
{
    "result": {
        "Default": [1502032366472],
        "Japanese::JLPT N3": [1502298036657, 1502298033753]
    },
    "error": null
}
```

## deleteDecks

Deletes decks with the given names. The argument cardsToo must be specified and set to true.
Sample request:

```
{
    "action": "deleteDecks",
    "version": 6,
    "params": {
        "decks": ["Japanese::JLPT N5", "Easy Spanish"],
        "cardsToo": true
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## getDeckConfig

Gets the configuration group object for the given deck.
Sample request:

```
{
    "action": "getDeckConfig",
    "version": 6,
    "params": {
        "deck": "Default"
    }
}
```

Sample result:

```
{
    "result": {
        "lapse": {
            "leechFails": 8,
            "delays": [10],
            "minInt": 1,
            "leechAction": 0,
            "mult": 0
        },
        "dyn": false,
        "autoplay": true,
        "mod": 1502970872,
        "id": 1,
        "maxTaken": 60,
        "new": {
            "bury": true,
            "order": 1,
            "initialFactor": 2500,
            "perDay": 20,
            "delays": [1, 10],
            "separate": true,
            "ints": [1, 4, 7]
        },
        "name": "Default",
        "rev": {
            "bury": true,
            "ivlFct": 1,
            "ease4": 1.3,
            "maxIvl": 36500,
            "perDay": 100,
            "minSpace": 1,
            "fuzz": 0.05
        },
        "timer": 0,
        "replayq": true,
        "usn": -1
    },
    "error": null
}
```

## saveDeckConfig

Saves the given configuration group, returning true on success or false if the ID of the configuration group is invalid (such as when it does not exist).
Sample request:

```
{
    "action": "saveDeckConfig",
    "version": 6,
    "params": {
        "config": {
            "lapse": {
                "leechFails": 8,
                "delays": [10],
                "minInt": 1,
                "leechAction": 0,
                "mult": 0
            },
            "dyn": false,
            "autoplay": true,
            "mod": 1502970872,
            "id": 1,
            "maxTaken": 60,
            "new": {
                "bury": true,
                "order": 1,
                "initialFactor": 2500,
                "perDay": 20,
                "delays": [1, 10],
                "separate": true,
                "ints": [1, 4, 7]
            },
            "name": "Default",
            "rev": {
                "bury": true,
                "ivlFct": 1,
                "ease4": 1.3,
                "maxIvl": 36500,
                "perDay": 100,
                "minSpace": 1,
                "fuzz": 0.05
            },
            "timer": 0,
            "replayq": true,
            "usn": -1
        }
    }
}
```

Sample result:

```
{
    "result": true,
    "error": null
}
```

## setDeckConfigId

Changes the configuration group for the given decks to the one with the given ID. Returns true on success or false if the given configuration group or any of the given decks do not exist.
Sample request:

```
{
    "action": "setDeckConfigId",
    "version": 6,
    "params": {
        "decks": ["Default"],
        "configId": 1
    }
}
```

Sample result:

```
{
    "result": true,
    "error": null
}
```

## cloneDeckConfigId

Creates a new configuration group with the given name, cloning from the group with the given ID, or from the default group if this is unspecified. Returns the ID of the new configuration group, or false if the specified group to clone from does not exist.
Sample request:

```
{
    "action": "cloneDeckConfigId",
    "version": 6,
    "params": {
        "name": "Copy of Default",
        "cloneFrom": 1
    }
}
```

Sample result:

```
{
    "result": 1502972374573,
    "error": null
}
```

## removeDeckConfigId

Removes the configuration group with the given ID, returning true if successful, or false if attempting to remove either the default configuration group (ID = 1) or a configuration group that does not exist.
Sample request:

```
{
    "action": "removeDeckConfigId",
    "version": 6,
    "params": {
        "configId": 1502972374573
    }
}
```

Sample result:

```
{
    "result": true,
    "error": null
}
```

## getDeckStats

Gets statistics such as total cards and cards due for the given decks.
Sample request:

```
{
    "action": "getDeckStats",
    "version": 6,
    "params": {
        "decks": ["Japanese::JLPT N5", "Easy Spanish"]
    }
}
```

Sample result:

```
{
    "result": {
        "1651445861967": {
            "deck_id": 1651445861967,
            "name": "Japanese::JLPT N5",
            "new_count": 20,
            "learn_count": 0,
            "review_count": 0,
            "total_in_deck": 1506
        },
        "1651445861960": {
            "deck_id": 1651445861960,
            "name": "Easy Spanish",
            "new_count": 26,
            "learn_count": 10,
            "review_count": 5,
            "total_in_deck": 852
        }
    },
    "error": null
}
```

# Media Actions
## storeMediaFile

Stores a file with the specified base64-encoded contents inside the media folder. Alternatively you can specify a absolute file path, or a url from where the file shell be downloaded. If more than one of data, path and url are provided, the data field will be used first, then path, and finally url. To prevent Anki from removing files not used by any cards (e.g. for configuration files), prefix the filename with an underscore. These files are still synchronized to AnkiWeb. Any existing file with the same name is deleted by default. Set deleteExisting to false to prevent that by letting Anki give the new file a non-conflicting name.
Sample request (relative path):

```
{
    "action": "storeMediaFile",
    "version": 6,
    "params": {
        "filename": "_hello.txt",
        "data": "SGVsbG8sIHdvcmxkIQ=="
    }
}
```

Content of _hello.txt:

Hello world!

Sample result (relative path):

```
{
    "result": "_hello.txt",
    "error": null
}
```

Sample request (absolute path):

```
{
    "action": "storeMediaFile",
    "version": 6,
    "params": {
        "filename": "_hello.txt",
        "path": "/path/to/file"
    }
}
```

Sample result (absolute path):

```
{
    "result": "_hello.txt",
    "error": null
}
```

Sample request (url):

```
{
    "action": "storeMediaFile",
    "version": 6,
    "params": {
        "filename": "_hello.txt",
        "url": "https://url.to.file"
    }
}
```

Sample result (url):

```
{
    "result": "_hello.txt",
    "error": null
}
```

## retrieveMediaFile

Retrieves the base64-encoded contents of the specified file, returning false if the file does not exist.
Sample request:

```
{
    "action": "retrieveMediaFile",
    "version": 6,
    "params": {
        "filename": "_hello.txt"
    }
}
```

Sample result:

```
{
    "result": "SGVsbG8sIHdvcmxkIQ==",
    "error": null
}
```

## getMediaFilesNames

Gets the names of media files matched the pattern. Returning all names by default.
Sample request:

```
{
    "action": "getMediaFilesNames",
    "version": 6,
    "params": {
        "pattern": "_hell*.txt"
    }
}
```

Sample result:

```
{
    "result": ["_hello.txt"],
    "error": null
}
```

## getMediaDirPath

Gets the full path to the collection.media folder of the currently opened profile.
Sample request:

```
{
    "action": "getMediaDirPath",
    "version": 6
}
```

Sample result:

```
{
    "result": "/home/user/.local/share/Anki2/Main/collection.media",
    "error": null
}
```

## deleteMediaFile

Deletes the specified file inside the media folder.
Sample request:

```
{
    "action": "deleteMediaFile",
    "version": 6,
    "params": {
        "filename": "_hello.txt"
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

# API Actions
## requestPermission

Requests permission to use the API exposed by this plugin. This method does not require the API key, and is the only one that accepts requests from any origin; the other methods only accept requests from trusted origins, which are listed under webCorsOriginList in the add-on config. localhost is trusted by default.

Calling this method from an untrusted origin will display a popup in Anki asking the user whether they want to allow your origin to use the API; calls from trusted origins will return the result without displaying the popup. When denying permission, the user may also choose to ignore further permission requests from that origin. These origins end up in the ignoreOriginList, editable via the add-on config.

The result always contains the permission field, which in turn contains either the string granted or denied, corresponding to whether your origin is trusted. If your origin is trusted, the fields requireApiKey (true if required) and version will also be returned.

This should be the first call you make to make sure that your application and Anki-Connect are able to communicate properly with each other. New versions of Anki-Connect are backwards compatible; as long as you are using actions which are available in the reported Anki-Connect version or earlier, everything should work fine.
Sample request:

```
{
    "action": "requestPermission",
    "version": 6
}
```

Sample results:

```
{
    "result": {
        "permission": "granted",
        "requireApiKey": false,
        "version": 6
    },
    "error": null
}
```

```
{
    "result": {
        "permission": "denied"
    },
    "error": null
}
```

## version

Gets the version of the API exposed by this plugin. Currently versions 1 through 6 are defined.
Sample request:

```
{
    "action": "version",
    "version": 6
}
```

Sample result:

```
{
    "result": 6,
    "error": null
}
```

## apiReflect

Gets information about the AnkiConnect APIs available. The request supports the following params:
    scopes - An array of scopes to get reflection information about. The only currently supported value is "actions".
    actions - Either null or an array of API method names to check for. If the value is null, the result will list all of the available API actions. If the value is an array of strings, the result will only contain actions which were in this array.

The result will contain a list of which scopes were used and a value for each scope. For example, the "actions" scope will contain a "actions" property which contains a list of supported action names.
Sample request:

```
{
    "action": "apiReflect",
    "version": 6,
    "params": {
        "scopes": ["actions", "invalidType"],
        "actions": ["apiReflect", "invalidMethod"]
    }
}
```

Sample result:

```
{
    "result": {
        "scopes": ["actions"],
        "actions": ["apiReflect"]
    },
    "error": null
}
```


## getProfiles

Retrieve the list of profiles.
Sample request:

```
{
    "action": "getProfiles",
    "version": 6
}
```

Sample result:

```
{
    "result": ["User 1"],
    "error": null
}
```

## getActiveProfile

Retrieve the active profile.
Sample request:

```
{
    "action": "getActiveProfile",
    "version": 6
}
```

Sample result:

```
{
    "result": "User 1",
    "error": null
}
```

## loadProfile

Selects the profile specified in request.
Sample request:

```
{
    "action": "loadProfile",
    "version": 6,
    "params": {
        "name": "user1"
    }
}
```

Sample result:

```
{
    "result": true,
    "error": null
}
```

## exportPackage

Exports a given deck in .apkg format. Returns true if successful or false otherwise. The optional property includeSched (default is false) can be specified to include the cards' scheduling data.
Sample request:

```
{
    "action": "exportPackage",
    "version": 6,
    "params": {
        "deck": "Default",
        "path": "/data/Deck.apkg",
        "includeSched": true
    }
}
```

Sample result:

```
{
    "result": true,
    "error": null
}
```

## importPackage

Imports a file in .apkg format into the collection. Returns true if successful or false otherwise. Note that the file path is relative to Anki's collection.media folder, not to the client.
Sample request:

```
{
    "action": "importPackage",
    "version": 6,
    "params": {
        "path": "/data/Deck.apkg"
    }
}
```

Sample result:

```
{
    "result": true,
    "error": null
}
```

## reloadCollection

Tells anki to reload all data from the database.
Sample request:

```
{
    "action": "reloadCollection",
    "version": 6
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

# Model Actions
## modelNames

Gets the complete list of model names for the current user.
Sample request:

```
{
    "action": "modelNames",
    "version": 6
}
```

Sample result:

```
{
    "result": ["Basic", "Basic (and reversed card)"],
    "error": null
}
```

## modelNamesAndIds

Gets the complete list of model names and their corresponding IDs for the current user.
Sample request:

```
{
    "action": "modelNamesAndIds",
    "version": 6
}
```

Sample result:

```
{
    "result": {
        "Basic": 1483883011648,
        "Basic (and reversed card)": 1483883011644,
        "Basic (optional reversed card)": 1483883011631,
        "Cloze": 1483883011630
    },
    "error": null
}
```

## findModelsById

Gets a list of models for the provided model IDs from the current user.
Sample request:

```
{
    "action": "findModelsById",
    "version": 6,
    "params": {
        "modelIds": [1704387367119, 1704387398570]
    }
}
```

Sample result:

```
{
    "result": [
      {
        "id": 1704387367119,
        "name": "Basic",
        "type": 0,
        "mod": 1704387367,
        "usn": -1,
        "sortf": 0,
        "did": null,
        "tmpls": [
          {
            "name": "Card 1",
            "ord": 0,
            "qfmt": "{{Front}}",
            "afmt": "{{FrontSide}}\n\n<hr id=answer>\n\n{{Back}}",
            "bqfmt": "",
            "bafmt": "",
            "did": null,
            "bfont": "",
            "bsize": 0,
            "id": 9176047152973362695
          }
        ],
        "flds": [
          {
            "name": "Front",
            "ord": 0,
            "sticky": false,
            "rtl": false,
            "font": "Arial",
            "size": 20,
            "description": "",
            "plainText": false,
            "collapsed": false,
            "excludeFromSearch": false,
            "id": 2453723143453745216,
            "tag": null,
            "preventDeletion": false
          },
          {
            "name": "Back",
            "ord": 1,
            "sticky": false,
            "rtl": false,
            "font": "Arial",
            "size": 20,
            "description": "",
            "plainText": false,
            "collapsed": false,
            "excludeFromSearch": false,
            "id": -4853200230425436781,
            "tag": null,
            "preventDeletion": false
          }
        ],
        "css": ".card {\n    font-family: arial;\n    font-size: 20px;\n    text-align: center;\n    color: black;\n    background-color: white;\n}\n",
        "latexPre": "\\documentclass[12pt]{article}\n\\special{papersize=3in,5in}\n\\usepackage[utf8]{inputenc}\n\\usepackage{amssymb,amsmath}\n\\pagestyle{empty}\n\\setlength{\\parindent}{0in}\n\\begin{document}\n",
        "latexPost": "\\end{document}",
        "latexsvg": false,
        "req": [
          [
            0,
            "any",
            [
              0
            ]
          ]
        ],
        "originalStockKind": 1
      },
      {
        "id": 1704387398570,
        "name": "Basic (and reversed card)",
        "type": 0,
        "mod": 1704387398,
        "usn": -1,
        "sortf": 0,
        "did": null,
        "tmpls": [
          {
            "name": "Card 1",
            "ord": 0,
            "qfmt": "{{Front}}",
            "afmt": "{{FrontSide}}\n\n<hr id=answer>\n\n{{Back}}",
            "bqfmt": "",
            "bafmt": "",
            "did": null,
            "bfont": "",
            "bsize": 0,
            "id": 1689886528158874152
          },
          {
            "name": "Card 2",
            "ord": 1,
            "qfmt": "{{Back}}",
            "afmt": "{{FrontSide}}\n\n<hr id=answer>\n\n{{Front}}",
            "bqfmt": "",
            "bafmt": "",
            "did": null,
            "bfont": "",
            "bsize": 0,
            "id": -7839609225644824587
          }
        ],
        "flds": [
          {
            "name": "Front",
            "ord": 0,
            "sticky": false,
            "rtl": false,
            "font": "Arial",
            "size": 20,
            "description": "",
            "plainText": false,
            "collapsed": false,
            "excludeFromSearch": false,
            "id": -7787837672455357996,
            "tag": null,
            "preventDeletion": false
          },
          {
            "name": "Back",
            "ord": 1,
            "sticky": false,
            "rtl": false,
            "font": "Arial",
            "size": 20,
            "description": "",
            "plainText": false,
            "collapsed": false,
            "excludeFromSearch": false,
            "id": 6364828289839985081,
            "tag": null,
            "preventDeletion": false
          }
        ],
        "css": ".card {\n    font-family: arial;\n    font-size: 20px;\n    text-align: center;\n    color: black;\n    background-color: white;\n}\n",
        "latexPre": "\\documentclass[12pt]{article}\n\\special{papersize=3in,5in}\n\\usepackage[utf8]{inputenc}\n\\usepackage{amssymb,amsmath}\n\\pagestyle{empty}\n\\setlength{\\parindent}{0in}\n\\begin{document}\n",
        "latexPost": "\\end{document}",
        "latexsvg": false,
        "req": [
          [
            0,
            "any",
            [
              0
            ]
          ],
          [
            1,
            "any",
            [
              1
            ]
          ]
        ],
        "originalStockKind": 1
      }
    ],
    "error": null
}
```

## findModelsByName

Gets a list of models for the provided model names from the current user.
Sample request:

```
{
    "action": "findModelsByName",
    "version": 6,
    "params": {
        "modelNames": ["Basic", "Basic (and reversed card)"]
    }
}
```

Sample result:

```
{
    "result": [
      {
        "id": 1704387367119,
        "name": "Basic",
        "type": 0,
        "mod": 1704387367,
        "usn": -1,
        "sortf": 0,
        "did": null,
        "tmpls": [
          {
            "name": "Card 1",
            "ord": 0,
            "qfmt": "{{Front}}",
            "afmt": "{{FrontSide}}\n\n<hr id=answer>\n\n{{Back}}",
            "bqfmt": "",
            "bafmt": "",
            "did": null,
            "bfont": "",
            "bsize": 0,
            "id": 9176047152973362695
          }
        ],
        "flds": [
          {
            "name": "Front",
            "ord": 0,
            "sticky": false,
            "rtl": false,
            "font": "Arial",
            "size": 20,
            "description": "",
            "plainText": false,
            "collapsed": false,
            "excludeFromSearch": false,
            "id": 2453723143453745216,
            "tag": null,
            "preventDeletion": false
          },
          {
            "name": "Back",
            "ord": 1,
            "sticky": false,
            "rtl": false,
            "font": "Arial",
            "size": 20,
            "description": "",
            "plainText": false,
            "collapsed": false,
            "excludeFromSearch": false,
            "id": -4853200230425436781,
            "tag": null,
            "preventDeletion": false
          }
        ],
        "css": ".card {\n    font-family: arial;\n    font-size: 20px;\n    text-align: center;\n    color: black;\n    background-color: white;\n}\n",
        "latexPre": "\\documentclass[12pt]{article}\n\\special{papersize=3in,5in}\n\\usepackage[utf8]{inputenc}\n\\usepackage{amssymb,amsmath}\n\\pagestyle{empty}\n\\setlength{\\parindent}{0in}\n\\begin{document}\n",
        "latexPost": "\\end{document}",
        "latexsvg": false,
        "req": [
          [
            0,
            "any",
            [
              0
            ]
          ]
        ],
        "originalStockKind": 1
      },
      {
        "id": 1704387398570,
        "name": "Basic (and reversed card)",
        "type": 0,
        "mod": 1704387398,
        "usn": -1,
        "sortf": 0,
        "did": null,
        "tmpls": [
          {
            "name": "Card 1",
            "ord": 0,
            "qfmt": "{{Front}}",
            "afmt": "{{FrontSide}}\n\n<hr id=answer>\n\n{{Back}}",
            "bqfmt": "",
            "bafmt": "",
            "did": null,
            "bfont": "",
            "bsize": 0,
            "id": 1689886528158874152
          },
          {
            "name": "Card 2",
            "ord": 1,
            "qfmt": "{{Back}}",
            "afmt": "{{FrontSide}}\n\n<hr id=answer>\n\n{{Front}}",
            "bqfmt": "",
            "bafmt": "",
            "did": null,
            "bfont": "",
            "bsize": 0,
            "id": -7839609225644824587
          }
        ],
        "flds": [
          {
            "name": "Front",
            "ord": 0,
            "sticky": false,
            "rtl": false,
            "font": "Arial",
            "size": 20,
            "description": "",
            "plainText": false,
            "collapsed": false,
            "excludeFromSearch": false,
            "id": -7787837672455357996,
            "tag": null,
            "preventDeletion": false
          },
          {
            "name": "Back",
            "ord": 1,
            "sticky": false,
            "rtl": false,
            "font": "Arial",
            "size": 20,
            "description": "",
            "plainText": false,
            "collapsed": false,
            "excludeFromSearch": false,
            "id": 6364828289839985081,
            "tag": null,
            "preventDeletion": false
          }
        ],
        "css": ".card {\n    font-family: arial;\n    font-size: 20px;\n    text-align: center;\n    color: black;\n    background-color: white;\n}\n",
        "latexPre": "\\documentclass[12pt]{article}\n\\special{papersize=3in,5in}\n\\usepackage[utf8]{inputenc}\n\\usepackage{amssymb,amsmath}\n\\pagestyle{empty}\n\\setlength{\\parindent}{0in}\n\\begin{document}\n",
        "latexPost": "\\end{document}",
        "latexsvg": false,
        "req": [
          [
            0,
            "any",
            [
              0
            ]
          ],
          [
            1,
            "any",
            [
              1
            ]
          ]
        ],
        "originalStockKind": 1
      }
    ],
    "error": null
}
```

## modelFieldNames

Gets the complete list of field names for the provided model name.
Sample request:

```
{
    "action": "modelFieldNames",
    "version": 6,
    "params": {
        "modelName": "Basic"
    }
}
```

Sample result:

```
{
    "result": ["Front", "Back"],
    "error": null
}
```

## modelFieldDescriptions

Gets the complete list of field descriptions (the text seen in the gui editor when a field is empty) for the provided model name.
Sample request:

```
{
    "action": "modelFieldDescriptions",
    "version": 6,
    "params": {
        "modelName": "Basic"
    }
}
```

Sample result:

```
{
    "result": ["", ""],
    "error": null
}
```

## modelFieldFonts

Gets the complete list of fonts along with their font sizes.
Sample request:

```
{
    "action": "modelFieldFonts",
    "version": 6,
    "params": {
        "modelName": "Basic"
    }
}
```

Sample result:

```
{
    "result": {
        "Front": {
            "font": "Arial",
            "size": 20
        },
        "Back": {
            "font": "Arial",
            "size": 20
        }
    },
    "error": null
}
```

## modelFieldsOnTemplates

Returns an object indicating the fields on the question and answer side of each card template for the given model name. The question side is given first in each array.
Sample request:

```
{
    "action": "modelFieldsOnTemplates",
    "version": 6,
    "params": {
        "modelName": "Basic (and reversed card)"
    }
}
```

Sample result:

```
{
    "result": {
        "Card 1": [["Front"], ["Back"]],
        "Card 2": [["Back"], ["Front"]]
    },
    "error": null
}
```

## createModel

Creates a new model to be used in Anki. User must provide the modelName, inOrderFields and cardTemplates to be used in the model. There are optional fields css and isCloze. If not specified, css will use the default Anki css and isCloze will be equal to false. If isCloze is true then model will be created as Cloze.

Optionally the Name field can be provided for each entry of cardTemplates. By default the card names will be Card 1, Card 2, and so on.
Sample request:

```
{
    "action": "createModel",
    "version": 6,
    "params": {
        "modelName": "newModelName",
        "inOrderFields": ["Field1", "Field2", "Field3"],
        "css": "Optional CSS with default to builtin css",
        "isCloze": false,
        "cardTemplates": [
            {
                "Name": "My Card 1",
                "Front": "Front html {{Field1}}",
                "Back": "Back html  {{Field2}}"
            }
        ]
    }
}
```

Sample result:

```
{
    "result":{
        "sortf":0,
        "did":1,
        "latexPre":"\\documentclass[12pt]{article}\n\\special{papersize=3in,5in}\n\\usepackage[utf8]{inputenc}\n\\usepackage{amssymb,amsmath}\n\\pagestyle{empty}\n\\setlength{\\parindent}{0in}\n\\begin{document}\n",
        "latexPost":"\\end{document}",
        "mod":1551462107,
        "usn":-1,
        "vers":[

        ],
        "type":0,
        "css":".card {\n font-family: arial;\n font-size: 20px;\n text-align: center;\n color: black;\n background-color: white;\n}\n",
        "name":"TestApiModel",
        "flds":[
            {
                "name":"Field1",
                "ord":0,
                "sticky":false,
                "rtl":false,
                "font":"Arial",
                "size":20,
                "media":[

                ]
            },
            {
                "name":"Field2",
                "ord":1,
                "sticky":false,
                "rtl":false,
                "font":"Arial",
                "size":20,
                "media":[

                ]
            }
        ],
        "tmpls":[
            {
                "name":"My Card 1",
                "ord":0,
                "qfmt":"",
                "afmt":"This is the back of the card {{Field2}}",
                "did":null,
                "bqfmt":"",
                "bafmt":""
            }
        ],
        "tags":[

        ],
        "id":1551462107104,
        "req":[
            [
                0,
                "none",
                [

                ]
            ]
        ]
    },
    "error":null
}
```

## modelTemplates

Returns an object indicating the template content for each card connected to the provided model by name.
Sample request:

```
{
    "action": "modelTemplates",
    "version": 6,
    "params": {
        "modelName": "Basic (and reversed card)"
    }
}
```

Sample result:

```
{
    "result": {
        "Card 1": {
            "Front": "{{Front}}",
            "Back": "{{FrontSide}}\n\n<hr id=answer>\n\n{{Back}}"
        },
        "Card 2": {
            "Front": "{{Back}}",
            "Back": "{{FrontSide}}\n\n<hr id=answer>\n\n{{Front}}"
        }
    },
    "error": null
}
```

## modelStyling

Gets the CSS styling for the provided model by name.
Sample request:

```
{
    "action": "modelStyling",
    "version": 6,
    "params": {
        "modelName": "Basic (and reversed card)"
    }
}
```

Sample result:

```
{
    "result": {
        "css": ".card {\n font-family: arial;\n font-size: 20px;\n text-align: center;\n color: black;\n background-color: white;\n}\n"
    },
    "error": null
}
```

## updateModelTemplates

Modify the templates of an existing model by name. Only specifies cards and specified sides will be modified. If an existing card or side is not included in the request, it will be left unchanged.
Sample request:

```
{
    "action": "updateModelTemplates",
    "version": 6,
    "params": {
        "model": {
            "name": "Custom",
            "templates": {
                "Card 1": {
                    "Front": "{{Question}}?",
                    "Back": "{{Answer}}!"
                }
            }
        }
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## updateModelStyling

Modify the CSS styling of an existing model by name.
Sample request:

```
{
    "action": "updateModelStyling",
    "version": 6,
    "params": {
        "model": {
            "name": "Custom",
            "css": "p { color: blue; }"
        }
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## findAndReplaceInModels

Find and replace string in existing model by model name. Customise to replace in front, back or css by setting to true/false.
Sample request:

```
{
    "action": "findAndReplaceInModels",
    "version": 6,
    "params": {
        "model": {
            "modelName": "",
            "findText": "text_to_replace",
            "replaceText": "replace_with_text",
            "front": true,
            "back": true,
            "css": true
        }
    }
}
```

Sample result:

```
{
    "result": 1,
    "error": null
}
```

## modelTemplateRename

Renames a template in an existing model.
Sample request:

```
{
    "action": "modelTemplateRename",
    "version": 6,
    "params": {
        "modelName": "Basic",
        "oldTemplateName": "Card 1",
        "newTemplateName": "Card 1 renamed"
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## modelTemplateReposition

Repositions a template in an existing model.

The value of index starts at 0. For example, an index of 0 puts the template in the first position, and an index of 2 puts the template in the third position.
Sample request:

```
{
    "action": "modelTemplateReposition",
    "version": 6,
    "params": {
        "modelName": "Basic",
        "templateName": "Card 1",
        "index": 1
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## modelTemplateAdd

Adds a template to an existing model by name. If you want to update an existing template, use updateModelTemplates.
Sample request:

```
{
    "action": "modelTemplateAdd",
    "version": 6,
    "params": {
        "modelName": "Basic",
        "template": {
            "Name": "Card 3",
            "Front": "Front html {{Field1}}",
            "Back": "Back html {{Field2}}"
        }
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## modelTemplateRemove

Removes a template from an existing model.
Sample request:

```
{
    "action": "modelTemplateRemove",
    "version": 6,
    "params": {
        "modelName": "Basic",
        "templateName": "Card 1"
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## modelFieldRename

Rename the field name of a given model.
Sample request:

```
{
    "action": "modelFieldRename",
    "version": 6,
    "params": {
        "modelName": "Basic",
        "oldFieldName": "Front",
        "newFieldName": "FrontRenamed"
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## modelFieldReposition

Reposition the field within the field list of a given model.

The value of index starts at 0. For example, an index of 0 puts the field in the first position, and an index of 2 puts the field in the third position.
Sample request:

```
{
    "action": "modelFieldReposition",
    "version": 6,
    "params": {
        "modelName": "Basic",
        "fieldName": "Back",
        "index": 0
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## modelFieldAdd

Creates a new field within a given model.

Optionally, the index value can be provided, which works exactly the same as the index in modelFieldReposition. By default, the field is added to the end of the field list.
Sample request:

```
{
    "action": "modelFieldAdd",
    "version": 6,
    "params": {
        "modelName": "Basic",
        "fieldName": "NewField",
        "index": 0
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## modelFieldRemove

Deletes a field within a given model.
Sample request:

```
{
    "action": "modelFieldRemove",
    "version": 6,
    "params": {
        "modelName": "Basic",
        "fieldName": "Front"
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## modelFieldSetFont

Sets the font for a field within a given model.
Sample request:

```
{
    "action": "modelFieldSetFont",
    "version": 6,
    "params": {
        "modelName": "Basic",
        "fieldName": "Front",
        "font": "Courier"
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## modelFieldSetFontSize

Sets the font size for a field within a given model.
Sample request:

```
{
    "action": "modelFieldSetFontSize",
    "version": 6,
    "params": {
        "modelName": "Basic",
        "fieldName": "Front",
        "fontSize": 10
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## modelFieldSetDescription

Sets the description (the text seen in the gui editor when a field is empty) for a field within a given model.

Older versions of Anki (2.1.49 and below) do not have field descriptions. In that case, this will return with false.
Sample request:

```
{
    "action": "modelFieldSetDescription",
    "version": 6,
    "params": {
        "modelName": "Basic",
        "fieldName": "Front",
        "description": "example field description"
    }
}
```

Sample result:

```
{
    "result": true,
    "error": null
}
```

# Note Actions
## addNote

Creates a note using the given deck and model, with the provided field values and tags. Returns the identifier of the created note created on success, and null on failure.

Anki-Connect can download audio, video, and picture files and embed them in newly created notes. The corresponding audio, video, and picture note members are optional and can be omitted. If you choose to include any of them, they should contain a single object or an array of objects with the mandatory filename field and one of data, path or url. Refer to the documentation of storeMediaFile for an explanation of these fields. The skipHash field can be optionally provided to skip the inclusion of files with an MD5 hash that matches the provided value. This is useful for avoiding the saving of error pages and stub files. The fields member is a list of fields that should play audio or video, or show a picture when the card is displayed in Anki. The allowDuplicate member inside options group can be set to true to enable adding duplicate cards. Normally duplicate cards can not be added and trigger exception.

The duplicateScope member inside options can be used to specify the scope for which duplicates are checked. A value of "deck" will only check for duplicates in the target deck; any other value will check the entire collection.

The duplicateScopeOptions object can be used to specify some additional settings:
    duplicateScopeOptions.deckName will specify which deck to use for checking duplicates in. If undefined or null, the target deck will be used.
    duplicateScopeOptions.checkChildren will change whether or not duplicate cards are checked in child decks. The default value is false.
    duplicateScopeOptions.checkAllModels specifies whether duplicate checks are performed across all note types. The default value is false.
Sample request:

```
{
    "action": "addNote",
    "version": 6,
    "params": {
        "note": {
            "deckName": "Default",
            "modelName": "Basic",
            "fields": {
                "Front": "front content",
                "Back": "back content"
            },
            "options": {
                "allowDuplicate": false,
                "duplicateScope": "deck",
                "duplicateScopeOptions": {
                    "deckName": "Default",
                    "checkChildren": false,
                    "checkAllModels": false
                }
            },
            "tags": [
                "yomichan"
            ],
            "audio": [{
                "url": "https://assets.languagepod101.com/dictionary/japanese/audiomp3.php?kanji=猫&kana=ねこ",
                "filename": "yomichan_ねこ_猫.mp3",
                "skipHash": "7e2c2f954ef6051373ba916f000168dc",
                "fields": [
                    "Front"
                ]
            }],
            "video": [{
                "url": "https://cdn.videvo.net/videvo_files/video/free/2015-06/small_watermarked/Contador_Glam_preview.mp4",
                "filename": "countdown.mp4",
                "skipHash": "4117e8aab0d37534d9c8eac362388bbe",
                "fields": [
                    "Back"
                ]
            }],
            "picture": [{
                "url": "https://upload.wikimedia.org/wikipedia/commons/thumb/c/c7/A_black_cat_named_Tilly.jpg/220px-A_black_cat_named_Tilly.jpg",
                "filename": "black_cat.jpg",
                "skipHash": "8d6e4646dfae812bf39651b59d7429ce",
                "fields": [
                    "Back"
                ]
            }]
        }
    }
}
```

Sample result:

```
{
    "result": 1496198395707,
    "error": null
}
```

## addNotes

Creates multiple notes using the given deck and model, with the provided field values and tags. Returns an array of identifiers of the created notes. In the event of any errors, all errors are gathered and returned.

Please see the documentation for addNote for an explanation of objects in the notes array.
Sample request:

```
{
   "action":"addNotes",
   "version":6,
   "params":{
      "notes":[
         {
            "deckName":"College::PluginDev",
            "modelName":"non_existent_model",
            "fields":{
               "Front":"front",
               "Back":"bak"
            }
         },
         {
            "deckName":"College::PluginDev",
            "modelName":"Basic",
            "fields":{
               "Front":"front",
               "Back":"bak"
            }
         }
      ]
   }
}
```

Sample result:

```
{
   "result":null,
   "error":"['model was not found: non_existent_model']"
}
```

## canAddNotes

Accepts an array of objects which define parameters for candidate notes (see addNote) and returns an array of booleans indicating whether or not the parameters at the corresponding index could be used to create a new note.
Sample request:

```
{
    "action": "canAddNotes",
    "version": 6,
    "params": {
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
            }
        ]
    }
}
```

Sample result:

```
{
    "result": [true],
    "error": null
}
```

## canAddNotesWithErrorDetail

Accepts an array of objects which define parameters for candidate notes (see addNote) and returns an array of objects with fields canAdd and error.
    canAdd indicates whether or not the parameters at the corresponding index could be used to create a new note.
    error contains an explanation of why a note cannot be added.
Sample request:

```
{
    "action": "canAddNotesWithErrorDetail",
    "version": 6,
    "params": {
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
    }
}
```

Sample result:

```
{
    "result": [
        {
            "canAdd": false,
            "error": "cannot create note because it is a duplicate"
        },
        {
            "canAdd": true
        }
    ],
    "error": null
}
```

## updateNoteFields

Modify the fields of an existing note. You can also include audio, video, or picture files which will be added to the note with an optional audio, video, or picture property. Please see the documentation for addNote for an explanation of objects in the audio, video, or picture array.

    Warning: You must not be viewing the note that you are updating on your Anki browser, otherwise the fields will not update. See this issue for further details.

Sample request:

```
{
    "action": "updateNoteFields",
    "version": 6,
    "params": {
        "note": {
            "id": 1514547547030,
            "fields": {
                "Front": "new front content",
                "Back": "new back content"
            },
            "audio": [{
                "url": "https://assets.languagepod101.com/dictionary/japanese/audiomp3.php?kanji=猫&kana=ねこ",
                "filename": "yomichan_ねこ_猫.mp3",
                "skipHash": "7e2c2f954ef6051373ba916f000168dc",
                "fields": [
                    "Front"
                ]
            }]
        }
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## updateNote

Modify the fields and/or tags of an existing note. In other words, combines updateNoteFields and updateNoteTags. Please see their documentation for an explanation of all properties.

Either fields or tags property can be omitted without affecting the other. Thus valid requests to updateNoteFields also work with updateNote. The note must have the fields property in order to update the optional audio, video, or picture objects.

If neither fields nor tags are provided, the method will fail. Fields are updated first and are not rolled back if updating tags fails. Tags are not updated if updating fields fails.

    Warning You must not be viewing the note that you are updating on your Anki browser, otherwise the fields will not update. See this issue for further details.

Sample request:

```
{
    "action": "updateNote",
    "version": 6,
    "params": {
        "note": {
            "id": 1514547547030,
            "fields": {
                "Front": "new front content",
                "Back": "new back content"
            },
            "tags": ["new", "tags"]
        }
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## updateNoteModel

Update the model, fields, and tags of an existing note. This allows you to change the note's model, update its fields with new content, and set new tags.
Sample request:

```
{
    "action": "updateNoteModel",
    "version": 6,
    "params": {
        "note": {
            "id": 1514547547030,
            "modelName": "NewModel",
            "fields": {
                "NewField1": "new field 1",
                "NewField2": "new field 2",
                "NewField3": "new field 3"
            },
            "tags": ["new", "updated", "tags"]
        }
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## updateNoteTags

Set a note's tags by note ID. Old tags will be removed.
Sample request:

```
{
    "action": "updateNoteTags",
    "version": 6,
    "params": {
        "note": 1483959289817,
        "tags": ["european-languages"]
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## getNoteTags

Get a note's tags by note ID.
Sample request:

```
{
    "action": "getNoteTags",
    "version": 6,
    "params": {
        "note": 1483959289817
    }
}
```

Sample result:

```
{
    "result": ["european-languages"],
    "error": null
}
```

## addTags

Adds tags to notes by note ID.
Sample request:

```
{
    "action": "addTags",
    "version": 6,
    "params": {
        "notes": [1483959289817, 1483959291695],
        "tags": "european-languages"
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## removeTags

Remove tags from notes by note ID.
Sample request:

```
{
    "action": "removeTags",
    "version": 6,
    "params": {
        "notes": [1483959289817, 1483959291695],
        "tags": "european-languages"
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## getTags

Gets the complete list of tags for the current user.
Sample request:

```
{
    "action": "getTags",
    "version": 6
}
```

Sample result:

```
{
    "result": ["european-languages", "idioms"],
    "error": null
}
```

## clearUnusedTags

Clears all the unused tags in the notes for the current user.
Sample request:

```
{
    "action": "clearUnusedTags",
    "version": 6
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## replaceTags

Replace tags in notes by note ID.
Sample request:

```
{
    "action": "replaceTags",
    "version": 6,
    "params": {
        "notes": [1483959289817, 1483959291695],
        "tag_to_replace": "european-languages",
        "replace_with_tag": "french-languages"
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## replaceTagsInAllNotes

Replace tags in all the notes for the current user.
Sample request:

```
{
    "action": "replaceTagsInAllNotes",
    "version": 6,
    "params": {
        "tag_to_replace": "european-languages",
        "replace_with_tag": "french-languages"
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## findNotes

Returns an array of note IDs for a given query. Query syntax is documented here.
Sample request:

```
{
    "action": "findNotes",
    "version": 6,
    "params": {
        "query": "deck:current"
    }
}
```

Sample result:

```
{
    "result": [1483959289817, 1483959291695],
    "error": null
}
```

## notesInfo

Returns a list of objects containing for each note ID the note fields, tags, note type, modification time,the cards belonging to the note and the profile where the note was created.
Sample request (note ids):

```
{
    "action": "notesInfo",
    "version": 6,
    "params": {
        "notes": [1502298033753]
    }
}
```

Sample request (query):

```
{
    "action": "notesInfo",
    "version": 6,
    "params": {
        "query": "deck:current"
    }
}
```

Sample result:

```
{
    "result": [
        {
            "noteId":1502298033753,
            "profile": "User_1",
            "modelName": "Basic",
            "tags":["tag","another_tag"],
            "fields": {
                "Front": {"value": "front content", "order": 0},
                "Back": {"value": "back content", "order": 1}
            },
            "mod": 1718377864,
            "cards": [1498938915662]
        }
    ],
    "error": null
}
```

s

## notesModTime

Returns a list of objects containings for each note ID the modification time.
Sample request:

```
{
    "action": "notesModTime",
    "version": 6,
    "params": {
        "notes": [1502298033753]
    }
}
```

Sample result:

```
{
    "result": [
        {
            "noteId": 1498938915662,
            "mod": 1629454092
        }
    ],
    "error": null
}
```

## deleteNotes

Deletes notes with the given ids. If a note has several cards associated with it, all associated cards will be deleted.
Sample request:

```
{
    "action": "deleteNotes",
    "version": 6,
    "params": {
        "notes": [1502298033753]
    }
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

## removeEmptyNotes

Removes all the empty notes for the current user.
Sample request:

```
{
    "action": "removeEmptyNotes",
    "version": 6
}
```

Sample result:

```
{
    "result": null,
    "error": null
}
```

