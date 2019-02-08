-- Put functions in this file to use them in several other scripts.
-- To get access to the functions, you need to put:
-- require "my_directory.my_file"
-- in any script using the functions.

--This file should be a way to create a method for displaying and running phonograms
--in the level scripts, and avoiding clutter and unnecessary repetition of code)

function comparisionOfPhonograms( letter1, letter2)
	Display letter1 and letter2
	Play sound letter 1
	If user selects letter 1
		Score += 1
		Play sound good job
	elseIf user selects letter 2
		Play sound woops
	Delay 6 seconds
end