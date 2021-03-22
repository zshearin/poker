# poker

Current state: testing failures - I have it crashing after a given amount of time to test AWS functionality with crashing program

This creates a poker game and evaluates what each hand is.  Right now it is a basic http server that returns the outcome of a hand for four players.  In the json return object currently:
1. The Board in order (list of 5 Card objects: 5 cards - first 3 are flop, 4th is turn and 5th is the river)
2. Hands dealt to each player (a list of the players hands.  Each player's hand is a list of 2 card objects)
3. Hand results (Ordered list in the winning hands with the player number and a description of their hand)
