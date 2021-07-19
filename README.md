# Crypto Bot

This bot gets a currency pair (cryptos currency and fiat currency) from Uphold open API (at https://api.uphold.com/v0/ticker/ [currency pair]) - more info here: https://uphold.com/en/developer/api/documentation/#tickers.

You can be aware of your favourites currencies pairs by setting alerts to them.

## How to install

### Pre-requisites

- You must have Golang installed to compile it (more about how to install Golang at https://www.golang.org/doc/install);
- You must have Git client installed to check out this repository (more about how to install Git client at https://git-scm.com/);
You must have internet access.

### How to install Crypto Bot

With your Git client installed and configured, you must clone the repository like this:

>git clone https://github.com/lgmontenegro/crypto_bot.git crypto_bot

The command above will create a folder called crypto_bot and will copy all repository code into it.

After you have cloned the repository, you will change to the cloned repository folder and install the code dependencies as below:

>cd cripto_bot 
>go build .

These commands will probably generate a file called crypto_bot (in Linux or Mac) or cripto_bot.exe (in Windows) into the repository folder. You can move this single file to any directory you wish.

### Configuring Crypto Bot

Crypto Bot can be configured with a JSON file or with a flag configuration at the command line.

#### JSON Configuration File

The configuration file has to be in the same folder where the crypto_bot file will be executed

The JSON configuration file has 5 entries as described below:

- times(integer): or frequency, or fetch interval, is the time in seconds between currency pair data fetch;
- url(string): the resource locator where the API is hosted;
- endpoint(string): the path to reach the API;
- pairs(array of string): the pair, or pairs you want to fetch;
- alerts (array of alerts): the alerts you want for each configured pair.

We will discuss some more about alerts below.

Here is an example of how a JSON configuration file looks like:

``` 
{
    "times": 5,
    "url": "https://api.uphold.com/",
    "endpoint": "v0/ticker/",
    "pairs": ["USD-ADA", "USD-BTC"],        
    "alerts": [
        {
            "pair": "ADA-USD",
            "float": 0.01
        }
    ]
}
```

You can find a functional example inside the repository folder and edit it.

Another way is setting these entries via command line with flags. Running `crypto_bot -h` at your command line you will see this:

```
A bot to watch crypto pair values

Usage:
  crypto_bot [flags]

Flags:
  -e, --endpoint string     EndPoint address for the ticker
  -h, --help                help for crypto_bot
  -i, --interval duration   Time in seconds for each ticker recovery
  -p, --pairs strings       coins pairs
  -u, --url string          URL to access the endpoint
  -v, --verbose             verbose output
With the flags below, you can skip these entries into the JSON configuration files.
```

You can also use both. In this case, the command line flags will subscribe to the JSON configuration file entries.

#### The Alerts

Unfortunately, Alerts can only be set at JSON configuration file by now.

Alerts are an array of Alert objects:

```
{
	"pair": "USD-BTC",
	"float": 0.01
}
```

Where `pair` is one of the currency pairs you choose to be fetched and `float` is the percentual difference between the first ticker fetched of this pair and the last one.

If this difference is reached, the bot will alert, writing a message with some details as:

```
{
	Ticker: the ticker that generates the alert
	Oscillation: the difference value
	CreatedAt: date and hour of the alarm
	Direction: the sign of the difference (up or down)
}
```

Any issues feel free to contact me!