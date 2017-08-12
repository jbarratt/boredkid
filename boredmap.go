package main

func boredMap() map[string][]string {
	bored := make(map[string][]string)
	bored["seed"] = []string{
		"play {{ outsidegame }} outside",
		"play outside",
		"read a book to yourself",
		"read a book out loud to someone else",
		"be creative",
		"{{ exercise }} for {{ minutes }} minutes",
		"{{ exercise }} for {{ minutes }} minutes",
		"do something helpful for {{ familymember }}",
		"{{ artproject }} for {{ person }}",
		"{{ artproject }} for {{ person }}",
		"{{ artproject }} for {{ person }}",
		"{{ artproject }} for {{ person }}",
		"{{ artproject }} for {{ person }}",
		"{{ artproject }} for {{ person }}",
		"{{ artproject }} for {{ person }}",
		"{{ artproject }} for {{ person }}",
		"make up new rules for one of your board games",
		"go on a pretend camping trip in the living room",
		"pretend the couch is {{ vehicle }}",
		"find a cool rock in the yard",
		"make an obstacle course and try and beat your own time going through it",
		"play with {{ toy }}",
		"clean up {{ location }}",
		"go for a walk in {{ location }}",
		"make a musical instrument from recycling",
		"cook something, if it is ok with your parents",
		"eat something you have never tried before",
		"memorize an important phone number",
		"have a dance party",
		"ask {{ familymember }} about {{ subject }}",
	}
	bored["toy"] = []string{
		"lego",
		"a ball",
		"your newest toy",
		"one of your oldest toys",
		"dress up clothes",
		"an outside toy",
		"the couch cushions",
		"a board game",
		"a card game",
		"folding paper",
		"a puzzle",
		"play dough",
		"the sprinklers",
		"squirt guns",
	}
	bored["vehicle"] = []string{
		"a race car",
		"a rocket ship",
		"a tugboat",
		"a cruise ship",
		"a motorhome",
		"an airplane",
	}
	bored["location"] = []string{
		"your bedroom",
		"the kitchen",
		"the place you keep your toys",
		"the laundry",
		"the garden",
	}
	bored["familymember"] = []string{
		"a parent",
		"your brother or sister",
		"the family",
		"a grandparent",
	}
	bored["person"] = []string{
		"an old friend",
		"your best friend",
		"a neigbhor",
		"someone who might need cheering up",
		"santa",
		"the tooth fairy",
		"the easter bunny",
		"your teacher",
		"an old teacher",
	}
	bored["person"] = append(bored["person"], bored["familymember"]...)
	bored["artproject"] = []string{
		"make a card",
		"make a book",
		"make a comic book",
		"{{ medium }} of {{ subject }}",
		"make up a song about {{ subject }}",
		"write a poem about {{ subject }}",
		"make up a play about {{ subject }}",
	}
	bored["subject"] = []string{
		"a dinosaur",
		"kittens",
		"the sunrise",
		"a tree",
		"a baby",
		"your favorite t.v. show",
		"a character you like",
		"{{ person }}",
		"a time you laughed really hard",
		"a thunderstorm",
		"your favorite food",
		"someone who loves you",
		"a time you were sad",
		"a time you felt loved",
	}
	bored["medium"] = []string{
		"paint a picture",
		"make a crayon drawing",
		"make a marker drawing",
		"make a pencil drawing",
		"make a watercolor painting",
		"make a sidewalk chalk drawing",
		"make a finger painting",
	}
	bored["exercise"] = []string{
		"do jumping jacks",
		"run in place",
		"hold a plank",
		"do squats",
		"run",
		"jump rope",
		"jump as far as you can",
		"run as fast as you can",
	}
	bored["outsidegame"] = []string{
		"soccer",
		"baseball",
		"tag",
	}
	bored["minutes"] = []string{
		"2", "5", "10", "20",
	}
	return bored
}
