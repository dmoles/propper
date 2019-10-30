# propper

A simplistic implementation of Vladimir Propp’s _Morphology of the Folktale_,
based on [this exegesis](http://bactra.org/reviews/propp-morphology.html) by
Cosma Shalizi, to generate a random folktale outline.

## Usage

```sh
propper [-s seed]
```

Pass the same seed to repeat generation of the same “tale”.

Example output:

```
seed: 1572475950978442000
functions: ABC↑DEFGHJIK↓PrRsoLQExTUW
```

```
One member of a family lacks something
Misfortune or lack is made known; the hero is approached with a request; he is allowed to go
The seeker agrees to counteraction
The hero leaves home
The hero is attacked, which prepares the way for receiving a magical agent
The hero reacts to the actions of the future donor
The hero acquires the use of a magical agent
The hero is led to the whereabouts of an object of search
The hero and the villain join in direct combat
The hero is branded or marked
The villain is defeated
The initial misfortune or lack is liquidated
The hero returns
The hero is pursued
Rescue of the hero from pursuit
The hero, unrecognized, in another country
A false hero presents unfounded claims
The hero is recognized
The false hero or villain is exposed
The hero is given a new appearance
The villain is punished
The hero is married and ascends the throne
```

## Simplifying assumptions

- where Propp presents an implicit choice, e.g. **_C_**, “The seeker agrees to or decides upon counteraction”,
  this is decomposed into branches of equal probability:
  1. “The seeker agrees to counteraction” (**_C₁_**)
  2. “The seeker decides upon counteraction” (**_C₂_**)
- for the explicit choice between **_A_** “The villain causes harm or injury to a member of a family”
  and **_a_** “One member of a family either lacks something or desires to have something”, we unpack
  the implicit choice and assume all branches have equal probability:
  1. “The villain causes harm or injury to a member of a family” (**_A_**)
  2. “One member of a family lacks something” (**_a₁_**)
  3. “One member of a family desires to have something” (**_a₂_**)

## Programmer laziness

Where we make these implicit choices, we don't bother to document which branch was taken
in the `functions:` output -- e.g. instead of **_A_**, **_a₁_**, and **_a₂_** we always output **_A_**.

## Open questions

- What is the significance of the em dash (**—**) in, e.g., **_ABC↑FH−IK↓LM−NQExUW_**?
- Where and how should we interpolate Propp's "less essential" functions **_β_** — **_θ_**?

## TO DO:

- Expand each "function" to its concrete sub-types, e.g. from “The hero is pursued” to
  “the pursuer flies after the hero”, “The puruser tries to gnaw through the tree in which the hero
  is taking refuge”, etc.