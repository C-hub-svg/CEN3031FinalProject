

## Work Completed

 - Implemented Search functionality for Amazon (more to come)
 - Implemented wishlist 
 - Implemented Login page
 ## Potential Additional features
 - add to cart on store's site i.e when you press add to cart in our app, it will be waiting for you in your cart in the relevant storefron
 - additional stores
 - filtering capabilities

## Frontend Unit Tests
- loads webpage
- adds item to wishlist
## Backend Unit Tests
- queries "white dress"
- registsers with faulty username
- registers with faulty password
- regsietrs with valid username and password 
## API Documentation  
 - pushProduct(n, u, i, p, d)
	 - pushes product info into the data JSON
	 - n=name; the product's name
	 - u=url; the product's page
	 - I=image; the picture of the product
	 - p=price; the price of the product
	 - d=data; the JSON that will be sent to the client
 - SearchAmazon(query)
	 - query is a string passed in by the user on the frontend
	 - loads webpage for the Amazon search for query
	 - reads price, name, and image for search results
	 - pushes product info into JSON
	 - returns JSON of product info
 - SearchNord(query)
- Search(query)
	- manipulates the entered string into the proper format by replacing delimiters i.e. "white dress" becomes "white%20dress"
	- runs Search functions for all supported stores e.g. SearchAmazon()
	- returns JSON of all results from lower search function. 
	
*note: may be possible to combine search functions, but there's so much variation between website's source code that it's difficult to see how*
