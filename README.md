# yaml2AnkiConnect
This is a first draft at a simple yaml -> Anki deck, using Go and AnkiConnect. Updates to follow.

## Setup
1. Install Anki for your platform https://apps.ankiweb.net/ (consider using Flatpak if you're using Linux)
2. Run Anki and install AnkiConnect (the addon code is 2055492159)
3. Compile the yaml2ankiconnect binary: `go build`

## YAML format
```
models:
  ClozeModel:
    fields: ["Text", "Extra"]
    templates:
      - name: "Cloze Deletion"
        front: "{{cloze:Text}}"
        back: "{{Text}}<br><br>{{Extra}}"

  ImageModel:
    fields: ["Question", "Answer", "Image"]
    templates:
      - name: "Image Card"
        front: "{{Question}}<br><img src='{{Image}}'>"
        back: "{{Answer}}"

decks:
  ScienceDeck:
    modelName: "ClozeModel"
    cards:
      - Text: "Water is composed of {{c1::hydrogen}} and {{c1::oxygen}}."
        Extra: "H₂O is the chemical formula."
        tags: ["chemistry"]

  GeographyDeck:
    modelName: "ImageModel"
    cards:
      - Question: "What landmark is this?"
        Answer: "Eiffel Tower, Paris."
        Image: "eiffel_tower.jpg"
        media:
          - "eiffel_tower.jpg"
        tags: ["geography"]

```

## To import your deck...
TODO
