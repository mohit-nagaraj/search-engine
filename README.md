# search-engine
Build a full-text search engine in Golang using an inverted index approach that scales well for small to mid-scale applications, allowing quick search through a large dataset.
get the data dump from here - https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz

## Notes
lolia search is very expensive it's a fully managed service and elastic search if you even if you host it on your own it's expensive and if you buy the managed service it's very expensive as well so if if you're not building something really huge this approach is going to be really really good for you okay uh now the way we'll build this project will be in a very modular way so that let's say if you've built an e-commerce platform in goang a blogging platform or a restaurant management application right

why do we need a to to create a full-blown full Tech search engine the problem with this is that this approach doesn't scale so it takes close to 2 seconds for 600,000 docs like this these two approaches but what happens when we have more than 10 million or or approximately 10 million docs right this won't scale at all it it it'll just keep the time will keep increasing to you know 30 seconds 40 sec

this is why the scalable approach that we have today is called the inverted Index this means that we'll take the text text basically meaning the complete data data set that we have and we'll pre-process it and we'll create something called as the inverted index from that text to show you an example of what this looks like imagine that we've taken now just three documents from the entire data set so this is what the three documents look like now this is like a very shrunken version of or or like a

## Three-step process

1. Tokenization splitting the text on a word boundary and removing punctuation marks

2. Filtering has three stages in it so the first stage is lower case so in order to make the search case insensitive because if you're searching, need to be normalized. Dropping your common words. Any document would match the stop words. Third one is stemming. Stemming we going to be using the snowball stem Library. 

3. Searching: just search through the index and we'll find where, arrays which says that the the word small is in all of these documents the word wild is in all these documents and the word cat is on all these documents and this is something called as a disjointed set,  intersection of these of these sets, get the doc which contains it