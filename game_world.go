package main

func GameWorld(hero *HeroSheet) func(Event) {
	// tried to do this async with channels but not working.
	// good enough for now
	return func(e Event) {
		if e.Message == "Won Fight" {
			hero.Xp = hero.Xp + 10;
		} else if e.Message == "Lost Fight" {
			// becomes next generation?
		}
	}
}
