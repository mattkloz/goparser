# GoParser Summary

This is a "discovery project" to possibly improve the performance a parser that I previously wrote in javascript and also learn to code in Go. I don't have any previous experience in Go so this is my first attempt to write anything usable with it.

## Background

As part of a larger mac-os application that I developed, I wrote an HTML and CSS parser in JavaScript that allowed users to feed in urls or example webapps or sites which were parsed and specific elements and their associated styles were indentified and recreated in the user's specific template. As an example, if you entered 'https://google.com,' the parser extracts all of the elements from the source and if an element (tag) matches a predefined tag (e.g. Text Input), then the element is recreated on the user's side, the styles sheets are applied to the element and then the specific styles associated with that element are extracted. This allowed users to recreate any site in their local design language by simply entering a url. The use case for this feature was to reduce the dev time needed to transform legacy PHP applications into Ruby-on-Rails apps running a Material Design frontend framework. 

## Goal
While the JS version works, it is very complex and not very fast. As the requirements were specific to my team's internal needs, there is a lot of complicated conditionals to translate specific elements to our design language (currently HAML). My goal with this Go project is to extract the parsing logic from the translation logic and speed up the parsing process significantly while also producing a parser that is user agnostic (applicable to any user or enterprise). I guess we'll find out how that turns out later on. Stay tuned!