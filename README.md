This is a basic implementation of a command line based snake game. The snake head is represented by the & character while the rest of the body is represented by # character. Food is represented by o character.

To run, clone the project using: git clone https://github.com/gkaps/skycoin_snake_game_golang.git

In the cloned project folder, run the game using: ./game for Mac or Linux or game for Windows

You can pass specify board size by passing command line arguments for height and width. E.g. to specify 20 by 30: game 20 30

This solution contains the functions (except from main) InitializeGameState, DisplayGameState, RequireNewMove, GetNewFoodLocation and UpdateGameState. 

The function InitializeGameState takes as inputs dimensions given from program execution arguments (if any). If there are no arguments, the default board dimensions are used. It returns the initial game state.

The function DisplayGameState takes snake game state as input and prints the state to terminal

The function RequireNewMove asks inputs from buffer that are matched to possible snake directions. The accepted characters are W for upwards, S for downwards, D for rightwards and A for leftwards. In any other case, the input character is translated to empty character "". The function returns string "W" "A" "S" "D" or "".

The function GetNewFoodLocation takes as input a snake game state and returns a random food location.

The function UpdateGameState takes as input the current state and move, it updates the SnakeGameState and returns a boolean called ok that is false for the Game Over move.

After running, the controls to play are presented on the right side of the board.

W - up | S - down | A - left | D - right and hit enter (case insensitive). If you hit enter without choosing or choosing other character, the snake continues on current direction Running into a wall or into snake body results in Game Over. 
The dead snake head is is represented by the X character

The task took me 3 days, working 2-4 hrs a day to complete.