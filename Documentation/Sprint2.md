# Sprint 2 Backend

## Primary Updates
We decided to implement a more "morally" constructive initiative within our program. The platform will still be used as a means of "gambling", of sorts, but it will have the added component of stock market connectivity. This way, users can develop skill sets while competing against each other, and also have the opportunity to gain experience in trading. This, we felt, allowed for an inherently more constructive and educative experience.

### Stock Market Connectivity
To facilitate trades from the eColosseum platform, we made use of Alpaca's Trade API. For the sake of testing, we assumed that the user already has an API key with Alpaca, but in the future, we will implement a way for users to create their own API keys. The user will then be able to connect their Alpaca account to the eColosseum platform, and then be able to trade stocks from the platform.
Implementation of the GUI for trading in eColosseum is simple: we will include input boxes to specify the stock, number of shares, price, and type of order (market or limit).
There should ideally be a way to view the user's current portfolio, but this is not a priority for the time being.
There should also be error handling for invalid stock ticker inputs, but for now we will assume the input ticker is the intended one. 

### Database API
We will be using a SQLite database to store player information. We need persistence for scores, "cash balance", and basic user information. This database will be seperate from that of the user's stock portfolio. 

Documentation for the E-Colosseum can be found at
* [openai.yaml](/backend/openai.yaml)
* [Swagger](https://app.swaggerhub.com/apis/b-cheek/E-Colosseum-API/1.0.0)

# Sprint 2 Frontend

## Primary Updates
To support the functionality of maintaining stock market trading, our opening page gives users the option to login or signup. There is also a prompt for the user to play as guest, if the user chooses to compete in cognitive brain games without the pressure of trading. After the user chooses one of the three options, they can access the homepage. This is where the user can access a variety of brain games and augment their understanding of stock market trading.

## Unit Tests / Cypress Tests
