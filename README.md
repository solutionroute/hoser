[![Build Status](https://travis-ci.org/solutionroute/hoser.svg?branch=master)](https://travis-ci.org/solutionroute/hoser) [![GoDoc](https://godoc.org/github.com/solutionroute/hoser?status.svg)](https://godoc.org/github.com/solutionroute/hoser) [![Coverage Status](https://coveralls.io/repos/github/solutionroute/hoser/badge.svg?branch=master)](https://coveralls.io/github/solutionroute/hoser?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/solutionroute/hoser)](https://goreportcard.com/report/github.com/solutionroute/hoser)

# hoser
Simple user and user rights management - a Go package

Package hoser impelements a Bolthold (BoltDB) store (easy to add other stores) implementation of a User
interface and a to-be growing variety of tools. The HTTP service and any related middleware is/will be compatible 
with the Golang standard http package and routers/muxes like Chi which follow the same signature.

## What's a "hoser"?

It's inappropriate to consider system users as *lusers*, but it happens, and a **hoser** is akin to the same idea. It's all in fun
and in short of a short name for this package. If you don't like it, **take off, eh**?

![Bob and Doug Mackenzie characters from SCTV show The Great White North](https://upload.wikimedia.org/wikipedia/en/2/28/Bob_and_Doug_McKenzie.jpg)

[Wikipedia](https://en.wikipedia.org/wiki/Bob_and_Doug_McKenzie) provides some
background; [this CBC
article](https://www.cbc.ca/news/entertainment/how-s-it-going-eh-bob-and-doug-mckenzie-help-raise-325k-in-special-show-1.4210544)
has a worthwhile quote:

> The Bob and Doug bit started out in the 1980s as a bit of a rebellion against
> a network order to include more Canadian content, but ended up becoming
> a pop-culture sensation  ensuring "eh" entered the Canadian lexicon.  
>
> For that, Thomas is sorry.  
> 
> "I apologize to the country, to everyone individually in Canada," Thomas said
> with a laugh. "I'm so sorry we brought hosers and 'take off' into the
> linguistic compendium of Canada."

