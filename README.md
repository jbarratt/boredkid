### I'm boooored

This package programatically attempts to answer the common complaint, "$PARENT?! I'm Boooored."

The idea is

- create a catalog of ideas for kids to enjoy
- ideally, hook it up to alexa so kids can self-serve boredom relief ("Alexa, I have a bored kid.")
- use a generative methodology to make the suggestions more specific

By more specific, it means the difference between

"go do some art"

and

"go paint a picture of a dinosaur hugging a penguin for a grantparent"



The architectural idea as a starting point is to leverage go templates. In the above example, we'd randomly select a seed activity from the available options.

	"Go {{ artproject }} of a {{ entity }} {{ interaction }} {{ entity }} for a {{ audience }} "

Each of those names would be functions, which might return a random number in a certain range, return a string value, or even return a string which contained more template code.

To evaluate these properly, all that needs to happen is

	finalString := getSeed()
	for  {
		result := evaluate(finalString)
		if result == finalString {
			break
		}
	}

This enables simple recursive template fragments. As long as it continues to change from evaluation to evaluation, it is re-evaluated.

Initial example outputs:
		
	go for a walk in your bedroom
	run as fast as you can for 5 minutes
	have a dance party
	read a book to yourself
	play outside
	make a watercolor painting of the sunrise for a parent
	cook something, if it is ok with your parents
	ask a parent about kittens
	memorize your parent&#39;s phone numbers
	play tag outside
	do something helpful for a parent
	clean up the garden
	be creative
	read a book to yourself
	pretend the couch is a motorhome
	make up a play about a baby for someone who might need cheering up
	find a cool rock in the yard
	go for a walk in the kitchen
	clean up the laundry
	eat something you have never tried before
	do something helpful for a grandparent
	make an obstacle course and try and beat your own time going through it
	find a cool rock in the yard
	make up new rules for one of your board games
	memorize an important phone number
	make a comic book for the easter bunny
	go on a pretend camping trip in the living room
	make a comic book for your teacher
	make a musical instrument from recycling
