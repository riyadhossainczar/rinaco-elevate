package main

var planTabs = []string{"Routine", "Phases", "Urge", "Body", "Mindset"}

var routine = []entry{
	{"5:30 AM", "Wake up", "Don't touch your phone. Just sit up. First 10 days are rough, then automatic."},
	{"5:30 – 6:00", "Cold water + sit quietly", "No meditation needed. 5 minutes of silence. Then one glass of cold water. Brain's signal — day has started."},
	{"6:00 – 6:30", "Walk outside", "30 minutes every day. No gym. This one thing resets dopamine naturally — no tricks."},
	{"6:30 – 7:30", "Cold shower + breakfast", "Cold shower reduces urges. Don't skip breakfast, whatever you have, eat it."},
	{"7:30 – 12:00", "Main work block", "Coding, learning, building. Phone on silent. Most valuable time of the day."},
	{"12:00 – 1:00", "Lunch + rest max 20 min", "More than 20 min lying down makes evenings restless — that's when urges hit hardest."},
	{"1:00 – 5:00", "Second work block", "If focus won't come, read. No phone scrolling — time drain and trigger at the same time."},
	{"5:00 – 7:00", "Go outside, not alone at home", "Evening 5–7 is the most dangerous window. Tea, people, anything outside the room."},
	{"7:00 – 9:00", "Dinner + light review", "Write 3 lines about your day. Not a diary, just 3 lines."},
	{"10:00 PM", "Phone goes to another room", "This one rule cuts nighttime relapse by 80%. Another room. Not face-down beside you."},
	{"10:30 PM", "Sleep", "Under 7 hours = 3x stronger urges next day. Sleep is your biggest weapon."},
}

var phases = []entry{
	{"Day 1–7", "hardest stretch",
		"Headache, restlessness, urges hitting hard.\n→ Don't be alone in the dark at night\n→ Phone in another room\n→ Walk 10 minutes if urge comes"},
	{"Day 8–21", "brain fights back",
		"Images pop in your head. Normal — brain losing grip.\n→ Start one new skill\n→ Go outside if bored\n→ Cut screen time"},
	{"Day 22–45", "first real relief",
		"Sleep better. Head clearer. Watch for 'just once'.\n→ Flatline may come — it passes\n→ Keep routine\n→ Small reward"},
	{"Day 46–90", "you'll see the change",
		"Focus returns. Memory improves. Overconfidence kills.\n→ Real goal — project or exam\n→ More people\n→ 10 PM rule"},
	{"Day 91–180", "40–50% recovered",
		"Discipline returning. Stress raises risk.\n→ Stress = walk, not sit alone\n→ Keep writing\n→ Look back 3 months"},
	{"Day 181–365+", "60–70% recovered",
		"Brain close to normal. Real joy from real things.\n→ Build something real\n→ Never forget where you came from"},
}

var urge = []entry{
	{"1", "Stand up", "Sitting makes worse. Stand up — step one."},
	{"2", "Cold water", "Drink. Splash face. Interrupts brain."},
	{"3", "Put phone down", "Away. Instagram, YouTube amplify trigger."},
	{"4", "Go outside", "5 minutes. Urges fade without feeding."},
	{"5", "Write something", "How you feel. Writing = prefrontal cortex."},
}

var body = []entry{
	{"Sleep", "10:30 PM – 5:30 AM", "7+ hours. Protect sleep."},
	{"Water", "8 glasses a day", "Dehydration = urge feeling."},
	{"Food", "Eggs, banana, nuts", "Zinc helps dopamine."},
	{"Walking", "30 min daily", "More than gym."},
	{"Screen curfew", "1 hour before bed", "Blue light ruins sleep."},
}

var mindset = []entry{
	{"Relapse doesn't mean it's over", "", "Start next day. Every day counts."},
	{"Stop shaming yourself", "", "Chemical issue. You're fixing it."},
	{"Track small progress", "", "Sleep, focus, feeling. Small wins."},
	{"One day at a time", "", "Get through today. Just today."},
	{"Who do you want to be", "", "Picture clearly. Do that now."},
}

var planSections = [][]entry{routine, phases, urge, body, mindset}
