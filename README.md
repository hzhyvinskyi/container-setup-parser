# `Container Setup Parser`

Container Setup Parser provides parsing function which takes a string representation of nested containers and checks if that string has a correctly defined structure.

###### Validation Rules:
- There can be 3 types of brackets: `(),[],{}`. Other characters should be ignored
- Container can contain other containers
- An opened container must always be closed
- Child container can't breach the border of a parent container